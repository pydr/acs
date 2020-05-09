package model

type Menu struct {
	Id       int64   `json:"id"`
	ParentId int64   `json:"parent_id"`
	Name     string  `orm:"unique" json:"name"`
	Path     string  `orm:"unique" json:"path"`
	Comment  string  `json:"comment"`
	Roles    []*Role `orm:"reverse(many)" json:"-"`
	BasicModel
}
