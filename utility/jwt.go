package utility

import (
	"fmt"
	"github.com/golang-jwt/jwt"
	"log"
	"questionnaire/config"
	"time"
)

type JwtData struct {
	ID string `json:"id"` //包含用户名称
	jwt.StandardClaims
}

// GenerateStandardJwt 生成标准的jwt
func GenerateStandardJwt(jwtData *JwtData) string {
	claims := jwtData

	// 为jwt加入标准claims
	claims.StandardClaims = jwt.StandardClaims{
		ExpiresAt: time.Now().Add(time.Duration(time.Duration(config.Config.Jwt.Expires) * time.Hour)).Unix(), // token过期时间 单位：小时
		Issuer:    config.Config.Jwt.Issuer,                                                                   // 签发人是谁
	}

	// 使用SHA256算法加密token生成签名，其中Claims以参数传入
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// 通过配置文件中的密钥对token进行签名，并获得该token的字符串形式
	token, err := tokenClaims.SignedString([]byte(config.Config.Jwt.Secret))

	// 如果token签名过程报错，则报错退出
	if err != nil {
		log.Fatalln("Jwt错误", err)
		panic(err)
	}

	return token
}

// ParseToken 解析jwt，并返回其中的ID
func ParseToken(token string) (string, error) {
	jwtSecret := []byte(config.Config.Jwt.Secret) // 获取jwt的secret

	// 通过jwt.ParseWithClaims()解码jwt，同时使用func(token *jwt.Token) (interface{}, error)函数对其中
	// 的载荷进行处理，返回载荷，所以传入除了token本身的function参数是nil
	// 如果解码不出来，则返回的token == nil，error != nil；否则，则返回错误信息nil
	tokenClaims, err := jwt.ParseWithClaims(token, &JwtData{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil // 返回的是用于结构化的fn，其中密钥==config.secret
	})

	// 如果能解析出jwt...
	if tokenClaims != nil {
		// 尝试将Claims转换为自己定义的JwtData类型
		if claims, ok := tokenClaims.Claims.(*JwtData); ok && tokenClaims.Valid {
			fmt.Println("解析出的ID是：", claims.ID)
			return claims.ID, err
		}
	}

	// 如果无法解析，则返回一个空值和错误信息err
	return "", err
}
