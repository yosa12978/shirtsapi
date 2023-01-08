package repos

import (
	"context"

	"github.com/yosa12978/MyShirts/internal/models"
	"github.com/yosa12978/MyShirts/internal/mongodb"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type AccountRepo interface {
	GetByID(id string) (models.Account, error)
	Create(account models.Account) error
	Delete(id string) error
	GetByToken(token string) (models.Account, error)
}

type AccountRepoMongo struct {
}

func NewAccountRepoMongo() AccountRepo {
	return new(AccountRepoMongo)
}

func (arm *AccountRepoMongo) GetByID(id string) (models.Account, error) {
	db := mongodb.GetDB()
	objid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return models.Account{}, mongo.ErrNoDocuments
	}
	var acc models.Account
	err = db.Collection("accounts").FindOne(context.TODO(), bson.M{"_id": objid}).Decode(&acc)
	if err != nil {
		return acc, err
	}
	return acc, nil
}

func (arm *AccountRepoMongo) Create(account models.Account) error {
	db := mongodb.GetDB()
	account.Id = primitive.NewObjectID()
	_, err := db.Collection("accounts").InsertOne(context.TODO(), account)
	return err
}

func (arm *AccountRepoMongo) Delete(id string) error {
	db := mongodb.GetDB()
	objid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return mongo.ErrNoDocuments
	}
	_, err = db.Collection("accounts").DeleteOne(context.TODO(), bson.M{"_id": objid})
	return err
}

func (arm *AccountRepoMongo) GetByToken(token string) (models.Account, error) {
	db := mongodb.GetDB()
	filter := bson.M{
		"token": token,
	}
	var account models.Account
	err := db.Collection("accounts").FindOne(context.TODO(), filter).Decode(&account)
	return account, err
}
