package main

import (
	"fmt"
	api "posp_api_go_v2/src"

	"github.com/ddadumitrescu/hellomod"
)

func main() {
	//Handle panic condition
	defer deferring()
	hellomod.Salut()
	api.Run()

}

func deferring() {
	if err := recover(); err != nil {
		fmt.Println("An error occurred:", err)
	}
}
