package mongodb

import (
	"context"
	"sync"

	"github.com/yosa12978/MyShirts/internal/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	db   *mongo.Database
	once sync.Once
)

func GetDB() *mongo.Database {
	once.Do(func() {
		c := config.GetConfig()
		client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(c.Mongo))
		if err != nil {
			panic(err)
		}
		if err := client.Ping(context.TODO(), nil); err != nil {
			panic(err)
		}
		db = client.Database(c.Database)
	})
	return db
}
