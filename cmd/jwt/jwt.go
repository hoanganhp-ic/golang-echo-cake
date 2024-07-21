package jwt

import (
	"fmt"
	"net/http"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

type (
	JWTConfig struct {
		Skipper    Skipper
		SigningKey interface{}
	}

	Skipper      func(c echo.Context) bool
	jwtExtractor func(echo.Context) (string, error)
)

var (
	ErrJWTMissing = echo.NewHTTPError(http.StatusUnauthorized, "missing or malformed jwt")
	ErrJWTInvalid = echo.NewHTTPError(http.StatusForbidden, "invalid or expired jwt")
)

func JWT(key interface{}) echo.MiddlewareFunc {
	c := JWTConfig{}
	c.SigningKey = key
	return JWTWithConfig(c)
}

func JWTWithConfig(config JWTConfig) echo.MiddlewareFunc {
	extractor := jwtHeader("Authorization", "Token")
	return func(hf echo.HandlerFunc) echo.HandlerFunc {
		return func(ctx echo.Context) error {
			auth, err := extractor(ctx)
			if err != nil {
				if config.Skipper != nil {
					if config.Skipper(ctx) {
						return hf(ctx)
					}
				}
				return ctx.JSON(http.StatusUnauthorized, "Unauthorized")
			}
			token, err := jwt.Parse(auth, func(t *jwt.Token) (interface{}, error) {
				if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
				}
				return config.SigningKey, nil
			})
			if err != nil {
				return ctx.JSON(http.StatusForbidden, "Forbidden")
			}
			if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
				userID := uint(claims["user"].(float64))
				ctx.Set("user", userID)
				return hf(ctx)
			} else {
				return ctx.JSON(http.StatusForbidden, "Forbidden")
			}
		}
	}
}

func jwtHeader(header string, authScheme string) jwtExtractor {
	return func(ctx echo.Context) (string, error) {
		auth := ctx.Request().Header.Get(header)
		l := len(authScheme)
		if len(auth) > l+1 && auth[:l] == authScheme {
			return auth[l+1:], nil
		}
		return "", ErrJWTMissing
	}
}
