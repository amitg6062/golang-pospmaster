package api

import (
	"net/http"

	"github.com/gin-gonic/gin"

	dbo "posp_api_go_v2/src/services/dbo"
	finance "posp_api_go_v2/src/services/finance"
)

func setupRouter() *gin.Engine {

	// gin.DisableConsoleColor()
	r := gin.Default()

	// Ping test
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	r.POST("/lead/insertUpdateAffiliateLeadDetails", dbo.InsertUpdateAffiliateLeadDetails)

	r.POST("/getQuotesFilterDetails", finance.GetQuotesFilterDetailsHandler)

	r.POST("/listFinancialGoalsQuotes", finance.ListFinancialGoalsQuotes)

	return r
}
