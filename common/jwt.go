package common

import (
	"SimpleBlog/model"
	"github.com/dgrijalva/jwt-go"
	"time"
)

var jwtKey = []byte("a_secret_key")

type Claims struct{
	UserId uint
	jwt.StandardClaims
}

//ReleaseToken 生成Token

func RleaseToken(user model.User) (string,error){
	expirationTime := time.Now().Add(7 *  24 * time.Hour)
	claims := &Claims{
		UserId: user.ID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
			IssuedAt: time.Now().Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,claims)
	tokenString, err:=token.SignedString(jwtKey)
	if err != nil {
		return "",err
	}

	return tokenString,nil
}
//用于解析Token

func ParseToken(tokenString string)(*jwt.Token,*Claims,error){
	claims := &Claims{}
	token,err:= jwt.ParseWithClaims(tokenString,claims,func(token *jwt.Token)(i interface{},err error){
	 	return jwtKey,nil
	})
	return token,claims,err
}
