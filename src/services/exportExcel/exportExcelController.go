package exportExcel

import (
	"database/sql"
	"fmt"
	"strconv"
	"time"

	"github.com/xuri/excelize/v2"
)

func ReadNewData(db *sql.DB, id string) JsonResponse {
	start := time.Now()
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

		//fmt.Printf("ID-newdata: %d, Name: %s, Location: %s\n", ep.Id, ep.Name, ep.Location)

		emp = append(emp, ep)

	}

	var response = JsonResponse{Type: "success", Data: emp}

	//	log.Fatal("Data is ", response)

	CreateExcel(emp)

	timeElapsed := time.Since(start)
	fmt.Printf("The `for` loop took %s", timeElapsed)

	return response
}

// Function for handling errors
func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

// Function for handling messages
// func printMessage(message string) {
// 	fmt.Println("")
// 	fmt.Println(message)
// 	fmt.Println("")
// }

func CreateExcel(data []Emp) {

	// wg := &sync.WaitGroup{}
	// mu := &sync.RWMutex{}
	f := excelize.NewFile()
	Sheet1 := "Sheet1"
	index := f.NewSheet(Sheet1)

	ch := make(chan *excelize.File)

	for i := 0; i < len(data); i = i + 1000 {
		fmt.Println(i)
		// wg.Add(1)
		go SetRowValue(data[i:i+1000], f, ch, i)
		// go SetRowValue(data[i:i+10], f, ch, i)

	}
	// wg.Wait()

	//SetRowValue(data, f, ch)

	f.SetActiveSheet(index)

	f.SaveAs("/tmp/abcd/test.xlsx")

}

func SetRowValue(data []Emp, f *excelize.File, ch chan *excelize.File, index int) {
	// mu.Lock()
	// index := 0
	for _, v := range data {
		f.SetCellValue(f.GetSheetName(f.GetActiveSheetIndex()), "A"+strconv.Itoa(index), v.Id)
		f.SetCellValue(f.GetSheetName(f.GetActiveSheetIndex()), "B"+strconv.Itoa(index), v.Name.String)
		f.SetCellValue(f.GetSheetName(f.GetActiveSheetIndex()), "C"+strconv.Itoa(index), v.Location.String)
		// f.SetCellValue(f.GetSheetName(f.GetActiveSheetIndex()), "D"+strconv.Itoa(index), v.Password.String)
		// f.SetCellValue(f.GetSheetName(f.GetActiveSheetIndex()), "E"+strconv.Itoa(index), v.UserName.String)
		// f.SetCellValue(f.GetSheetName(f.GetActiveSheetIndex()), "F"+strconv.Itoa(index), v.IsActive.String)
		index++
	}
	// defer mu.Unlock()
	// defer wg.Done()

	//ch <- f
}
