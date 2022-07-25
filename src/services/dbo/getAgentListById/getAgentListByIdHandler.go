package getAgentListById

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	lib "github.com/amitg6062/golang-posp-dbconnection"
	hf "github.com/amitg6062/golang-posp-helpers"
	"github.com/go-playground/validator"
)

func CallAmit(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	//Handle panic condition
	defer hf.Deferring()

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
		response = JsonResponse{Error: true, Data: ret, Message: err.Error()}

	} else {
		Conn := lib.InitialMigration()
		defer Conn.Close()
		response = ReadData(Conn, agentListRequest)
	}

	json.NewEncoder(w).Encode(response)

}
