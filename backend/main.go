package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var collection *mongo.Collection

type Todo struct {
	ID        primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Title     string             `json:"title"`
	Summery   string             `json:"summery"`
	Completed bool               `json:"completed"`
}

func main() {
	fmt.Println("hello wolrd")
	err := godotenv.Load()

	if err != nil {
		log.Fatal("dotenv is not found")
	}

	mongouUrl := os.Getenv("MONGO_URL")

	///connect

	monogOptions := options.Client().ApplyURI(mongouUrl)

	client, err := mongo.Connect(context.Background(), monogOptions)

	if err != nil {
		panic(err)
	}
	///databae todo_app /// coollection

	// /
	collection = client.Database("todo_app").Collection("todo")

	fmt.Println("mongodb connected")

	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:5173/",
		AllowHeaders: "Content-Type,Accept,origin",
	}))

	app.Get("api/todos", GetTodos)
	app.Post("api/create", CreateTodo)
	app.Delete("/api/delete/:id", DeleteTodo)
	app.Put("/api/update/:id", UpdateTodo)

	///cdh
	log.Fatal(app.Listen(":3000"))
}

func GetTodos(c *fiber.Ctx) error {
	var todos []Todo

	///find data

	result, err := collection.Find(context.Background(), bson.M{})

	if err != nil {
		return err
	}

	///loop

	for result.Next(context.Background()) {
		var todo Todo
		if err := result.Decode(&todo); err != nil {
			return err
		}

		todos = append(todos, todo)
	}

	return c.Status(200).JSON(todos)

}

///create

func CreateTodo(c *fiber.Ctx) error {
	var todo = new(Todo)

	if err := c.BodyParser(todo); err != nil {
		return err
	}
	////quran title
	///descrioption:summery saba

	if todo.Title == "" {
		return c.Status(401).JSON(fiber.Map{"success": false, "meesage": "Please enter title"})
	}

	///create
	insert, err := collection.InsertOne(context.Background(), todo)
	if err != nil {
		log.Fatal(err)
	}

	todo.ID = insert.InsertedID.(primitive.ObjectID)

	return c.Status(200).JSON(fiber.Map{"success": true, "message": "user has been created"})

}

///delete

func DeleteTodo(c *fiber.Ctx) error {
	///params
	params := c.Params("id")

	//bson

	id, err := primitive.ObjectIDFromHex(params)

	if err != nil {
		return c.JSON("something wrong the id")
	}
	///filter id

	// /delete
	filter := bson.M{"_id": id}

	_, err = collection.DeleteOne(context.Background(), filter)

	if err != nil {
		return c.JSON("user already deleted")
	}

	return c.Status(200).JSON(fiber.Map{"success": true, "message": "user has been deleted"})

}

///update

func UpdateTodo(c *fiber.Ctx) error {
	// /params
	params := c.Params("id")

	id, err := primitive.ObjectIDFromHex(params)

	if err != nil {
		return err
	}

	fiter := bson.M{"_id": id}

	update := bson.M{"$set": bson.M{"completed": true}}

	///find and upate

	_, err = collection.UpdateOne(context.Background(), fiter, update)

	if err != nil {
		panic(err)
	}

	return c.Status(200).JSON(fiber.Map{"success": true, "message": "todo has been updated"})

}
