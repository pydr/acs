package acs

import (
	"github.com/pydr/acs/model"

	"github.com/pydr/acs/internal/action"
)

// 添加权限到权限池
func (c *Client) AddPermission(per *model.Permission) (int64, error) {
	return action.InsertPermission(per)
}

// 添加多个权限
func (c *Client) AddPermissions(pers []*model.Permission) error {
	return action.InsertPermissions(pers)
}

// 获取所有权限
func (c *Client) GetPermission(offset, page int32) ([]*model.Permission, int64, error) {
	return action.GetPermissions(offset, page)
}

// 删除权限
func (c *Client) DelPermission(id int64) error {
	return action.DelPermission(id)
}

// 更新权限
func (c *Client) UpdatePermission(per *model.Permission) (*model.Permission, error) {
	return action.UpdatePermission(per)
}

// 根据id查询权限
func (c *Client) GetPermissionById(id int64) (*model.Permission, error) {
	return action.GetPermissionById(id)
}

// 检查权限是否存在
func (c *Client) IsPermissionExist(path, method string) bool {
	return action.IsPermissionExist(path, method)
}
