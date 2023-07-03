package handler

import (
	"net/http"

	"audit-service/model"
	"audit-service/repository"

	"github.com/gin-gonic/gin"
)

type AuditLogHandler struct {
	repository *repository.AuditRepository
}

func NewAuditLogHandler() *AuditLogHandler {
	return &AuditLogHandler{}
}

func (handler *AuditLogHandler) CreateAuditLog(c *gin.Context) {
	initializeRepository(handler)
	var req model.AuditLog

	if err := c.ShouldBindJSON(&req); err != nil {
		// Handle error, e.g., return a bad request response
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	handler.repository.SaveAuditLog(req)

	c.JSON(http.StatusCreated, gin.H{
		"message": "User created successfully",
	})
}

func (handler *AuditLogHandler) GetAllAuditlogs(c *gin.Context) {
	initializeRepository(handler)
	response, err := handler.repository.GetAllAuditlogs()

	if err != nil {
		c.JSON(http.StatusInternalServerError, nil)
	}

	c.JSON(http.StatusOK, response)
}

func initializeRepository(handler *AuditLogHandler) {
	if handler.repository == nil {
		handler.repository = repository.NewAuditRepository()
	}
}
