package block

import (
	"Todo_back_end/DB/structs"
	Handling "Todo_back_end/api"
	"Todo_back_end/jwt"
	"Todo_back_end/route/block/get"
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

var all Handling.Handling

func Create(ctx *gin.Context) {
	type reqd struct {
		Type int `json:"type" binding:"required"`
		Data string `json:"data" binding:"required"`
		Top_id string `json:"top_id"`
		Final_Date string `json:"final_date" binding:"required"`
	}	
	token,_:= ctx.Cookie("Token")
	token_data :=jwt.Get().ValidateToken(token)

	req := &reqd{}
	if ctx.ShouldBind(req) != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"code" : -1,
		})
		return
	}
	top := sql.NullString{req.Top_id,true}
	if req.Top_id == "" {
		top.Valid = false
	}
	code,block := all.Block.Create(structs.Block{Typed: req.Type,Data: req.Data,Top_id: top,Final_Date: req.Final_Date,User_id: token_data.User_id})
	if code == 0 {
		ctx.JSON(200,gin.H{
			"code" : 0,
		})
	} else {
		ctx.JSON(200,gin.H{
			"code":200,
			"block":block,
		})
	}
}
func Get(ctx *gin.Context) {
	token,_:= ctx.Cookie("Token")
	token_data :=jwt.Get().ValidateToken(token)
	var Dan []get.Blocks
	for i,data := range all.Block.Get.Custom(token_data.User_id,"type = ?","1") {
		Dan = append(Dan, *get.Change(&data))
		get.Create(token_data.User_id,&Dan[i])
	}
	ctx.JSON(200,gin.H{
		"code" : 200,
		"data" : Dan,
	})
}
func Edit(ctx *gin.Context) {
	type reqd struct {
		Id string `json:"id" binding:"required"`
		Tpye int `json:"type" binding:"required"`
		Data string `json:"data" binding:"required"`
	}
	req := &reqd{}
	if ctx.ShouldBind(req) != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"code" : -1,
		})
		return
	}

	token,_:= ctx.Cookie("Token")
	token_data :=jwt.Get().ValidateToken(token)	
	fmt.Println(token_data.User_id)
	if 0 != all.Block.Edit(token_data.User_id,req.Id,req.Tpye,req.Data) {
		ctx.JSON(200,gin.H{
			"code":200,
		})
	}else {
		ctx.JSON(200,gin.H{
			"code":0,
		})
	}
}
func Delete(ctx *gin.Context) {
	type reqd struct {
		Id string `json:"id" binding:"required"`
	}
	req := &reqd{}
	if ctx.ShouldBind(req) != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"code" : -1,
		})
		return
	}
	token,_:= ctx.Cookie("Token")
	token_data :=jwt.Get().ValidateToken(token)		
	if all.Block.Delete(token_data.User_id,req.Id) != 0 {
		ctx.JSON(200,gin.H{
			"code":200,
		})	
	}else {
		ctx.JSON(200,gin.H{
			"code":0,
		})
	}
}