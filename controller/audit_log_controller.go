package controller

import (
	"audit-service/handler"

	"github.com/gin-gonic/gin"
)

type AuditLogController struct {
	gin *gin.Engine
}

type ControllerInterface interface {
	CreateController()
}

func (controller *AuditLogController) CreateController() {
	controller.gin = gin.Default()

	controller.gin.POST("/newAudit", handler.NewAuditLogHandler().CreateAuditLog)

	controller.gin.Run(":8080")
}

func CreateAuditLogController(controller ControllerInterface) {
	controller.CreateController()
}
