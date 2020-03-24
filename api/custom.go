package api

import (
	"goApi/repositories"
	"strconv"

	"github.com/gin-gonic/gin"
)

// CustomName 通过id获得name
// func CustomName(c *gin.Context) {
// 	var service service.UserRegisterService
// 	if err := c.ShouldBind(&service); err == nil {
// 		res := service.Register()
// 		c.JSON(200, res)
// 	} else {
// 		c.JSON(200, ErrorResponse(err))
// 	}
// }

// CreateCustom 创建普通用户输入id,name
func CreateCustom(c *gin.Context) {
	// 转换为int值
	custid, err := strconv.Atoi(c.PostForm("id"))
	if err != nil {
		// handle the error in some way
		c.JSON(500, err)
	}
	name := c.PostForm("name")

	// 调用业务层接口，实体化接口
	var repository repositories.CustomInterface
	// 填入用户业务层需要用到的参数，实体化
	customrepository := repositories.CustomRepository{
		ID: custid, CustomName: name,
	}
	// 将实力化参数存储到接口中
	repository = &customrepository
	// 调用接口的方法
	repository.CreateCustoms()
	c.JSON(200, "res")

}
