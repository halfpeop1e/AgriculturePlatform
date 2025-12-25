package gorm

import (
	"encoding/json"
	"fmt"
	"go-agriculture/internal/config"
	"go-agriculture/internal/dao"
	"go-agriculture/internal/dto/request"
	"go-agriculture/internal/dto/respond"
	"go-agriculture/internal/model"
	"go-agriculture/internal/util"
	"log"
	"math/rand"
	"strconv"
	"time"
)

type productServer struct{}

var ProductServer = &productServer{}

func (p *productServer) PostProduct(req request.PostProductRequest) (string, int) {
	if len(req.Name) > 50 || len(req.Category) > 200 || len(req.Saler) > 20 {
		return "参数错误", -1
	}
	if req.Price < 0 || req.Stock < 0 {
		return "数量错误", -1
	}
	var newProduct model.Product
	// var count int64
	// res := dao.GormDB.Model(&model.Product{}).Where("name = ? AND saler_id = ?", req.Name, req.SalerId).Count(&count)
	// if res.Error != nil {
	// 	log.Printf("Database error: %v", res.Error)
	// 	return "查询失败", -1
	// }
	// if count > 0 {
	// 	return "商品已存在", -1
	// }
	allImages, _ := util.FileReceiverBase64Util(req.Images, util.IsFileType)
	imagesJSON, err := json.Marshal(allImages)
	if err != nil {
		return "图片序列化失败", -1
	}
	newProduct.Uuid = util.GenerateUUID()
	newProduct.Name = req.Name
	newProduct.Images = string(imagesJSON)
	newProduct.Category = req.Category
	newProduct.Description = req.Description
	newProduct.Price = req.Price
	newProduct.Stock = req.Stock
	newProduct.Saler = req.Saler
	newProduct.SalerId = req.SalerId
	newProduct.Status = "1"
	newProduct.Email = req.ContactEmail
	newProduct.CreatAt = time.Now()
	newProduct.UpdateAt = time.Now()
	dao.GormDB.Create(&newProduct)
	return "提交成功", 0
}

func (p *productServer) GetProductList(page, pageSize int, salerId string) (string, *respond.ProductListRespond, int) {
	var total int64
	if res := dao.GormDB.Model(&model.Product{}).Where("status = ?", "1").Count(&total); res.Error != nil {
		log.Printf("Database error: %v", res.Error)
		return "查询失败", nil, -1
	}

	productList := &respond.ProductListRespond{
		ProductList: make([]respond.ProductRespond, 0),
		Total:       total,
		Page:        page,
		PageSize:    pageSize,
		HasMore:     int64(page*pageSize) < total,
	}

	if total == 0 {
		return "获取成功", productList, 0
	}

	offset := (page - 1) * pageSize
	if offset < 0 {
		offset = 0
	}
	log.Print("收到的salerId:" + salerId)
	var products []model.Product
	if salerId != "all" {
		if res := dao.GormDB.Where("status = ? AND saler_id = ?", "1", salerId).Order("create_at DESC").Limit(pageSize).Offset(offset).Find(&products); res.Error != nil {
			log.Printf("Database error: %v", res.Error)
			return "查询失败", nil, -1
		}
	} else {
		if res := dao.GormDB.Where("status = ?", "1").Order("create_at DESC").Limit(pageSize).Offset(offset).Find(&products); res.Error != nil {
			log.Printf("Database error: %v", res.Error)
			return "查询失败", nil, -1
		}
	}

	for _, product := range products {
		images := make([]string, 0)
		if product.Images == "" {
			images = []string{}
		} else if err := json.Unmarshal([]byte(product.Images), &images); err != nil {
			images = []string{}
		}
		for i, image := range images {
			images[i] = "http://" + fmt.Sprint(config.GetConfig().MainConfig.ServerIP) + ":" + fmt.Sprint(config.GetConfig().MainConfig.Port) + "/static/files/" + image
		}
		newProduct := respond.ProductRespond{
			Id:          product.Id,
			Name:        product.Name,
			Image:       images,
			Price:       product.Price,
			Stock:       product.Stock,
			Category:    product.Category,
			Saler:       product.Saler,
			SalerId:     product.SalerId,
			Description: product.Description,
		}
		productList.ProductList = append(productList.ProductList, newProduct)
	}

	if len(productList.ProductList) == 0 {
		productList.HasMore = false
	} else {
		productList.HasMore = int64(page*pageSize) < total
	}
	return "获取成功", productList, 0
}

