package repository

import (
	"audit-service/model"
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
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

func (repo *AuditRepository) GetAllAuditlogs() []model.AuditLog {
	cursor, err := repo.client.Database("Audit").Collection("audit-logs").Find(context.TODO(), bson.M{})

	if err != nil {
		log.Fatalln("Couldn't retrieve logs from database")
	}

	return buildArrayOfAuditLogs(cursor)
}

func buildArrayOfAuditLogs(cursor *mongo.Cursor) []model.AuditLog {
	var auditLogs []model.AuditLog
	for cursor.Next(context.TODO()) {
		var auditLog model.AuditLog

		if err := cursor.Decode(&auditLog); err != nil {
			log.Fatalln("Couldn't retrieve logs from database")
		}

		auditLogs = append(auditLogs, auditLog)
	}
	return auditLogs
}
