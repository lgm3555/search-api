package model

// prod index feild
type ProdSearchEs struct {
	PRODUCTCODE  string `json:"productCode"`
	PRODUCTNAME  string `json:"productName"`
	PRICE        int    `json:"price"`
	STOCKYN      string `json:"stockYN"`
	REGISTERDATE string `json:"registerDate"`
}
