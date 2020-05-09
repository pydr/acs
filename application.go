package acs

import (
	"github.com/pydr/acs/internal/action"
	"github.com/pydr/acs/model"
)

func (c *Client) AddApplication(app *model.Application) (int64, error) {
	return action.InsertApplication(app)
}

func (c *Client) SwitchApplicationStatus(id int64) error {
	return action.SwitchApplicationStatus(id)
}

func (c *Client) UpdateApplication(app *model.Application) (*model.Application, error) {
	return action.UpdateApplication(app)
}

func (c *Client) UpdateApplicationPermission(app *model.Application) (*model.Application, error) {
	return action.UpdateApplicationPermissions(app)
}

func (c *Client) GetApplicationByAppid(appid string) (*model.Application, error) {
	return action.GetApplicationByAppid(appid)
}

func (c *Client) GetApplicationById(id int64) (*model.Application, error) {
	return action.GetApplicationById(id)
}

func (c *Client) GetApplications(offset, page int32) ([]*model.Application, int64, error) {
	return action.GetApplications(offset, page)
}
