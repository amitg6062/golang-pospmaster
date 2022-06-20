package reports

import (
	"fmt"

	"github.com/gorilla/mux"
)

func Router() *mux.Router {

	fmt.Println("Report Router")
	r := mux.NewRouter()

	r.HandleFunc("/amit", CallAmit).Methods("GET")
	r.HandleFunc("/v1/business/getLeadsByAffiliate", GetPartnerLeads).Methods("POST")
	r.HandleFunc("/users/{id}", GetUsers).Methods("GET")
	r.HandleFunc("/createEmp", CreateEmp).Methods("POST")
	return r

}
