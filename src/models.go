package api

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

// type JsonResponse struct {
// 	Type    string `json:"type"`
// 	Data    []Emp  `json:"data"`
// 	Message string `json:"message"`
// }

type Emp struct {
	Id       int    `json:"id"`
	Name     string `json:"username"`
	Location string `json:"location"`
}

type JsonResponse struct {
	Error   bool                     `json:"error,bool"`
	Data    []map[string]interface{} `json:"data"`
	Message string                   `json:"message"`
}
