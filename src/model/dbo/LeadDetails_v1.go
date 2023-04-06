package dbo

type LeadDetails_v1 struct {
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
