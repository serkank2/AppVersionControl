package jwt

import (
	"fmt"
	"github.com/golang-jwt/jwt"
	"time"
)

func GetTime() []byte {
	return []byte(time.Now().String())
}

func CreateJWT() string {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()
	tokenStr, err := token.SignedString(GetTime())
	if err != nil {
		fmt.Println(err)
		return ""
	}
	return tokenStr
}

var LoginState = map[string]string{}

func SetLoginState(username string) string {
	LoginState[username] = fmt.Sprintf(CreateJWT())
	return LoginState[username]
}
func GetLoginState(key string) string {
	return LoginState[key]

}
