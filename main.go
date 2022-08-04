package main

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoInstance struct{
	Client
	Db
}

var mg MongoInstance
// var dbName, mongoURI string
// dbName := goDotEnvVariable("DBNAME")

type Employee struct{
	ID		string
	Name	string
	Salary	float64
	Age		float64
	Sex		string
}

func Connect() error{
	//load env 
	godotenv.Load(".env")
	dbName := os.Getenv("DBNAME")
	mongoURI := os.Getenv("MONG0URI") + dbName

	client, err := mongo.NewClient(options.Client().ApplyURI(mongoURI))
	context.WithTimeout(context.Background(), 25*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	db : client.Database(dbName)
	
	if err != nil {
		return err
	}

	mg = MongoInstance{Client : client,
	Db: }
}

func cancel() {
	panic("unimplemented")
}


func main(){
	if err := Connect(); err != nil{
		log.Fatal(err)
	}
	app := fiber.New()
	app.Get("/employee", func(c *fiber.Ctx) error{})
	app.Post("/employee")
	app.Put("/emplyee/:id")
	app.Delete("emplyee/:id")
}