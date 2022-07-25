package partnerUpgrade

type RequestParam struct {
	OldAffiliateCode string `validate:"required"`
	NewAffiliateCode string `validate:"required"`
	Type             int16  `validate:"required"`
	ParentCode       string `validate:"required"`
}

type JsonResponse struct {
	Error   bool                     `json:"error,bool"`
	Data    []map[string]interface{} `json:"data"`
	Message string                   `json:"message"`
}
