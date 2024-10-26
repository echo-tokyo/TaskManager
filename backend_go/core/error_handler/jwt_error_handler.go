package error_handler

import (
	"net/http"

	echo "github.com/labstack/echo/v4"
	echojwt "github.com/labstack/echo-jwt/v4"
)


// настройка обработчика ошибок для JWT middleware
func CustomJWTErrorHandler(context echo.Context, err error) error {
	// ошибка валидации токена (кривой/истёк)
	tokenParsingError, ok := err.(*echojwt.TokenParsingError)
	if ok {
		httpError := &echo.HTTPError{
			Code: http.StatusUnauthorized,
			Message: map[string]string{"token": tokenParsingError.Error()},
		}
		return httpError
	}

	// токен не был отправллен в заголовке
	tokenExtractionError, ok := err.(*echojwt.TokenExtractionError)
	if ok {
		httpError := &echo.HTTPError{
			Code: http.StatusBadRequest,
			Message: map[string]string{"token": tokenExtractionError.Error()},
		}
		return httpError
	}

	return err
}
