package getBookingListDumpV1

import (
	"database/sql"
	"time"
)

type RequestParam struct {
	ToDate      string
	FromDate    string
	ProductId   string
	PartnerCode string
	ParentId    int16
	PageNum     int16
	PageSize    int16
}

type JsonResponse struct {
	Error   bool   `json:"error,bool"`
	Data    []dict `json:"data"`
	Message string `json:"message"`
}

type OutputData struct {
	LeadId          int            `json:"LeadId"`
	InsuredName     sql.NullString `json:"InsuredName"`
	Dob             time.Time      `json:"DOB"`
	Insurer         sql.NullString `json:"Insurer"`
	InsurerFullName sql.NullString `json:"InsurerFullName"`
	// Ape              sql.NullInt64  `json:"APE"`
	ActualLeadSource sql.NullString `json:"ActualLeadSource"`
	Address          sql.NullString `json:"Address"`
	ApplicationNo    sql.NullString `json:"ApplicationNo"`
	BasicPremium     sql.NullString `json:"BasicPremium,omitempty"`
	// BookingDate      sql.NullTime   `json:"BookingDate"`
	BookingMode      sql.NullString `json:"BookingMode"`
	BusinessType     sql.NullString `json:"BusinessType"`
	ChatStatus       sql.NullString `json:"ChatStatus"`
	Circle           sql.NullString `json:"Circle"`
	City             sql.NullString `json:"City"`
	CubicCapacity    sql.NullInt16  `json:"CubicCapacity"`
	CustomerID       sql.NullInt16  `json:"CustomerId"`
	Discount         sql.NullString `json:"Discount"`
	FuleType         sql.NullString `json:"FuleType"`
	Istp             bool           `json:"ISTP"`
	InstallmentsPaid sql.NullInt16  `json:"InstallmentsPaid"`

	IsE2E           sql.NullInt16 `json:"IsE2E"`
	IssuanceRejDate time.Time     `json:"Issuance/Rej Date"`
	LeadDate        time.Time     `json:"LeadDate"`

	LeadRank      sql.NullString `json:"LeadRank"`
	MakeName      sql.NullString `json:"MakeName"`
	MaritalStatus sql.NullString `json:"MaritalStatus"`
	// NetPremium             sql.NullString `json:"NetPremium"`
	NoOfSeats              sql.NullInt16  `json:"NoOfSeats"`
	Noofwheels             sql.NullInt16  `json:"Noofwheels"`
	ODPremium              sql.NullInt16  `json:"ODPremium"`
	PGType                 sql.NullString `json:"PGType"`
	ParentID               sql.NullString `json:"ParentId"`
	ParentLeadCreationDate sql.NullString `json:"ParentLeadCreationDate"`
	ParentLeadSource       sql.NullString `json:"ParentLeadSource"`
	PartnerID              sql.NullString `json:"PartnerId"`
	PaymentPeriodicity     sql.NullString `json:"PaymentPeriodicity"`
	PaymentSubStatus       sql.NullInt16  `json:"PaymentSubStatus"`
	PersonalAccidentCover  sql.NullString `json:"PersonalAccidentCover"`
	PinCode                sql.NullInt16  `json:"PinCode"`
	PlanName               sql.NullString `json:"PlanName"`
	PolicyNo               sql.NullString `json:"PolicyNo"`
	PolicyType             sql.NullString `json:"PolicyType"`
	Premium                sql.NullInt16  `json:"Premium"`
	Product                sql.NullString `json:"Product"`
	RMCode                 sql.NullString `json:"RMCode"`
	RMName                 sql.NullString `json:"RMName"`
	RegistrationDate       time.Time      `json:"RegistrationDate"`
	RegistrationNo         sql.NullString `json:"RegistrationNo"`
	Source                 sql.NullString `json:"Source"`
	State                  sql.NullString `json:"State"`
	Status                 sql.NullString `json:"Status"`
	StpNstp                sql.NullString `json:"StpNstp"`
	SumInsured             sql.NullString `json:"SumInsured"`
	TPPremium              sql.NullString `json:"TPPremium"`
	UtmMedium              sql.NullString `json:"Utm_Medium"`
	UtmCampaign            sql.NullString `json:"Utm_campaign"`
	UtmSource              sql.NullString `json:"Utm_source"`
	UtmTerm                sql.NullString `json:"Utm_term"`
	VechicleCarrier        sql.NullString `json:"VechicleCarrier"`
	VehicleAge             sql.NullInt16  `json:"VehicleAge"`
	VehicleModelName       sql.NullString `json:"VehicleModelName"`
	VehicleSubClass        sql.NullString `json:"VehicleSubClass"`
	Grossvehicleweight     sql.NullInt16  `json:"grossvehicleweight"`
	KaliPili               sql.NullString `json:"kaliPili"`
}
