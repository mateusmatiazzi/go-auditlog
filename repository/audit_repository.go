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

var COLLECTION_NAME = "audit-logs"
var DATABASE_NAME = "Audit"
var CONNECTION_URL = "mongodb://localhost:27017"

type IAuditCollection interface {
	InsertOne(ctx context.Context, document interface{}, opts ...*options.InsertOneOptions) (*mongo.InsertOneResult, error)
	Find(ctx context.Context, document interface{}, opts ...*options.FindOptions) (*mongo.Cursor, error)
}

type AuditRepository struct {
	Collection IAuditCollection
}

func NewAuditRepository() *AuditRepository {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(CONNECTION_URL))

	validateConnection(err, client)

	fmt.Println("Connected to MongoDB!")

	return &AuditRepository{
		Collection: client.Database(DATABASE_NAME).Collection(COLLECTION_NAME),
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

func (repo *AuditRepository) SaveAuditLog(auditLog interface{}) error {
	fmt.Println("Saving in database")
	_, err := repo.Collection.InsertOne(context.TODO(), auditLog)

	if err != nil {
		return fmt.Errorf("Couldn't add log in database")
	}

	return nil
}

func (repo *AuditRepository) GetAllAuditlogs() ([]model.AuditLog, error) {
	cursor, err := repo.Collection.Find(context.TODO(), bson.M{})

	if err != nil {
		return nil, fmt.Errorf("Couldn't retrieve logs from database")
	}

	return buildArrayOfAuditLogs(cursor)
}

func buildArrayOfAuditLogs(cursor *mongo.Cursor) ([]model.AuditLog, error) {
	var auditLogs []model.AuditLog
	for cursor.Next(context.TODO()) {
		var auditLog model.AuditLog

		if err := cursor.Decode(&auditLog); err != nil {
			return nil, err
		}

		auditLogs = append(auditLogs, auditLog)
	}
	return auditLogs, nil
}
