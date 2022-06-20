package jwtauth

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt"
)

type Error struct {
	IsError bool   `json:"isError"`
	Message string `json:"message"`
}

var (
	secretkey string = "secretkeyjwt"
)

//Generate JWT token
func GenerateJWT(email, role string) (string, error) {
	var mySigningKey = []byte(secretkey)
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["authorized"] = true
	claims["email"] = email
	claims["role"] = role
	claims["exp"] = time.Now().Add(time.Minute * 30).Unix()

	tokenString, err := token.SignedString(mySigningKey)
	if err != nil {
		fmt.Errorf("Something went Wrong: %s", err.Error())
		return "", err
	}

	return tokenString, nil
}

//check whether user is authorized or not
func IsAuthorized(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		tokenString := r.Header["Token"][0]

		if r.Header["Token"] == nil {
			var err Error
			err = SetError(err, "No Token Found")
			json.NewEncoder(w).Encode(err)
			return
		}

		var mySigningKey = []byte(secretkey)

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("There was an error in parsing token.")
			}
			return mySigningKey, nil
		})

		if err != nil {
			var err Error
			err = SetError(err, "Your Token has been expired.")
			json.NewEncoder(w).Encode(err)
			return
		}

		if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {

			fmt.Println("hello amit gupta Function 1")
			//r.Header.Set("Role", "admin")
			fmt.Println("Admin Authorised Success")
			handler.ServeHTTP(w, r)
			return
		}

		var reserr Error
		reserr = SetError(reserr, "Not Authorized.")
		json.NewEncoder(w).Encode(err)
	}
}

//set error message in Error struct
func SetError(err Error, message string) Error {
	err.IsError = true
	err.Message = message
	return err
}
