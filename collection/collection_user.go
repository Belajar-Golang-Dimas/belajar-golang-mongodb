package collection

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

func CollectionUser(ctx context.Context, db *mongo.Database, name string) {
	result := db.Collection("users").FindOne(ctx, bson.D{{"name", name}})
	var user any
	err := result.Decode(&user)
	if err != nil {
		switch err {
		case mongo.ErrNoDocuments:
			log.Println("Data tidak ditemukan.")
		default:
			log.Println(err)
		}

	}

	fmt.Println(user)

}
