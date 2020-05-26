package model

import (
	"context"
	"fmt"
	"log"

	"../db"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Char struct {
	ID   primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name string             `json:"Name"`
	Game string             `json:"Game"`
}

// get all task from the DB and return it
func GetAll() []primitive.M {
	cur, err := db.ColObj.Find(context.Background(), bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}

	var results []primitive.M
	for cur.Next(context.Background()) {
		var result bson.M
		e := cur.Decode(&result)
		if e != nil {
			log.Fatal(e)
		}
		// fmt.Println("cur..>", cur, "result", reflect.TypeOf(result), reflect.TypeOf(result["_id"]))
		results = append(results, result)

	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	cur.Close(context.Background())
	return results
}

func DeleteOne(char string) {
	fmt.Println(char)
	id, _ := primitive.ObjectIDFromHex(char)
	filter := bson.M{"_id": id}
	d, err := db.ColObj.DeleteOne(context.Background(), filter)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Deleted Document", d.DeletedCount)
}

func GetOne(char string) []primitive.M {
	id, _ := primitive.ObjectIDFromHex(char)
	filter := bson.M{"_id": id}

	filterCursor, err := db.ColObj.Find(context.Background(), filter)
	if err != nil {
		// ErrNoDocuments means that the filter did not match any documents in the collection
		if err == mongo.ErrNoDocuments {
			log.Println("no existe")

		}
		log.Fatal(err)
	}
	var result []primitive.M
	if err = filterCursor.All(context.Background(), &result); err != nil {
		log.Fatal(err)
	}
	fmt.Println(result)
	return result
}

// Insert one task in the DB
func InsertOne(char Char) {
	insertResult, err := db.ColObj.InsertOne(context.Background(), char)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Inserted a Single Record ", insertResult.InsertedID)
}

func Edit(char string, res Char) (bson.M, error) {
	id, _ := primitive.ObjectIDFromHex(char)
	filter := bson.M{"_id": id}
	update := bson.M{
		"$set": bson.M{"name": res.Name, "game": res.Game},
	}

	//options
	upsert := true
	after := options.After
	opt := options.FindOneAndUpdateOptions{
		ReturnDocument: &after,
		Upsert:         &upsert,
	}

	//find and update
	result := db.ColObj.FindOneAndUpdate(context.Background(), filter, update, &opt)
	if result.Err() != nil {
		return nil, result.Err()
	}

	//decode
	doc := bson.M{}
	decodeErr := result.Decode(&doc)

	log.Println(doc)
	return doc, decodeErr
}
