package exception

import (
	"fmt"
	"reflect"

	"github.com/qytela/example-project-layout/internal/pkg/logger"

	"github.com/labstack/echo/v4"
)

type ErrorResponse struct {
	Status         bool        `json:"status"`
	Code           int         `json:"code"`
	Message        string      `json:"message"`
	InvalidRequest interface{} `json:"invalid_request"`
	Exceptions     interface{} `json:"exceptions"`
}

func (e *ErrorResponse) Error() string {
	return fmt.Sprintf("%v", e.Exceptions)
}

func HTTPErrorHandler(err error, c echo.Context) {
	if err.Error() != "<nil>" {
		logger.MakeLogEntry(c).Error(err)
	}

	he, ok := err.(*ErrorResponse)
	if ok {
		reflectKind := reflect.ValueOf(he.Exceptions).Kind()
		if reflectKind == reflect.Ptr {
			statusCode := reflect.ValueOf(he.Exceptions).
				Elem().
				FieldByName("Code").
				Interface().(int)
			c.JSON(statusCode, he.Exceptions)
		} else {
			if he.Code == 0 {
				he.Code = 400
			}

			c.JSON(he.Code, he)
		}
	} else {
		c.JSON(500, err)
	}
}
