package controllers

import (
	"unraid-rest-api/services"

	"github.com/gin-gonic/gin"
)

type MemoryController struct {
}

func (ctrl *MemoryController) GetMemoryInfo(c *gin.Context, service *services.MemoryService) {
	c.JSON(200, service.GetMemoryInfo())
}

func (ctrl *MemoryController) GetMemoryHistory(c *gin.Context, service *services.MemoryService) {
	c.JSON(200, service.GetMemoryHistory())
}
