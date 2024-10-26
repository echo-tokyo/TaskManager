package services

import (
	jwt "github.com/golang-jwt/jwt/v5"
	echo "github.com/labstack/echo/v4"

    "github.com/echo-tokyo/TaskManager/core/db/models"
    
    coreDB "github.com/echo-tokyo/TaskManager/core/db"
    coreErrors "github.com/echo-tokyo/TaskManager/core/errors"
)


// получение ID юзера из запроса
func GetUserIDFromRequest(context echo.Context) (int, error) {
	var contextUserID int

	// достаём map значений JWT-токена из контекста context
    token, ok := context.Get("user").(*jwt.Token)
    if !ok {
        return contextUserID, coreErrors.GetTokenUserIdError
    }
    tokenClaims, ok := token.Claims.(jwt.MapClaims)
    if !ok {
        return contextUserID, coreErrors.GetTokenUserIdError
    }

    // приведение значения id юзера к float64
    floatContextUserID, ok := tokenClaims["user_id"].(float64)
    if !ok {
        return contextUserID, coreErrors.GetTokenUserIdError
    }
    contextUserID = int(floatContextUserID)

    return contextUserID, nil
}


// получение записи юзера в БД из запроса
func GetUserFromRequest(context echo.Context, user *models.User) (error) {
    // достаём map значений JWT-токена из контекста context
    userID, err := GetUserIDFromRequest(context)
    if err != nil {
        return err
    }

    dbConnect, err := coreDB.GetConnection()
    if err != nil {
        return coreErrors.DBConnectError
    }

    findResult := dbConnect.Where("id = ?", userID).First(user)
    // если юзер с таким ID не найден
    if err = findResult.Error; err != nil {
        return err
    }

    return nil
}
