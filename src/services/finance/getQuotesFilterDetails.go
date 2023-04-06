package finance

import (
	"fmt"
	"log"
	"net/http"
	h "posp_api_go_v2/src/helpers"

	lib "github.com/amitg6062/golang-posp-dbconnection"
	"github.com/gin-gonic/gin"
)

type FilterDetailsRequest struct {
	// GoalId is the ID of the goal to filter quotes by
	// Example: 1
	GoalId int `json:"GoalId" binding:"required"`
}

// FilterDetailsResponse represents the response body for the getQuotesFilterDetails endpoint
type FilterDetailsResponse struct {
	// Your response fields here
}

// @Summary Get filter details for quotes
// @Description Get filter details for quotes based on a goal ID
// @Tags quotes
// @ID get-quotes-filter-details
// @Accept json
// @Produce json
// @Param   GoalId   path      int  true  "1"
// @Success 200 {object} h.JsonResponse
// @Router /getQuotesFilterDetails [post]
func GetQuotesFilterDetailsHandler(c *gin.Context) {

	var requestBody interface{}
	var response h.JsonResponse
	// var response string

	err := c.BindJSON(&requestBody)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response = Service_getQuotesFilterDetailsHandler(requestBody)

	c.JSON(200, response)

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

	var response = h.JsonResponse{Data: ret}

	return response

}
