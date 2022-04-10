package models

// import "go.mongodb.org/mongo-driver/bson/primitive"

type UserInfo struct {
	Username			string							`json:"username,omitempty"`
	Password			string							`json:"password,omitempty"`
	Payment				int									`json:"payment,omitempty"`
}