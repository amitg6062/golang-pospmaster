package rnd

import (
	"database/sql"

	"gorm.io/gorm"
)

type Emp struct {
	Id       int    `json:"id"`
	Name     string `json:"username"`
	Location string `json:"location"`
}

type Emp2 struct {
	Id       int    `json:"id"`
	Name     string `json:"username" validate:"required"`
	Location string `json:"location" validate:"required"`
}

type JsonResponse struct {
	Type    string `json:"type"`
	Data    []Emp  `json:"data"`
	Message string `json:"message"`
}

type PartnerLeadsSchema struct {
	RegistrationNo string
	FromDate       string
	ToDate         string
	AgentID        string
	ParentID       string
	RMID           string
	SearchBy       string
	SearchByLead   string
	ProductGroupId string
	SortDir        string
	PageNum        string
	PageSize       string
}

type PartnerLeadsFinalResponse struct {
	Type    string                 `json:"type"`
	Data    []PartnerLeadsResponse `json:"data"`
	Message string                 `json:"message"`
}

type PartnerLeadsResponse struct {
	RowNum             string         `json:"RowNum"`
	LeadID             string         `json:"LeadID"`
	Product            string         `json:"Product"`
	MobileNo           string         `json:"MobileNo"`
	CustomerName       string         `json:"CustomerName"`
	LeadCreationDate   string         `json:"LeadCreationDate"`
	BookingDate        string         `json:"BookingDate"`
	SupplierName       string         `json:"SupplierName"`
	PlanName           string         `json:"PlanName"`
	PremiumAmount      string         `json:"PremiumAmount"`
	PaymentMode        string         `json:"PaymentMode"`
	CurrentStatus      string         `json:"CurrentStatus"`
	PolicyLink         string         `json:"PolicyLink"`
	SelectionPlanID    string         `json:"SelectionPlanID"`
	SumInsured         string         `json:"SumInsured"`
	ProductID          string         `json:"ProductID"`
	Utm_term           string         `json:"Utm_term"`
	UTM_Medium         string         `json:"UTM_Medium"`
	Utm_campaign       string         `json:"Utm_campaign"`
	ODPremium          string         `json:"ODPremium"`
	APE                string         `json:"APE"`
	AgentName          string         `json:"AgentName"`
	ExitPointURL       string         `json:"ExitPointURL"`
	RegistrationNumber sql.NullString `json:"RegistrationNumber"`
	InspectionStatusId string         `json:"InspectionStatusId"`
	TotalCount         string         `json:"TotalCount"`
}

//------------STRUCTS---------------------

type User struct {
	gorm.Model
	Name     string `json:"name"`
	Email    string `gorm:"unique" json:"email"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

type Authentication struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Token struct {
	Role        string `json:"role"`
	Email       string `json:"email"`
	TokenString string `json:"token"`
}

type Error struct {
	IsError bool   `json:"isError"`
	Message string `json:"message"`
}

var (
	secretkey string = "secretkeyjwt"
)
