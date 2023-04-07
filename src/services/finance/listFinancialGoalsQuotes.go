package finance

import (
	"fmt"
	"log"
	"net/http"
	"posp_api_go_v2/src/helpers"

	lib "github.com/amitg6062/golang-posp-dbconnection"
	"github.com/gin-gonic/gin"
)

// @Summary ListFinancialGoalsQuotes
// @Description ListFinancialGoalsQuotes
// @Tags finance
// @ID ListFinancialGoalsQuotes
// @Accept json
// @Produce json
// @Param   GoalId   path      int  true  "1"
// @Success 200 {object} helpers.JsonResponse
// @Router /listFinancialGoalsQuotes [post]
func ListFinancialGoalsQuotes(c *gin.Context) {

	var requestBody interface{}
	var response helpers.JsonResponse
	// var response string

	err := c.BindJSON(&requestBody)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response = Service_listFinancialGoalsQuotes(requestBody)

	c.JSON(200, response)

}

func Service_listFinancialGoalsQuotes(requestBody interface{}) helpers.JsonResponse {

	fmt.Println("nested data")
	fmt.Println(requestBody)
	db := lib.InitialMigration()

	// Use a type assertion to convert the interface to a map[string]interface{}
	requestMap, ok := requestBody.(map[string]interface{})
	if !ok {
		// Handle the case where the interface doesn't represent a map[string]interface{}
	}

	GoalId := requestMap["GoalId"]

	tsql := fmt.Sprint("EXEC [finance].[ListFinancialGoalsQuotes] @GoalId=?")
	rows, err := db.Query(tsql, GoalId)

	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	ret := helpers.RenderData(rows)

	var response = helpers.JsonResponse{Data: ret}

	return response

}
