package repository

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
)

type AuditRepository struct {
	client *mongo.Client
}

func NewAuditRepository(client *mongo.Client) *AuditRepository {
	return &AuditRepository{
		client: client,
	}
}

func (repo *AuditRepository) SaveAuditLog(auditLog interface{}) {
	fmt.Println("Salvando no banco")
	_, err := repo.client.Database("Audit").Collection("audit-logs").InsertOne(context.TODO(), auditLog)

	if err != nil {
		log.Fatalln("Couldn't add log in database")
	}
}
