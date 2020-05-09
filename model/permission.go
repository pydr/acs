package model

// 权限表
type Permission struct {
	Id           int64          `json:"id"`
	Path         string         `json:"path"`
	Method       string         `json:"method"`
	Comment      string         `json:"comment"`
	Applications []*Application `orm:"reverse(many)" json:"-"`
	Roles        []*Role        `orm:"reverse(many)" json:"-"`
	BasicModel
}
