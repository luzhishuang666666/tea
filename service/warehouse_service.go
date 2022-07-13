package service

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"tea/model"
	"tea/serializer"
)

type WareHouseListRequst struct {
	Page     int `form:"page" json:"page"`
	PageSize int `form:"pageSize" json:"pageSize"`
}

type WareHouseListRespose struct {
	Id         int    `json:"id"`
	Name       string `json:"name"`
	Createtime string `json:"createtime"`
	Remark     string `json:"remark"`
}

func (service *WareHouseListRequst) WareHouseList(c *gin.Context) serializer.PageResponse {
	var count int64
	var warehouse []model.Warehouse

	if err := model.DB.Model(&warehouse).Count(&count).Error; err != nil {
		var res = serializer.PageResponse{
			Code: 500,
			Msg:  string("仓库列表查询失败"),
		}
		return res
	}
	offsets := (service.Page - 1) * service.PageSize
	limits := service.PageSize
	if err := model.DB.Model(&warehouse).Select("id,name,createtime,remark").Limit(limits).Offset(offsets).Find(&warehouse).Error; err != nil {
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
		List:  warehouse,
		Count: count16,
		Msg:   string("sucess"),
	}
	return res
}
