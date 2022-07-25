package insertUpdateBmsLeadDetails

import (
	"database/sql"
	"fmt"

	hf "github.com/amitg6062/golang-posp-helpers"
)

func ReadData(db *sql.DB, requestParam RequestParam) JsonResponse {
	//Handle panic condition
	defer hf.Deferring()

	tsql := fmt.Sprint("EXEC [dbo].[InsertUpdateAffiliateLeadDetails] ")

	tsql = fmt.Sprint(tsql, "@LeadID = '", requestParam.LeadID, "', ")
	tsql = fmt.Sprint(tsql, "@SessionID = '", requestParam.SessionID, "', ")
	tsql = fmt.Sprint(tsql, "@Name = '", requestParam.Name, "', ")
	tsql = fmt.Sprint(tsql, "@Gender = '", requestParam.Gender, "', ")
	tsql = fmt.Sprint(tsql, "@MobileNo = '", requestParam.MobileNo, "', ")
	tsql = fmt.Sprint(tsql, "@AltPhoneNo = '", requestParam.AltPhoneNo, "', ")
	tsql = fmt.Sprint(tsql, "@EmailID = '", requestParam.EmailID, "', ")
	tsql = fmt.Sprint(tsql, "@Address = '", requestParam.Address, "', ")
	tsql = fmt.Sprint(tsql, "@CityID = '", requestParam.CityID, "', ")
	tsql = fmt.Sprint(tsql, "@StateID = '", requestParam.StateID, "', ")
	tsql = fmt.Sprint(tsql, "@PostCode = '", requestParam.PostCode, "', ")
	tsql = fmt.Sprint(tsql, "@Country = '", requestParam.Country, "', ")
	tsql = fmt.Sprint(tsql, "@MaritalStatus = '", requestParam.MaritalStatus, "', ")
	tsql = fmt.Sprint(tsql, "@AnnualIncome = '", requestParam.AnnualIncome, "', ")
	tsql = fmt.Sprint(tsql, "@HighestBookingStep = '", requestParam.HighestBookingStep, "', ")
	tsql = fmt.Sprint(tsql, "@ReferralId = '", requestParam.ReferralId, "', ")
	tsql = fmt.Sprint(tsql, "@ExitPointURL = '", requestParam.ExitPointURL, "', ")
	tsql = fmt.Sprint(tsql, "@Utm_source = '", requestParam.Utm_source, "', ")
	tsql = fmt.Sprint(tsql, "@UTM_Medium = '", requestParam.UTM_Medium, "', ")
	tsql = fmt.Sprint(tsql, "@Utm_term = '", requestParam.Utm_term, "', ")
	tsql = fmt.Sprint(tsql, "@Utm_campaign = '", requestParam.Utm_campaign, "', ")
	tsql = fmt.Sprint(tsql, "@ProductID = '", requestParam.ProductID, "', ")
	tsql = fmt.Sprint(tsql, "@CustomerID = '", requestParam.CustomerID, "', ")
	tsql = fmt.Sprint(tsql, "@HasAddon = '", requestParam.HasAddon, "', ")
	tsql = fmt.Sprint(tsql, "@DateOfBirth = '", requestParam.DateOfBirth, "', ")
	tsql = fmt.Sprint(tsql, "@SupplierId = '", requestParam.SupplierId, "', ")
	tsql = fmt.Sprint(tsql, "@PlanId = '", requestParam.PlanID, "', ")
	tsql = fmt.Sprint(tsql, "@SupplierName = '", requestParam.SupplierName, "', ")
	tsql = fmt.Sprint(tsql, "@PlanName = '", requestParam.PlanName, "', ")
	tsql = fmt.Sprint(tsql, "@EnquiryId = '", requestParam.EnquiryId, "', ")
	tsql = fmt.Sprint(tsql, "@LeadSource = '", requestParam.LeadSource, "', ")
	tsql = fmt.Sprint(tsql, "@Source = 'Online-CJ', ")
	tsql = fmt.Sprint(tsql, "@PreviousPolicyExpiryDate = '", requestParam.PreviousPolicyExpiryDate, "'; ")

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
