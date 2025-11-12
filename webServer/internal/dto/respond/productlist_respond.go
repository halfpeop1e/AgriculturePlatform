package respond

type ProductRespond struct {
	Id          int64    `json:"id"`
	Name        string   `json:"name"`
	Image       []string `json:"image"`
	Description string   `json:"description"`
	Price       float64  `json:"price"`
	Stock       int64    `json:"stock"`
	Category    string   `json:"category"`
	Saler       string   `json:"saler"`
	SalerId     string   `json:"salerId"`
}

type ProductListRespond struct {
	ProductList []ProductRespond `json:"productList"`
}
