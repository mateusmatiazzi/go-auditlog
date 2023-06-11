package main

import (
	"context"
	"fmt"
	"log"

	"audit-service/model"
	"audit-service/repository"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")
	auditRepository := repository.NewAuditRepository(client)
	auditRepository.SaveAuditLog(model.NewAuditLog(123, "Account", "12-12-2023"))
}
