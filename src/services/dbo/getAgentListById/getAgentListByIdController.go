package getAgentListById

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

/*

func ReadData(db *sql.DB, agentList AgentListRequest) JsonResponse {
	//Handle panic condition
	defer deferring()

	var emp []AgentListResponse

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
		fmt.Println("Error reading rows: " + err.Error())

	}
	defer rows.Close()

	for rows.Next() {
		var ep AgentListResponse
		err := rows.Scan(&ep.TotalCount, &ep.RowNum, &ep.Type, &ep.AffiliateID, &ep.RMCode, &ep.ParentCode, &ep.AffiliateCode, &ep.Name, &ep.MobileNo, &ep.EmailId, &ep.BankName, &ep.BankBranchAddress1, &ep.BankBranchAddress2, &ep.BankBranchStateId, &ep.BankBranchCityId, &ep.BankBranchPinCode, &ep.BankAccountHolderName, &ep.BankAccountNumber, &ep.NEFTIFSCCode, &ep.RTGSIFSCCode, &ep.UpdatedOn, &ep.CreatedOn, &ep.Status, &ep.IsVerified, &ep.AffiliateStatusId, &ep.AffiliateStatus, &ep.BranchName, &ep.EmployeeCode, &ep.IsPanVerified, &ep.ReviewerName, &ep.IsCertified, &ep.IsGreenChannel)

		checkErr(err)

		emp = append(emp, ep)

	}

	var response = JsonResponse{Error: false, Data: emp}

	//	log.Fatal("Data is ", response)

	//using map starts

	//using map ends

	return response

}

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

*/

func deferring() {
	if err := recover(); err != nil {
		fmt.Println("An error occurred:", err)
	}
}