func (p *productServer) BuyProduct(req request.BuyProductRequest, user_id string) (string, int) {
	var product model.Product
	productId, _ := strconv.Atoi(req.ProductId)
	if res := dao.GormDB.Where("id = ?", productId).First(&product); res.Error != nil {
		log.Printf("Database error: %v", res.Error)
		return "查询失败", -1
	}
	if product.Stock < int64(req.Quantity) {
		return "库存不足", -1
	}
	var buyUser model.User
	if user_id == "" {
		if res := dao.GormDB.Where("name = ?", req.Buyer).First(&buyUser); res.Error != nil {
			log.Printf("Database error: %v", res.Error)
			return "用户查询失败", -1
		}
	} else {
		if res := dao.GormDB.Where("uuid = ?", user_id).First(&buyUser); res.Error != nil {
			log.Printf("Database error: %v", res.Error)
			return "用户查询失败", -1
		}
	}
	order := model.Order{
		Uuid:       util.GenerateUUID(),
		Name:       product.Name,
		Quantity:   int64(req.Quantity),
		Totalprice: float64(req.Totalprice),
		Status:     "待支付",
		Ordertype:  "buy",
		ProductId:  product.Id,
		UserId:     user_id,
		Buyer:      buyUser.Name,
		Saler:      product.Saler,
		CreatAt:    time.Now(),
	}
	dao.GormDB.Create(&order)
	order = model.Order{
		Uuid:       util.GenerateUUID(),
		Name:       product.Name,
		Quantity:   int64(req.Quantity),
		Totalprice: float64(req.Totalprice),
		Status:     "待支付",
		Ordertype:  "sell",
		ProductId:  product.Id,
		UserId:     product.SalerId,
		Buyer:      buyUser.Name,
		Saler:      product.Saler,
		CreatAt:    time.Now(),
	}
	dao.GormDB.Create(&order)
	product.Stock -= int64(req.Quantity)
	dao.GormDB.Save(&product)
	return "购买成功", 0
}

func (p *productServer) GetOrderList() (string, int, *respond.OrderListRespond) {
	var orders []model.Order
	if res := dao.GormDB.Find(&orders); res.Error != nil {
		log.Printf("Database error: %v", res.Error)
		return "查询失败", -1, nil
	}
	orderList := &respond.OrderListRespond{}
	for _, order := range orders {
		newOrder := respond.Order{
			OrderId:    order.Uuid,
			Name:       order.Name,
			Quantity:   int(order.Quantity),
			Totalprice: int(order.Totalprice),
			Status:     order.Status,
			Type:       order.Ordertype,
			Buyer:      order.Buyer,
			Saler:      order.Saler,
		}
		orderList.Orders = append(orderList.Orders, newOrder)
	}
	return "获取成功", 0, orderList
}

func (p *productServer) EditerProduct(req request.EditerProductRequest, product_id string) (string, int) {
	var product model.Product
	if res := dao.GormDB.Where("id = ?", product_id).First(&product); res.Error != nil {
		log.Printf("Database error: %v", res.Error)
		return "查询失败", -1
	}
	if req.Name != "" {
		product.Name = req.Name
	}
	if req.Category != "" {
		product.Category = req.Category
	}
	if req.Description != "" {
		product.Description = req.Description
	}
	if req.Price != 0 {
		product.Price = req.Price
	}
	if req.Stock != 0 {
		product.Stock = req.Stock
	}
	if req.Saler != "" {
		product.Saler = req.Saler
	}
	if req.SalerId != "" {
		product.SalerId = req.SalerId
	}
	if req.ContactEmail != "" {
		product.Email = req.ContactEmail
	}
	if len(req.Images) > 0 {
		allImages, _ := util.FileReceiverBase64Util(req.Images, util.IsFileType)
		imagesJSON, _ := json.Marshal(allImages)
		product.Images = string(imagesJSON)
	}
	dao.GormDB.Save(&product)
	return "编辑成功", 0
}

