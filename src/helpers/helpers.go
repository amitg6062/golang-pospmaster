package helpers

import (
	"database/sql"
	"fmt"
	"strings"
)

type Emp struct {
	Id       int    `json:"id"`
	Name     string `json:"username"`
	Location string `json:"location"`
}

type JsonResponse struct {
	Error   bool                     `json:"error,bool"`
	Data    []map[string]interface{} `json:"data"`
	Message string                   `json:"message"`
}

type Response struct {
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Errors  interface{} `json:"errors"`
	Data    interface{} `json:"data"`
}

//Empty object is used when data doesnot want to be null on json
type EmptyObj struct{}

//BuildResponse method is used toinject data value to dynamic success response
func BuildResponse(status bool, message string, data interface{}) Response {
	res := Response{
		Status:  status,
		Message: message,
		Errors:  nil,
		Data:    data,
	}
	return res
}

//BuildErrorResponse method is used toinject data value to dynamic error response
func BuildErrorResponse(message string, err string, data interface{}) Response {
	splittedError := strings.Split(err, "\n")
	res := Response{
		Status:  false,
		Message: message,
		Errors:  splittedError,
		Data:    data,
	}

	return res
}

//Function to get data from db.
func RenderData(rows *sql.Rows) []map[string]interface{} {
	var err error
	cols, _ := rows.Columns()
	ret := make([]map[string]interface{}, 0)
	for rows.Next() {
		colVals := make([]interface{}, len(cols))
		for i := range colVals {
			colVals[i] = new(interface{})
		}
		err = rows.Scan(colVals...)
		CheckErr(err)
		colNames, err := rows.Columns()
		CheckErr(err)
		these := make(map[string]interface{})
		for idx, name := range colNames {
			these[name] = *colVals[idx].(*interface{})
		}
		ret = append(ret, these)
	}
	if err = rows.Err(); err != nil {
		CheckErr(err)
	}
	return ret
}

//Function to handle Panic
func Deferring() {
	if err := recover(); err != nil {
		fmt.Println("An error occurred:", err)
	}
}

// Function for handling errors
func CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}
