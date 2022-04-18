package router

import (
	"github.com/adyang94/circle-hackathon1/controller"
	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/api/login", controller.Login).Methods("POST", "OPTIONS")

	router.HandleFunc("/api/getListOfPayments", controller.GetListOfPayments).Methods("GET", "OPTIONS")

	router.HandleFunc("/api/createPayment", controller.CreatePayment).Methods("POST", "OPTIONS")

	router.HandleFunc("/api/addNewUser", controller.AddSingleUser).Methods("POST", "OPTIONS")

	return router
}
