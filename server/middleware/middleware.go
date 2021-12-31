package middleware

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/adyang94/react-go-todo-app/models"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"go.mongo.org/mongo-driver/mongo"
	"go.mongo.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/tools/go/analysis/passes/nilfunc"
)

var collection *mongo.Collection

func init () {
	loadTheEnv()
	createDBInstance()
}

func loadTheEnv () {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading the .env file")
	}
}

func createDBInstance () {
	connectionString := os.Getenv("DB_URI")
	dbName := os.Getenv("DB_NAME")
	collName := os.Getenv("DB_COLLECTION_NAME")

	clientOptions := options.Client().ApplyURL(connectionString)

	client, err :=mongo.Connect(context.TODO(), clientOptions)

	if err!=nil {
		log.Fatal(err)
	}

	err = client.Ping(context.TODO(), nil)
	if err!=nil{
		log.Fatal(err)
	}

	fmt.Println("connected to mongodb!")

	collection = client.Database(dbName).Collection(collName)
	fmt.Println("collection instance created")
}

func GetAllTasks (w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	payload := getAllTasks()
	json.NewEncoder(w).Encode(payload)
}

func CreateTask (w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	var task models.ToDoList

	json.NewDecoder(r.Body).Decode(&task)
	insertOneTask(task)
	json.NewEncoder(w).Encode(task)
}

func TaskComplete (w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "PUT")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	params := mux.Vars(r)
	TaskComplete(params["id"])
	json.NewEncoder(w).Encode(params["id"])
}

func UndoTask (w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "applications/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "PUT")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	
	params := mux.Vars(r)
	UndoTask(params["id"])
	json.NewEncoder(w).Encode(params["id"])
}

func DeleteTask (w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "applications/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	params := mux.Vars(r)
	deleteOneTask(params["id"])
}

func DeleteAllTasks (w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "applications/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	count := deleteAllTasks()
	json.NewDecoder(w).Encode(count)
}

func getAllTasks () {

}

func TaskComplete(task string) {

}

func insertOneTask () {

}

func UndoTask () {

}

func deleteOneTask () {

}

func deleteAllTasks () {

}