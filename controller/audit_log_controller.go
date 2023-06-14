package controller

import (
	"audit-service/handler"

	"github.com/gin-gonic/gin"
)

func CreateAuditLogController() {
	gin := gin.Default()
	handler := handler.NewAuditLogHandler()

	gin.POST("/newAudit", handler.CreateAuditLog)
	gin.GET("/getAll", handler.GetAllAuditlogs)

	gin.Run(":8080")
}
