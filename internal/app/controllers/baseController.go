package controllers

import "storm-center-wod-service/internal/domain/models"

type BaseController struct {
	app *models.Application
}

func NewBaseController(a *models.Application) *BaseController {
	return &BaseController{app: a}
}
