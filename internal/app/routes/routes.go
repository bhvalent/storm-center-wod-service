package routes

import (
	"storm-center-wod-service/internal/app/controllers"
	"storm-center-wod-service/internal/domain/models"

	"github.com/gin-gonic/gin"
)

func Routes(a *models.Application) *gin.Engine {
	baseController := controllers.NewBaseController(a)

	routes := gin.Default()
	routes.GET("/wod/:date", baseController.GetWodHandler)
	return routes
}
