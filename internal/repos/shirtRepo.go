package repos

import (
	"context"

	"github.com/yosa12978/MyShirts/internal/helpers"
	"github.com/yosa12978/MyShirts/internal/models"
	"github.com/yosa12978/MyShirts/internal/mongodb"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type ShirtRepo interface {
	GetAll() ([]models.Shirt, error)
	GetByID(id string) (models.Shirt, error)
	Create(shirt models.Shirt) error
	Update(shirt models.Shirt) error
	Delete(id string) error
}

type shirtRepoMongo struct {
}

func NewShirtRepoMongo() ShirtRepo {
	return new(shirtRepoMongo)
}

func (sr *shirtRepoMongo) GetAll() ([]models.Shirt, error) {
	db := mongodb.GetDB()
	filterOpts := options.Find().SetSort(bson.M{"_id": -1})
	var shirts []models.Shirt
	cursor, err := db.Collection("shirts").Find(context.TODO(), bson.M{}, filterOpts)
	if err != nil {
		return shirts, err
	}
	err = cursor.All(context.TODO(), &shirts)
	return shirts, err
}

func (sr *shirtRepoMongo) GetByID(id string) (models.Shirt, error) {
	var shirt models.Shirt
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return shirt, helpers.ErrNotFound
	}
	db := mongodb.GetDB()
	err = db.Collection("shirts").FindOne(context.TODO(), bson.M{"_id": objID}).Decode(&shirt)
	if err == mongo.ErrNoDocuments {
		return shirt, helpers.ErrNotFound
	}
	return shirt, err
}

func (sr *shirtRepoMongo) Create(shirt models.Shirt) error {
	db := mongodb.GetDB()
	shirt.Id = primitive.NewObjectID()
	_, err := db.Collection("shirts").InsertOne(context.TODO(), shirt)
	return err
}

func (sr *shirtRepoMongo) Update(shirt models.Shirt) error {
	db := mongodb.GetDB()
	_, err := db.Collection("shirts").ReplaceOne(context.TODO(), bson.M{"_id": shirt.Id}, shirt)
	if err == mongo.ErrNoDocuments {
		return helpers.ErrNotFound
	}
	return err
}

func (sr *shirtRepoMongo) Delete(id string) error {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return helpers.ErrNotFound
	}
	db := mongodb.GetDB()
	_, err = db.Collection("shirts").DeleteOne(context.TODO(), bson.M{"_id": objID})
	if err == mongo.ErrNoDocuments {
		return helpers.ErrNotFound
	}
	return err
}
