package action

import (
	"time"

	"github.com/astaxie/beego/orm"
	"github.com/pydr/acs/model"
)

func registerUserModels() {
	orm.RegisterModel(new(model.User))
}

func UserStatusSwitcher(status int32) string {
	switch status {
	case model.UnVerify:
		return "未认证"
	case model.Verified:
		return "已认证"
	default:
		return "未知状态"
	}
}

// 创建用户
func InsertUser(user *model.User) (int64, error) {
	o := orm.NewOrm()
	err := o.Begin()
	if err != nil {
		return 0, err
	}
	id, err := o.Insert(user)
	if err != nil {
		o.Rollback()
		return 0, err
	}
	m2m := o.QueryM2M(user, "Roles")
	for _, r := range user.Roles {
		if _, err = m2m.Add(r); err != nil {
			o.Rollback()
			return 0, err
		}
	}
	m2m = o.QueryM2M(user, "Applications")
	for _, a := range user.Applications {
		if _, err = m2m.Add(a); err != nil {
			o.Rollback()
			return 0, err
		}
	}

	return id, o.Commit()
}

// 删除用户
func DelUser(id int64) error {
	o := orm.NewOrm()
	err := o.Begin()
	if err != nil {
		return err
	}
	user := &model.User{Id: id}
	if err = o.Read(user); err != nil {
		o.Rollback()
		return err
	}

	user.Deleted = true
	if _, err = o.Update(user, "Deleted"); err != nil {
		o.Rollback()
		return err
	}
	if err = o.Read(user); err != nil {
		o.Rollback()
		return err
	}
	m2m := o.QueryM2M(user, "Roles")
	if _, err = m2m.Clear(); err != nil {
		o.Rollback()
		return err
	}

	return o.Commit()
}

// 更新用户信息
func UpdateUser(user *model.User) (*model.User, error) {
	o := orm.NewOrm()
	if _, err := o.Update(user, "Password", "Email", "Mobile"); err != nil {
		return nil, err
	}
	return user, nil
}

// 更新用户的角色
func UpdateUserRoles(user *model.User) (*model.User, error) {
	o := orm.NewOrm()
	err := o.Begin()
	if err != nil {
		return nil, err
	}
	m2m := o.QueryM2M(user, "Roles")
	_, err = m2m.Clear()
	if err != nil {
		o.Rollback()
		return nil, err
	}
	for _, r := range user.Roles {
		if _, err = m2m.Add(r); err != nil {
			o.Rollback()
			return nil, err
		}
	}
	if err = o.Commit(); err != nil {
		o.Rollback()
		return nil, err
	}
	return user, nil
}

// 更新用户的终端
func UpdateUserApplications(user *model.User) (*model.User, error) {
	o := orm.NewOrm()
	err := o.Begin()
	if err != nil {
		return nil, err
	}
	m2m := o.QueryM2M(user, "Applications")
	_, err = m2m.Clear()
	if err != nil {
		o.Rollback()
		return nil, err
	}
	for _, a := range user.Applications {
		if _, err = m2m.Add(a); err != nil {
			o.Rollback()
			return nil, err
		}
	}
	if err = o.Commit(); err != nil {
		o.Rollback()
		return nil, err
	}
	return user, nil
}

// 根据用户名查询
func GetUserByUsername(username string) (ret *model.User, err error) {
	ret = new(model.User)
	o := orm.NewOrm()
	err = o.QueryTable(new(model.User)).
		Filter("Username", username).
		Filter("Deleted", false).
		RelatedSel().
		One(ret)
	if err != nil {
		return
	}
	if _, err = o.LoadRelated(ret, "Roles"); err != nil {
		return
	}
	if _, err = o.LoadRelated(ret, "Applications"); err != nil {
		return
	}
	return
}

// 根据手机号查询
func GetUserByMobile(mobile string) (ret *model.User, err error) {
	ret = new(model.User)
	o := orm.NewOrm()
	err = o.QueryTable(new(model.User)).
		Filter("Mobile", mobile).
		Filter("Deleted", false).
		RelatedSel().
		One(ret)
	if err != nil {
		return
	}
	if _, err = o.LoadRelated(ret, "Roles"); err != nil {
		return
	}
	if _, err = o.LoadRelated(ret, "Applications"); err != nil {
		return
	}
	return
}

// 根据邮箱查询
func GetUserByMail(mail string) (ret *model.User, err error) {
	ret = new(model.User)
	o := orm.NewOrm()
	err = o.QueryTable(new(model.User)).
		Filter("Email", mail).
		Filter("Deleted", false).
		RelatedSel().
		One(ret)
	if err != nil {
		return
	}
	if _, err = o.LoadRelated(ret, "Roles"); err != nil {
		return
	}
	if _, err = o.LoadRelated(ret, "Applications"); err != nil {
		return
	}
	return
}

func GetUserLikeUsernameAndMobile(username string, mobile string) (ret []*model.User, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(model.User))
	if username != "" {
		qs = qs.Filter("username__contains", username)
	}
	if mobile != "" {
		qs = qs.Filter("mobile__contains", mobile)
	}
	_, err = qs.All(ret)
	if err != nil {
		return
	}
	if _, err = o.LoadRelated(ret, "Roles"); err != nil {
		return
	}
	if _, err = o.LoadRelated(ret, "Applications"); err != nil {
		return
	}
	return
}

// 根据id查询用户
func GetUserById(id int64) (ret *model.User, err error) {
	ret = new(model.User)
	o := orm.NewOrm()
	err = o.QueryTable(new(model.User)).
		Filter("Id", id).
		Filter("Deleted", false).
		RelatedSel().
		One(ret)
	if err != nil {
		return
	}
	if _, err = o.LoadRelated(ret, "Roles"); err != nil {
		return
	}
	if _, err = o.LoadRelated(ret, "Applications"); err != nil {
		return
	}
	return
}

// 查询所有用户
func GetUsers(offset, page int32) ([]*model.User, int64, error) {
	var ret []*model.User
	o := orm.NewOrm()
	page--
	_, err := o.QueryTable(new(model.User)).
		Filter("Deleted", false).
		OrderBy("-Created").
		Limit(offset, offset*page).
		RelatedSel().
		All(&ret)
	if err != nil {
		return nil, 0, err
	}
	total, err := o.QueryTable(new(model.User)).
		Filter("Deleted", false).
		Count()
	return ret, total, err
}

// 查询用户名是否存在
func IsUsernameExist(username string) bool {
	o := orm.NewOrm()
	return o.QueryTable(new(model.User)).Filter("Username", username).Exist()
}

// 查询mobile是否存在
func IsMobileExist(mobile string) bool {
	o := orm.NewOrm()
	return o.QueryTable(new(model.User)).Filter("Mobile", mobile).Exist()
}

// 更新用户状态
func UpdateUserStatus(id int64, status int32) (err error) {
	o := orm.NewOrm()
	_, err = o.Raw("UPDATE user SET status = ?, updated = ? WHERE id = ?", status, time.Now(), id).Exec()
	return
}
