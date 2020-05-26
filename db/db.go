package db

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// const connectionString = "mongodb://localhost:27017"
const ConnectionString = "mongodb://localhost"

// Database Name
const DbName = "nintendo"

// Collection name
const CollName = "chars"

// collection object/instance
var ColObj *mongo.Collection

func Connect() {
	// Database Config
	clientOptions := options.Client().ApplyURI(ConnectionString)
	client, err := mongo.NewClient(clientOptions)
	//Set up a context required by mongo.Connect
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	//To close the connection at the end
	defer cancel()

	// Ping our db connection
	err = client.Ping(context.Background(), readpref.Primary())
	if err != nil {
		log.Fatal("Couldn't connect to the database", err)
	} else {
		log.Println("Connected!")
	}
	// Create the database called *go_mongo*
	db := client.Database(DbName)
	ColObj = db.Collection(CollName)

	return
}
