package partnerUpgrade

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	lib "github.com/amitg6062/golang-posp-dbconnection"
	hf "github.com/amitg6062/golang-posp-helpers"
	"github.com/go-playground/validator"
)

func HandleRequest(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	//Handle panic condition
	defer hf.Deferring()

	var response JsonResponse

	reqBody, _ := ioutil.ReadAll(r.Body)
	var requestParam RequestParam

	json.Unmarshal(reqBody, &requestParam)

	//Validation
	validate := validator.New()
	err := validate.Struct(requestParam)

	// check for validation
	if err != nil {
		ret := make([]map[string]interface{}, 0)
		response = JsonResponse{Error: true, Data: ret, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	Conn := lib.InitialMigration()
	response = ReadData(Conn, requestParam)

	json.NewEncoder(w).Encode(response)

}
