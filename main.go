package main

import (
	"audit-service/controller"
)

func main() {
	controller.CreateAuditLogController(&controller.AuditLogController{})
}
