package api

import (
	"net/http"

	"posp_api_go_v2/docs"
	dbo "posp_api_go_v2/src/services/dbo"
	finance "posp_api_go_v2/src/services/finance"

	"github.com/gin-gonic/gin"

	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func setupRouter() *gin.Engine {

	// gin.DisableConsoleColor()
	r := gin.Default()

	// Serve Swagger documentation

	docs.SwaggerInfo.BasePath = ""

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	// Ping test
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	r.POST("/lead/insertUpdateAffiliateLeadDetails", dbo.InsertUpdateAffiliateLeadDetails)

	r.POST("/getQuotesFilterDetails", finance.GetQuotesFilterDetailsHandler)

	r.POST("/listFinancialGoalsQuotes", finance.ListFinancialGoalsQuotes)

	return r
}
