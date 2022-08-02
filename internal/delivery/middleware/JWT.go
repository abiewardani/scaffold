package middleware

import (
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
)

// JWTClaims data structure for claims
type JWTClaims struct {
	UserID   string `json:"user_id"`
	Username string `json:"user_name"`
	jwt.StandardClaims
}

// JSONFailed ...
type JSONFailed struct {
	Success bool   `json:"success"`
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// VerifyExpiresAt function, will override original VerifyExpiresAt function
func (c *JWTClaims) VerifyExpiresAt(cmp int64, req bool) bool {
	var leeway int64 = 60 // one minutes
	return c.StandardClaims.VerifyExpiresAt(cmp-leeway, req)
}

func VerifyToken(key string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			response := new(JSONFailed)

			tokenStr := strings.Replace(c.Request().Header.Get("Authorization"), "Bearer ", "", -1)
			if len(tokenStr) == 0 {
				response.Success = false
				response.Message = "Not Found Authorization"
				response.Code = http.StatusUnauthorized

				return c.JSON(http.StatusNonAuthoritativeInfo, response)
			}

			claims := jwt.MapClaims{}
			token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
				return []byte(key), nil
			})

			if err != nil {
				return echo.NewHTTPError(http.StatusUnauthorized, err.Error())
			}

			// do something with decoded claims
			if claims["authorised"] == false {
				response.Success = false
				response.Message = "Unauthorized"
				response.Code = http.StatusUnauthorized

				return c.JSON(http.StatusNonAuthoritativeInfo, response)
			}

			if tokenClaims, ok := token.Claims.(*JWTClaims); token.Valid && ok {
				c.Set("UserID", tokenClaims.UserID)
				c.Set("Username", tokenClaims.Username)
			}
			return next(c)
		}
	}
}
