package insertUpdateBmsLeadDetails

type RequestParam struct {
	LeadID                   int `validate:"required"`
	SessionID                string
	Name                     string
	Gender                   string
	MobileNo                 string
	AltPhoneNo               string
	EmailID                  string
	Address                  string
	CityID                   string
	StateID                  string
	PostCode                 string
	Country                  string
	MaritalStatus            string
	AnnualIncome             string
	HighestBookingStep       string
	ReferralId               string
	ExitPointURL             string
	Utm_source               string `validate:"required"`
	UTM_Medium               string
	Utm_term                 string
	Utm_campaign             string
	ProductID                string
	CustomerID               string
	HasAddon                 string
	DateOfBirth              string
	SupplierId               string
	PlanID                   string
	SupplierName             string
	PlanName                 string
	EnquiryId                string
	LeadSource               string
	Source                   string
	PreviousPolicyExpiryDate string
}

type JsonResponse struct {
	Error   bool                     `json:"error,bool"`
	Data    []map[string]interface{} `json:"data"`
	Message string                   `json:"message"`
}
