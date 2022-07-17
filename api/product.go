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
