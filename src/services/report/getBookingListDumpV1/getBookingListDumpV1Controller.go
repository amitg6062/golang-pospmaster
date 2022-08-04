package getBookingListDumpV1

import (
	"database/sql"
	"fmt"
	"strconv"
	"time"

	hf "github.com/amitg6062/golang-posp-helpers"

	excelize "github.com/xuri/excelize/v2"
)

type dict map[string]interface{}

func (d dict) d(k string) dict {
	return d[k].(map[string]interface{})
}

func (d dict) s(k string) string {
	return d[k].(string)
}

func ReadData(db *sql.DB, requestParam RequestParam) JsonResponse {
	//Handle panic condition
	defer hf.Deferring()
	start := time.Now()

	tsql := fmt.Sprint("exec [report].[GetBookingListDump_v1] ")
	tsql = fmt.Sprint(tsql, "@ToDate = '", requestParam.ToDate, "', ")
	tsql = fmt.Sprint(tsql, "@FromDate = '", requestParam.FromDate, "', ")
	tsql = fmt.Sprint(tsql, "@ProductId = '", requestParam.ProductId, "', ")

	if requestParam.ProductId != "" {
		tsql = fmt.Sprint(tsql, "@ParentId = ", requestParam.ParentId, ", ")
	}

	if requestParam.ProductId != "" {
		tsql = fmt.Sprint(tsql, "@PartnerCode = '", requestParam.PartnerCode, "', ")
	}

	tsql = fmt.Sprint(tsql, "@PageNum = ", requestParam.PageNum, ", ")
	tsql = fmt.Sprint(tsql, "@PageSize = ", requestParam.PageSize, "; ")

	fmt.Println(tsql)

	rows, err := db.Query(tsql)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	//Scan Rows and Get Response in interface.
	var ret []dict
	ret = GetDBResponse(rows, ret)

	var response = JsonResponse{Error: false, Data: ret}

	CreateExcel(ret)

	timeElapsed := time.Since(start)
	fmt.Printf("The `for` loop took %s", timeElapsed)
	return response

}

func GetDBResponse(rows *sql.Rows, ret []dict) []dict {
	cols, _ := rows.Columns()

	for rows.Next() {
		colVals := make([]interface{}, len(cols))
		for i := range colVals {
			colVals[i] = new(interface{})
		}
		err := rows.Scan(colVals...)
		if err != nil {
			hf.CheckErr(err)
		}
		colNames, err := rows.Columns()
		if err != nil {
			hf.CheckErr(err)
		}
		these := make(map[string]interface{})
		for idx, name := range colNames {
			these[name] = *colVals[idx].(*interface{})
		}
		ret = append(ret, these)
	}
	if err := rows.Err(); err != nil {
		hf.CheckErr(err)
	}

	return ret

}

func CreateExcel(data []dict) {

	f := excelize.NewFile()
	Sheet1 := "Sheet1"
	index := f.NewSheet(Sheet1)

	excelCellName := [...]string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z", "AA", "AB", "AC", "AD", "AE", "AF", "AG", "AH", "AI", "AJ", "AK", "AL", "AM", "AN", "AO", "AP", "AQ", "AR", "AS", "AT", "AU", "AV", "AW", "AX", "AY", "AZ", "BA", "BB", "BC", "BD", "BE", "BF", "BG", "BH", "BI", "BJ", "BK", "BL", "BM", "BN", "BO", "BP", "BQ", "BR", "BS", "BT", "BU", "BV", "BW", "BX", "BY", "BZ"}

	headerColumnName := [...]string{"BookingDate", "LeadId", "InsuredName", "DOB", "Insurer", "InsurerFullName", "Product", "PlanName", "SumInsured", "BasicPremium", "NetPremium", "Premium", "ODPremium", "APE", "Status", "City", "ApplicationNo", "PolicyNo", "PolicyType", "PaymentPeriodicity", "IsE2E", "PGType", "CustomerId", "Address", "State", "PinCode", "ActualLeadSource", "Utm_source", "Utm_term", "Utm_Medium", "Utm_campaign", "RMCode", "RMName", "Circle", "LeadRank", "ParentId", "ParentLeadCreationDate", "ParentLeadSource", "MaritalStatus", "LeadDate", "ChatStatus", "Issuance/Rej Date", "PaymentSubStatus", "InstallmentsPaid", "Source", "PartnerId", "VehicleModelName", "RegistrationNo", "RegistrationDate", "ISTP", "FuleType", "grossvehicleweight", "MakeName", "VehicleSubClass", "VehicleAge", "BookingMode", "VechicleCarrier", "Noofwheels", "BusinessType", "CubicCapacity", "kaliPili", "PersonalAccidentCover", "StpNstp", "NoOfSeats", "Discount", "TPPremium"}

	itr := 0

	for _, v := range data {
		for key, columnVal := range headerColumnName {
			f.SetCellValue(f.GetSheetName(f.GetActiveSheetIndex()), excelCellName[key]+strconv.Itoa(itr), v[columnVal])
		}

		itr++
	}

	f.SetActiveSheet(index)

	f.SaveAs("/tmp/abcd/test.xlsx")

}
