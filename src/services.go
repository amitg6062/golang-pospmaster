package api

import (
	"database/sql"
	"fmt"
	"log"
)

func ReadNewData(db *sql.DB, id string) JsonResponse {

	tsql := fmt.Sprintf("EXEC [TestSchema].[GetEmployee] @id = $1;")
	rows, err := db.Query(tsql, id)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	ret := RenderData(rows)

	var response = JsonResponse{Error: false, Data: ret}

	return response

}

/*
func ReadNewData(db *sql.DB, id string) JsonResponse {

	var emp []Emp
	tsql := fmt.Sprintf("EXEC [TestSchema].[GetEmployee] @id = $1;")
	rows, err := db.Query(tsql, id)

	if err != nil {
		fmt.Println("Error reading rows: " + err.Error())

	}
	defer rows.Close()

	for rows.Next() {
		var ep Emp
		err := rows.Scan(&ep.Id, &ep.Name, &ep.Location)

		checkErr(err)

		fmt.Printf("ID-newdata: %d, Name: %s, Location: %s\n", ep.Id, ep.Name, ep.Location)

		emp = append(emp, ep)

	}

	var response = JsonResponse{Type: "success", Data: emp}

	//	log.Fatal("Data is ", response)

	return response

}
*/
