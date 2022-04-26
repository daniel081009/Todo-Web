package jwt

import (
	"Todo_back_end/utils"
	"errors"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type JwtWrapper struct {
	SecretKey       string
	Issuer          string
	ExpirationHours int64
}

type JwtClaim struct {
	User_id string
	jwt.StandardClaims
}
func Get() *JwtWrapper {
	return &JwtWrapper{"superDan","Dan",24}
}

func (j *JwtWrapper) GenerateToken(id string) (string,error) {
	return jwt.NewWithClaims(jwt.SigningMethodHS256, &JwtClaim{
		User_id: id,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(j.ExpirationHours)).Unix(),
			Issuer:    j.Issuer,
		},
	}).SignedString([]byte(j.SecretKey))
}

func (j *JwtWrapper) ValidateToken(Token string) (*JwtClaim) {
	token, err := jwt.ParseWithClaims(
		Token,
		&JwtClaim{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(j.SecretKey), nil
		},
	)
	claims, ok := token.Claims.(*JwtClaim)
	if err != nil || !ok || claims.ExpiresAt < time.Now().Local().Unix() {
		err = errors.New("Couldn't parse claims")
		return nil
	}
	return claims
}

func Check(ctx *gin.Context) {
	token,err := ctx.Cookie("Token")
	utils.ErrCheck(err,"쿠키에 토큰이 없음",ctx)

	token_data := Get().ValidateToken(token)
	if token_data == nil {
		ctx.JSON(200,gin.H{
			"code" : 0,
			"msg" : "토큰이 잘못됨",
		})
		ctx.Abort()
	}
	if token_data == nil {
		ctx.Next()
	} 
}