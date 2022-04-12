package models

import (
	"github.com/google/uuid"
)

type UserInfo struct {
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
	Payment  int    `json:"payment,omitempty"`
	CardDetails
}

type CardDetails struct {
	BillingDetails
	Metadata
	IdempotencyKey 		uuid.Domain
	EncryptedData  		string `json:"encryptedData,omitempty"`
	ExpMonth       		int    `json:"expMonth,omitempty"`
	ExpYear        		int    `json:"expYear,omitempty"`
	keyId							string
}

type BillingDetails struct {
	Name           		string `json:"name,omitempty"`
	City           		string `json:"city,omitempty"`
	District       		string `json:"state,omitempty"`
	Country        		string `json:"country,omitempty"`
	Address        		string `json:"address,omitempty"`
	ZipCode        		int    `json:"zipcode,omitempty"`
}

type Metadata struct {
	Email          		string `json:"email,omitempty"`
	SessionId      		string
	IpAddress      		string
}

type Reponse struct {
	Data							map[string] interface{}		`json:"-"`
}