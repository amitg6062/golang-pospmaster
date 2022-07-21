package insertUpdateBmsBookingDetails

import (
	"database/sql"
	"fmt"
	"log"
)

func InsertUpdateLeadProductDetails(db *sql.DB, requestParam RequestParam) JsonResponse {
	//Handle panic condition
	defer deferring()

	if requestParam.PaymentSTATUS != 300 && requestParam.PaymentSTATUS != 4002 {
		requestParam.OfferCreatedON = ""
	}

	tsql := fmt.Sprint("EXEC exec [dbo].[InsertUpdateAffiliateBookingDetail] ")

	tsql = fmt.Sprint(tsql, "@LeadID = '", requestParam.LeadID, "', ")
	tsql = fmt.Sprint(tsql, "@ProductID = '", requestParam.ProductID, "', ")
	tsql = fmt.Sprint(tsql, "@ProductName = '", requestParam.ProductName, "', ")
	tsql = fmt.Sprint(tsql, "@SelectedPlanName = '", requestParam.SelectedPlanName, "', ")
	tsql = fmt.Sprint(tsql, "@BasicPremium = '", requestParam.BasicPremium, "', ")
	tsql = fmt.Sprint(tsql, "@TotalPremium = '", requestParam.TotalPremium, "', ")
	tsql = fmt.Sprint(tsql, "@OfferCreatedON = '", requestParam.OfferCreatedON, "', ")
	tsql = fmt.Sprint(tsql, "@ServiceTax = '", requestParam.ServiceTax, "', ")
	tsql = fmt.Sprint(tsql, "@PaymentSTATUS = '", requestParam.PaymentSTATUS, "', ")
	tsql = fmt.Sprint(tsql, "@PolicyTypeName = '", requestParam.PolicyTypeName, "', ")
	tsql = fmt.Sprint(tsql, "@SupplierId = '", requestParam.SupplierId, "', ")
	tsql = fmt.Sprint(tsql, "@SupplierName = '", requestParam.SupplierName, "', ")
	tsql = fmt.Sprint(tsql, "@TransRefNo = '", requestParam.TransRefNo, "', ")
	tsql = fmt.Sprint(tsql, "@PaymentPeriodicity = '", requestParam.PaymentPeriodicity, "', ")
	tsql = fmt.Sprint(tsql, "@BookingType = '", requestParam.BookingType, "', ")
	tsql = fmt.Sprint(tsql, "@IsEMI = '", requestParam.IsEMI, "', ")
	tsql = fmt.Sprint(tsql, "@PolicyNo = '", requestParam.PolicyNo, "', ")
	tsql = fmt.Sprint(tsql, "@ApplicationNo = '", requestParam.ApplicationNo, "', ")
	tsql = fmt.Sprint(tsql, "@PlanType = '", requestParam.PlanType, "', ")
	tsql = fmt.Sprint(tsql, "@IsSTP = '", requestParam.IsSTP, "', ")
	tsql = fmt.Sprint(tsql, "@IsBooked = '", requestParam.IsBooked, "', ")
	tsql = fmt.Sprint(tsql, "@PolicyTenure = '", requestParam.PolicyTerm, "', ")
	tsql = fmt.Sprint(tsql, "@IssuanceDate = '", requestParam.IssuanceDate, "', ")
	tsql = fmt.Sprint(tsql, "@PolicyStartDate = '", requestParam.PolicyStartDateDt, "', ")
	tsql = fmt.Sprint(tsql, "@PolicyEndDate = '", requestParam.PolicyStartDateDt, "', ")
	tsql = fmt.Sprint(tsql, "@PolicyPdfUrl = '", requestParam.PolicyPdfUrl, "', ")
	tsql = fmt.Sprint(tsql, "@PaymentSubStatus = '", requestParam.PaymentSubStatus, "', ")
	tsql = fmt.Sprint(tsql, "@BookingSource = '", requestParam.BookingSource, "', ")
	// tsql = fmt.Sprint(tsql, "@PaymentDate = '", requestParam.PaymentDate, "', ")
	tsql = fmt.Sprint(tsql, "@PaymentDate = '2022-01-01', ")
	tsql = fmt.Sprint(tsql, "@PolicyTenureType = '", requestParam.PolicyTenureType, "'; ")

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

func insertUpdateVehicleDetails_v2(db *sql.DB, requestParam RequestParam) JsonResponse {
	//Handle panic condition
	defer deferring()

	tsql := fmt.Sprint("EXEC [dbo].[InsertUpdateVehicleDetails] ")

	tsql = fmt.Sprint(tsql, "@LeadID = '", requestParam.LeadID, "', ")
	tsql = fmt.Sprint(tsql, "@SubProduct = '", requestParam.SubProduct, "', ")
	tsql = fmt.Sprint(tsql, "@PlanType = '", requestParam.PlanType, "', ")
	tsql = fmt.Sprint(tsql, "@BasicPremium = '", requestParam.BasicPremium, "', ")
	tsql = fmt.Sprint(tsql, "@FinalPremium = '", requestParam.FinalPremium, "', ")
	tsql = fmt.Sprint(tsql, "@BasicOdPremium = '", requestParam.BasicOdPremium, "', ")
	tsql = fmt.Sprint(tsql, "@TotalOdPremium = '", requestParam.TotalOdPremium, "', ")
	tsql = fmt.Sprint(tsql, "@RTO = '", requestParam.RTO, "', ")
	tsql = fmt.Sprint(tsql, "@Zone = '", requestParam.Zone, "', ")
	tsql = fmt.Sprint(tsql, "@CPA = '", requestParam.CPA, "', ")
	tsql = fmt.Sprint(tsql, "@PACoverPremium = '", requestParam.PACoverPremium, "', ")
	tsql = fmt.Sprint(tsql, "@Discount = '", requestParam.Discount, "', ")
	tsql = fmt.Sprint(tsql, "@IDV = '", requestParam.IDV, "', ")
	tsql = fmt.Sprint(tsql, "@RegistrationNumber = '", requestParam.RegistrationNumber, "', ")
	tsql = fmt.Sprint(tsql, "@NumberOfWheels = '", requestParam.NumberOfWheels, "', ")
	tsql = fmt.Sprint(tsql, "@GVW = '", requestParam.GVW, "', ")
	tsql = fmt.Sprint(tsql, "@CubicCapacity = '", requestParam.CubicCapacity, "', ")
	tsql = fmt.Sprint(tsql, "@CarryingCapacity = '", requestParam.CarryingCapacity, "', ")
	tsql = fmt.Sprint(tsql, "@SeatingCapacity = '", requestParam.SeatingCapacity, "', ")
	tsql = fmt.Sprint(tsql, "@CarryingCapacity = '", requestParam.CarryingCapacity, "', ")
	tsql = fmt.Sprint(tsql, "@RegistrationDate = '", requestParam.RegistrationDate, "', ")
	tsql = fmt.Sprint(tsql, "@ManufacturingDate = '", requestParam.ManufacturingDate, "', ")
	tsql = fmt.Sprint(tsql, "@KaliPili = '", requestParam.KaliPili, "', ")
	tsql = fmt.Sprint(tsql, "@VehicleState = '", requestParam.VehicleState, "', ")
	tsql = fmt.Sprint(tsql, "@VehicleCity = '", requestParam.VehicleCity, "', ")
	tsql = fmt.Sprint(tsql, "@FuelType = '", requestParam.FuelType, "', ")
	tsql = fmt.Sprint(tsql, "@VehicleSubClass = '", requestParam.VehicleSubClass, "', ")
	tsql = fmt.Sprint(tsql, "@VehicleAge = '", requestParam.VehicleAge, "', ")
	tsql = fmt.Sprint(tsql, "@VehicleMake = '", requestParam.VehicleMake, "', ")
	tsql = fmt.Sprint(tsql, "@VehicleModel = '", requestParam.VehicleModel, "', ")
	tsql = fmt.Sprint(tsql, "@VariantName = '", requestParam.VariantName, "', ")
	tsql = fmt.Sprint(tsql, "@VariantID = '", requestParam.VariantID, "', ")
	tsql = fmt.Sprint(tsql, "@VehicleCarrier = '", requestParam.VehicleCarrier, "', ")
	tsql = fmt.Sprint(tsql, "@BusinessType = '", requestParam.BusinessType, "', ")
	tsql = fmt.Sprint(tsql, "@NCB = '", requestParam.NCB, "', ")
	tsql = fmt.Sprint(tsql, "@BasicTPPremium = '", requestParam.BasicTPPremium, "', ")
	tsql = fmt.Sprint(tsql, "@TotalTPPremium = '", requestParam.TotalTPPremium, "', ")
	tsql = fmt.Sprint(tsql, "@ZDTaken = '", requestParam.ZDTaken, "', ")
	tsql = fmt.Sprint(tsql, "@IsInspectionFirstCase = '", requestParam.IsInspectionFirstCase, "', ")
	tsql = fmt.Sprint(tsql, "@IsPortability = '", requestParam.IsPortability, "', ")
	tsql = fmt.Sprint(tsql, "@AgeOfEldestMember = '", requestParam.AgeOfEldestMember, "', ")
	tsql = fmt.Sprint(tsql, "@NoOfAdults = '", requestParam.NoOfAdults, "', ")
	tsql = fmt.Sprint(tsql, "@NoOfChildren = '", requestParam.NoOfChildren, "', ")
	tsql = fmt.Sprint(tsql, "@ProfileMemberCombo = '", requestParam.FuelType, "', ")
	tsql = fmt.Sprint(tsql, "@NatureOfBusiness = '", requestParam.NatureOfBusiness, "', ")
	tsql = fmt.Sprint(tsql, "@NoOfEmployees = '", requestParam.NoOfEmployees, "', ")
	tsql = fmt.Sprint(tsql, "@NoOfLives = '", requestParam.NoOfLives, "', ")
	tsql = fmt.Sprint(tsql, "@Structure = '", requestParam.Structure, "', ")
	tsql = fmt.Sprint(tsql, "@Parental = '", requestParam.Parental, "', ")
	tsql = fmt.Sprint(tsql, "@SumInsuredType = '", requestParam.SumInsuredType, "', ")
	tsql = fmt.Sprint(tsql, "@ZDPremium = '", requestParam.ZDPremium, "', ")
	tsql = fmt.Sprint(tsql, "@PayTerm = '", requestParam.PayTerm, "', ")
	tsql = fmt.Sprint(tsql, "@PlanVariant = '", requestParam.PlanVariant, "'; ")

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
func CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}

// Function for handling messages
func PrintMessage(message string) {
	fmt.Println("")
	fmt.Println(message)
	fmt.Println("")
}

func deferring() {
	if err := recover(); err != nil {
		fmt.Println("An error occurred:", err)
	}
}

func GetLeadDetailsByIdAndProductId_v1(db *sql.DB, LeadID int64) JsonResponse {
	//Handle panic condition
	defer deferring()

	tsql := fmt.Sprint("[dbo].[GetLeadDetailsByIdAndProductId_v1] ")
	tsql = fmt.Sprint(tsql, "@LeadID = '", LeadID, "'; ")

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

func getSupplierPlanDetails(ProductId, SupplierId, PlanId int64) {

}
