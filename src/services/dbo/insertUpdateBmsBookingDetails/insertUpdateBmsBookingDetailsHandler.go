package insertUpdateBmsBookingDetails

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	lib "github.com/amitg6062/golang-posp-dbconnection"
	"github.com/go-playground/validator"
)

func CallAmit2(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	//Handle panic condition
	defer deferring()

	var response JsonResponse

	reqBody, _ := ioutil.ReadAll(r.Body)
	var agentListRequest AgentListRequest

	json.Unmarshal(reqBody, &agentListRequest)

	//Validation
	validate := validator.New()
	err := validate.Struct(agentListRequest)

	// check errors
	if err != nil {
		ret := make([]map[string]interface{}, 0)
		response = JsonResponse{Error: false, Data: ret, Message: err.Error()}

	} else {
		Conn := lib.InitialMigration()
		response = ReadData(Conn, agentListRequest)
	}

	json.NewEncoder(w).Encode(response)

}
