package model

type AuditLog struct {
	ID         int    `json:"id"`
	Entity     string `json:"entity"`
	CreateDate string `json:"createDate"`
}

func NewAuditLog(id int, entity string, date string) *AuditLog {
	return &AuditLog{
		ID:         id,
		Entity:     entity,
		CreateDate: date,
	}
}
