package acs

import (
	"github.com/pydr/acs/internal/action"
	"github.com/pydr/acs/model"
)

func (c *Client) AddUser(user *model.User) (int64, error) {
	return action.InsertUser(user)
}

// 删除用户
func (c *Client) DelUser(id int64) error {
	return action.DelUser(id)
}

// 更新用户信息
func (c *Client) UpdateUser(user *model.User) (*model.User, error) {
	return action.UpdateUser(user)
}

// 更新用户的角色
func (c *Client) UpdateUserRole(user *model.User) (*model.User, error) {
	return action.UpdateUserRoles(user)
}

// 更新用户所属终端
func (c *Client) UpdateUserApplications(user *model.User) (*model.User, error) {
	return action.UpdateUserApplications(user)
}

// 根据id查询用户
func (c *Client) GetUser(id int64) (*model.User, error) {
	return action.GetUserById(id)
}

// 根据用户名查询用户
func (c *Client) GetUserByUsername(username string) (*model.User, error) {
	return action.GetUserByUsername(username)
}

// 根据手机号查询用户
func (c *Client) GetUserByMobile(mobile string) (*model.User, error) {
	return action.GetUserByMobile(mobile)
}

// 根据邮箱查询用户
func (c *Client) GetUserByMail(mail string) (*model.User, error) {
	return action.GetUserByMail(mail)
}

func (c *Client)GetUsersLikeUsernameAndMobile(username string , mobile string )([]*model.User,error){
	return action.GetUserLikeUsernameAndMobile(username,mobile)
}

// 查询所有用户
func (c *Client) GetUsers(offset, page int32) ([]*model.User, int64, error) {
	return action.GetUsers(offset, page)
}

// 查询用户名是否存在
func (c *Client) IsUsernameExist(username string) bool {
	return action.IsUsernameExist(username)
}

// 查询手机号是否存在
func (c *Client) IsMobileExist(mobile string) bool {
	return action.IsMobileExist(mobile)
}

// 更新用户状态
func (c *Client) UpdateUserStatus(id int64, status int32) error {
	return action.UpdateUserStatus(id, status)
}

func (c *Client) UserStatusSwitcher(status int32) string {
	return action.UserStatusSwitcher(status)
}
