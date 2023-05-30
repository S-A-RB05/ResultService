package main

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type ResultCSV struct {
	CSV      string    `json:"csv"`
}

// global variable mongodb connection client
var mongoClient mongo.Client = newClient()

// ----Create----
func insertResult(csv string) {
	var csvObject ResultCSV;
	csvObject.CSV = csv;
	resultCollection := mongoClient.Database("results").Collection("csv")
	_, err := resultCollection.InsertOne(context.TODO(), csvObject)
	if err != nil {
		panic(err)
		//return
	}

	// return the ID of the newly inserted csv
	log.Println("!!!!!!!!!!!!inserted: " + csvObject.CSV)
}


func readAllResults() (values []primitive.M) {
	resultCollection := mongoClient.Database("results").Collection("csv")
	// retrieve all the documents (empty filter)
	cursor, err := resultCollection.Find(context.TODO(), bson.D{})
	// check for errors in the finding
	if err != nil {
		panic(err)
	}

	// convert the cursor result to bson
	var results []bson.M
	// check for errors in the conversion
	if err = cursor.All(context.TODO(), &results); err != nil {
		panic(err)
	}

	// display the documents retrieved
	fmt.Println("displaying all results from the search query")
	for _, result := range results {
		fmt.Println(result)
	}

	values = results
	return
}

// other
func newClient() (value mongo.Client) {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb+srv://stockbrood:admin@stockbrood.sifn3lq.mongodb.net/test"))
	if err != nil {
		panic(err)
	}
	if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
		panic(err)
	}
	value = *client

	return
}
