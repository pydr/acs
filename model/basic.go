package model

import "time"

type (
	BasicModel struct {
		Status  int32     `json:"-"`
		Deleted bool      `json:"-"`
		Created time.Time `orm:"auto_now_add; type(datetime)" json:"-"`
		Updated time.Time `orm:"auto_now; type(datetime)" json:"-"`
	}
)
