package app_jwt

import (
	"github.com/dgrijalva/jwt-go"
	"log"
	"simple-video-server/config"
	"simple-video-server/pkg/business_code"
	"time"
)

type appJwt struct {
}

var AppJwt = &appJwt{}

type JwtPayload struct {
	jwt.StandardClaims
	UID uint
}

// func (j *appJwt) Create(claims *JwtPayload) (signedString string, err error) {
func (j *appJwt) Create(uid uint) (signedString string, err error) {
	now := time.Now()
	//expiresAt := now + config.Jwt.EffectiveTime
	//var n = 60
	// TODO: time.Second不能直接乘变量int
	// 参考: https://blog.csdn.net/u010389253/article/details/106693715
	//second :=

	expiresAt := now.Add(time.Second * time.Duration(config.Jwt.EffectiveTime))

	claims := &JwtPayload{
		StandardClaims: jwt.StandardClaims{
			//ExpiresAt: expiresAt,
			ExpiresAt: expiresAt.Unix(),
			//ExpiresAt: config.Jwt.EffectiveTime,
			//IssuedAt:  now,
		},
		UID: uid,
	}

	bytes := []byte(config.Jwt.SecretKey)

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedString, err = jwtToken.SignedString(bytes)

	return
}

func (j *appJwt) Parse(token string) (claims *JwtPayload, err error) {
	jwtToken, err := jwt.ParseWithClaims(token, &JwtPayload{}, func(token *jwt.Token) (interface{}, error) {

		return []byte(config.Jwt.SecretKey), nil
	})

	if err != nil {
		if validationError, ok := err.(*jwt.ValidationError); ok {
			/*
				当validationError中的错误信息由错误的token结构引起时，
				**************************************************
				源码vErr.Errors |= ValidationErrorExpired，
				或运算，只有都为0才为0，0000 0000|0000 0101 = 0000 0101
				由于vErr.Errors的初始值为0，所以等价于将ValidationErrorMalformed赋值给validationError的Errors，
				*****************************************************
				如果没有赋值，Errors的初始值为0，那么validationError.Errors&jwt.ValidationErrorMalformed = 0，
				赋值后造成validationError.Errors不为0，那么validationError.Errors&jwt.ValidationErrorMalformed != 0
			*/
			if validationError.Errors&jwt.ValidationErrorMalformed != 0 {
				err = business_code.TokenMalformed
			} else if validationError.Errors&jwt.ValidationErrorExpired != 0 {
				err = business_code.TokenExpired
			} else if validationError.Errors&jwt.ValidationErrorNotValidYet != 0 {
				err = business_code.TokenNotValidYet
			} else {
				err = business_code.TokenInvalid
			}

		}

		return nil, err
	}

	log.Println("Parse")

	claims, ok := jwtToken.Claims.(*JwtPayload)
	if !ok && jwtToken.Valid {

		return nil, business_code.TokenInvalid
	}

	return
}
