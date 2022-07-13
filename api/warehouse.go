package api

import (
	"github.com/gin-gonic/gin"
	"tea/service"
)

func WarehouseList(c *gin.Context) {
	var wareHouseListService service.WareHouseListRequst
	if err := c.ShouldBind(&wareHouseListService); err == nil {
		res := wareHouseListService.WareHouseList(c)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}
