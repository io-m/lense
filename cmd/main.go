package main

import (
	// "fmt"
	// "log"
	"context"
	"time"

	"github.com/io-m/lenses/pkg/endpoints"
	"github.com/io-m/lenses/pkg/models/storages"
)

// var(
// 	dbname = "lenses"
// 	collection = "users"
// )


func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	// mongoDB, _ := storages.Connect()
	// mongoClient := storages.Client
	// mongoClient.Connect(ctx, mongoconn)
	// mongodb := mongoClient.NewMongoDatabase(dbname)
	// if mongodb.Collection(collection) == nil {
	// 	if err := mongodb.CreateCollection(ctx, collection); err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	log.Printf("Collection %s is created", collection)
	// }
	// log.Printf("COLLECTION : %s", collection)
	defer storages.Client.Disconnect(ctx)
	defer cancel()

	endpoints.RunApp()
}
