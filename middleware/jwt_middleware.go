package middleware

import (
	emw "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

func JWTMiddleware(secret string) echo.MiddlewareFunc {
	return emw.WithConfig(emw.Config{
		SigningKey:    []byte(secret),
		SigningMethod: "HS256",
	})
}
