package model

// user表
type User struct {
	Id           int64          `json:"id"`
	Username     string         `orm:"unique" json:"username"`
	Password     string         `json:"-"`
	Email        *string        `orm:"unique; null" json:"email"`
	Mobile       string         `orm:"unique" json:"mobile"`
	Roles        []*Role        `orm:"rel(m2m)" json:"roles"`
	Group        *Group         `orm:"rel(fk); null; on_delete(set_null)" json:"group"`
	Applications []*Application `orm:"reverse(many)" json:"-"`
	BasicModel
}

// 用户组表
type Group struct {
	Id           int64          `json:"id"`
	Name         string         `orm:"unique" json:"name"`
	Users        []*User        `orm:"reverse(many)" json:"-"`
	Applications []*Application `orm:"reverse(many)" json:"-"`
	BasicModel
}

// 用户状态
const (
	UnVerify = iota + 1
	Verified
)
