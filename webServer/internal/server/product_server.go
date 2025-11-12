package gorm

import (
	"encoding/json"
	"go-agriculture/internal/config"
	"go-agriculture/internal/dao"
	"go-agriculture/internal/dto/request"
	"go-agriculture/internal/dto/respond"
	"go-agriculture/internal/model"
	"go-agriculture/internal/util"
	"io"
	"log"
	"os"
	"path/filepath"
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
	allImages := make([]string, len(req.Images))
	for _, file := range req.Images {
		if file != nil {
			beforePath := config.GetConfig().StaticFilePath
			ext := filepath.Ext(file.Filename)
			if ext != ".jpg" && ext != ".jpeg" && ext != ".png" && ext != ".webp" {
				log.Printf("图片格式不支持")
				continue
			}
			if file.Size > 1920*1080*10 {
				log.Printf("图片大小超过限制")
				continue
			}
			filename := util.GenerateUUID() + ext
			file.Filename = filename
			filePath := beforePath + "/" + filename
			// 打开上传的文件
			src, err := file.Open()
			if err != nil {
				continue
			}
			defer src.Close()

			// 创建目标文件
			dst, err := os.Create(filePath)
			if err != nil {
				continue
			}
			defer dst.Close()
			// 保存文件
			if _, err := io.Copy(dst, src); err != nil {
				continue
			}
			savePath := "/static/files/" + filename
			allImages = append(allImages, savePath)
		}
	}
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
		newProduct := respond.ProductRespond{
			Id:       product.Id,
			Name:     product.Name,
			Image:    images,
			Price:    product.Price,
			Stock:    product.Stock,
			Category: product.Category,
			Saler:    product.Saler,
			SalerId:  product.SalerId,
		}
		productList.ProductList = append(productList.ProductList, newProduct)
	}
	return "获取成功", productList, 0
}
