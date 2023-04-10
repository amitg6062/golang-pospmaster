package dbo

import (
	"fmt"

	"posp_api_go_v2/src/helpers"

	lib "github.com/amitg6062/golang-posp-dbconnection"
	"github.com/gin-gonic/gin"
)

type model_getCityByStateId struct {
	StateId uint16 `json:"StateId" binding:"required"`
}

// @Summary InsertUpdateAffiliateLeadDetails
// @Description InsertUpdateAffiliateLeadDetails
// Inserts or Updates affiliate lead details. This API is used to insert or update the affiliate lead details in the database. The API accepts a JSON payload containing the required parameters.
// @ID InsertUpdateAffiliateLeadDetails
// @Tags dbo
// @Param LeadID body int true "Lead ID"
// @Param SessionID body int true "Session ID"
// @Param Name body string true "Name of the lead"
// @Param PosType body int true "Position Type"
// @Param Gender body string false "Gender of the lead"
// @Param MobileNo body string false "Mobile number of the lead"
// @Param AltPhoneNo body string false "Alternate phone number of the lead"
// @Param EmailID body string false "Email ID of the lead"
// @Param Address body string false "Address of the lead"
// @Param CityID body int false "City ID"
// @Param StateID body int false "State ID"
// @Param PostCode body string false "Postal code"
// @Param Country body string false "Country"
// @Param MaritalStatus body string false "Marital status"
// @Param AnnualIncome body int false "Annual income of the lead"
// @Param HighestBookingStep body int false "Highest booking step"
// @Param ReferralId body int false "Referral ID"
// @Param ExitPointURL body string false "Exit point URL"
// @Param Utm_source body string false "UTM source"
// @Param UTM_Medium body string false "UTM medium"
// @Param Utm_term body string false "UTM term"
// @Param Utm_campaign body string false "UTM campaign"
// @Param ProductID body int false "Product ID"
// @Param CustomerID body int false "Customer ID"
// @Param HasAddon body bool false "Has addon"
// @Param DateOfBirth body string false "Date of birth of the lead"
// @Param SupplierId body int false "Supplier ID"
// @Param PlanId body int false "Plan ID"
// @Param SupplierName body string false "Supplier name"
// @Param PlanName body string false "Plan name"
// @Param EnquiryId body int false "Enquiry ID"
// @Param LeadSource body string false "Lead source"
// @Param Source body string false "Source"
// @Param PreviousPolicyExpiryDate body string false "Previous policy expiry date"
// @Param PrevPolicyNo body string false "Previous policy number"
// @Success 200 {object} helpers.JsonResponse "Success"
// @Router /lead/insertUpdateAffiliateLeadDetails [post]
func GetCityByStateId(c *gin.Context) {

	helpers.Deferring()

	var modelgetCityByStateStruct model_getCityByStateId
	if err := c.ShouldBindUri(&modelgetCityByStateStruct); err != nil {
		c.JSON(400, gin.H{"msg": err})
		return
	}
	helpers.Deferring()
	var response helpers.JsonResponse
	response = modelgetCityByStateStruct.Service_getCityByStateId()

	c.JSON(200, response)

}

func (data model_getCityByStateId) Service_getCityByStateId() helpers.JsonResponse {
	helpers.Deferring()
	db := lib.InitialMigration()
	defer db.Close()
	fmt.Println(data)
	//tsql := fmt.Sprint("exec [dbo].[GET_State_Master]")
	tsql := fmt.Sprint("exec [dbo].[GET_City_Master] @StateId=?")

	//log.Fatal("tsql", tsql)
	rows, err := db.Query(tsql, data.StateId)

	helpers.CheckErr(err)
	defer rows.Close()

	ret := helpers.RenderData(rows)

	var response = helpers.JsonResponse{Error: false, Data: ret}

	return response

}
