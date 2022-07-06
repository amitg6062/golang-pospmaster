package exportExcel

import (
	"encoding/json"
	"fmt"
	"net/http"

	lib "github.com/amitg6062/golang-posp-dbconnection"
	"github.com/gorilla/mux"
)

func CallAmit(w http.ResponseWriter, r *http.Request) {
	fmt.Println("HEllo amit gupta")
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var response JsonResponse
	params := mux.Vars(r)
	id := params["id"]

	Conn := lib.InitialMigration()

	response = ReadNewData(Conn, id)

	json.NewEncoder(w).Encode(response)
}
