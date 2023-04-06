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

	// // Monolithic
	r.GET("/api/rnd/users/:id", GetUsers)

	// // microservice
	r.GET("/api/rnd/users2/:id", GetUsers2)

	r.POST("/lead/insertUpdateAffiliateLeadDetails2", dbo.Handler_insertUpdateAffiliateLeadDetails)

	r.POST("/getQuotesFilterDetails", finance.Handler_getQuotesFilterDetailsHandler)

	return r
}
