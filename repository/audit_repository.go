package repository

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type AuditRepository struct {
	client *mongo.Client
}

func NewAuditRepository() *AuditRepository {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017"))

	validateConnection(err, client)

	fmt.Println("Connected to MongoDB!")

	return &AuditRepository{
		client: client,
	}
}

func validateConnection(err error, client *mongo.Client) {
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}
}

func (repo *AuditRepository) SaveAuditLog(auditLog interface{}) {
	fmt.Println("Salvando no banco")
	_, err := repo.client.Database("Audit").Collection("audit-logs").InsertOne(context.TODO(), auditLog)

	if err != nil {
		log.Fatalln("Couldn't add log in database")
	}
}
