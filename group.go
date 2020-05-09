package acs

import (
	"github.com/pydr/acs/internal/action"
	"github.com/pydr/acs/model"
)

func (c *Client) AddGroup(group *model.Group) (int64, error) {
	return action.InsertGroup(group)
}

func (c *Client) DelGroup(id int64) error {
	return action.DelGroup(id)
}

func (c *Client) UpdateGroup(group *model.Group) (*model.Group, error) {
	return action.UpdateGroup(group)
}

func (c *Client) GetGroups(offset, page int32) ([]*model.Group, int64, error) {
	return action.GetGroups(offset, page)
}

func (c *Client) GetGroupById(id int64) (*model.Group, error) {
	return action.GetGroupById(id)
}
