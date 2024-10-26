package errors

import (
	echo "github.com/labstack/echo/v4"
)


// ошибка подключения к БД
var DBConnectError *echo.HTTPError = echo.NewHTTPError(500, map[string]string{"dbConnect": "Failed to connect to DB"})

// ошибка парсинга токена (его нет или он неверный)
var InvalidTokenError *echo.HTTPError = echo.NewHTTPError(401, map[string]string{"token": "JWT token missing or invalid"})

// ошибка при получении map'а с информацией из токена
var GetTokenClaimsError *echo.HTTPError = echo.NewHTTPError(400, map[string]string{"token": "Failed to get claims from token"})

// ошибка соответствия токена типу (access/refresh)
var TokenTypeMatchingError *echo.HTTPError = echo.NewHTTPError(401, map[string]string{"token": "Invalid JWT token"})

// ошибка при получении типа токена
var GetTokenTypeError *echo.HTTPError = echo.NewHTTPError(401, map[string]string{"token": "Invalid payload of JWT token"})

// ошибка при получении id юзера из токена
var GetTokenUserIdError *echo.HTTPError = echo.NewHTTPError(400, map[string]string{"token": "Failed to get user id from token"})
