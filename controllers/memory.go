package controllers

import (
	"github.com/gin-gonic/gin"
	"unraid-rest-api/services"
)

type MemoryController struct {
}

func (ctrl *MemoryController) GetMemoryInfo(c *gin.Context, service *services.MemoryService) {

	c.JSON(200, service.GetMemoryInfo())
}
