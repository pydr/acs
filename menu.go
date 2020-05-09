package acs

import (
	"github.com/pydr/acs/internal/action"
	"github.com/pydr/acs/model"
)

func (c *Client) AddMenu(menu *model.Menu) (int64, error) {
	return action.InsertMenu(menu)
}

func (c *Client) AddMenus(menus []*model.Menu) error {
	return action.InsertMenus(menus)
}

func (c *Client) DelMenu(id int64) error {
	return action.DelMenu(id)
}

func (c *Client) UpdateMenu(menu *model.Menu) (*model.Menu, error) {
	return action.UpdateMenu(menu)
}

func (c *Client) GetMenuById(id int64) (*model.Menu, error) {
	return action.GetMenuById(id)
}

func (c *Client) GetMenus() ([]*model.Menu, error) {
	return action.GetMenus()
}
