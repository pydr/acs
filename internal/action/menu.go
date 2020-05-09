package action

import (
	"github.com/astaxie/beego/orm"
	"github.com/pydr/acs/model"
)

func registerMenuModels() {
	orm.RegisterModel(new(model.Menu))
}

// 添加菜单
func InsertMenu(menu *model.Menu) (int64, error) {
	o := orm.NewOrm()
	return o.Insert(menu)
}

// 添加多个菜单
func InsertMenus(menus []*model.Menu) error {
	o := orm.NewOrm()
	_, err := o.InsertMulti(len(menus), menus)
	return err
}

// 删除菜单
func DelMenu(id int64) (err error) {
	o := orm.NewOrm()
	if err = o.Begin(); err != nil {
		return
	}
	m := model.Menu{Id: id}
	if err = o.Read(&m); err != nil {
		o.Rollback()
		return
	}
	m.Deleted = true
	if _, err = o.Update(&m); err != nil {
		o.Rollback()
		return
	}
	m2m := o.QueryM2M(m, "Roles")
	if _, err = m2m.Clear(); err != nil {
		o.Rollback()
		return
	}
	return
}

// 更新菜单
func UpdateMenu(menu *model.Menu) (*model.Menu, error) {
	o := orm.NewOrm()
	_, err := o.Update(menu, "Name", "Path", "Comment")
	return menu, err
}

// 根据id查询菜单
func GetMenuById(id int64) (ret *model.Menu, err error) {
	ret = new(model.Menu)
	o := orm.NewOrm()
	err = o.QueryTable(new(model.Menu)).Filter("Id", id).Filter("Deleted", false).One(ret)
	return
}

// 查询所有菜单
func GetMenus() ([]*model.Menu, error) {
	var ret []*model.Menu
	o := orm.NewOrm()
	_, err := o.QueryTable(new(model.Menu)).Filter("Deleted", false).All(&ret)
	return ret, err
}
