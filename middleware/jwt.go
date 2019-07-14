package middleware

import (
	"time"
	"net/http"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

var key []byte = []byte("123r12r12r1e")

func CreateToken() (result string,err error) {
	claims := &jwt.StandardClaims{
		NotBefore: int64(time.Now().Unix()),
		ExpiresAt: int64(time.Now().Unix() + 1000),
		Issuer:    "Bitch",
	}
 
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	result, err = token.SignedString(key)
	return
}

func CheckToken(c *gin.Context)  {
	_, err := jwt.Parse(c.Request.Header.Get("Token"), func(*jwt.Token) (interface{}, error) {
		return key, nil
	})
	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
	}
}