/**
 * @Description JWT中间件
 **/
package internal

import (
	"errors"
	"github.com/golang-jwt/jwt"
	"go.uber.org/zap"
	"shershon1991/fund-analye-system/global"
	"shershon1991/fund-analye-system/model/request"
	"time"
)

// 创建Jwt
func CreateToken(uid uint) (string, error) {
	newWithClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, &request.UserClaims{
		StandardClaims: &jwt.StandardClaims{
			ExpiresAt: time.Now().Add(global.GvaConfig.Jwt.Expire).Unix(), // 有效期
			Issuer:    global.GvaConfig.Jwt.Issuer,                        // 签发人
			IssuedAt:  time.Now().Unix(),                                  // 签发时间
		},
		Uid: uid,
	})
	return newWithClaims.SignedString([]byte(global.GvaConfig.Jwt.Secret))
}

// 验证JWT
func ParseToken(tokenString string) (*request.UserClaims, error) {
	var err error
	var token *jwt.Token
	token, err = jwt.ParseWithClaims(tokenString, &request.UserClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(global.GvaConfig.Jwt.Secret), nil
	})
	if err != nil {
		global.GvaLogger.Error("解析JWT失败", zap.String("error", err.Error()))
		return nil, err
	}
	// 断言
	userClaims, ok := token.Claims.(*request.UserClaims)
	// 验证
	if !ok || !token.Valid {
		return nil, errors.New("JWT验证失败")
	}
	return userClaims, nil
}
