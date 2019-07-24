
package dao

import (
	"context"
	"github.com/mongodb/mongo-go-driver/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/x/bsonx"
	"log"
	"webserver/models"
	"github.com/nu7hatch/gouuid"
)
const CONNECTIONSTRING = "mongodb://localhost:27017"
const DBNAME = "shoutit"
const COLLECTIONNAME = "users"

var db *mongo.Database;

func init() {client, err := mongo.NewClient(options.Client().ApplyURI(CONNECTIONSTRING))
if err != nil {
log.Fatal(err)
}
err = client.Connect(context.Background())
if err != nil {
log.Fatal(err)
}
db = client.Database(DBNAME)
index := mongo.IndexModel{
		Keys:bsonx.Doc{{"email", bsonx.Int32(1)}},
		Options:options.Index().SetUnique(true),
	}

_,err = db.Collection(COLLECTIONNAME).Indexes().CreateOne(context.Background(),index,)
if err != nil {
	log.Println(err)
}
}

func InsertOneValue(user models.User) error {
	uid,err := uuid.NewV4()
	user.Id = uid.String()
	_, err = db.Collection(COLLECTIONNAME).InsertOne(context.Background(), user)
	if err != nil {
		return err
	}
	return nil
}

func FindBy(email string) (models.User,error){
	var user models.User
	filter := bson.D{{
		"email",email,
	}}
	err := db.Collection(COLLECTIONNAME).FindOne(context.Background(),filter).Decode(&user)
	if err != nil {
		return models.User{},err
	}
	return user,nil
}