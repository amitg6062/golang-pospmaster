package main

import (
	"fmt"
	api "posp_api_go_v2/src"

	hf "github.com/amitg6062/golang-posp-helpers"
	"github.com/ddadumitrescu/hellomod"
)

func main() {
	//Handle panic condition
	defer deferring()
	hellomod.Salut()

	/*
		//Send Http request.
		reader := strings.NewReader(`{"ProductId":123,"SupplierId":5673,"PlanId":98}`)
		request, _ := http.NewRequest("POST", "https://reqres.in/api/users", reader)
		//request.Header.Set("Content-Type: application/json", "'token:Content-Length:")
		// TODO: check err
		client := &http.Client{}
		resp, _ := client.Do(request)
		body, _ := ioutil.ReadAll(resp.Body)
		bodyString := string(body)
		fmt.Println(resp)
		fmt.Println(bodyString)
		// TODO: check err
	*/

	hf.TestGolangHelper()

	api.Run()

}

func deferring() {
	if err := recover(); err != nil {
		fmt.Println("An error occurred:", err)
	}
}
