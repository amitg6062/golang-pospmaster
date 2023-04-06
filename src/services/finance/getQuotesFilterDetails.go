package finance

import (
	"fmt"
	"log"
	"net/http"
	h "posp_api_go_v2/src/helpers"

	lib "github.com/amitg6062/golang-posp-dbconnection"
	"github.com/gin-gonic/gin"
)

func Handler_getQuotesFilterDetailsHandler(c *gin.Context) {

	var requestBody interface{}
	var response h.JsonResponse
	// var response string

	err := c.BindJSON(&requestBody)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response = Service_getQuotesFilterDetailsHandler(requestBody)

	c.JSON(200, gin.H{
		"message":  "User created successfully!",
		"response": response,
	})

}

func Service_getQuotesFilterDetailsHandler(requestBody interface{}) h.JsonResponse {

	fmt.Println("nested data")
	fmt.Println(requestBody)
	db := lib.InitialMigration()

	// Use a type assertion to convert the interface to a map[string]interface{}
	requestMap, ok := requestBody.(map[string]interface{})
	if !ok {
		// Handle the case where the interface doesn't represent a map[string]interface{}
	}

	GoalId := requestMap["GoalId"]

	tsql := fmt.Sprint("EXEC [finance].[QuotesFilterDetails] @GoalId=?")
	rows, err := db.Query(tsql, GoalId)

	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	ret := h.RenderData(rows)

	var response = h.JsonResponse{Error: false, Data: ret}

	return response

}
