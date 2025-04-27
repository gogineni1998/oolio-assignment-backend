package database

import (
	"context"
	"testing"

	"github.com/gogineni1998/oolio-assignment-backend/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MockCollection struct {
	mock.Mock
}

func (m *MockCollection) Find(ctx context.Context, filter interface{}, opts ...*options.FindOptions) (*mongo.Cursor, error) {
	args := m.Called(ctx, filter)
	cursor, _ := args.Get(0).(*mongo.Cursor)
	return cursor, args.Error(1)
}

func (m *MockCollection) FindOne(ctx context.Context, filter interface{}, opts ...*options.FindOptions) *mongo.SingleResult {
	args := m.Called(ctx, filter)
	return args.Get(0).(*mongo.SingleResult)
}

func (m *MockCollection) InsertOne(ctx context.Context, document interface{}, opts ...*options.FindOptions) (*mongo.InsertOneResult, error) {
	args := m.Called(ctx, document)
	result, _ := args.Get(0).(*mongo.InsertOneResult)
	return result, args.Error(1)
}

func (m *MockCollection) CountDocuments(ctx context.Context, filter interface{}, opts ...*options.FindOptions) (int64, error) {
	args := m.Called(ctx, filter)
	return int64(args.Int(0)), args.Error(1)
}

func (m *MockCollection) InsertMany(ctx context.Context, documents []interface{}, opts ...*options.FindOptions) (*mongo.InsertManyResult, error) {
	args := m.Called(ctx, documents)
	result, _ := args.Get(0).(*mongo.InsertManyResult)
	return result, args.Error(1)
}

func TestGetAllProducts(t *testing.T) {
	mockColl := new(MockCollection)
	ctx := context.Background()

	mockCursor := new(mongo.Cursor)
	mockColl.On("Find", ctx, bson.D{}).Return(mockCursor, nil)

	_, err := mockColl.Find(ctx, bson.D{})
	assert.NoError(t, err)
}

func TestInsertOrder(t *testing.T) {
	mockColl := new(MockCollection)
	ctx := context.Background()

	order := models.Order{
		Username: "testuser",
		Items: []models.Item{
			{ProductID: "p1", Quantity: 2},
			{ProductID: "p2", Quantity: 1},
		},
		Products: []*models.Product{
			{ID: "p1", Name: "Product 1", Price: 10.0, Category: "Category 1"},
			{ID: "p2", Name: "Product 2", Price: 20.0, Category: "Category 2"},
		},
	}

	mockResult := &mongo.InsertOneResult{
		InsertedID: "mocked_order_id",
	}
	mockColl.On("InsertOne", mock.Anything, order).Return(mockResult, nil)

	orderID, err := mockColl.InsertOne(ctx, order)
	assert.NoError(t, err)
	assert.Equal(t, "mocked_order_id", orderID.InsertedID)
}

func TestCheckUser_UserFound(t *testing.T) {
	mockColl := new(MockCollection)
	ctx := context.Background()

	expectedUser := &models.Credentials{
		Username: "testuser",
		Password: "testpass",
	}

	mockSingleResult := mongo.NewSingleResultFromDocument(expectedUser, nil, nil)

	mockColl.On("FindOne", ctx, bson.D{{Key: "username", Value: "testuser"}}).Return(mockSingleResult)

	result := mockColl.FindOne(ctx, bson.D{{Key: "username", Value: "testuser"}})
	assert.NotNil(t, result)
}

func TestInsertUser(t *testing.T) {
	mockColl := new(MockCollection)
	ctx := context.Background()

	credentials := models.Credentials{
		Username: "newuser",
		Password: "newpass",
	}

	mockResult := &mongo.InsertOneResult{
		InsertedID: "mocked_user_id",
	}
	mockColl.On("InsertOne", mock.Anything, credentials).Return(mockResult, nil)

	insertedID, err := mockColl.InsertOne(ctx, credentials)
	assert.NoError(t, err)
	assert.Equal(t, "mocked_user_id", insertedID.InsertedID)
}
