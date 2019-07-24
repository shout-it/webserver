
package dao

import (
	"context"
	"github.com/mongodb/mongo-go-driver/mongo"
	"github.com/nu7hatch/gouuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/x/bsonx"
	"log"
	"webserver/config"
	"webserver/models"
)
const CollectionName = "users"

var db *mongo.Database;

func init() {
	config := config.GetConfig()
	client, err := mongo.NewClient(options.Client().ApplyURI(config.DataBaseHost))
	if err != nil {
		log.Fatal(err)
	}
	err = client.Connect(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	db = client.Database(config.DataBaseName)
	index := mongo.IndexModel{
		Keys:bsonx.Doc{{"email", bsonx.Int32(1)}},
		Options:options.Index().SetUnique(true),
	}

	_,err = db.Collection(CollectionName).Indexes().CreateOne(context.Background(),index,)
	if err != nil {
		log.Println(err)
	}
}

func InsertOneValue(user models.User) error {
	uid,err := uuid.NewV4()
	user.Id = uid.String()
	_, err = db.Collection(CollectionName).InsertOne(context.Background(), user)
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
	err := db.Collection(CollectionName).FindOne(context.Background(),filter).Decode(&user)
	if err != nil {
		return models.User{},err
	}
	return user,nil
}