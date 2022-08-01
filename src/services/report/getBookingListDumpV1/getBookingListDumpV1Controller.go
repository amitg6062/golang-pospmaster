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

	tsql := fmt.Sprint("exec [report].[GetBookingListDump_v1-amit] ")
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
	itr := 0

	for _, v := range data {
		// fmt.Println(v["LeadId"])
		//println(position, v, v.d("args").s("LeadId"))
		// println(position, v, v.d("args").s("LeadId"))
		f.SetCellValue(f.GetSheetName(f.GetActiveSheetIndex()), "A"+strconv.Itoa(itr), v["LeadId"])
		f.SetCellValue(f.GetSheetName(f.GetActiveSheetIndex()), "B"+strconv.Itoa(itr), v["InsuredName"])
		f.SetCellValue(f.GetSheetName(f.GetActiveSheetIndex()), "C"+strconv.Itoa(itr), v["ApplicationNo"])
		f.SetCellValue(f.GetSheetName(f.GetActiveSheetIndex()), "D"+strconv.Itoa(itr), v["LeadDate"])
		f.SetCellValue(f.GetSheetName(f.GetActiveSheetIndex()), "E"+strconv.Itoa(itr), v["BusinessType"])
		f.SetCellValue(f.GetSheetName(f.GetActiveSheetIndex()), "F"+strconv.Itoa(itr), v["Circle"])
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

/*
func SetRowValue(data []map[string]interface{}, f *excelize.File) {
	// mu.Lock()
	// defer wg.Done()
	// defer mu.Unlock()
	index := 0

	for _, v := range data {

		f.SetCellValue(f.GetSheetName(f.GetActiveSheetIndex()), "A"+strconv.Itoa(index), v.LeadId)
		f.SetCellValue(f.GetSheetName(f.GetActiveSheetIndex()), "B"+strconv.Itoa(index), v.InsuredName.String)
		f.SetCellValue(f.GetSheetName(f.GetActiveSheetIndex()), "C"+strconv.Itoa(index), v.Dob)
		// f.SetCellValue(f.GetSheetName(f.GetActiveSheetIndex()), "D"+strconv.Itoa(index), v.ApplicationNo)
		// f.SetCellValue(f.GetSheetName(f.GetActiveSheetIndex()), "E"+strconv.Itoa(index), v.BasicPremium)
		// f.SetCellValue(f.GetSheetName(f.GetActiveSheetIndex()), "F"+strconv.Itoa(index), v.IsActive.String)
		index++
	}
	// time.Sleep(time.Second)
}
*/

/*
func ReadData(db *sql.DB, requestParam RequestParam) []OutputData {
	//Handle panic condition
	defer hf.Deferring()
	start := time.Now()
	var emp []OutputData

	tsql := fmt.Sprint("exec [report].[GetBookingListDump_v1-amit] ")
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

	for rows.Next() {
		var ep OutputData
		//err := rows.Scan(&ep.LeadId, &ep.InsuredName, &ep.Dob, &ep.Insurer, &ep.InsurerFullName)



			err := rows.Scan(&ep.Ape, &ep.ActualLeadSource, &ep.Address, &ep.ApplicationNo, &ep.BasicPremium, &ep.BookingDate, &ep.BookingMode, &ep.BusinessType, &ep.ChatStatus, &ep.Circle, &ep.City, &ep.CubicCapacity, &ep.CustomerID, &ep.Dob, &ep.Discount, &ep.FuleType, &ep.Istp, &ep.InstallmentsPaid, &ep.InsuredName, &ep.Insurer, &ep.InsurerFullName, &ep.IsE2E, &ep.IssuanceRejDate, &ep.LeadDate, &ep.LeadId, &ep.LeadRank, &ep.MakeName, &ep.MaritalStatus, &ep.NetPremium, &ep.NoOfSeats, &ep.Noofwheels, &ep.ODPremium, &ep.PGType, &ep.ParentID, &ep.ParentLeadCreationDate, &ep.ParentLeadSource, &ep.PartnerID, &ep.PaymentPeriodicity, &ep.PaymentSubStatus, &ep.PersonalAccidentCover, &ep.PinCode, &ep.PlanName, &ep.PolicyNo, &ep.PolicyType, &ep.Premium, &ep.Product, &ep.RMCode, &ep.RMName, &ep.RegistrationDate, &ep.RegistrationNo, &ep.Source, &ep.State, &ep.Status, &ep.StpNstp, &ep.SumInsured, &ep.TPPremium, &ep.UtmMedium, &ep.UtmCampaign, &ep.UtmSource, &ep.UtmTerm, &ep.VechicleCarrier, &ep.VehicleAge, &ep.VehicleModelName, &ep.VehicleSubClass, &ep.Grossvehicleweight, &ep.KaliPili)


		//err := rows.Scan(ep...)

		if err != nil {
			panic(err)
		}

		//fmt.Printf("ID-newdata: %d, Name: %s, Location: %s\n", ep.Id, ep.Name, ep.Location)

		emp = append(emp, ep)

	}

	//var response = JsonResponse{Type: "success", Data: emp}

	//	log.Fatal("Data is ", response)

	CreateExcel(emp)

	timeElapsed := time.Since(start)
	fmt.Printf("The `for` loop took %s", timeElapsed)

	return emp
}
*/

// var mu sync.Mutex
