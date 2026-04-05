package orders


type orderItem struct {
	ProductId int64 `json:"product_id"`
	Quantity int `json:"quantity"`
}

type createOrderParams struct {
	CustomerId int64 `json:"customer_id"`
	Items []orderItem `json:"items"`
}