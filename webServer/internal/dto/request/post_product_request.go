package request

type PostProductRequest struct {
	Name        string   `form:"name" json:"name"`
	Images      []string `form:"images" json:"images"`
	Category    string   `form:"category" json:"category"`
	Description string   `form:"description" json:"description"`
	Price       float64  `form:"price" json:"price"`
	Stock       int64    `form:"stock" json:"stock"`
	Saler       string   `form:"saler" json:"saler"`
	SalerId     string   `form:"salerId" json:"salerId"`
}
