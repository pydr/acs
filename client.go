package acs

import (
	"fmt"

	gormadapter "github.com/casbin/gorm-adapter/v2"

	"github.com/casbin/casbin/v2"
	"github.com/pydr/acs/internal/action"
)

type (
	Client struct {
		CasbinCli *casbin.Enforcer
	}

	// mysql 配置项
	Options struct {
		MysqlIP   string
		MysqlPort int
		MysqlUser string
		MysqlPwd  string
		MysqlDB   string
	}
)

func initCasbin(op *Options) (e *casbin.Enforcer, err error) {
	a, err := gormadapter.NewAdapter(
		"mysql",
		fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", op.MysqlUser, op.MysqlPwd, op.MysqlIP, op.MysqlPort, op.MysqlDB),
		true)
	if err != nil {
		return nil, err
	}

	e, err = casbin.NewEnforcer("conf/rbac-with_domains_model.conf", a)
	if err != nil {
		return
	}

	err = e.LoadPolicy()
	return
}

func NewAcs(op *Options) (*Client, error) {
	action.RegisterAcsModels()

	casbinCli, err := initCasbin(op) // 初始化casbin
	if err != nil {
		return nil, err
	}

	return &Client{CasbinCli: casbinCli}, nil
}
