package model

type AuditLog struct {
	Entity     string `json:"entity"`
	CreateDate string `json:"createDate"`
	Value      string `jason:"value"`
}

func NewAuditLog(entity string, date string, value string) *AuditLog {
	return &AuditLog{
		Entity:     entity,
		CreateDate: date,
		Value:      value,
	}
}
