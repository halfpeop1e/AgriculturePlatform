package request

type BuyProductRequest struct {
	OrderId    string `json:"orderId"`
	ProductId  int    `json:"productId"`
	Quantity   int    `json:"quantity"`
	Totalprice int    `json:"totalprice"`
}
