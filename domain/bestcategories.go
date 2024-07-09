package domain

type BestCategoriesData struct {
	CategoryName   string `json:"categoryName"`
	IsRocket       bool   `json:"isRocket"`
	IsFreeShipping bool   `json:"isFreeShipping"`
	ProductId      int    `json:"productId"`
	ProductImage   string `json:"productImage"`
	ProductName    string `json:"productName"`
	ProductPrice   int    `json:"productPrice"`
	ProductUrl     string `json:"productUrl"`
}

type BestCategoriesResponseBody struct {
	RCode    string               `json:"rCode"`
	RMessage string               `json:"rMessage"`
	Data     []BestCategoriesData `json:"data"`
}
