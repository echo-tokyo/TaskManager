package handlers

import (
	"net/http"

	echo "github.com/labstack/echo/v4"

    coreServices "github.com/echo-tokyo/TaskManager/core/services"
)


// эндпоинт без авторизации
func Test(context echo.Context) error {
	return context.String(http.StatusOK, "Test endpoint")

}

// эндпоинт с авторизацией
func Auth(context echo.Context) error {
	userID, err := coreServices.GetUserIDFromRequest(context)
	if err != nil {
		return err
	}
	return context.JSON(http.StatusOK, map[string]int{"user_id": userID})
}


// эндпоинт с проверкой refresh токена
func RefreshTest(context echo.Context) error {
	return context.String(http.StatusOK, "Refresh test endpoint")
}
