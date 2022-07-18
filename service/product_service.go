package service

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"tea/model"
	"tea/serializer"
)

type ProductAddRequest struct {
	Name   string `form:"name" json:"name"`
	Number string `form:"number" json:"number"`
	Unit   string `form:"unit" json:"unit"`
}

type SearchByNameRequest struct {
	Name string `form:"name" json:"name"`
}

type ProductListRequst struct {
	Page     int `form:"page" json:"page"`
	PageSize int `form:"pageSize" json:"pageSize"`
}

func (service *ProductAddRequest) ProductAdd(c *gin.Context) serializer.NormalResponse {

	var Name = service.Name
	var Number = service.Number
	var Unit = service.Unit

	var productModel []model.Product

	var product = model.Product{
		Name:   Name,
		Number: Number,
		Unit:   Unit,
	}

	if err := model.DB.Model(&productModel).Create(&product).Error; err != nil {
		var res = serializer.NormalResponse{
			Code: 500,
			Msg:  string("仓库新增失败,产品名称是:" + Name),
		}
		return res
	}

	var res = serializer.NormalResponse{
		Code: 200,
		Msg:  string("sucess"),
	}

	return res
}

func (service *SearchByNameRequest) SearchByNameService(c *gin.Context) serializer.NormalResponse {
	var Name = service.Name
	var productModel []model.Product

	if err := model.DB.Model(&productModel).Where("name LIKE ?", "%"+Name+"%").Find(&productModel).Error; err != nil {
		var res = serializer.NormalResponse{
			Code: 500,
			Msg:  string("产品模糊匹配失败，模糊名称是:" + Name),
		}
		return res
	}

	var res = serializer.NormalResponse{
		Code: 200,
		Data: productModel,
		Msg:  string("sucess"),
	}

	return res
}

func (service *ProductListRequst) ProductList(c *gin.Context) serializer.PageResponse {
	var count int64
	var product []model.Product

	if err := model.DB.Model(&product).Count(&count).Error; err != nil {
		var res = serializer.PageResponse{
			Code: 500,
			Msg:  string("仓库列表查询失败"),
		}
		return res
	}
	offsets := (service.Page - 1) * service.PageSize
	limits := service.PageSize
	if err := model.DB.Model(&product).Limit(limits).Offset(offsets).Find(&product).Error; err != nil {
		var res = serializer.PageResponse{
			Code: 500,
			Msg:  string("仓库列表查询失败"),
		}
		return res
	}

	strInt64 := strconv.FormatInt(count, 10)
	count16, _ := strconv.Atoi(strInt64)
	var res = serializer.PageResponse{
		Code:  200,
		List:  product,
		Count: count16,
		Msg:   string("sucess"),
	}
	return res
}
