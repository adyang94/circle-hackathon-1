package middleware

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/adyang94/circle-hackathon1/models"
	"github.com/golang-jwt/jwt"
)

func ValidateAndRefreshToken1(tokenStr string) bool {

	return true
}

func ValidateAndRefreshToken(w http.ResponseWriter, r *http.Request) http.ResponseWriter {
	log.Println("Validate and Refresh Token")

	cookie, err := r.Cookie("token")

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Please login to get list of payments.")
		return w
	}

	log.Println("Cookie:  ", cookie, err)
	tokenStr := cookie.Value

	var claims = &models.Claims{}

	var jwtKey = []byte(os.Getenv("JWT_KEY"))
	log.Println("refresh token str: ", tokenStr, "claims: ", claims)

	tkn, err := jwt.ParseWithClaims(tokenStr, claims,
		func(t *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})

	log.Println("refresh token: ", tkn)
	log.Println("refresh token1: ", err)

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			log.Println("refresh token2")
			w.WriteHeader(http.StatusUnauthorized)
			return w
		}
		log.Println("refresh token3")
		w.WriteHeader(http.StatusBadRequest)
		return w
	}
	if !tkn.Valid {
		log.Println("refresh token4")
		w.WriteHeader(http.StatusUnauthorized)
		return w
	}

	// if time.Unix(claims.ExpiresAt, 0).Sub(time.Now()) > 30*time.Second {
	// 	w.WriteHeader(http.StatusBadRequest)
	// 	return
	// }

	expirationTime := time.Now().Add(time.Minute * 5)

	claims.ExpiresAt = expirationTime.Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return w
	}

	http.SetCookie(w,
		&http.Cookie{
			Name:    "refresh_token",
			Value:   tokenString,
			Expires: expirationTime,
		})

	return w
}
