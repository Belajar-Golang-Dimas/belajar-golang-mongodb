package main

import (
	"context"
	"github.com/Belajar-Golang-Dimas/belajar-golang-mongodb.git/app"
	"github.com/Belajar-Golang-Dimas/belajar-golang-mongodb.git/collection"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	client := app.MongoConnection(ctx)
	database := client.Database("elearning")

	collection.CollectionUser(ctx, database, "dimasasdasd")

}
