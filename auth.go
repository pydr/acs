package acs

import (
	"errors"

	"github.com/pydr/acs/model"
)

// 检查权限
func (c *Client) CheckPermission(role, tenant, path, method string) (bool, error) {
	return c.CasbinCli.Enforce(role, tenant, path, method)
}

// 添加角色规则
func (c *Client) addRolePolicy(role *model.Role) error {
	for _, per := range role.Permissions {
		_, _ = c.CasbinCli.AddPolicy(role.Name, role.Application.Name, per.Path, per.Method)
	}

	if err := c.CasbinCli.SavePolicy(); err != nil {
		return err
	}

	return c.CasbinCli.LoadPolicy()
}

// 更新角色的规则
func (c *Client) updateRolePolicy(role *model.Role) error {
	// 清空已存在的规则
	if ok, err := c.CasbinCli.RemoveFilteredNamedPolicy("p", 0, role.Name); !ok || err != nil {
		return errors.New("clear policy failed")
	}

	for _, per := range role.Permissions {
		_, _ = c.CasbinCli.AddPolicy(role.Name, role.Application.Name, per.Path, per.Method)
	}

	if err := c.CasbinCli.SavePolicy(); err != nil {
		return err
	}

	return c.CasbinCli.LoadPolicy()
}
