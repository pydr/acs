package action

import (
	"github.com/astaxie/beego/orm"
	"github.com/pydr/acs/model"
)

func registerApplicationModels() {
	orm.RegisterModel(new(model.Application))
}

// 创建新终端
func InsertApplication(app *model.Application) (int64, error) {
	o := orm.NewOrm()
	if err := o.Begin(); err != nil {
		return 0, err
	}
	id, err := o.Insert(app)
	if err != nil {
		o.Rollback()
		return 0, err
	}
	m2m := o.QueryM2M(app, "Permissions")
	for _, p := range app.Permissions {
		if _, err = m2m.Add(p); err != nil {
			o.Rollback()
			return 0, err
		}
	}

	return id, o.Commit()
}

// 切换终端状态
func SwitchApplicationStatus(id int64) error {
	o := orm.NewOrm()
	a := &model.Application{Id: id}
	if err := o.Read(a); err != nil {
		a.Deleted = !a.Deleted
		if _, err = o.Update(a, "Deleted"); err != nil {
			return err
		}
	}

	return nil
}

// 更新终端
func UpdateApplication(app *model.Application) (*model.Application, error) {
	o := orm.NewOrm()
	_, err := o.Update(app, "Name", "Comment")
	return app, err
}

// 根据id查询终端
func GetApplicationById(id int64) (ret *model.Application, err error) {
	ret = new(model.Application)
	o := orm.NewOrm()
	err = o.QueryTable(new(model.Application)).
		Filter("id", id).
		Filter("Deleted", false).
		One(ret)
	if err != nil {
		return
	}
	if _, err = o.LoadRelated(ret, "Groups"); err != nil {
		return
	}
	if _, err = o.LoadRelated(ret, "Roles"); err != nil {
		return
	}
	if _, err = o.LoadRelated(ret, "Permissions"); err != nil {
		return
	}
	return
}

// 根据appid查询终端
func GetApplicationByAppid(appid string) (ret *model.Application, err error) {
	ret = new(model.Application)
	o := orm.NewOrm()
	err = o.QueryTable(new(model.Application)).
		Filter("Appid", appid).
		Filter("Deleted", false).
		One(ret)
	return
}

// 获取所有终端
func GetApplications(offset, page int32) ([]*model.Application, int64, error) {
	var ret []*model.Application
	page--
	o := orm.NewOrm()
	_, err := o.QueryTable(new(model.Application)).
		Limit(offset, offset*page).
		All(&ret)
	if err != nil {
		return nil, 0, err
	}

	total, err := o.QueryTable(new(model.Application)).Count()
	return ret, total, err
}

// 更新终端权限
func UpdateApplicationPermissions(app *model.Application) (*model.Application, error) {
	o := orm.NewOrm()
	if err := o.Begin(); err != nil {
		return nil, err
	}
	m2m := o.QueryM2M(app, "Permissions")
	_, err := m2m.Clear()
	if err != nil {
		o.Rollback()
		return nil, err
	}
	for _, p := range app.Permissions {
		if _, err := m2m.Add(p); err != nil {
			o.Rollback()
			return nil, err
		}
	}

	return app, o.Commit()
}
