package middleware

import (
	"bitroom/config"
	"bitroom/types"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
)

// generate new csrf token
func SetCsrfTokenMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// generate
		// token, err := utils.GenerateCSRFToken()
		// if err != nil {
		// 	return echo.NewHTTPError(http.StatusInternalServerError, "could not generate csrf token")
		// }

		// // use jwt as key of csrf value in cache
		// jwt := c.Request().Header.Get("Authorization")
		// if jwt == "" {
		// 	return echo.NewHTTPError(http.StatusUnauthorized, "missing token")
		// }

		// // save in cach
		// ch := utils.GetCache()
		// ch.Add(jwt, token, 60*time.Minute)

		// cookie := new(http.Cookie)
		// cookie.Name = constants.CsrfCookieName
		// cookie.Value = token
		// cookie.Path = "/"
		// cookie.HttpOnly = true
		// cookie.Secure = true
		// cookie.SameSite = http.SameSiteLaxMode
		// cookie.Expires = time.Now().Add(2 * time.Hour)

		// http.SetCookie(c.Response().Writer, cookie)

		// c.Response().Header().Set("X-CSRF-Token", token)

		return next(c)
	}
}

// validate csrf token
func CsrfVerifyMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// check cookie exist
		// cookie, err := c.Cookie(constants.CsrfCookieName)
		// if err != nil {
		// 	return echo.NewHTTPError(http.StatusForbidden, "CSRF token missing")
		// }

		// // use jwt as key of csrf value in cache
		// jwt := c.Request().Header.Get("Authorization")
		// if jwt == "" {
		// 	return echo.NewHTTPError(http.StatusUnauthorized, "missing token")
		// }
		// // get from cach
		// ch := utils.GetCache()
		// csrfToken, exists := ch.Get(jwt)
		// if !exists {
		// 	return echo.NewHTTPError(http.StatusUnauthorized, "token expired")
		// }

		// // validate cookie
		// if csrfToken != cookie {
		// 	return echo.NewHTTPError(http.StatusUnauthorized, "invalid csrf")
		// }
		return next(c)
	}
}

// validating jwt tokens
func JwtMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		token := c.Request().Header.Get("Authorization")
		if token == "" {
			return echo.NewHTTPError(http.StatusUnauthorized, "missing token")
		}

		claims := &types.JwtCustomClaims{}
		tkn, err := jwt.ParseWithClaims(token, claims, func(t *jwt.Token) (interface{}, error) {
			return []byte(config.ServerConfig.JWT_SECRET), nil
		})
		if err != nil || !tkn.Valid {
			return echo.NewHTTPError(http.StatusUnauthorized, "invalid jwt")
		}

		c.Set("user", claims)
		return next(c)
	}
}

// rolebase access middleware
func RoleBaseMiddleware(roles []string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			userClaims, ok := c.Get("user").(*types.JwtCustomClaims)
			if !ok {
				return echo.NewHTTPError(http.StatusUnauthorized, "invalid user claims")
			}

			for _, role := range roles {
				if userClaims.Role == role {
					return next(c)
				}
			}
			return echo.NewHTTPError(http.StatusForbidden, "access denied")
		}
	}
}
