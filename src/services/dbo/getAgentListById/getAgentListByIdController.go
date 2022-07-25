package getAgentListById

import (
	"database/sql"
	"fmt"

	hf "github.com/amitg6062/golang-posp-helpers"
)

func ReadData(db *sql.DB, agentList AgentListRequest) JsonResponse {
	//Handle panic condition
	defer hf.Deferring()

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
		hf.CheckErr(err)
	}
	defer rows.Close()

	//Scan Rows and Get Response in interface.
	ret := make([]map[string]interface{}, 0)
	ret = hf.GetDBResponse(rows, ret)

	var response = JsonResponse{Error: false, Data: ret}

	return response

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
*/
