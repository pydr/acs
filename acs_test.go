package acs

import (
	"testing"

	"github.com/pydr/acs/model"
)

func TestAcs(t *testing.T) {
	p := &Options{
		MysqlIP:   "127.0.0.1",
		MysqlPort: 3306,
		MysqlUser: "root",
		MysqlPwd:  "123456",
		MysqlDB:   "test",
	}

	cli, err := NewAcs(p)
	if err != nil {
		t.Fatal(err)
	}

	per := &model.Permission{
		Path:    "/api/v1",
		Method:  "POST",
		Comment: "",
	}
	_, err = cli.AddPermission(per)
	if err != nil {
		t.Fatal(err)
	}

	app := &model.Application{
		Name:        "test002",
		Appid:       "hello",
		SecretKey:   "ppp",
		Comment:     "yes",
		Permissions: []*model.Permission{per},
	}
	_, err = cli.AddApplication(app)
	if err != nil {
		t.Fatal(err)
	}

	role := &model.Role{
		Name:        "test",
		Application: app,
		Permissions: app.Permissions,
	}

	_, err = cli.AddRole(role)
	if err != nil {
		t.Fatal()
	}

	user := &model.User{
		Username: "admin",
		Password: "123456",
		//Email:        "hello@me.com",
		Mobile:       "159000888",
		Roles:        []*model.Role{role},
		Applications: []*model.Application{app},
	}

	_, err = cli.AddUser(user)
	if err != nil {
		t.Fatal()
	}
}
