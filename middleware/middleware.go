package middleware

import (
	"fmt"
	"log"
	"net/http"
	"nwd/shared/response"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

var JwtSecret = "gogomcall"

type JwtClaims struct {
	jwt.StandardClaims
	Username string `json:"username"` //user name
	Password string `json:"password"` //password
}

//generate jwt
func GenerateJwt(username, password string, timeout int) (string, bool) {
	expiresAt := time.Now().Add(time.Hour * time.Duration(timeout)).Unix()
	claims := JwtClaims{}
	claims.Username = username
	claims.Password = password
	claims.ExpiresAt = expiresAt
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(JwtSecret))
	if err != nil {
		fmt.Printf("GenerateJwt error:%v", err)
		return "", false
	}
	return tokenString, true
}

//valid jwt
func ValidateJwt(tokenString string) (*JwtClaims, bool) {
	token, err := jwt.ParseWithClaims(tokenString, &JwtClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(JwtSecret), nil
	})

	if err != nil {
		log.Println(err)
		return nil, false
	}

	claims, ok := token.Claims.(*JwtClaims)
	if ok && token.Valid {
		return claims, true
	}

	return nil, false
}

//ctx.Abort need return
func Jwt() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		code := http.StatusOK
		authorization := ctx.Request.Header.Get("Authorization")
		if authorization == "" {
			code = http.StatusUnauthorized
			response.Response(ctx.Writer, http.StatusUnauthorized, "Unauthorized", "")
			ctx.Abort()
			return
		}

		strSlice := strings.SplitN(authorization, " ", 2)
		if len(strSlice) != 2 && strSlice[0] != "Bearer" {
			code = http.StatusUnauthorized
			response.Response(ctx.Writer, http.StatusUnauthorized, "Unauthorized", "")
			ctx.Abort()
			return
		}

		claim, ok := ValidateJwt(strSlice[1])
		if !ok {
			fmt.Println("JWT auth failed")
			code = http.StatusUnauthorized
			response.Response(ctx.Writer, http.StatusUnauthorized, "Unauthorized", "")
			ctx.Abort()
			return
		}

		if time.Now().Unix() > claim.ExpiresAt {
			code = http.StatusUnauthorized
			response.Response(ctx.Writer, http.StatusUnauthorized, "Unauthorized", "")
			ctx.Abort()
			return
		}

		//ctx.JSON(http.StatusOK, gin.H{"code": code, "message": "Auth success"})
		fmt.Printf("Auth success. Code = %v", code)

		ctx.Set("username", claim.Username)
		ctx.Next()
	}
}
