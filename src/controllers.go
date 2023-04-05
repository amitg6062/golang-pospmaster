package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	lib "github.com/amitg6062/golang-posp-dbconnection"
	"github.com/gin-gonic/gin"
)

func helloController(c *gin.Context) {

	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.AbortWithStatus(400)
		return
	}

	var user User
	err = json.Unmarshal(body, &user)
	if err != nil {
		c.AbortWithStatus(400)
		return
	}

	// Do something with the user

	fmt.Printf("Received user: %+v\n", user)

	c.JSON(200, gin.H{
		"message":  "User created successfully!",
		"username": user.Name,
	})
}

func GetUsers(c *gin.Context) {

	Conn := lib.InitialMigration()

	// response = ReadNewData(Conn, id)

	id := c.Param("id")
	// c.String(http.StatusOK, "Hello %s", id)
	var response JsonResponse
	response = ReadNewData(Conn, id)

	c.JSON(200, gin.H{
		"message":  "User created successfully!",
		"response": response,
	})
}

// func ReadNewData(db *sql.DB, id string) JsonResponse {

// 	var emp []Emp
// 	tsql := fmt.Sprintf("EXEC [TestSchema].[GetEmployee] @id = $1;")
// 	rows, err := db.Query(tsql, id)

// 	if err != nil {
// 		fmt.Println("Error reading rows: " + err.Error())

// 	}
// 	defer rows.Close()

// 	for rows.Next() {
// 		var ep Emp
// 		err := rows.Scan(&ep.Id, &ep.Name, &ep.Location)

// 		checkErr(err)

// 		fmt.Printf("ID-newdata: %d, Name: %s, Location: %s\n", ep.Id, ep.Name, ep.Location)

// 		emp = append(emp, ep)

// 	}

// 	var response = JsonResponse{Type: "success", Data: emp}

// 	//	log.Fatal("Data is ", response)

// 	return response

// }

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

// Function for handling messages
func printMessage(message string) {
	fmt.Println("")
	fmt.Println(message)
	fmt.Println("")
}

/*
func helloController(c *gin.Context) {

	username := c.PostForm("username")
	fmt.Println(username)

	// id := c.Query("id")
	// page := c.DefaultQuery("page", "0")
	// name := c.PostForm("username")
	// message := c.PostForm("message")
	// fmt.Printf("id: %s; page: %s; name: %s; message: %s", id, page, name, message)
	// password := c.PostForm("password")

	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.AbortWithStatus(400)
		return
	}

	var user User
	err = json.Unmarshal(body, &user)
	if err != nil {
		c.AbortWithStatus(400)
		return
	}

	// Do something with the user

	fmt.Printf("Received user: %+v\n", user)

	jsonData, err := json.Marshal(user)
	if err != nil {
		fmt.Println("Error marshalling JSON:", err)
		return
	}

	fmt.Println(string(jsonData))

	c.JSON(200, gin.H{
		"message": "User created successfully!",
		"data":    jsonData,
	})

	// c.JSON(200, gin.H{
	// 	"message": username,
	// 	"data":    "sdfasdf",
	// })
}
*/
