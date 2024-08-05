package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Connect() {
	MONGO_URI := os.Getenv("MONGO_URI")

	clientOptions := options.Client().ApplyURI(MONGO_URI)
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("✅ Connected to mongoDB.")
	}
}

func mainHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!")
}

func main() {
	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "80"
	}

	Connect()

	http.HandleFunc("/", mainHandler)
	err := http.ListenAndServe(":"+PORT, nil)

	if err != nil {
		fmt.Println("Error starting server")
		return
	}
}
