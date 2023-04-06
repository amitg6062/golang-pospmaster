package api

import (
	"database/sql"
	"fmt"
	"log"

	lib "github.com/amitg6062/golang-posp-dbconnection"
	"github.com/gin-gonic/gin"
)

type JsonResponse2 struct {
	Error   bool                     `json:"error,bool"`
	Data    []map[string]interface{} `json:"data"`
	Message string                   `json:"message"`
}

func GetUsers2(c *gin.Context) {

	Conn := lib.InitialMigration()

	// response = ReadNewData(Conn, id)

	id := c.Param("id")
	// c.String(http.StatusOK, "Hello %s", id)
	var response JsonResponse2
	response = ReadNewData2(Conn, id)

	c.JSON(200, gin.H{
		"message":  "User created successfully!",
		"response": response,
	})
}

func ReadNewData2(db *sql.DB, id string) JsonResponse2 {

	tsql := fmt.Sprintf("EXEC [TestSchema].[GetEmployee] @id = $1;")
	rows, err := db.Query(tsql, id)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	ret := RenderData(rows)

	var response = JsonResponse2{Error: false, Data: ret}

	return response

}
