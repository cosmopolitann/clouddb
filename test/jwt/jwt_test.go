package jwt

import (
	"encoding/base64"
	"testing"
	"time"

	"github.com/dgrijalva/jwt-go"
)

/*
token, err := utils.GenerateToken(
		user.StariverUserId,
		user.NickName,
		user.Mobile,
		user.Email,
		"", 30*24*60*60)
*/
type LoginClaims struct {
	UserId string
	jwt.StandardClaims
}

const (
	tokenStr = "adsfa#^$%#$fgrf" //houxu fengzhuang dao nacos
)

func GenerateToken(userId string, expireDuration int64) (string, error) {
	// 将 uid，用户角色， 过期时间作为数据写入 token 中
	calim := LoginClaims{
		UserId:         userId,
		StandardClaims: jwt.StandardClaims{},
	}
	if expireDuration != -1 {
		calim.StandardClaims = jwt.StandardClaims{
			ExpiresAt: time.Now().Unix() + expireDuration,
		}
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, calim)
	strBase, _ := base64.URLEncoding.DecodeString(tokenStr)

	return token.SignedString(strBase)
}

//func ParseToken(strGen string) (*jwt.Token, error) {
//	strBase, _ := base64.URLEncoding.DecodeString(tokenStr)
//	return jwt.Parse(strGen, func(*jwt.Token) (interface{}, error) {
//		return strBase, nil
//	})
//}
func TestJwt(t *testing.T) {
	//token,err:=GenerateToken("10001",30*24*60*60)
	token, err := GenerateToken("123", 60*60*60)

	if err != nil {
		t.Log("jwt is failed.")
	}
	t.Log("Token = ", token)

}

//eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySWQiOiI0MTY5ODQ1NDUwNjIwMzEzNjAiLCJleHAiOjE2MjYzNTUxMTl9.Ko9C6ojPzShQ3BSP_ASa602EUjD27trRO_11zaV4hCY
