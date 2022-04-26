package main

import (
	"Todo_back_end/jwt"
	block_route "Todo_back_end/route/block"
	user_route "Todo_back_end/route/user"

	"github.com/gin-gonic/gin"
)

type Bind struct {
    Name string `json:"name" binding:"required"`
}

func main() {
    r := gin.Default()
	user := r.Group("/user")
	{
		user.POST("login",user_route.Login)
		user.POST("create",user_route.Create)
		user.Use(jwt.Check)
		user.PUT("edit",user_route.Edit)
		user.DELETE("delete",user_route.Delete)
	}
	r.Use(jwt.Check)
	block := r.Group("/block")
	{
		block.GET("get",block_route.Get)
		block.POST("create",block_route.Create)
		block.PUT("edit",block_route.Edit)
		block.DELETE("delete")
	}
	r.Use(func(ctx *gin.Context) {
		ctx.JSON(404,gin.H{
			"code":404,
		})
	})
	r.Run(":80")
}