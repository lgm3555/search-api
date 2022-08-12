package response

// Response API Response값
type Response struct {
	Desc   string `json:"desc"  example:"null"`
	Status int    `json:"status" example:"200"`
	Result `json:"result,omitempty" example:"null"`
}

// Result 검색엔진 결과값
type Result struct {
	SearchList interface{} `json:"searchList,omitempty"`
}
