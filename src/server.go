package api

import (
	"fmt"
	"log"

	// "posp_api_go_v2/src/lib"

	lib "github.com/amitg6062/golang-posp-dbconnection"
	"github.com/joho/godotenv"
)

func init() {
	// loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		log.Print("sad .env file found")
	}
}

func Run() {

	//var err error
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error getting env, %v", err)
	} else {
		// s3Bucket := os.Getenv("S3_BUCKET")
		fmt.Println("We are getting the env values")
	}

	//Load DB
	lib.InitialMigration()

	//InitizeRouter for routing
	// initializeRouter()

	r := setupRouter()
	// Listen and Server in 0.0.0.0:8080
	r.Run(":8081")

}
