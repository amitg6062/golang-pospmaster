package exportExcel

import (
	"fmt"

	"github.com/gorilla/mux"
)

func Router() *mux.Router {

	fmt.Println("Export Excel Router")
	r := mux.NewRouter()

	r.HandleFunc("/ge", GetUsers).Methods("GET")

	return r

}
