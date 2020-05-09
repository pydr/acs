package model

// 终端表
type Application struct {
	Id          int64         `json:"id"`
	Name        string        `orm:"unique" json:"name"`
	Appid       string        `orm:"size(64)" json:"appid"`
	SecretKey   string        `orm:"size(64)" json:"secret_key"`
	Comment     string        `json:"comment"`
	Groups      []*Group      `orm:"rel(m2m)"`
	Roles       []*Role       `orm:"reverse(many)"`
	Users       []*User       `orm:"rel(m2m)"`
	Permissions []*Permission `orm:"rel(m2m)"`
	BasicModel
}
