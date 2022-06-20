package rnd

import (
	"fmt"
	"net/http"

	// jwtauth "posp_api_go_v2/src/middleware/jwtAuth"

	jwtauth "github.com/amitg6062/golang-posp-jwtMiddleware"

	"github.com/gorilla/mux"
)

type Middleware func(http.HandlerFunc) http.HandlerFunc

func Middlewares(h http.HandlerFunc, m ...Middleware) http.HandlerFunc {

	if len(m) < 1 {
		return h
	}

	wrapped := h

	// loop in reverse to preserve middleware order
	for i := len(m) - 1; i >= 0; i-- {
		wrapped = m[i](wrapped)
	}

	return wrapped

}

func Router() *mux.Router {

	fmt.Println("RND Router")
	r := mux.NewRouter()

	r.HandleFunc("/amit", CallAmit).Methods("GET")
	r.HandleFunc("/v1/business/getLeadsByAffiliate", GetPartnerLeads).Methods("POST")
	r.HandleFunc("/users/{id}", GetUsers).Methods("GET")
	r.HandleFunc("/createEmp", CreateEmp).Methods("POST")

	r.HandleFunc("/jwt-get-token", JwtGetToken).Methods("GET")
	r.HandleFunc("/jwt-validiate", Middlewares(SuccessFunction, jwtauth.IsAuthorized)).Methods("GET")

	r.HandleFunc("/middleware", Middlewares(SuccessFunction, jwtauth.IsAuthorized, HiAmit))

	return r

}
