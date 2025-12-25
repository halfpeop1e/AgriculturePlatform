package respond

// MarketCategoryData 大类走势数据
type MarketCategoryData struct {
	Name  string   `json:"name"`
	Data  []string `json:"data"`
	Color string   `json:"color"`
}

// MarketProductData 单品详情数据
type MarketProductData struct {
	Name   string   `json:"name"`
	Price  string   `json:"price"`
	Change string   `json:"change"`
	Data   []string `json:"data"`
}

// MarketDataRespond 市场数据响应
type MarketDataRespond struct {
	Dates      []string             `json:"dates"`
	Categories []MarketCategoryData `json:"categories"`
	Products   []MarketProductData  `json:"products"`
}
