package main

import (
	pb "github.com/ChenHanZhang/microservices-in-golang-proto/user"
	"github.com/dgrijalva/jwt-go"
	"time"
)

var (
	key = []byte("mySuperSecretKey")
)

// CustomClaims 是一个自定义的元数据，它的哈希值会被当作JWT的第二个部分
type CustomClaims struct {
	User *pb.User
	jwt.StandardClaims
}

type Authable interface {
	Decode(token string) (*CustomClaims, error)
	Encode(user *pb.User) (string, error)
}

type TokenService struct {
	repo Repostory
}

func (s *TokenService) Decode(tokenString string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (i interface{}, e error) {
		return key, nil
	})

	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, err
	}
}

func (s *TokenService) Encode(user *pb.User) (string, error)  {

	expireToken := time.Now().Add(time.Hour * 72).Unix()
	// 生成 claims
	claims := CustomClaims{
		user,
		jwt.StandardClaims{
			ExpiresAt: expireToken,
			Issuer:    "go.micro.srv.user",
		},
	}
	// 生成 token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// 返回签名后的 token
	return token.SignedString(key)
}