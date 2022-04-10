package models

// import "go.mongodb.org/mongo-driver/bson/primitive"

type UserInfo struct {
	Username			string							`json:"username,omitempty"`
	Password			string							`json:"password,omitempty"`
	Payment				int									`json:"payment,omitempty"`
	CardDetails
}

type CardDetails struct {
	IdempotencyKey				string
	EncryptedData					string			`json:"encryptedData,omitempty"`
	Name									string			`json:"name,omitempty"`
	City									string			`json:"city,omitempty"`
	Country								string			`json:"country,omitempty"`
	Address								string			`json:"address,omitempty"`
	ZipCode								int					`json:"zipcode,omitempty"`
	ExpMonth							int					`json:"expMonth,omitempty"`
	ExpYear								int					`json:"expYear,omitempty"`
	Email									string			`json:"email,omitempty"`
	SessionId							string
	IpAddress							string
}