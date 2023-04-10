package dbo

import (
	"fmt"

	"posp_api_go_v2/src/helpers"

	lib "github.com/amitg6062/golang-posp-dbconnection"
	"github.com/gin-gonic/gin"
)

type model_getstatemaster struct{}

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
func GetStateMaster(c *gin.Context) {
	helpers.Deferring()

	var modelgetStatestruct model_getstatemaster
	// if err := c.ShouldBindUri(&modelstruct); err != nil {
	// 	c.JSON(400, gin.H{"msg": err})
	// 	return
	// }
	//c.JSON(200, gin.H{"name": modelstruct.AssociationTypeId, "uuid": modelstruct.IsActive})

	// AssociationTypeId_, err1 := strconv.ParseUint(c.Param("AssociationTypeId"), 0, 0)
	// IsActive_, err2 := strconv.ParseUint(c.Param("IsActive"), 0, 0)

	// if err1 != nil {
	// 	res := helpers.BuildErrorResponse("Invalid AssociationTypeId", err1.Error(), helpers.EmptyObj{})
	// 	c.AbortWithStatusJSON(http.StatusBadRequest, res)
	// }

	// if err2 != nil {
	// 	res := helpers.BuildErrorResponse("Invalid AssociationTypeId", err2.Error(), helpers.EmptyObj{})
	// 	c.AbortWithStatusJSON(http.StatusBadRequest, res)
	// }

	//var modelstruct model_getDistributionByAssociationTypeId
	// errDto := c.ShouldBind(&modelstruct)

	// if errDto != nil {
	// 	res := helpers.BuildErrorResponse("faild to process request", errDto.Error(), helpers.EmptyObj{})
	// 	c.JSON(http.StatusBadRequest, res)
	// 	return
	// }

	var response helpers.JsonResponse
	response = modelgetStatestruct.Service_getStateMaster()

	c.JSON(200, response)

}

func (data model_getstatemaster) Service_getStateMaster() helpers.JsonResponse {
	helpers.Deferring()
	db := lib.InitialMigration()
	defer db.Close()

	fmt.Println(data)

	tsql := fmt.Sprint("exec [dbo].[GET_State_Master]")

	//log.Fatal("tsql", tsql)
	rows, err := db.Query(tsql)

	helpers.CheckErr(err)
	defer rows.Close()

	ret := helpers.RenderData(rows)

	var response = helpers.JsonResponse{Error: false, Data: ret}

	return response

}
