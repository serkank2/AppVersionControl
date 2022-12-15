package config

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

func ConnectDB(mongoUri string) *mongo.Client {
	client, err := mongo.NewClient(options.Client().ApplyURI(mongoUri))

	if err != nil {
		log.Fatalln(err)
	}

	ctx, _ := context.WithTimeout(context.Background(), 20*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("Connected to MongoDB!")
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("Ping to MongoDB!")

	return client
}

func GetCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	return client.Database("AppVersionControl").Collection(collectionName)
}
