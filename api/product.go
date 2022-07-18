package api

import (
	"github.com/gin-gonic/gin"
	"tea/service"
)

func ProductAdd(c *gin.Context) {
	var productAddService service.ProductAddRequest
	if err := c.ShouldBind(&productAddService); err == nil {
		res := productAddService.ProductAdd(c)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

func SearchByName(c *gin.Context) {
	var SearchByNameService service.SearchByNameRequest
	if err := c.ShouldBind(&SearchByNameService); err == nil {
		res := SearchByNameService.SearchByNameService(c)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

func ProductList(c *gin.Context) {
	var productService service.ProductListRequst
	if err := c.ShouldBind(&productService); err == nil {
		res := productService.ProductList(c)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}
