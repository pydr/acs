package acs

import (
	"github.com/pydr/acs/internal/action"
	"github.com/pydr/acs/model"
)

// 添加角色到角色池
func (c *Client) AddRole(role *model.Role) (int64, error) {
	id, err := action.InsertRole(role)
	if err != nil {
		return 0, err
	}
	err = c.addRolePolicy(role)
	return id, err
}

// 更新角色信息
func (c *Client) UpdateRole(role *model.Role) (*model.Role, error) {
	return action.UpdateRole(role)
}

// 更新角色的权限
func (c *Client) UpdateRolePermissions(role *model.Role) (*model.Role, error) {
	r, err := action.UpdateRolePermissions(role)
	if err != nil {
		return nil, err
	}
	err = c.updateRolePolicy(r)
	return r, err
}

// 更新角色菜单
func (c *Client) UpdateRoleMenus(role *model.Role) (*model.Role, error) {
	return action.UpdateRoleMenus(role)
}

// 获取所有角色
func (c *Client) GetRoles(offset, page int32) ([]*model.Role, int64, error) {
	return action.GetRoles(offset, page)
}

// 删除角色
func (c *Client) DelRole(id int64) error {
	return action.DelRole(id)
}

// 根据id查询角色
func (c *Client) GetRoleById(id int64) (*model.Role, error) {
	return action.GetRoleById(id)
}

// 根据角色名查询角色
func (c *Client) GetRoleByName(name string) (*model.Role, error) {
	return action.GetRoleByName(name)
}

// 查询角色是否存在
func (c *Client) IsRoleExist(name string) bool {
	return action.IsRoleExist(name)
}
