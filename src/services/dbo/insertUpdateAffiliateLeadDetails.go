package dbo

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	h "posp_api_go_v2/src/helpers"

	lib "github.com/amitg6062/golang-posp-dbconnection"
	"github.com/gin-gonic/gin"
)

type model_insertUpdateAffiliateLeadDetails struct {
	LeadID                   int    `json:"LeadID"`
	SessionID                int    `json:"SessionID"`
	Name                     string `json:"Name"`
	PosType                  int    `json:"PosType"`
	Gender                   string `json:"Gender"`
	MobileNo                 string `json:"MobileNo"`
	AltPhoneNo               string `json:"AltPhoneNo"`
	EmailID                  string `json:"EmailID"`
	Address                  string `json:"Address"`
	CityID                   int    `json:"CityID"`
	StateID                  int    `json:"StateID"`
	PostCode                 string `json:"PostCode"`
	Country                  string `json:"Country"`
	MaritalStatus            string `json:"MaritalStatus"`
	AnnualIncome             int    `json:"AnnualIncome"`
	HighestBookingStep       int    `json:"HighestBookingStep"`
	ReferralId               int    `json:"ReferralId"`
	ExitPointURL             string `json:"ExitPointURL"`
	Utm_source               string `json:"Utm_source"`
	UTM_Medium               string `json:"UTM_Medium"`
	Utm_term                 string `json:"Utm_term"`
	Utm_campaign             string `json:"Utm_campaign"`
	ProductID                int    `json:"ProductID"`
	CustomerID               int    `json:"CustomerID"`
	HasAddon                 bool   `json:"HasAddon"`
	DateOfBirth              string `json:"DateOfBirth"`
	SupplierId               int    `json:"SupplierId"`
	PlanId                   int    `json:"PlanId"`
	SupplierName             string `json:"SupplierName"`
	PlanName                 string `json:"PlanName"`
	EnquiryId                int    `json:"EnquiryId"`
	LeadSource               string `json:"LeadSource"`
	Source                   string `json:"Source"`
	PreviousPolicyExpiryDate string `json:"PreviousPolicyExpiryDate"`
	PrevPolicyNo             string `json:"PrevPolicyNo"`
}

type ValidateAffiliateLeadDetail struct {
	LeadID    int    `json:"LeadID" binding:"required"`
	SessionID int    `json:"SessionID" binding:"required"`
	Name      string `json:"Name" binding:"required"`
	PosType   int    `json:"PosType" binding:"required"`
}

func InsertUpdateAffiliateLeadDetails(c *gin.Context) {

	//	v := validator.New()

	var validateAffiliateLeadDetail ValidateAffiliateLeadDetail
	errDto := c.ShouldBind(&validateAffiliateLeadDetail)
	if errDto != nil {
		res := h.BuildErrorResponse("faild to process request", errDto.Error(), h.EmptyObj{})
		c.JSON(http.StatusBadRequest, res)
		return
	}

	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		fmt.Println("error found", err)
		// c.AbortWithStatus(400)
		return
	}

	var modelstruct model_insertUpdateAffiliateLeadDetails
	err = json.Unmarshal(body, &modelstruct)
	if err != nil {
		c.AbortWithStatus(400)
		return
	}

	var response h.JsonResponse

	response = modelstruct.Service_insertUpdateAffiliateLeadDetails()

	c.JSON(200, response)

}

func (data model_insertUpdateAffiliateLeadDetails) Service_insertUpdateAffiliateLeadDetails() h.JsonResponse {
	db := lib.InitialMigration()
	fmt.Println(data)

	tsql := fmt.Sprint("EXEC dbo.InsertUpdateAffiliateLeadDetails @LeadID=?, @SessionID=?, @Name=?, @PosType=?, @ProductID = 180")
	rows, err := db.Query(tsql, data.LeadID, data.SessionID, data.Name, data.PosType)

	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	ret := h.RenderData(rows)

	var response = h.JsonResponse{Error: false, Data: ret}

	return response

}
