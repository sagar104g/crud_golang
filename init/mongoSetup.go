package initSetup

import (
	"context"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var mongoDatabaseMap = make(map[string]*mongo.Database)

func initMongo() {

	bookCollection := createDatabaseConnection("books")
	mongoDatabaseMap["books"] = bookCollection
	log.Println("mongo setup done")
}

func createDatabaseConnection(databaseName string) *mongo.Database {
	mongoConnectionString := os.Getenv("mongoUrl")
	if mongoConnectionString == "" {
		panic("mongo connection string should be present")
	}
	clientOptions := options.Client().ApplyURI(mongoConnectionString)
	client, err := mongo.Connect(context.Background(), clientOptions)

	if err != nil {
		log.Fatalln(err)
	}
	return client.Database(databaseName)
}

func GetCollection(databaseName string, collectionName string) *mongo.Collection {

	if mongoDatabase, ok := mongoDatabaseMap[databaseName]; ok {
		return mongoDatabase.Collection(collectionName)
	}
	return nil
}
