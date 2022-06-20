package rnd

import (
	"database/sql"
	"fmt"
	"log"
)

// ReadEmployees read all employees
func ReadEmployees(db *sql.DB) (int, error) {

	tsql := fmt.Sprintf("SELECT Id, Name, Location FROM TestSchema.Employees;")
	rows, err := db.Query(tsql)

	if err != nil {
		fmt.Println("Error reading rows: " + err.Error())
		return -1, err
	}
	defer rows.Close()
	count := 0
	for rows.Next() {
		var name, location string
		var id int
		err := rows.Scan(&id, &name, &location)
		if err != nil {
			fmt.Println("Error reading rows: " + err.Error())
			return -1, err
		}
		fmt.Printf("ID: %d, Name: %s, Location: %s\n", id, name, location)
		count++
	}
	return count, nil
}

// ReadEmployees read all employees
func ReadNewDataOLD8june(db *sql.DB, id string) JsonResponse {
	tsql := fmt.Sprintf("EXEC [TestSchema].[GetEmployee] @id = $1;")
	rows, err := db.Query(tsql, id)

	if err != nil {
		fmt.Println("Error reading rows: " + err.Error())

	}
	defer rows.Close()

	// var response []JsonResponse
	var emp []Emp

	for rows.Next() {
		var name, location string
		var id int
		err := rows.Scan(&id, &name, &location)

		// check errors
		checkErr(err)

		fmt.Printf("ID-newdata: %d, Name: %s, Location: %s\n", id, name, location)

		emp = append(emp, Emp{Id: id, Name: name, Location: location})

	}

	var response = JsonResponse{Type: "success", Data: emp}

	//	log.Fatal("Data is ", response)

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

// GetPartnerLeads read all employees
func GetPartnerLeadswithmap(db *sql.DB, id string) map[string]string {
	tsql := fmt.Sprintf("EXEC [TestSchema].[GetEmployee] @id = $1;")
	rows, _ := db.Query(tsql, id)

	cols, _ := rows.Columns()
	data := make(map[string]string)

	if rows.Next() {
		columns := make([]string, len(cols))
		columnPointers := make([]interface{}, len(cols))
		for i, _ := range columns {
			columnPointers[i] = &columns[i]
		}

		rows.Scan(columnPointers...)

		for i, colName := range cols {
			data[colName] = columns[i]
		}
	}

	log.Fatal("Data is ", data)

	return data

}

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

func GetPartnerLeadsC(db *sql.DB, pl *PartnerLeadsSchema) PartnerLeadsFinalResponse {

	// log.Fatal("agent id is ", pl.AgentID)
	tsql := fmt.Sprintf(" EXEC [business].[GetPartnerLeads_v1]   @FromDate= '" + pl.FromDate + "' , @ToDate='" + pl.ToDate + "' , @AgentID='" + pl.AgentID + "' , @ParentID='" + pl.ParentID + "' , @RMID='" + pl.RMID + "' ,@RegistrationNo='', @ProductGroupId='" + pl.ProductGroupId + "', @PageNum='" + pl.PageNum + "', @SortDir='" + pl.SortDir + "',  @PageSize='" + pl.PageSize + "'")

	rows, err := db.Query(tsql)

	if err != nil {
		fmt.Println("Error reading rows: " + err.Error())
	}

	defer rows.Close()

	// var response []JsonResponse
	var emp []PartnerLeadsResponse

	for rows.Next() {
		var ep PartnerLeadsResponse

		err := rows.Scan(&ep.RowNum, &ep.LeadID, &ep.Product, &ep.MobileNo, &ep.CustomerName, &ep.LeadCreationDate, &ep.BookingDate, &ep.SupplierName, &ep.PlanName, &ep.PremiumAmount, &ep.PaymentMode, &ep.CurrentStatus, &ep.PolicyLink, &ep.SelectionPlanID, &ep.SumInsured, &ep.ProductID, &ep.Utm_term, &ep.UTM_Medium, &ep.Utm_campaign, &ep.ODPremium, &ep.APE, &ep.AgentName, &ep.ExitPointURL, &ep.RegistrationNumber, &ep.InspectionStatusId, &ep.TotalCount)

		// check errors
		checkErr(err)

		//fmt.Printf("ID-newdata: %d, Name: %s, Location: %s\n", id, name, location)

		emp = append(emp, ep)

	}

	var response = PartnerLeadsFinalResponse{Type: "success", Data: emp}

	//	log.Fatal("Data is ", response)

	return response

}
