package insertUpdateBmsBookingDetails

type RequestParam struct {
	LeadID         int64 `json:"LeadID",validate:"required"`
	SubProduct     string
	PlanType       string
	FinalPremium   string
	BasicOdPremium string
	TotalOdPremium string
	RTO            string
	Zone           string

	CPA string

	PACoverPremium        string
	Discount              string
	IDV                   string
	RegistrationNumber    string
	NumberOfWheels        string
	GVW                   string
	CubicCapacity         string
	CarryingCapacity      string
	SeatingCapacity       string
	RegistrationDate      string
	ManufacturingDate     string
	KaliPili              string
	VehicleState          string
	VehicleCity           string
	FuelType              string
	VehicleSubClass       string
	VehicleAge            string
	VehicleMake           string
	VehicleModel          string
	VariantName           string
	VariantID             string
	VehicleCarrier        string
	BusinessType          string
	NCB                   string
	BasicTPPremium        string
	TotalTPPremium        string
	ZDTaken               string
	IsInspectionFirstCase string
	IsPortability         string
	AgeOfEldestMember     string
	NoOfAdults            string
	NoOfChildren          string
	NatureOfBusiness      string
	Structure             string
	Parental              string
	SumInsuredType        string
	ZDPremium             string
	PlanVariant           string

	SupplierId       string
	IsBooked         string
	IssuanceDate     string
	PolicyPdfUrl     string
	BookingSource    string
	PolicyTenureType string

	IsSTP                     bool    `json:"IsSTP"`
	IsLinkPB                  bool    `json:"IsLinkPB"`
	IsStandingInstruction     bool    `json:"IsStandingInstruction"`
	IsEMI                     bool    `json:"IsEMI"`
	PGType                    int     `json:"PGType"`
	PGHoldRelease             int     `json:"PGHoldRelease"`
	SumInsured                int     `json:"SumInsured"`
	TotalPremium              int     `json:"TotalPremium"`
	CustomerAmount            int     `json:"CustomerAmount"`
	PaidPremium               int     `json:"PaidPremium"`
	ServiceTax                int     `json:"ServiceTax"`
	ODPremium                 int     `json:"ODPremium"`
	RefundAmount              int     `json:"RefundAmount"`
	DeductionAmount           int     `json:"DeductionAmount"`
	BasicPremium              int     `json:"BasicPremium"`
	OfferCreatedON            string  `json:"OfferCreatedON"`
	OfferCreatedOnDt          string  `json:"OfferCreatedOnDt"`
	CreatedOn                 string  `json:"CreatedOn"`
	CreatedOnDt               string  `json:"CreatedOnDt"`
	ProductID                 int     `json:"ProductID"`
	PlanID                    int     `json:"PlanId"`
	PaymentSTATUS             int     `json:"PaymentSTATUS"`
	PolicyTypeID              int     `json:"PolicyTypeId"`
	SupplierID                int     `json:"SupplierId"`
	PaymentSubStatus          int     `json:"PaymentSubStatus"`
	PolicyTerm                int     `json:"PolicyTerm"`
	PayTerm                   int     `json:"PayTerm"`
	RefundBy                  int     `json:"RefundBy"`
	UserID                    int     `json:"UserID"`
	IsHealthEx                int     `json:"IsHealthEx"`
	CoverTypeID               int     `json:"CoverTypeId"`
	VehicleTypeID             int     `json:"VehicleTypeId"`
	NoOfEmployees             int     `json:"NoOfEmployees"`
	NoOfLives                 int     `json:"NoOfLives"`
	OfferPack                 int     `json:"OfferPack"`
	PolicyType                int     `json:"PolicyType"`
	RiderSI                   int     `json:"RiderSI"`
	InstallmentsPaid          int     `json:"InstallmentsPaid"`
	Portability               int     `json:"Portability"`
	CreditReceived            int     `json:"CreditReceived"`
	ItemID                    int     `json:"ItemId"`
	PaymentDate               string  `json:"PaymentDate"`
	DateOfInspection          string  `json:"DateOfInspection"`
	OfferNumber               string  `json:"OfferNumber"`
	ProductName               string  `json:"ProductName"`
	SelectedPlanName          string  `json:"SelectedPlanName"`
	ChequeNo                  string  `json:"ChequeNo"`
	BankNameBranch            string  `json:"BankNameBranch"`
	PolicyTypeName            string  `json:"PolicyTypeName"`
	SupplierName              string  `json:"SupplierName"`
	TransRefNo                string  `json:"TransRefNo"`
	PaymentPeriodicity        string  `json:"PaymentPeriodicity"`
	BookingType               string  `json:"BookingType"`
	OrderNo                   string  `json:"OrderNo"`
	InsuredName               string  `json:"InsuredName"`
	ApplicationNo             string  `json:"ApplicationNo"`
	PolicyNo                  string  `json:"PolicyNo"`
	Rider                     string  `json:"Rider"`
	PreviousBookingNo         string  `json:"PreviousBookingNo"`
	CourierName               string  `json:"CourierName"`
	CourierAddress            string  `json:"CourierAddress"`
	TrackingNumber            string  `json:"TrackingNumber"`
	RefundChequeTxnID         string  `json:"RefundChequeTxnID"`
	RefundBankName            string  `json:"RefundBankName"`
	SalesAgent                string  `json:"SalesAgent"`
	ReferenceNo               string  `json:"ReferenceNo"`
	InspectionStatus          string  `json:"InspectionStatus"`
	PolicyStartDate           string  `json:"PolicyStartDate"`
	PolicyStartDateDt         string  `json:"PolicyStartDateDt"`
	PolicyEndDate             string  `json:"PolicyEndDate"`
	PolicyEndDateDt           string  `json:"PolicyEndDateDt"`
	Bank                      string  `json:"Bank"`
	PaymentMode               string  `json:"PaymentMode"`
	PaymentType               string  `json:"PaymentType"`
	PolicyLink                string  `json:"PolicyLink"`
	PaymentSource             string  `json:"PaymentSource"`
	PrevPolicyNo              string  `json:"PrevPolicyNo"`
	ProposalNo                string  `json:"ProposalNo"`
	IseMandate                bool    `json:"IseMandate"`
	IseNACH                   bool    `json:"IseNACH"`
	CubicCapacityID           int     `json:"CubicCapacityID"`
	IsExclusiveBooking        bool    `json:"IsExclusiveBooking"`
	IsROC                     bool    `json:"IsROC"`
	IsDiscounted              bool    `json:"IsDiscounted"`
	IsS3URL                   bool    `json:"IsS3URL"`
	IsPACover                 bool    `json:"IsPACover"`
	RegistrationNo            string  `json:"RegistrationNo"`
	PolicyTenure              int     `json:"PolicyTenure"`
	PolicyExtension           int     `json:"PolicyExtension"`
	ExtPremium                float64 `json:"ExtPremium"`
	ExtDate                   string  `json:"ExtDate"`
	TPPremium                 int     `json:"TPPremium"`
	TPStartDate               string  `json:"TPStartDate"`
	TPEndDate                 string  `json:"TPEndDate"`
	InsurerPolicyNo           string  `json:"InsurerPolicyNo"`
	InspectionType            int     `json:"InspectionType"`
	CurrencyID                int     `json:"CurrencyId"`
	CardlessEMI               string  `json:"CardlessEMI"`
	IsVehicleVerified         bool    `json:"IsVehicleVerified"`
	InspectionFirst           bool    `json:"InspectionFirst"`
	VehicleVerificationStatus int     `json:"VehicleVerificationStatus"`
	DocUploadID               string  `json:"DocUploadId"`
	MobileNo                  int64   `json:"MobileNo"`
	CountryID                 int     `json:"CountryId"`
	EncEmailID                string  `json:"EncEmailId"`
}

type JsonResponse struct {
	Error   bool                     `json:"error,bool"`
	Data    []map[string]interface{} `json:"data"`
	Message string                   `json:"message"`
}
