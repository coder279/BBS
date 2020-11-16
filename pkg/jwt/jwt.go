package jwt

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"time"
)

var mySecret = []byte("惊雷惊雷")
var TokenExpireDuration = time.Hour * 2
type MyClaims struct {
	UserID int64 `json:"user_id"`
	jwt.StandardClaims
}

func GenToken(userID int64) (aToken, rToken string, err error) {
	// 创建一个我们自己的声明
	c := MyClaims{
		userID, // 自定义字段
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(TokenExpireDuration).Unix(), // 过期时间
			Issuer:    "bluebell",                                 // 签发人
		},
	}
	// 加密并获得完整的编码后的字符串token
	aToken, err = jwt.NewWithClaims(jwt.SigningMethodHS256, c).SignedString(mySecret)

	// refresh token 不需要存任何自定义数据
	rToken, err = jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		ExpiresAt: time.Now().Add(time.Second * 30).Unix(), // 过期时间
		Issuer:    "bluebell",                              // 签发人
	}).SignedString(mySecret)
	// 使用指定的secret签名并获得完整的编码后的字符串token
	return
}

// ParseToken 解析JWT
func ParseToken(tokenString string) (claims *MyClaims, err error) {
	// 解析token
	var token *jwt.Token
	claims = new(MyClaims)
	token, err = jwt.ParseWithClaims(tokenString, claims, keyFunc)
	if err != nil {
		return
	}
	if !token.Valid { // 校验token
		err = errors.New("invalid token")
	}
	return
}

func keyFunc(_ *jwt.Token) (i interface{}, err error) {
	return mySecret, nil
}

// RefreshToken 刷新AccessToken
func RefreshToken(aToken, rToken string) (newAToken, newRToken string, err error) {
	// refresh token无效直接返回
	if _, err = jwt.Parse(rToken, keyFunc); err != nil {
		return
	}

	// 从旧access token中解析出claims数据
	var claims MyClaims
	_, err = jwt.ParseWithClaims(aToken, &claims, keyFunc)
	v, _ := err.(*jwt.ValidationError)

	// 当access token是过期错误 并且 refresh token没有过期时就创建一个新的access token
	if v.Errors == jwt.ValidationErrorExpired {
		return GenToken(claims.UserID)
	}
	return
}


