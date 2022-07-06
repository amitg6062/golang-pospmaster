package exportExcel

import "database/sql"

type Emp struct {
	Id       int            `json:"id"`
	Name     sql.NullString `json:"username,omitempty"`
	Location sql.NullString `json:"location,omitempty"`
}

type JsonResponse struct {
	Type    string `json:"type"`
	Data    []Emp  `json:"data"`
	Message string `json:"message"`
}
