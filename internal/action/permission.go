package action

import (
	"github.com/astaxie/beego/orm"
	"github.com/pydr/acs/model"
)

func registerPermissionModels() {
	orm.RegisterModel(new(model.Permission))
}

// 创建权限
func InsertPermission(permission *model.Permission) (int64, error) {
	o := orm.NewOrm()
	return o.Insert(permission)
}

// 添加多个权限
func InsertPermissions(permissions []*model.Permission) error {
	o := orm.NewOrm()
	_, err := o.InsertMulti(len(permissions), permissions)
	return err
}

// 删除权限
func DelPermission(id int64) (err error) {
	o := orm.NewOrm()
	if err = o.Begin(); err != nil {
		return
	}

	p := model.Permission{Id: id}
	if err = o.Read(&p); err != nil {
		o.Rollback()
		return
	}

	p.Deleted = true
	if _, err = o.Update(&p, "Deleted"); err != nil {
		o.Rollback()
		return
	}
	m2m := o.QueryM2M(p, "Applications")
	if _, err = m2m.Clear(); err != nil {
		o.Rollback()
		return
	}

	m2m = o.QueryM2M(p, "Roles")
	if _, err = m2m.Clear(); err != nil {
		o.Rollback()
		return
	}

	return o.Commit()
}

// 更新权限
func UpdatePermission(permission *model.Permission) (*model.Permission, error) {
	o := orm.NewOrm()
	_, err := o.Update(permission, "Name", "Path", "Method")
	return permission, err
}

// 根据id查询权限
func GetPermissionById(id int64) (ret *model.Permission, err error) {
	ret = new(model.Permission)
	o := orm.NewOrm()
	err = o.QueryTable(new(model.Permission)).
		Filter("Id", id).
		Filter("Deleted", false).
		One(ret)
	return
}

// 查询所有权限
func GetPermissions(offset, page int32) ([]*model.Permission, int64, error) {
	var ret []*model.Permission
	o := orm.NewOrm()
	page--
	_, err := o.QueryTable(new(model.Permission)).
		Filter("Deleted", false).
		OrderBy("Created").
		Limit(offset, offset*page).
		All(&ret)
	if err != nil {
		return nil, 0, err
	}
	total, err := o.QueryTable(new(model.Permission)).
		Filter("Deleted", false).
		Count()
	return ret, total, err
}

// 检查权限是否存在
func IsPermissionExist(path, method string) bool {
	o := orm.NewOrm()
	return o.QueryTable(new(model.Permission)).Filter("Path", path).Filter("Method", method).Filter("Deleted", false).Exist()
}
