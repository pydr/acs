package action

import (
	"github.com/astaxie/beego/orm"
	"github.com/pydr/acs/model"
)

func registerGroupModels() {
	orm.RegisterModel(new(model.Group))
}

// 添加分组
func InsertGroup(group *model.Group) (int64, error) {
	o := orm.NewOrm()
	id, err := o.Insert(&group)
	if err != nil {
		o.Rollback()
		return 0, err
	}
	m2m := o.QueryM2M(group, "Applications")
	for _, a := range group.Applications {
		if _, err = m2m.Add(a); err != nil {
			o.Rollback()
			return 0, err
		}
	}
	return id, o.Commit()
}

// 删除分组
func DelGroup(id int64) (err error) {
	o := orm.NewOrm()
	if err = o.Begin(); err != nil {
		return
	}
	g := &model.Group{Id: id}
	if err = o.Read(g); err != nil {
		o.Rollback()
		return
	}
	g.Deleted = true
	if _, err = o.Update(g); err != nil {
		o.Rollback()
		return
	}
	m2m := o.QueryM2M(g, "Users")
	if _, err = m2m.Clear(); err != nil {
		o.Rollback()
		return
	}
	m2m = o.QueryM2M(g, "Applications")
	if _, err = m2m.Clear(); err != nil {
		o.Rollback()
		return
	}
	return o.Commit()
}

// 更新分组信息
func UpdateGroup(role *model.Group) (*model.Group, error) {
	o := orm.NewOrm()
	_, err := o.Update(role, "Name")
	return role, err
}

// 根据id获取分组
func GetGroupById(id int64) (ret *model.Group, err error) {
	ret = new(model.Group)
	o := orm.NewOrm()
	err = o.QueryTable(new(model.Group)).
		Filter("Id", id).
		Filter("Deleted", false).
		One(ret)
	if err != nil {
		return
	}

	if _, err = o.LoadRelated(ret, "Users"); err != nil {
		return
	}
	return
}

// 获取所有分组
func GetGroups(offset, page int32) ([]*model.Group, int64, error) {
	var ret []*model.Group
	o := orm.NewOrm()
	page--
	_, err := o.QueryTable(new(model.Group)).
		Filter("Deleted", false).
		Limit(offset, page*offset).
		All(&ret)
	if err != nil {
		return nil, 0, err
	}

	total, err := o.QueryTable(new(model.Group)).
		Filter("Deleted", false).
		Count()
	return ret, total, err
}