func (p *productServer) DeleteProduct(productId string) (string, int) {
	product := model.Product{}
	if res := dao.GormDB.Where("id = ?", productId).First(&product); res.Error != nil {
		log.Printf("Database error: %v", res.Error)
		return "查询失败", -1
	}
	dao.GormDB.Delete(&product)
	return "删除成功", 0
}

func (p *productServer) GetDateAnlazy() (string, *respond.MarketDataRespond, int) {
	var orders []model.Order
	if res := dao.GormDB.Where("status = ? AND ordertype = ?", "已支付", "buy").Find(&orders); res.Error != nil {
		log.Printf("Database error: %v", res.Error)
		return "查询失败", nil, -1
	}

	// 生成前7天加上后2天的日期（后2天作为预测）
	dates := make([]string, 9)
	for i := 0; i < 9; i++ {
		date := time.Now().AddDate(0, 0, -7+i)
		dates[i] = date.Format("01-02")
	}

	// 统计各分类的平均价格
	categoryMap := make(map[string][]float64)
	productMap := make(map[string]*respond.MarketProductData)

	// 从订单中统计价格数据
	for _, order := range orders {
		var product model.Product
		if res := dao.GormDB.Where("id = ?", order.ProductId).First(&product); res.Error != nil {
			log.Printf("Database error: %v", res.Error)
			continue
		}

		// 统计分类数据
		if _, exists := categoryMap[product.Category]; !exists {
			categoryMap[product.Category] = make([]float64, 9)
		}

		// 根据订单日期计算价格（简化处理，使用订单价格作为当日价格）
		orderDay := order.CreatAt.Day()
		index := orderDay - time.Now().Day() + 7
		if index >= 0 && index < 7 { // 0-6表示前7天，7-8是后2天的预测
			categoryMap[product.Category][index] += order.Totalprice / float64(order.Quantity)
		}

		// 统计产品数据
		if _, exists := productMap[product.Name]; !exists {
			productMap[product.Name] = &respond.MarketProductData{
				Name:   product.Name,
				Price:  fmt.Sprintf("%.2f", product.Price),
				Change: fmt.Sprintf("%.1f", (rand.Float64()*4 - 2)),
				Data:   make([]string, 9),
			}
		}

		// 为产品生成价格走势数据（前7天基于真实价格，后2天基于预测）
		for i := 0; i < 7; i++ {
			priceVariation := product.Price * (1 + (rand.Float64()-0.5)*0.2)
			productMap[product.Name].Data[i] = fmt.Sprintf("%.2f", priceVariation)
		}
		// 明天的预测价格（基于当前价格的小幅波动）
		tomorrowPrice := product.Price * (1 + (rand.Float64()-0.5)*0.1)
		productMap[product.Name].Data[7] = fmt.Sprintf("%.2f", tomorrowPrice)
		// 后天的预测价格（基于明天预测价格的进一步波动）
		dayAfterTomorrowPrice := tomorrowPrice * (1 + (rand.Float64()-0.5)*0.15)
		productMap[product.Name].Data[8] = fmt.Sprintf("%.2f", dayAfterTomorrowPrice)
	}

	// 构建分类数据
	var categories []respond.MarketCategoryData
	categoryColors := map[string]string{
		"grain":     "#f59e0b",
		"vegetable": "#10b981",
		"fruit":     "#ef4444",
		"seedling":  "#8b5cf6",
	}

	// 如果没有真实数据，生成模拟数据
	if len(categoryMap) == 0 {
		categories = []respond.MarketCategoryData{
			{Name: "粮食", Data: []string{"2.5", "2.8", "2.6", "2.7", "2.9", "2.4", "2.6", "2.8", "2.9"}, Color: "#f59e0b"},
			{Name: "蔬菜", Data: []string{"3.5", "4.2", "3.8", "4.0", "3.9", "4.1", "4.0", "4.2", "4.3"}, Color: "#10b981"},
			{Name: "水果", Data: []string{"8.5", "9.2", "8.8", "9.0", "8.7", "9.1", "9.0", "9.3", "9.4"}, Color: "#ef4444"},
			{Name: "种苗", Data: []string{"1.5", "1.8", "1.6", "1.7", "1.9", "1.4", "1.6", "1.7", "1.8"}, Color: "#8b5cf6"},
		}
	} else {
		for category, prices := range categoryMap {
			data := make([]string, 9)
			var recentPrices []float64
			for i := 0; i < 7; i++ { // 前7天数据
				if prices[i] > 0 {
					data[i] = fmt.Sprintf("%.1f", prices[i])
					recentPrices = append(recentPrices, prices[i])
				} else {
					// 如果某天没有数据，生成模拟数据
					basePrice := 2.0
					switch category {
					case "grain":
						basePrice = 3.0
					case "vegetable":
						basePrice = 4.0
					case "fruit":
						basePrice = 9.0
					case "seedling":
						basePrice = 1.5
					}
					simPrice := basePrice + (rand.Float64()-0.5)*2
					data[i] = fmt.Sprintf("%.1f", simPrice)
					recentPrices = append(recentPrices, simPrice)
				}
			}
			// 为明天和后天生成预测价格（基于最近7天的平均价格）
			if len(recentPrices) > 0 {
				sum := 0.0
				for _, price := range recentPrices {
					sum += price
				}
				avgPrice := sum / float64(len(recentPrices))
				// 明天的预测（基于7天均值的±5%波动）
				tomorrowPredict := avgPrice * (1 + (rand.Float64()-0.5)*0.1)
				data[7] = fmt.Sprintf("%.1f", tomorrowPredict)
				// 后天的预测（基于明天预测价格的±7.5%波动）
				dayAfterTomorrowPredict := tomorrowPredict * (1 + (rand.Float64()-0.5)*0.15)
				data[8] = fmt.Sprintf("%.1f", dayAfterTomorrowPredict)
			} else {
				data[7] = "2.0" // 默认预测值
				data[8] = "2.1" // 默认预测值
			}

			// 将英文分类名转换为中文
			categoryCN := category
			switch category {
			case "grain":
				categoryCN = "粮食"
			case "vegetable":
				categoryCN = "蔬菜"
			case "fruit":
				categoryCN = "水果"
			case "seedling":
				categoryCN = "种苗"
			}
			categories = append(categories, respond.MarketCategoryData{
				Name:  categoryCN,
				Data:  data,
				Color: categoryColors[category],
			})
		}
	}

	// 构建产品数据
	var products []respond.MarketProductData
	for _, product := range productMap {
		products = append(products, *product)
	}

	// 如果没有产品数据，生成模拟数据
	if len(products) == 0 {
		products = []respond.MarketProductData{
			{
				Name:   "红富士苹果 (特级)",
				Price:  fmt.Sprintf("%.2f", 14+rand.Float64()),
				Change: fmt.Sprintf("%.1f", rand.Float64()*4-2),
				Data:   []string{"13.2", "13.5", "14.2", "13.8", "14.5", "14.1", "14.3", "14.0", "14.2"},
			},
			{
				Name:   "有机胡萝卜",
				Price:  fmt.Sprintf("%.2f", 5.8+rand.Float64()),
				Change: fmt.Sprintf("%.1f", rand.Float64()*4-2),
				Data:   []string{"5.2", "5.5", "6.2", "5.8", "6.0", "5.9", "6.1", "6.0", "6.2"},
			},
			{
				Name:   "优质大米",
				Price:  fmt.Sprintf("%.2f", 4.5+rand.Float64()),
				Change: fmt.Sprintf("%.1f", rand.Float64()*4-2),
				Data:   []string{"4.2", "4.3", "4.8", "4.6", "4.9", "4.7", "4.4", "4.6", "4.5"},
			},
		}
	}
	return "获取成功", &respond.MarketDataRespond{
		Dates:      dates,
		Categories: categories,
		Products:   products,
	}, 0
}
