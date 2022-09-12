package controller

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"net/http"
	"strings"
	"time"
	"wangStoreServer/app/admin/models"
	"wangStoreServer/app/admin/service"
	"wangStoreServer/config"
)

type Auth struct {
	Username string
	Password string
	jwt.RegisteredClaims
}

func (a Auth) SetAuthToken(c *gin.Context, username string, password string) string {
	jwtObj := config.GetConfigEnv().Jwt
	a.Username = username
	a.Password = password
	a.RegisteredClaims = jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Second * time.Duration(jwtObj.Timeout))),
	}
	secret := []byte(jwtObj.Secret)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, a)
	tokenString, err := token.SignedString(secret)
	if err != nil {
		ErrResponseBody(c, http.StatusInternalServerError, "error", err.Error())
		return ""
	}
	return tokenString
}

func (a Auth) ParseAuthToken(tokenString string) (*Auth, error) {
	tokenObj, err := jwt.ParseWithClaims(tokenString, &a, keyFun())
	if err != nil {
		if e, ok := err.(*jwt.ValidationError); ok {
			if e.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, errors.New("token 格式不正确")
			} else if e.Errors&jwt.ValidationErrorExpired != 0 {
				return nil, errors.New("token 已过期")
			} else if e.Errors&jwt.ValidationErrorUnverifiable != 0 {
				return nil, errors.New("token 无法验证")
			} else if e.Errors&jwt.ValidationErrorClaimsInvalid != 0 {
				return nil, errors.New("token 签名无效")
			} else if e.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, errors.New("token 尚未生效")
			} else {
				return nil, errors.New("无法处理此 token")
			}
		}
	}
	if myClaims, ok := tokenObj.Claims.(*Auth); ok && tokenObj.Valid {
		return myClaims, nil
	}
	return nil, errors.New("无法处理此 token")
}

func keyFun() jwt.Keyfunc {
	return func(token *jwt.Token) (interface{}, error) {
		return []byte(config.GetConfigEnv().Jwt.Secret), nil
	}
}

func (a Auth) ValidAuthToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		authorization := c.GetHeader("Authorization")
		if authorization == "" || !strings.HasPrefix(authorization, "Bearer ") {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"msg":    "权限不足",
				"status": "fail",
			})
			return
		}
		token := authorization[7:]
		claims, err := a.ParseAuthToken(token)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"msg":    err.Error(),
				"status": "fail",
			})
			return
		}
		var adminUser = models.AdminUser{
			Password: claims.Password,
			Username: claims.Username,
		}
		if !service.GetAdminUserInfoSer(&adminUser, 0, "") {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"msg":    "权限不足",
				"status": "fail",
			})
			return
		}
		c.Next()
	}
}
