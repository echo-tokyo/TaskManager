package handlers

import (
	// "fmt"
	"net/http"

	echo "github.com/labstack/echo/v4"

)


// эндпоинт без авторизации
func Test(context echo.Context) error {
	return context.String(http.StatusOK, "Test endpoint")

}

// эндпоинт с авторизацией
func Auth(context echo.Context) error {
	return context.String(http.StatusOK, "Auth endpoint")
}


// эндпоинт с проверкой refresh токена
func RefreshTest(context echo.Context) error {
	return context.String(http.StatusOK, "Refresh test endpoint")
}
