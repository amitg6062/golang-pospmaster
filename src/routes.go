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

	route := gin.Default()

	route.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	docs.SwaggerInfo.BasePath = ""

	route.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	route.POST("/lead/insertUpdateAffiliateLeadDetails", dbo.InsertUpdateAffiliateLeadDetails)

	route.POST("/getQuotesFilterDetails", finance.GetQuotesFilterDetailsHandler)

	route.POST("/listFinancialGoalsQuotes", finance.ListFinancialGoalsQuotes)

	return route
}
