package initialize

import (
	"errors"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

// JWT 规定了7个官方字段，提供使用:
// - iss (issuer)：发布者
// - sub (subject)：主题
// - iat (Issued At)：生成签名的时间
// - exp (expiration time)：签名过期时间
// - aud (audience)：观众，相当于接受者
// - nbf (Not Before)：生效时间
// - jti (JWT ID)：编号

const TokenExpireDuration = time.Hour * 6 //过期时间

var myKey = []byte("adb1234556") //加盐值

type MyClaims struct {
	UserID   int64  `json:"user_id"`
	Username string `json:"username"`
	jwt.RegisteredClaims
}

func GenToken(userID int64, username string) (string, error) {
	c := MyClaims{
		UserID:   userID,
		Username: username, //自定义字段
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(TokenExpireDuration)), //过期时间
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    "my-projet", //签发人
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	signingString, err := token.SignedString(myKey)
	return signingString, err
}

// ParseToken 解析JWT
func ParseToken(tokenString string) (*MyClaims, error) {

	// 解析token
	var mc = new(MyClaims)
	token, err := jwt.ParseWithClaims(tokenString, mc, func(token *jwt.Token) (i interface{}, err error) {
		return myKey, nil
	})
	if err != nil {
		return nil, err
	}
	if token.Valid { // 校验token
		return mc, nil
	}
	return nil, errors.New("invalid token")
}

// RefreshToken 更新 Token，用以提供 refresh token 接口
func RefreshToken(mc *MyClaims) (string, error) {
	token, err := GenToken(mc.UserID, mc.Username)
	return token, err
}
