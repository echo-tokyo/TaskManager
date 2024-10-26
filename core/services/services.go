package services

import (
	jwt "github.com/golang-jwt/jwt/v5"
	echo "github.com/labstack/echo/v4"

    coreErrors "github.com/echo-tokyo/TaskManager/core/errors"
)


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
