package service

import (
	"github.com/gin-gonic/gin"
	"tea/model"
	"tea/serializer"
)

type ProductAddRequest struct {
	Name   string `form:"name" json:"name"`
	Number string `form:"number" json:"number"`
	Unit   string `form:"unit" json:"unit"`
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
