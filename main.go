package main

import (
	"fmt"
	api "posp_api_go_v2/src"
)

func main() {
	//Handle panic condition
	defer deferring()

	api.Run()

}

func deferring() {
	if err := recover(); err != nil {
		fmt.Println("An error occurred:", err)
	}
}
