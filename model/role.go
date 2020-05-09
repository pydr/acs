package model

// 角色表
type Role struct {
	Id          int64         `json:"id"`
	Name        string        `orm:"unique" json:"name"`
	Application *Application  `orm:"rel(fk)" json:"application"`
	Menus       []*Menu       `orm:"rel(m2m)" json:"menus"`
	Users       []*User       `orm:"reverse(many)" json:"-"`
	Permissions []*Permission `orm:"rel(m2m)" json:"permissions"`
	BasicModel
}
