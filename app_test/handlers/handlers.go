package handlers

import (
	"fmt"
	"net/http"

	echo "github.com/labstack/echo/v4"

    
    "github.com/echo-tokyo/TaskManager/core/db/models"
    
    coreDB "github.com/echo-tokyo/TaskManager/core/db"
    coreErrors "github.com/echo-tokyo/TaskManager/core/errors"
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

	dbConnect, err := coreDB.GetConnection()
	if err != nil {
		return coreErrors.DBConnectError
	}

	var requestUser models.User
	findResult := dbConnect.Where("id = ?", userID).First(&requestUser)
	// если юзер с таким ID не найден
	if err = findResult.Error; err != nil {
		return err
	}

	fmt.Println(requestUser)

	return context.JSON(http.StatusOK, requestUser)
}


// эндпоинт с проверкой refresh токена
func RefreshTest(context echo.Context) error {
	return context.String(http.StatusOK, "Refresh test endpoint")
}
