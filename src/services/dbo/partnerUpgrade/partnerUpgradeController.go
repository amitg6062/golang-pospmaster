package partnerUpgrade

import (
	"database/sql"
	"fmt"
	"log"

	hf "github.com/amitg6062/golang-posp-helpers"
)

func ReadData(db *sql.DB, requestParam RequestParam) JsonResponse {
	//Handle panic condition
	defer hf.Deferring()

	tsql := fmt.Sprint("EXEC [dbo].[UpgradeUserRoles] ")
	tsql = fmt.Sprint(tsql, "@OldAffiliateCode = '", requestParam.OldAffiliateCode, "', ")
	tsql = fmt.Sprint(tsql, "@NewAffiliateCode = '", requestParam.NewAffiliateCode, "', ")
	tsql = fmt.Sprint(tsql, "@Type = '", requestParam.Type, "', ")
	tsql = fmt.Sprint(tsql, "@ParentCode = '", requestParam.ParentCode, "'; ")

	fmt.Println(tsql)

	rows, err := db.Query(tsql)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	//Scan Rows and Get Response in interface.
	ret := make([]map[string]interface{}, 0)
	ret = hf.GetDBResponse(rows, ret)

	//Call Another Request.
	if requestParam.Type == 2 || requestParam.Type == 5 {
		moveParentCode(db, requestParam.NewAffiliateCode)
	}

	var response = JsonResponse{Error: false, Data: ret}

	return response

}

func moveParentCode(db *sql.DB, parencode string) {
	//Handle panic condition
	defer hf.Deferring()

	tsql := fmt.Sprint("[dbo].[UpgradeInnerHirearchyPartners] ")
	tsql = fmt.Sprint(tsql, "'", parencode, "'; ")

	fmt.Println(tsql)

	rows, err := db.Query(tsql)
	if err != nil {
		hf.CheckErr(err)
	}
	defer rows.Close()
}
