package repository_test

import (
	"errors"
	"fmt"
	"testing"

	"audit-service/model"
	"audit-service/repository"
	"audit-service/repository/mocks"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var ENTITY = "Order"
var CREATE_DATE = "31-10-1994"
var VALUE = "123"

func TestGetAllAuditlogs(t *testing.T) {
	repo, mockAuditCollection := mockAuditCollection(t)
	mockAuditCollection.EXPECT().Find(gomock.Any(), gomock.Any(), gomock.Any()).Return(buildMockedAuditLogCursor(), nil).Times(1)

	result, _ := repo.GetAllAuditlogs()

	assert.Equal(t, ENTITY, result[0].Entity)
}

func TestErrorInGetAllAuditlogs(t *testing.T) {
	repo, mockAuditCollection := mockAuditCollection(t)
	mockAuditCollection.EXPECT().Find(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil, errors.New("Cannot access the database")).Times(1)

	_, err := repo.GetAllAuditlogs()

	assert.Error(t, err, "Couldn't retrieve logs from database")
}

func TestSaveAuditLogAuditlogs(t *testing.T) {
	repo, mockAuditCollection := mockAuditCollection(t)
	mockAuditCollection.EXPECT().InsertOne(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil, nil).Times(1)

	err := repo.SaveAuditLog(buildMockedAuditLog())

	assert.NoError(t, err)
}

func TestErrorInSaveAuditLogAuditlogs(t *testing.T) {
	repo, mockAuditCollection := mockAuditCollection(t)
	mockAuditCollection.EXPECT().InsertOne(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil, fmt.Errorf("Couldn't add log in database")).Times(1)

	err := repo.SaveAuditLog(buildMockedAuditLog())

	assert.Error(t, err, "Couldn't add log in database")
}

func mockAuditCollection(t *testing.T) (*repository.AuditRepository, *mocks.MockIAuditCollection) {
	mockCollection := gomock.NewController(t)
	defer mockCollection.Finish()

	mockAuditCollection := mocks.NewMockIAuditCollection(mockCollection)

	repo := &repository.AuditRepository{
		Collection: mockAuditCollection,
	}

	return repo, mockAuditCollection
}

func buildMockedAuditLogCursor() *mongo.Cursor {

	doc := bson.M{"entity": ENTITY, "createDate": CREATE_DATE, "value": VALUE}

	cursor, _ := mongo.NewCursorFromDocuments([]interface{}{doc}, nil, nil)

	return cursor
}

func buildMockedAuditLog() *model.AuditLog {
	return model.NewAuditLog(ENTITY, CREATE_DATE, VALUE)
}
