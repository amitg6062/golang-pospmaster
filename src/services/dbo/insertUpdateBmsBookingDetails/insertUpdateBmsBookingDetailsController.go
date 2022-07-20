package insertUpdateBmsBookingDetails

import (
	"database/sql"
	"fmt"
	"log"
)

func ReadData(db *sql.DB, agentList AgentListRequest) JsonResponse {
	//Handle panic condition
	defer deferring()

	tsql := fmt.Sprint("EXEC [dbo].[GetMyAgents] ")

	if agentList.IsVerified != "" {
		tsql = fmt.Sprint(tsql, "@IsVerified = '", agentList.IsVerified, "', ")
	}
	if agentList.SortDir != 0 {
		tsql = fmt.Sprint(tsql, "@SortDir = ", agentList.SortDir, ", ")
	}
	if agentList.SearchBy != "" {
		tsql = fmt.Sprint(tsql, "@SearchBy = '", agentList.SearchBy, "', ")
	}

	if agentList.Circle != "" {
		tsql = fmt.Sprint(tsql, "@Circle = '", agentList.Circle, "', ")
	}

	if agentList.ParentId != 0 {
		tsql = fmt.Sprint(tsql, "@ParentId = '", agentList.ParentId, "', ")
	}
	if agentList.IsPanVerified != "" {
		tsql = fmt.Sprint(tsql, "@IsPanVerified = '", agentList.IsPanVerified, "', ")
	}
	if agentList.IsGreenChannel != "" {
		tsql = fmt.Sprint(tsql, "@IsGreenChannel = ", agentList.IsGreenChannel, ", ")
	}
	if agentList.FromDate != "" && agentList.ToDate != "" {
		tsql = fmt.Sprint(tsql, "@FromDate = '", agentList.FromDate, "', ")
		tsql = fmt.Sprint(tsql, "@ToDate = '", agentList.ToDate, "', ")
	}

	if agentList.AffilateId != 0 {
		tsql = fmt.Sprint(tsql, "@AffiliateId = '", agentList.AffilateId, "', ")
	}

	if agentList.PageSize != "" {
		tsql = fmt.Sprint(tsql, "@PageSize = ", agentList.PageSize, ", ")
	}

	if agentList.Type != "" {
		tsql = fmt.Sprint(tsql, "@Type = ", agentList.Type, ", ")
	}

	if agentList.PageNum != 0 {
		tsql = fmt.Sprint(tsql, "@PageNum = ", agentList.PageNum, "; ")
	}

	fmt.Println(tsql)

	rows, err := db.Query(tsql)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
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

	var response = JsonResponse{Error: false, Data: ret}

	return response

}

// Function for handling errors
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

func deferring() {
	if err := recover(); err != nil {
		fmt.Println("An error occurred:", err)
	}
}
