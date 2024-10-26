package middlewares

import (
	echo "github.com/labstack/echo/v4"
    echoJWT "github.com/labstack/echo-jwt/v4"

    jwt "github.com/golang-jwt/jwt/v5"

    coreErrorHandler "github.com/echo-tokyo/TaskManager/core/error_handler"
    coreErrors "github.com/echo-tokyo/TaskManager/core/errors"
    "github.com/echo-tokyo/TaskManager/settings"
)


// middleware для распаковки содержимого токена в содержимое context'а запроса и валидации токена
var ValidateJWTMiddleware echo.MiddlewareFunc = echoJWT.WithConfig(echoJWT.Config{
    SigningKey: []byte(settings.SecretForJWT),
    ErrorHandler: coreErrorHandler.CustomJWTErrorHandler,
})

// проверка соответствия типа токена из контекста типу tokenType
func checkTokenType(context echo.Context, tokenType string) error {
    // достаём map значений JWT-токена из контекста context
    token, ok := context.Get("user").(*jwt.Token)
    if !ok {
        return coreErrors.InvalidTokenError
    }
    tokenClaims, ok := token.Claims.(jwt.MapClaims)
    if !ok {
        return coreErrors.GetTokenClaimsError
    }

    // приведение значения типа токена к string
    contextTokenType, ok := tokenClaims["token_type"].(string)
    if !ok {
        return coreErrors.GetTokenTypeError
    }
    // проверка, что тип токена соответствует tokenType
    if contextTokenType != tokenType {
        return coreErrors.TokenTypeMatchingError
    }

    return nil
}


// middleware для проверки соответствия токена типу access
func TokenIsAccessMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
    return func(context echo.Context) error {
        // проверка типа токена
        if err := checkTokenType(context, "access"); err != nil {
            return err
        }
        return next(context)
    }
}

// middleware для проверки соответствия токена типу refresh
func TokenIsRefreshMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
    return func(context echo.Context) error {
        // проверка типа токена
        if err := checkTokenType(context, "refresh"); err != nil {
            return err
        }
        return next(context)
    }
}
