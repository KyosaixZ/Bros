package main

import (
	"context"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Person struct {
	Picture string `json:"picture" bson:"picture"`
	Name    string `json:"name" bson:"name"`
	Email   string `json:"email" bson:"email"`
	Age     int    `json:"age" bson:"age"`
}

func main() {
	// MongoDB connection string
	clientOptions := options.Client().ApplyURI("mongodb+srv://mangkornko:1234@cluster0.fjtujrv.mongodb.net/")

	// Connect to MongoDB
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}

	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowHeaders:     "Origin,Content-Type,Accept,Content-Length,Accept-Language,Accept-Encoding,Connection,Access-Control-Allow-Origin",
		AllowOrigins:     "*",
		AllowCredentials: true,
		AllowMethods:     "GET,POST,HEAD,PUT,DELETE,PATCH,OPTIONS",
	}))

	// Define an endpoint to retrieve all data from MongoDB
	app.Get("/", func(c *fiber.Ctx) error {
		// Create a MongoDB client
		collection := client.Database("1stcow").Collection("card")

		// Define a filter to match all documents (empty filter)
		filter := bson.D{{}}

		// Find all documents in the collection
		cur, err := collection.Find(context.Background(), filter)
		if err != nil {
			log.Fatal(err)
			return c.Status(500).SendString("Internal Server Error")
		}

		defer cur.Close(context.Background())

		// Iterate through the cursor and collect the results in a slice
		var results []bson.M
		for cur.Next(context.Background()) {
			var result bson.M
			err := cur.Decode(&result)
			if err != nil {
				log.Fatal(err)
				return c.Status(500).SendString("Internal Server Error")
			}
			results = append(results, result)
		}

		// Return the results as JSON
		return c.JSON(results)
	})

	app.Post("/add", func(c *fiber.Ctx) error {
		// Parse the request body as a JSON object
		var person Person
		if err := c.BodyParser(&person); err != nil {
			return c.Status(400).SendString("Bad Request")
		}

		// Create a MongoDB client
		collection := client.Database("1stcow").Collection("card")

		// Insert the new person data into the database
		_, err := collection.InsertOne(context.Background(), person)
		if err != nil {
			log.Fatal(err)
			return c.Status(500).SendString("Internal Server Error")
		}

		return c.SendString("Data added successfully")
	})

	app.Listen(":8000")
}
