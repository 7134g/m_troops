package model

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Result struct {
	//Status int
	Name   string
	Age    int
	Salary int
}

func main() {
	var MongoClient *mongo.Client
	uri := "mongodb://localhost:27017"
	option := options.Client().ApplyURI(uri)
	MongoClient, err := mongo.Connect(context.Background(), option)
	if err != nil {
		fmt.Println("error mongo connect")
		return
	}

	database := MongoClient.Database("test")
	collection := database.Collection("sales")

	ctx := context.TODO()
	f := bson.D{{"name", "abc"}}
	r := Result{}
	err = collection.FindOne(ctx, f).Decode(&r)
	if err != nil {
		fmt.Println("error mongo find")
		return
	}
	fmt.Println("done", r)

	r = Result{
		Name:   "xcs",
		Age:    24,
		Salary: 3255,
	}
	insertR, err := collection.InsertOne(context.Background(), r)
	if err != nil {
		fmt.Println("error mongo insert")
	}
	fmt.Println(insertR)

}
