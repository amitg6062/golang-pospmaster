package insertUpdateBmsBookingDetails

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	lib "github.com/amitg6062/golang-posp-dbconnection"
	"github.com/go-playground/validator"
)

func HandleRequest(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	//Handle panic condition
	defer deferring()

	var response JsonResponse
	var dataArr JsonResponse

	reqBody, _ := ioutil.ReadAll(r.Body)
	var requestParam RequestParam

	json.Unmarshal(reqBody, &requestParam)

	//Validation
	validate := validator.New()
	err := validate.Struct(requestParam)

	// check for validation
	if err != nil {
		ret := make([]map[string]interface{}, 0)
		response = JsonResponse{Error: false, Data: ret, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	Conn := lib.InitialMigration()
	dataArr = GetLeadDetailsByIdAndProductId_v1(Conn, requestParam.LeadID)

	//log.Fatal()

	if dataArr.Data[0]["LeadID"] == "" {

		ret := make([]map[string]interface{}, 0)
		response = JsonResponse{Error: true, Data: ret, Message: "Bypass"}

	} else {

		Conn := lib.InitialMigration()
		insertUpdateVehicleDetails_v2(Conn, requestParam)
		response = insertUpdateVehicleDetails_v2(Conn, requestParam)

	}

	json.NewEncoder(w).Encode(response)

}
