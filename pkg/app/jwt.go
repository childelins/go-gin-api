package app

import (
	"time"

	"github.com/dgrijalva/jwt-go"

	"github.com/childelins/go-gin-api/global"
)

type Claims struct {
	CustomClaims
	jwt.StandardClaims
}

type CustomClaims struct {
	C int
	U int
}

func getSignatureKey() []byte {
	return []byte(global.ServerConfig.JWTInfo.SignatureKey)
}

func GenerateToken(customClaims CustomClaims) (string, error) {
	now := time.Now()
	claims := Claims{
		CustomClaims: customClaims,
		StandardClaims: jwt.StandardClaims{
			NotBefore: now.Unix(),                                         // 签名的生效时间
			ExpiresAt: now.Unix() + global.ServerConfig.JWTInfo.ExpiresAt, // 过期时间
			Issuer:    global.ServerConfig.JWTInfo.Issuer,                 // 颁发公司
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(getSignatureKey())
}

func ParseToken(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return getSignatureKey(), nil
	})
	if err != nil {
		return nil, err
	}
	if token != nil {
		claims, ok := token.Claims.(*Claims)
		if ok && token.Valid {
			return claims, nil
		}
	}

	return nil, err
}
