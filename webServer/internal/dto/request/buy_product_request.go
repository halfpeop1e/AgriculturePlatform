package request

type BuyProductRequest struct {
	OrderId    string `json:"orderId"`
	ProductId  string `json:"productId"`
	Quantity   int    `json:"quantity"`
	Totalprice int    `json:"totalprice"`
	Buyer      string `json:"buyer"`
	Saler      string `json:"saler"`
}

type BuyProductRequestWrapper struct {
	Order BuyProductRequest `json:"order"`
}
