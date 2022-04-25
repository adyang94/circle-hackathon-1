package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"reflect"
	"strings"
	"time"

	"github.com/adyang94/circle-hackathon1/middleware"
	"github.com/adyang94/circle-hackathon1/models"
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"

	// "go.mongodb.org/mongo-driver/bson"
	// "go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

/* ------------------------- Variables ------------------------- */
var collection *mongo.Collection

var circleToken string = "Bearer QVBJX0tFWTpjYTdmODZlNTNjN2ZmNDdmNjA5ZDRkNjg1ZmFlOTg3Nzo1NTUyYzk3NzkwOTczZmM2M2I5ZTNiNmFlYTgxOTI3Mg=="

var users = []models.UserInfo{
	{Username: "client1", Password: "client1", Payment: 123456},
	{Username: "client3", Password: "client3", Payment: 131313},
	{Username: "client2", Password: "client2", Payment: 123455},
}

var testCards = []string{
	"4757140000000001", "5102420000000006",
}

var jwtKey = []byte(os.Getenv("JWT_KEY"))

/* ------------------------- Code ------------------------- */
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

		// log.Println("1")

		expirationTime := time.Now().Add(time.Minute * 5)

		var claims models.Claims

		// log.Println("2")

		claims.Username = user.Username
		claims.StandardClaims = jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		}

		// log.Println("3: ", claims)

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		tokenString, err := token.SignedString(jwtKey)

		// log.Println("4", tokenString)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		http.SetCookie(w,
			&http.Cookie{
				Name:    "token",
				Value:   tokenString,
				Expires: expirationTime,
			})

		json.NewEncoder(w).Encode(user)
		return
	}

	w.WriteHeader(http.StatusBadRequest)
	json.NewEncoder(w).Encode("Login Unsuccessful.  Invalid credentials.")
	return
}

func GetListOfPayments(w http.ResponseWriter, r *http.Request) {

	validToken := middleware.ValidateAndRefreshToken(w, r)

	if !validToken {
		log.Println("INVALID TOKEN: ", validToken)
	}

	//  Check if user is logged in or has valid JWT.  If not, alert user to login.
	cookie, err := r.Cookie("token")
	log.Println("Cookie:  ", cookie)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Please login to get list of payments.")
		return
	}

	// log.Println("Cookie1:  ", cookie, err)

	// tokenStr := cookie.Value
	// var claims = &models.Claims{}
	// // log.Println("token string1: ", tokenStr, "claims: ", claims)

	// tkn, err := jwt.ParseWithClaims(tokenStr, claims,
	// 	func(t *jwt.Token) (interface{}, error) {
	// 		return jwtKey, nil
	// 	},
	// )

	// log.Println("token parsed with claims: ", tokenStr, "\ntkn: ", tkn, "\nerror: ", err)

	// if err != nil {
	// 	if err == jwt.ErrSignatureInvalid {
	// 		w.WriteHeader(http.StatusUnauthorized)
	// 		return
	// 	}
	// 	w.WriteHeader(http.StatusBadRequest)
	// 	return
	// }
	// if !tkn.Valid {
	// 	w.WriteHeader(http.StatusUnauthorized)
	// 	return
	// }

	// expirationTime := time.Now().Add(time.Minute * 5)

	// claims.ExpiresAt = expirationTime.Unix()

	// token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// tokenString, err := token.SignedString(jwtKey)

	// if err != nil {
	// 	w.WriteHeader(http.StatusInternalServerError)
	// 	return
	// }

	// http.SetCookie(w,
	// 	&http.Cookie{
	// 		Name:    "refresh_token",
	// 		Value:   tokenString,
	// 		Expires: expirationTime,
	// 	})

	//  Note::  Circle API returns all payments from then token specified, whether from same or different users.
	//  We will use logged in users payment id to parse the list of payments, and return those that match.

	//  Check if user has a payment stored

	//  Making call to Circle API

	fmt.Println("Get list of payments!")

	url := "https://api-sandbox.circle.com/v1/payments"

	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Authorization", circleToken)

	res, _ := http.DefaultClient.Do(req)

	// fmt.Println("1: ", res)

	// fmt.Println("2: ", res.Body)
	// fmt.Println("2.1: ", *res)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	// fmt.Println("3: ", res)
	// fmt.Println("4: ", res.Body)
	// fmt.Println("5: ", string(body))
	// fmt.Println("6: ", body)

	//  Response back to client
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	// w.Header().Set("Access-Control-Allow-Origin", "*")

	//  This reponse is a stringified JSON.  Need to make into JSON.
	json.NewEncoder(w).Encode(string(body))
}

func CreatePayment(w http.ResponseWriter, r *http.Request) {
	//  Retrieving data from request
	var PaymentDetails models.PaymentDetails
	json.NewDecoder(r.Body).Decode(&PaymentDetails)
	log.Println("REQBODY:  ", PaymentDetails)
	log.Println("REQBODY LEN:  ", len(PaymentDetails.Metadata))

	if len(PaymentDetails.Metadata) == 0 || len(PaymentDetails.Source) == 0 || len(PaymentDetails.Amount) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Payment Unsuccessful.  Please check submitted information and try again.")
		return
	}

	validPaymentMethod := checkPaymentMethod(PaymentDetails)
	log.Println("Valid Payment Method:  ", validPaymentMethod)

	if !validPaymentMethod {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Payment Unsuccessful.  Make sure to use test credit cards.")
		return
	}

	log.Println("REQBODY.DATA:  ", PaymentDetails.Metadata)
	log.Println("REQBODY.DATA1:  ", reflect.TypeOf(PaymentDetails.Metadata["email"]))

	//  Creating idempotencyKey
	idKey := uuid.New()
	log.Println(idKey)

	url := "https://api-sandbox.circle.com/v1/payments"

	payload := strings.NewReader("{\"metadata\":{\"email\":\"" + PaymentDetails.Metadata["email"].(string) + "\",\"sessionId\":\"DE6FA86F60BB47B379307F851E238617\",\"ipAddress\":\"244.28.239.130\"},\"amount\":{\"amount\":\"" + PaymentDetails.Amount["amount"].(string) + "\",\"currency\":\"USD\"},\"autoCapture\":true,\"source\":{\"id\":\"b8627ae8-732b-4d25-b947-1df8f4007a29\",\"type\":\"card\"},\"idempotencyKey\":\"" + idKey.String() + "\",\"keyId\":\"key1\",zxx\"verification\":\"none\"}")

	log.Println("PAYLOAD: ", payload)

	req, _ := http.NewRequest("POST", url, payload)

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", circleToken)

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println("RESPONSE:  ", res)
	fmt.Println("RESPONSE BODY:  ", string(body))
}

func checkUserInfo(user models.UserInfo) bool {
	for _, profile := range users {
		if (user.Username == profile.Username) && (user.Password == profile.Password) {
			return true
		}
	}
	return false

}

func checkPaymentMethod(payment models.PaymentDetails) bool {
	for _, card := range testCards {
		if payment.Source["id"] == card {
			return true
		}
	}
	return false
}

func AddSingleUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	var user models.UserInfo

	json.NewDecoder(r.Body).Decode(&user)

	insertResult, err := collection.InsertOne(context.Background(), user)

	if err != nil {
		log.Fatal(err)
	}

	log.Println("Added one new user:  ", insertResult)
}

/*
-----------------------------------------------------------
*/

func TaskComplete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "PUT")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	params := mux.Vars(r)
	// taskComplete(params["id"])
	json.NewEncoder(w).Encode(params["id"])
}
