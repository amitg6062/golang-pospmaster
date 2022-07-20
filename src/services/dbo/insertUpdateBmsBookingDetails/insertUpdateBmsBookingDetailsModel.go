package insertUpdateBmsBookingDetails

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

type JsonResponse struct {
	Error   bool                     `json:"error,bool"`
	Data    []map[string]interface{} `json:"dataAgents"`
	Message string                   `json:"message"`
}

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
