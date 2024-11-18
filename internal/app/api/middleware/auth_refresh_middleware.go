package middleware

import (
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/qytela/example-project-layout/internal/pkg/auth"
	"github.com/qytela/example-project-layout/internal/pkg/exception"
)

func AuthRefresh(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		clientToken := c.Request().Header.Get("Authorization")
		if clientToken == "" {
			return c.JSON(400, &exception.ErrorResponse{
				Status:  false,
				Code:    400,
				Message: "No Authorization header provided",
			})
		}

		extractedToken := strings.Split(clientToken, "Bearer ")
		if len(extractedToken) == 2 {
			clientToken = strings.TrimSpace(extractedToken[1])
		} else {
			return c.JSON(400, &exception.ErrorResponse{
				Status:  false,
				Code:    400,
				Message: "Incorrect Format of Authorization Token",
			})
		}

		claims, err := auth.ValidateRefreshToken(clientToken)
		if err != nil {
			return c.JSON(401, &exception.ErrorResponse{
				Status:  false,
				Code:    401,
				Message: "Token is Invalid / Expired",
			})
		}

		c.Set("userId", claims.Sub)

		return next(c)
	}
}
