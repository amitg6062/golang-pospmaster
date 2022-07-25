package getAgentListById

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	lib "github.com/amitg6062/golang-posp-dbconnection"
	"github.com/go-playground/validator"
)

func CallAmit(w http.ResponseWriter, r *http.Request) {
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
		response = JsonResponse{Error: true, Data: ret, Message: err.Error()}

	} else {
		Conn := lib.InitialMigration()
		response = ReadData(Conn, agentListRequest)
	}

	json.NewEncoder(w).Encode(response)

}

/*
func GetUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var response JsonResponse
	params := mux.Vars(r)
	id := params["id"]

	Conn := lib.InitialMigration()

	response = ReadNewData(Conn, id)

	json.NewEncoder(w).Encode(response)
}


// use a single instance of Validate, it caches struct info
var validate *validator.Validate

func CreateEmp(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	validate = validator.New()
	var emp Emp2
	reqBody, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(reqBody, &emp)

	// returns nil or ValidationErrors ( []FieldError )
	err := validate.Struct(emp)
	if err != nil {

		// this check is only needed when your code could produce
		// an invalid value for validation such as interface with nil
		// value most including myself do not usually have code like this.
		// if _, ok := err.(*validator.InvalidValidationError); ok {
		// 	fmt.Println(err)
		// 	return
		// }

		var ThisNote string

		for _, err := range err.(validator.ValidationErrors) {

			errMsg := "Error Occured on field " + err.Field() + " with validation " + err.Tag() + "\n"

			fmt.Println("====== Start ==========")
			fmt.Printf("Error Occured on field " + err.Field())
			fmt.Printf(" with validation " + err.Tag())
			fmt.Printf(" with value ")
			fmt.Println(err.Value())
			fmt.Println("====== End ==========")

			//fmt.Println(err.Namespace())
			//fmt.Println(err.StructNamespace())
			//fmt.Println(err.StructField())
			//fmt.Println(err.Tag())
			//fmt.Println(err.ActualTag())
			//fmt.Println(err.Kind())
			//fmt.Println(err.Type())
			//fmt.Println(err.Value())
			//fmt.Println(err.Param())
			// fmt.Println()

			ThisNote += errMsg
		}

		json.NewEncoder(w).Encode(ThisNote)

		return
	}

	fmt.Println("Created Work")

	json.NewEncoder(w).Encode(emp)

}

func CreateEmpOld(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	//name := r.FormValue("username")
	//location := r.FormValue("location")

	reqBody, _ := ioutil.ReadAll(r.Body)
	var emp Emp
	json.Unmarshal(reqBody, &emp)

	json.NewEncoder(w).Encode(emp)

	newData, err := json.Marshal(emp)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(string(newData))
	}

}

func GetPartnerLeads(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	//validate = validator.New()
	var pl PartnerLeadsSchema //var emp m.Emp2
	var response PartnerLeadsFinalResponse
	reqBody, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(reqBody, &pl)

	// returns nil or ValidationErrors ( []FieldError )
	//err := validate.Struct(emp)

	Conn := lib.InitialMigration()
	response = GetPartnerLeadsC(Conn, &pl)

	json.NewEncoder(w).Encode(response)
}
*/
