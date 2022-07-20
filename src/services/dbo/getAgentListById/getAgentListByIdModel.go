package getAgentListById

import (
	"database/sql"
	"encoding/json"
)

type AgentListRequest struct {
	AffilateId     int
	Type           string
	IsVerified     string
	SortDir        int8
	PageNum        int8   `validate:"required"`
	PageSize       string `validate:"required"`
	SearchBy       string
	Circle         string
	FromDate       string
	ToDate         string
	ParentId       int16
	IsPanVerified  string
	IsGreenChannel string
}

/*
type AgentListResponse struct {
	TotalCount            int16
	RowNum                int16
	Type                  int16  `json:",omitempty"`
	AffiliateID           int16  `json:",omitempty"`
	RMCode                string `json:",omitempty"`
	ParentCode            string `json:",omitempty"`
	AffiliateCode         string `json:",omitempty"`
	Name                  string `json:",omitempty"`
	MobileNo              string `json:",omitempty"`
	EmailId               string `json:",omitempty"`
	BankName              string `json:",omitempty"`
	BankBranchAddress1    string `json:",omitempty"`
	BankBranchAddress2    string `json:",omitempty"`
	BankBranchStateId     int16  `json:",omitempty"`
	BankBranchCityId      int16  `json:",omitempty"`
	BankBranchPinCode     string `json:",omitempty"`
	BankAccountHolderName string `json:",omitempty"`
	BankAccountNumber     string `json:",omitempty"`
	NEFTIFSCCode          string `json:",omitempty"`
	RTGSIFSCCode          string `json:",omitempty"`
	UpdatedOn             string `json:",omitempty"`
	CreatedOn             string `json:",omitempty"`
	Status                bool   `json:",omitempty"`
	isVerified            int16  `json:",omitempty"`
	AffiliateStatusId     int16  `json:",omitempty"`
	AffiliateStatus       string `json:",omitempty"`
	BranchName            string `json:",omitempty"`
	EmployeeCode          string `json:",omitempty"`
	IsPanVerified         int16  `json:",omitempty"`
	ReviewerName          string `json:",omitempty"`
	IsCertified           int16  `json:",omitempty"`
	IsGreenChannel        int16  `json:",omitempty"`
}
*/

type JsonResponse struct {
	Error   bool                     `json:"error,bool"`
	Data    []map[string]interface{} `json:"dataAgents"`
	Message string                   `json:"message"`
}

// type JsonResponse struct {
// 	Error   bool                `json:"error,bool"`
// 	Data    []AgentListResponse `json:"dataAgents"`
// 	Message string              `json:"message"`
// }

type MyNullString struct {
	sql.NullString
}

func (s MyNullString) MarshalJSON() ([]byte, error) {
	if s.Valid {
		return json.Marshal(s.String)
	}
	return []byte(`null`), nil
}

type MyNullInt16 struct {
	sql.NullInt16
}

func (s MyNullInt16) MarshalJSON() ([]byte, error) {
	if s.Valid {
		return json.Marshal(s.Int16)
	}
	return []byte(`null`), nil
}

type MyNullBool struct {
	sql.NullBool
}

func (s MyNullBool) MarshalJSON() ([]byte, error) {
	if s.Valid {
		return json.Marshal(s.NullBool)
	}
	return []byte(`null`), nil
}

type AgentListResponse struct {
	TotalCount            int16
	RowNum                int16
	Type                  int16
	AffiliateID           int16        `json:",omitempty"`
	RMCode                MyNullString `json:",omitempty"`
	ParentCode            MyNullString `json:",omitempty"`
	AffiliateCode         MyNullString `json:",omitempty"`
	Name                  MyNullString `json:",omitempty"`
	MobileNo              MyNullString `json:",omitempty"`
	EmailId               MyNullString `json:",omitempty"`
	BankName              MyNullString `json:",omitempty"`
	BankBranchAddress1    MyNullString `json:",omitempty"`
	BankBranchAddress2    MyNullString `json:",omitempty"`
	BankBranchStateId     MyNullInt16  `json:",omitempty"`
	BankBranchCityId      MyNullInt16  `json:",omitempty"`
	BankBranchPinCode     MyNullString `json:",omitempty"`
	BankAccountHolderName MyNullString `json:",omitempty"`
	BankAccountNumber     MyNullString `json:",omitempty"`
	NEFTIFSCCode          MyNullString `json:",omitempty"`
	RTGSIFSCCode          MyNullString `json:",omitempty"`
	UpdatedOn             MyNullString `json:",omitempty"`
	CreatedOn             MyNullString `json:",omitempty"`
	Status                MyNullString `json:",omitempty"`
	IsVerified            MyNullInt16  `json:"isVerified,omitempty"`
	AffiliateStatusId     MyNullInt16  `json:",omitempty"`
	AffiliateStatus       MyNullString `json:",omitempty"`
	BranchName            MyNullString `json:",omitempty"`
	EmployeeCode          MyNullString `json:",omitempty"`
	IsPanVerified         MyNullInt16  `json:",omitempty"`
	ReviewerName          MyNullString `json:",omitempty"`
	IsCertified           MyNullString `json:",omitempty"`
	IsGreenChannel        MyNullInt16  `json:",omitempty"`
}
