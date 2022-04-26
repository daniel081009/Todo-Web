package utils

import (
	"math/rand"
	"time"

	"github.com/gin-gonic/gin"
)
const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
func RandStringBytes(n int) string {
    b := make([]byte, n)
    for i := range b {
		rand.Seed(time.Now().UnixNano())
        b[i] = letterBytes[rand.Intn(len(letterBytes))]
    }
    return string(b)
}
func ErrCheck(err interface{},SendMsg string,ctx *gin.Context) int {
   if err != nil {
      ctx.JSON(200,gin.H{
          "code" : 0,
          "msg" : SendMsg,
      }) 
      ctx.Abort()
      return 0
   } 
   return 1
}