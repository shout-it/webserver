
package dao

import (
	"context"
	"github.com/mongodb/mongo-go-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"webserver/models"
)
const CONNECTIONSTRING = "mongodb://localhost:27017"
const DBNAME = "shoutit"
const COLLECTIONNAME = "users"

var db *mongo.Database;

func init() {
client, err := mongo.NewClient(options.Client().ApplyURI(CONNECTIONSTRING))
if err != nil {
log.Fatal(err)
}
err = client.Connect(context.Background())
if err != nil {
log.Fatal(err)
}
db = client.Database(DBNAME)
}

func InsertOneValue(user models.User) {
	_, err := db.Collection(COLLECTIONNAME).InsertOne(context.Background(), user)
	if err != nil {
		log.Fatal(err)
	}
}