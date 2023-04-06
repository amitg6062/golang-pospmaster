package helpers

import (
	"database/sql"
	"log"
)

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

// type JsonResponse struct {
// 	Type    string `json:"type"`
// 	Data    []Emp  `json:"data"`
// 	Message string `json:"message"`
// }

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
		if err != nil {
			log.Fatal(err)
		}
		colNames, err := rows.Columns()
		if err != nil {
			log.Fatal(err)
		}
		these := make(map[string]interface{})
		for idx, name := range colNames {
			these[name] = *colVals[idx].(*interface{})
		}
		ret = append(ret, these)
	}
	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}
	return ret
}
