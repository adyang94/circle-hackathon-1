package middleware

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/adyang94/circle-hackathon1/models"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"

	// "go.mongodb.org/mongo-driver/bson"
	// "go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var collection *mongo.Collection

var token string = "Bearer QVBJX0tFWTpjYTdmODZlNTNjN2ZmNDdmNjA5ZDRkNjg1ZmFlOTg3Nzo1NTUyYzk3NzkwOTczZmM2M2I5ZTNiNmFlYTgxOTI3Mg=="

var users = []models.UserInfo{
	{Username: "client1", Password: "client1", Payment: 123456},
	{Username: "client3", Password: "client3", Payment: 131313},
	{Username: "client2", Password: "client2", Payment: 123455},
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
	validUser := checkUserInfo(user)
	log.Println("ValidUser:  ", validUser)

	if validUser {
		json.NewEncoder(w).Encode(user)
		return
	}

	w.WriteHeader(http.StatusBadRequest)
	json.NewEncoder(w).Encode("Login Unsuccessful.  Invalid credentials.")
	return
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

	fmt.Println("1: ", res, "\n")

	fmt.Println("2: ", res.Body, "\n")
	fmt.Println("2.1: ", *res, "\n")

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println("3: ", res, "\n")
	fmt.Println("4: ", res.Body, "\n")
	fmt.Println("5: ", string(body))
	fmt.Println("6: ", body)

	//  Response back to client
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	// w.Header().Set("Access-Control-Allow-Origin", "*")

	//  This reponse is a stringified JSON.  Need to make into JSON.
	json.NewEncoder(w).Encode(string(body))
}

func CreatePayment(w http.ResponseWriter, r *http.Request) {

	//  Retrieving data from request
	var reqBody models.Response
	json.NewDecoder(r.Body).Decode(&reqBody.Data)
	log.Println("REQBODY:  ", reqBody)
	log.Println("REQBODY:  ", reqBody)

	//  Creating idempotencyKey
	idKey := uuid.New()
	log.Println(idKey)

	url := "https://api-sandbox.circle.com/v1/payments"

	payload := strings.NewReader("{\"metadata\":{\"email\":\"satoshi@circle.com\",\"sessionId\":\"DE6FA86F60BB47B379307F851E238617\",\"ipAddress\":\"244.28.239.130\"},\"amount\":{\"amount\":\"312\",\"currency\":\"USD\"},\"autoCapture\":true,\"source\":{\"id\":\"b8627ae8-732b-4d25-b947-1df8f4007a29\",\"type\":\"card\"},\"idempotencyKey\":\"" + idKey.String() + "\",\"keyId\":\"key1\",\"verification\":\"none\"}")

	log.Println("PAYLOAD: ", payload)

	req, _ := http.NewRequest("POST", url, payload)

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", token)

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println("RESPONSE:  ", res)
	fmt.Println("RESPONSE BODY:  ", string(body))

}

func CreateCard(w http.ResponseWriter, r *http.Request) {

}

func checkUserInfo(user models.UserInfo) bool {
	for _, profile := range users {
		if (user.Username == profile.Username) && (user.Password == profile.Password) {
			return true
		}
	}
	return false

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
