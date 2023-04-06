package api

import (
	"net/http"

	"github.com/gin-gonic/gin"

	api "posp_api_go_v2/src/services"
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

	r.POST("/lead/insertUpdateAffiliateLeadDetails", api.Handler_insertUpdateAffiliateLeadDetails)
	r.POST("/getQuotesFilterDetails", api.Handler_getQuotesFilterDetailsHandler)

	return r
}
