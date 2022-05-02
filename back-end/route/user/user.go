package user_route

import (
	"Todo_back_end/DB/structs"
	Handling "Todo_back_end/api"
	"Todo_back_end/jwt"
	"Todo_back_end/route/block/get"
	"Todo_back_end/utils"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)
var all Handling.Handling

func Login(ctx *gin.Context) {
	type login_req struct {
		Email string `json:"Email" binding:"required"`
		Password string `json:"Password" binding:"required"`
	}	
	req := &login_req{}
	if ctx.ShouldBind(req) != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"code" : -1,
		})
		return
	}

	user := all.User.Login(req.Email,req.Password)
	if  user != nil {
		token ,err:= jwt.Get().GenerateToken(user.Id)
		utils.ErrCheck(err,"토큰 생성 과정에서 문제생김",ctx)
		ctx.SetCookie("Token",token,24,"/","todo.5origin.io",false,false)
		ctx.JSON(200, gin.H{
			"code" : 200,
			"token":token,
		})
	}else {
		ctx.JSON(200, gin.H{
			"code" : 0,
		})
	}
}

func Create(ctx *gin.Context) {
	type user struct {
		Email string `json:"Email" binding:"required"`
		Name string `json:"Name" binding:"required"` 
		Password string `json:"Password" binding:"required"`
	}
	req := &user{}
	if ctx.ShouldBind(req) != nil {
		ctx.JSON(200, gin.H{
			"code" : -1,
		})
		return
	}
	if all.User.Create(structs.User{"Non",req.Name,req.Email,req.Password}) != 0 {
		ctx.JSON(200, gin.H{
			"code" : 200,
		})
	}else {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code" : 0,
		})
	}
}

func Edit(ctx *gin.Context) {
	token,_:= ctx.Cookie("Token")
	token_data :=jwt.Get().ValidateToken(token)
	type reqd struct {
		Typed int `json:"type" binding:"required"`
		Data string `json:"data" binding:"required"`
	}
	req := &reqd{}
	if utils.ErrCheck(ctx.ShouldBind(req),"무언가를 빼먹음",ctx) == 0 {
		return
	}

	if all.User.Edit(req.Typed,req.Data,token_data.User_id) != 0 {
		ctx.JSON(200,gin.H{
			"code" : 200,
		})
	}else {
		ctx.JSON(200,gin.H{
			"code" : 200,
			"msg" : "뭐지?",
		})
	}
}

func Delete(ctx *gin.Context) {
	token,_:= ctx.Cookie("Token")
	token_data :=jwt.Get().ValidateToken(token)
	d := all.User.Delete(token_data.User_id)
	if d == 0 {
		ctx.JSON(200,gin.H{
			"code": 0,
			"msg" : "? 이미 삭제한듯",
		})
		ctx.Abort()
		return
	}
	ctx.JSON(200,gin.H{
		"code": 200,
	})	
}
func Check(ctx *gin.Context) {
	token,_:= ctx.Cookie("Token")
	token_data :=jwt.Get().ValidateToken(token)
	data :=  all.User.Check(token_data.User_id)
	if data != nil {
		var origin []get.Blocks
		for i,datad := range all.Block.Get.Custom(token_data.User_id,"type = ?","1") {
			origin = append(origin, *get.Change(&datad))
			get.Create(token_data.User_id,&origin[i])
		}
		fmt.Println(origin)
		ctx.JSON(200,gin.H{
			"code" : 200,
			"user" : data,
			"block":origin,
		})
	}else {
		ctx.JSON(200,gin.H{
			"code" : 200,
			"msg" : "뭐지?",
		})
	}
}
