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

	// mu := &sync.RWMutex{}
	//var wg sync.WaitGroup
	f := excelize.NewFile()
	Sheet1 := "Sheet1"
	index := f.NewSheet(Sheet1)

	//ch := make(chan *excelize.File)
	//SetRowValue(data, f)

	// excelCellName := [...]string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z", "AA", "AB", "AC", "AD", "AE", "AF", "AG", "AH", "AI", "AJ", "AK", "AL", "AM", "AN", "AO", "AP", "AQ", "AR", "AS", "AT", "AU", "AV", "AW", "AX", "AY", "AZ", "BA", "BB", "BC", "BD", "BE", "BF", "BG", "BH", "BI", "BJ", "BK", "BL", "BM", "BN", "BO", "BP", "BQ", "BR", "BS", "BT", "BU", "BV", "BW", "BX", "BY", "BZ"}

	//log.Fatal("data is ", excelCellName)

	itr := 0

	for _, v := range data {
		// fmt.Println(v["LeadId"])
		//println(position, v, v.d("args").s("LeadId"))
		// println(position, v, v.d("args").s("LeadId"))
		f.SetCellValue(f.GetSheetName(f.GetActiveSheetIndex()), "A"+strconv.Itoa(itr), v["BookingDate"])
		f.SetCellValue(f.GetSheetName(f.GetActiveSheetIndex()), "B"+strconv.Itoa(itr), v["LeadId"])
		f.SetCellValue(f.GetSheetName(f.GetActiveSheetIndex()), "C"+strconv.Itoa(itr), v["InsuredName"])
		f.SetCellValue(f.GetSheetName(f.GetActiveSheetIndex()), "D"+strconv.Itoa(itr), v["DOB"])
		f.SetCellValue(f.GetSheetName(f.GetActiveSheetIndex()), "E"+strconv.Itoa(itr), v["Insurer"])
		f.SetCellValue(f.GetSheetName(f.GetActiveSheetIndex()), "F"+strconv.Itoa(itr), v["InsurerFullName"])
		f.SetCellValue(f.GetSheetName(f.GetActiveSheetIndex()), "G"+strconv.Itoa(itr), v["Product"])
		f.SetCellValue(f.GetSheetName(f.GetActiveSheetIndex()), "H"+strconv.Itoa(itr), v["PlanName"])
		f.SetCellValue(f.GetSheetName(f.GetActiveSheetIndex()), "I"+strconv.Itoa(itr), v["SumInsured"])
		f.SetCellValue(f.GetSheetName(f.GetActiveSheetIndex()), "J"+strconv.Itoa(itr), v["BasicPremium"])
		f.SetCellValue(f.GetSheetName(f.GetActiveSheetIndex()), "K"+strconv.Itoa(itr), v["NetPremium"])
		f.SetCellValue(f.GetSheetName(f.GetActiveSheetIndex()), "L"+strconv.Itoa(itr), v["Premium"])
		f.SetCellValue(f.GetSheetName(f.GetActiveSheetIndex()), "M"+strconv.Itoa(itr), v["ODPremium"])
		f.SetCellValue(f.GetSheetName(f.GetActiveSheetIndex()), "N"+strconv.Itoa(itr), v["APE"])
		f.SetCellValue(f.GetSheetName(f.GetActiveSheetIndex()), "O"+strconv.Itoa(itr), v["Status"])
		f.SetCellValue(f.GetSheetName(f.GetActiveSheetIndex()), "P"+strconv.Itoa(itr), v["City"])
		f.SetCellValue(f.GetSheetName(f.GetActiveSheetIndex()), "Q"+strconv.Itoa(itr), v["ApplicationNo"])
		f.SetCellValue(f.GetSheetName(f.GetActiveSheetIndex()), "R"+strconv.Itoa(itr), v["PolicyNo"])
		f.SetCellValue(f.GetSheetName(f.GetActiveSheetIndex()), "S"+strconv.Itoa(itr), v["PolicyType"])
		f.SetCellValue(f.GetSheetName(f.GetActiveSheetIndex()), "T"+strconv.Itoa(itr), v["PaymentPeriodicity"])
		f.SetCellValue(f.GetSheetName(f.GetActiveSheetIndex()), "U"+strconv.Itoa(itr), v["IsE2E"])
		f.SetCellValue(f.GetSheetName(f.GetActiveSheetIndex()), "V"+strconv.Itoa(itr), v["PGType"])
		f.SetCellValue(f.GetSheetName(f.GetActiveSheetIndex()), "W"+strconv.Itoa(itr), v["CustomerId"])
		f.SetCellValue(f.GetSheetName(f.GetActiveSheetIndex()), "X"+strconv.Itoa(itr), v["Address"])
		f.SetCellValue(f.GetSheetName(f.GetActiveSheetIndex()), "Y"+strconv.Itoa(itr), v["State"])
		f.SetCellValue(f.GetSheetName(f.GetActiveSheetIndex()), "Z"+strconv.Itoa(itr), v["PinCode"])
		f.SetCellValue(f.GetSheetName(f.GetActiveSheetIndex()), "AA"+strconv.Itoa(itr), v["ActualLeadSource"])
		f.SetCellValue(f.GetSheetName(f.GetActiveSheetIndex()), "AB"+strconv.Itoa(itr), v["Utm_source"])
		f.SetCellValue(f.GetSheetName(f.GetActiveSheetIndex()), "AC"+strconv.Itoa(itr), v["Utm_term"])
		f.SetCellValue(f.GetSheetName(f.GetActiveSheetIndex()), "AD"+strconv.Itoa(itr), v["Utm_Medium"])
		f.SetCellValue(f.GetSheetName(f.GetActiveSheetIndex()), "AE"+strconv.Itoa(itr), v["Utm_campaign"])
		f.SetCellValue(f.GetSheetName(f.GetActiveSheetIndex()), "AF"+strconv.Itoa(itr), v["RMCode"])
		f.SetCellValue(f.GetSheetName(f.GetActiveSheetIndex()), "AG"+strconv.Itoa(itr), v["RMName"])
		f.SetCellValue(f.GetSheetName(f.GetActiveSheetIndex()), "AH"+strconv.Itoa(itr), v["Circle"])
		f.SetCellValue(f.GetSheetName(f.GetActiveSheetIndex()), "AI"+strconv.Itoa(itr), v["LeadRank"])
		f.SetCellValue(f.GetSheetName(f.GetActiveSheetIndex()), "AJ"+strconv.Itoa(itr), v["ParentId"])
		f.SetCellValue(f.GetSheetName(f.GetActiveSheetIndex()), "AK"+strconv.Itoa(itr), v["ParentLeadCreationDate"])
		f.SetCellValue(f.GetSheetName(f.GetActiveSheetIndex()), "AL"+strconv.Itoa(itr), v["ParentLeadSource"])
		f.SetCellValue(f.GetSheetName(f.GetActiveSheetIndex()), "AM"+strconv.Itoa(itr), v["MaritalStatus"])
		f.SetCellValue(f.GetSheetName(f.GetActiveSheetIndex()), "AN"+strconv.Itoa(itr), v["LeadDate"])
		f.SetCellValue(f.GetSheetName(f.GetActiveSheetIndex()), "AO"+strconv.Itoa(itr), v["ChatStatus"])
		f.SetCellValue(f.GetSheetName(f.GetActiveSheetIndex()), "AP"+strconv.Itoa(itr), v["Issuance/Rej Date"])
		f.SetCellValue(f.GetSheetName(f.GetActiveSheetIndex()), "AQ"+strconv.Itoa(itr), v["PaymentSubStatus"])
		f.SetCellValue(f.GetSheetName(f.GetActiveSheetIndex()), "AR"+strconv.Itoa(itr), v["InstallmentsPaid"])
		f.SetCellValue(f.GetSheetName(f.GetActiveSheetIndex()), "AS"+strconv.Itoa(itr), v["Source"])
		f.SetCellValue(f.GetSheetName(f.GetActiveSheetIndex()), "AT"+strconv.Itoa(itr), v["PartnerId"])
		f.SetCellValue(f.GetSheetName(f.GetActiveSheetIndex()), "AU"+strconv.Itoa(itr), v["VehicleModelName"])
		f.SetCellValue(f.GetSheetName(f.GetActiveSheetIndex()), "AV"+strconv.Itoa(itr), v["RegistrationNo"])
		f.SetCellValue(f.GetSheetName(f.GetActiveSheetIndex()), "AW"+strconv.Itoa(itr), v["RegistrationDate"])
		f.SetCellValue(f.GetSheetName(f.GetActiveSheetIndex()), "AX"+strconv.Itoa(itr), v["ISTP"])
		f.SetCellValue(f.GetSheetName(f.GetActiveSheetIndex()), "AY"+strconv.Itoa(itr), v["FuleType"])
		f.SetCellValue(f.GetSheetName(f.GetActiveSheetIndex()), "AZ"+strconv.Itoa(itr), v["grossvehicleweight"])
		f.SetCellValue(f.GetSheetName(f.GetActiveSheetIndex()), "BA"+strconv.Itoa(itr), v["MakeName"])
		f.SetCellValue(f.GetSheetName(f.GetActiveSheetIndex()), "BB"+strconv.Itoa(itr), v["VehicleSubClass"])
		f.SetCellValue(f.GetSheetName(f.GetActiveSheetIndex()), "BC"+strconv.Itoa(itr), v["VehicleAge"])
		f.SetCellValue(f.GetSheetName(f.GetActiveSheetIndex()), "BD"+strconv.Itoa(itr), v["BookingMode"])
		f.SetCellValue(f.GetSheetName(f.GetActiveSheetIndex()), "BE"+strconv.Itoa(itr), v["VechicleCarrier"])
		f.SetCellValue(f.GetSheetName(f.GetActiveSheetIndex()), "BF"+strconv.Itoa(itr), v["Noofwheels"])
		f.SetCellValue(f.GetSheetName(f.GetActiveSheetIndex()), "BG"+strconv.Itoa(itr), v["BusinessType"])
		f.SetCellValue(f.GetSheetName(f.GetActiveSheetIndex()), "BH"+strconv.Itoa(itr), v["CubicCapacity"])
		f.SetCellValue(f.GetSheetName(f.GetActiveSheetIndex()), "BI"+strconv.Itoa(itr), v["kaliPili"])
		f.SetCellValue(f.GetSheetName(f.GetActiveSheetIndex()), "BJ"+strconv.Itoa(itr), v["PersonalAccidentCover"])
		f.SetCellValue(f.GetSheetName(f.GetActiveSheetIndex()), "BK"+strconv.Itoa(itr), v["StpNstp"])
		f.SetCellValue(f.GetSheetName(f.GetActiveSheetIndex()), "BL"+strconv.Itoa(itr), v["NoOfSeats"])
		f.SetCellValue(f.GetSheetName(f.GetActiveSheetIndex()), "BM"+strconv.Itoa(itr), v["Discount"])
		f.SetCellValue(f.GetSheetName(f.GetActiveSheetIndex()), "BN"+strconv.Itoa(itr), v["TPPremium"])
		itr++
	}
	/*
		for i := 0; i < len(data); i = i + 1 {
			fmt.Println(i)
			//wg.Add(1)

			SetRowValue(&data[i:i+1], f, i)
			// SetRowValue(data[i:i+1], f, ch, wg, i)
			// go SetRowValue(data[i:i+10], f, ch, i)

		}*/

	//wg.Wait()

	//SetRowValue(data, f, ch)

	f.SetActiveSheet(index)

	f.SaveAs("/tmp/abcd/test.xlsx")

}
