package respond

type Order struct {
	OrderId    string `json:"orderId"`
	Name       string `json:"name"`
	Quantity   int    `json:"quantity"`
	Totalprice int    `json:"totalprice"`
	Status     string `json:"status"`
	Type       string `json:"type"`
	Buyer      string `json:"buyer"`
	Saler      string `json:"saler"`
}

type OrderListRespond struct {
	Orders []Order `json:"orders"`
}
