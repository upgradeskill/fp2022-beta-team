package web

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

// ResponseFormatter returning formatted JSON response
func ResponseFormatter(ctx echo.Context, code int, message string, body interface{}, err error) error {
	var response map[string]interface{}

	if err != nil {
		response = map[string]interface{}{
			"message": message,
			"data":    nil,
			"error":   err.Error(),
		}
	} else {
		response = map[string]interface{}{
			"message": message,
			"data":    body,
			"error":   nil,
		}
	}

	return ctx.JSON(code, response)
}

// ResponseErrValidation returning formatted JSON response
func ResponseErrValidation(ctx echo.Context, message string, errMap map[string]interface{}) error {

	var b strings.Builder
	for k, v := range errMap {
		b.WriteString(fmt.Sprintf("%s : %v, ", k, v))
	}
	errorString := strings.TrimRight(b.String(), ", ")

	response := map[string]interface{}{
		"message":          message,
		"data":             nil,
		"error_validation": errMap,
		"error":            errorString,
	}

	return ctx.JSON(http.StatusBadRequest, response)
}

func ResponseErrValidationWithCode(ctx echo.Context, message string, errMap map[string]interface{}, code int) error {

	var msg string

	if len(errMap) == 0 {
		msg = message
	} else {
		for _, value := range errMap {
			msg = fmt.Sprintf("%s", value)
			break
		}
	}

	response := map[string]interface{}{
		"message":          msg,
		"data":             nil,
		"error_validation": errMap,
	}

	return ctx.JSON(code, response)
}

// ResponseFormatter returning formatted JSON response with meta
func ResponseFormatterWithMeta(ctx echo.Context, code int, message string, body interface{}, meta interface{}, err error) error {
	var response map[string]interface{}

	if err != nil {
		response = map[string]interface{}{
			"message": message,
			"data":    body,
			"error":   err.Error(),
			"meta":    meta,
		}
	} else {
		response = map[string]interface{}{
			"message": message,
			"data":    body,
			"error":   nil,
			"meta":    meta,
		}
	}

	return ctx.JSON(code, response)
}

func Containts(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}
