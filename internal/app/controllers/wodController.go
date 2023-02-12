package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (bc *BaseController) GetWodHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"Message": "Hello world"})
}
