package middleware

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/adyang94/react-go-todo-app/models"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"

	// "go.mongodb.org/mongo-driver/bson"
	// "go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var collection *mongo.Collection

var token string = "Bearer QVBJX0tFWTpiNmFjMzFmMzAzOGJjNWRhMjNkY2ViMzk3NDBkNTk5MTozMGViYjliNTM5NDQ5M2I4YzU5M2NiODQzMDFkOTI1MA=="

var users = []models.UserInfo{
	{Username: "client1", Password: "client1", Payment: 123456},
}

func init() {
	loadTheEnv()
	createDBInstance()
}

func loadTheEnv() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading the .env file")
	}
}

func createDBInstance() {
	connectionString := os.Getenv("DB_URI")
	dbName := os.Getenv("DB_NAME")
	collName := os.Getenv("DB_COLLECTION_NAME")

	clientOptions := options.Client().ApplyURI(connectionString)

	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("connected to mongodb!")

	collection = client.Database(dbName).Collection(collName)
	fmt.Println("collection instance created")
}

func Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	var user models.UserInfo

	json.NewDecoder(r.Body).Decode(&user)
	var validUser := checkUserInfo(user)
	json.NewEncoder(w).Encode(user)
}

func GetListOfPayments(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get list of payments!")
	 
	//  Note::  Circle API returns all payments from then token specified, whether from same or different users.
	//  We will use logged in users payment id to parse the list of payments, and return those that match.

	//  Check if user is logged in or has valid JWT.  If not, alert user to login.

	//  Check if user has a payment stored

	//  Making call to Circle API
	url := "https://api-sandbox.circle.com/v1/payments"

	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Authorization", token)

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))

	//  Response back to client
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
}

func CreatePayment (w http.ResponseWriter, r *http.Request) {
	var card models.CardDetails
}

func checkUserInfo(user models.UserInfo) {
	return true
}

/*
-----------------------------------------------------------
*/

func GetAllTasks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	fmt.Println("Get all tasks start!")

	// payload := getAllTasks()
	// json.NewEncoder(w).Encode(payload)
}

func CreateTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	var task models.UserInfo

	json.NewDecoder(r.Body).Decode(&task)
	// insertOneTask(task)
	json.NewEncoder(w).Encode(task)
}

func TaskComplete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "PUT")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	params := mux.Vars(r)
	// taskComplete(params["id"])
	json.NewEncoder(w).Encode(params["id"])
}
