package router

import (
	"github.com/adyang94/react-go-todo-app/middleware"
	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/api/login", middleware.Login).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/getListOfPayments", middleware.GetListOfPayments).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/createPayment", middleware.CreatePayment).Methods("POST", "OPTIONS")

	/*
		router.HandleFunc("/api/task", middleware.GetAllTasks).Methods("GET", "OPTIONS")
		router.HandleFunc("/api/task", middleware.CreateTask).Methods("POST", "OPTIONS")
		router.HandleFunc("/api/tasks/{id}", middleware.TaskComplete).Methods("PUT", "OPTIONS")
		router.HandleFunc("/api/undoTask/{id}", middleware.UndoTask).Methods("PUT", "OPTIONS")
		router.HandleFunc("api/deleteTask/{id}", middleware.DeleteTask).Methods("DELETE", "OPTIONS")
		router.HandleFunc("/api/deleteAllTasks", middleware.DeleteAllTasks).Methods("DELETE", "OPTIONS")
	*/

	return router
}
