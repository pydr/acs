package action

import (
	"github.com/astaxie/beego/orm"
	"github.com/pydr/acs/model"
)

func registerRoleModels() {
	orm.RegisterModel(new(model.Role))
}

// 创建角色
func InsertRole(role *model.Role) (int64, error) {
	o := orm.NewOrm()
	err := o.Begin()
	if err != nil {
		return 0, err
	}

	id, err := o.Insert(role)
	if err != nil {
		o.Rollback()
		return 0, err
	}
	m2m := o.QueryM2M(role, "Menus")
	for _, m := range role.Menus {
		if _, err = m2m.Add(m); err != nil {
			o.Rollback()
			return 0, err
		}
	}
	m2m = o.QueryM2M(role, "Permissions")
	for _, p := range role.Permissions {
		if _, err = m2m.Add(p); err != nil {
			o.Rollback()
			return 0, err
		}
	}

	return id, o.Commit()
}

// 删除角色
func DelRole(id int64) (err error) {
	o := orm.NewOrm()
	if err = o.Begin(); err != nil {
		return
	}
	r := model.Role{Id: id}
	if err = o.Read(&r); err != nil {
		o.Rollback()
		return
	}
	r.Deleted = true
	if _, err = o.Update(&r); err != nil {
		o.Rollback()
		return
	}
	m2m := o.QueryM2M(r, "Menus")
	if _, err = m2m.Clear(); err != nil {
		o.Rollback()
		return
	}
	m2m = o.QueryM2M(r, "Users")
	if _, err = m2m.Clear(); err != nil {
		o.Rollback()
		return
	}
	m2m = o.QueryM2M(r, "Applications")
	if _, err = m2m.Clear(); err != nil {
		o.Rollback()
		return
	}
	m2m = o.QueryM2M(r, "Permissions")
	if _, err = m2m.Clear(); err != nil {
		o.Rollback()
		return
	}

	return o.Commit()
}

// 更新角色信息
func UpdateRole(role *model.Role) (*model.Role, error) {
	o := orm.NewOrm()
	_, err := o.Update(role, "Name")
	return role, err
}

// 更新角色菜单
func UpdateRoleMenus(role *model.Role) (*model.Role, error) {
	o := orm.NewOrm()
	if err := o.Begin(); err != nil {
		return nil, err
	}
	m2m := o.QueryM2M(role, "Menus")
	_, err := m2m.Clear()
	if err != nil {
		o.Rollback()
		return nil, err
	}
	for _, r := range role.Menus {
		if _, err = m2m.Add(r); err != nil {
			o.Rollback()
			return nil, err
		}
	}
	if err = o.Commit(); err != nil {
		return nil, err
	}
	return role, nil
}

// 更新角色权限
func UpdateRolePermissions(role *model.Role) (*model.Role, error) {
	o := orm.NewOrm()
	if err := o.Begin(); err != nil {
		return nil, err
	}
	m2m := o.QueryM2M(role, "Permissions")
	_, err := m2m.Clear()
	if err != nil {
		o.Rollback()
		return nil, err
	}
	for _, p := range role.Permissions {
		if _, err = m2m.Add(p); err != nil {
			o.Rollback()
			return nil, err
		}
	}
	if err = o.Commit(); err != nil {
		return nil, err
	}
	return role, nil
}

// 根据id查询角色
func GetRoleById(id int64) (ret *model.Role, err error) {
	ret = new(model.Role)
	o := orm.NewOrm()
	err = o.QueryTable(new(model.Role)).
		Filter("Id", id).
		Filter("Deleted", false).
		One(ret)
	if err != nil {
		return
	}
	if _, err = o.LoadRelated(ret, "Menus"); err != nil {
		return
	}
	if _, err = o.LoadRelated(ret, "Permissions"); err != nil {
		return
	}
	return
}

// 根据名称查询角色
func GetRoleByName(name string) (ret *model.Role, err error) {
	ret = new(model.Role)
	o := orm.NewOrm()
	err = o.QueryTable(new(model.Role)).
		Filter("Name", name).
		Filter("Deleted", false).
		One(ret)
	if err != nil {
		return
	}
	if _, err = o.LoadRelated(ret, "Menus"); err != nil {
		return
	}
	_, err = o.LoadRelated(ret, "Permissions")
	return
}

// 查询所有角色
func GetRoles(offset, page int32) ([]*model.Role, int64, error) {
	var ret []*model.Role
	o := orm.NewOrm()
	page--
	_, err := o.QueryTable(new(model.Role)).
		Filter("Deleted", false).
		Limit(offset, offset*page).
		All(&ret)
	if err != nil {
		return nil, 0, err
	}
	total, err := o.QueryTable(new(model.Role)).
		Filter("Deleted", false).
		Count()
	return ret, total, err
}

// 检查角色是否存在
func IsRoleExist(roleName string) bool {
	o := orm.NewOrm()
	return o.QueryTable(new(model.Role)).Filter("Name", roleName).Filter("Deleted", false).Exist()
}
