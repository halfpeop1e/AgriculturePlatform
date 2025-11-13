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
	if res := dao.GormDB.Where("name = ? AND saler_id = ?", req.Name, req.SalerId).First(&newProduct); res.Error == nil {
		return "商品已存在", -1
	}
	allImages := util.FileReceiverUtil(req.Images, util.IsFileType)
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
	newProduct.CreatAt = time.Now()
	newProduct.UpdateAt = time.Now()
	dao.GormDB.Create(&newProduct)
	return "提交成功", 0
}

func (p *productServer) GetProductList() (string, *respond.ProductListRespond, int) {
	var products []model.Product
	if res := dao.GormDB.Where("status = ?", "1").Find(&products); res.Error != nil {
		log.Printf("Database error: %v", res.Error)
		return "查询失败", nil, -1
	}
	productList := &respond.ProductListRespond{}
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
	return "获取成功", productList, 0
}

func (p *productServer) BuyProduct(req request.BuyProductRequest, user_id string) (string, int) {
	var product model.Product
	if res := dao.GormDB.Where("id = ?", req.ProductId).First(&product); res.Error != nil {
		log.Printf("Database error: %v", res.Error)
		return "查询失败", -1
	}
	if product.Stock < int64(req.Quantity) {
		return "库存不足", -1
	}
	var buyUser model.User
	if res := dao.GormDB.Where("uuid = ?", user_id).First(&buyUser); res.Error != nil {
		log.Printf("Database error: %v", res.Error)
		return "用户查询失败", -1
	}
	order := model.Order{
		Uuid:       util.GenerateUUID(),
		Name:       product.Name,
		Quantity:   int64(req.Quantity),
		Totalprice: float64(req.Totalprice),
		Status:     "已完成",
		Ordertype:  "buy",
		ProductId:  product.Id,
		UserId:     user_id,
		CreatAt:    time.Now(),
	}
	dao.GormDB.Create(&order)
	order = model.Order{
		Uuid:       util.GenerateUUID(),
		Name:       product.Name,
		Quantity:   int64(req.Quantity),
		Totalprice: float64(req.Totalprice),
		Status:     "已完成",
		Ordertype:  "sell",
		ProductId:  product.Id,
		UserId:     product.SalerId,
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
		}
		orderList.Orders = append(orderList.Orders, newOrder)
	}
	return "获取成功", 0, orderList
}
