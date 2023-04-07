package main

import (
	api "posp_api_go_v2/src"
	"posp_api_go_v2/src/helpers"
)

func main() {
	//Handle panic condition
	defer helpers.Deferring()

	api.Run()

}
