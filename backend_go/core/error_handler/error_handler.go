package error_handler

import (
	"net/http"
	"time"

	echo "github.com/labstack/echo/v4"

	"github.com/echo-tokyo/TaskManager/settings"
)


// настройка обработчика ошибок
func CustomErrorHandler(echoApp *echo.Echo) {
	echoApp.HTTPErrorHandler = func(err error, context echo.Context) {
		// проверка, является ли ошибка err структурой *echo.HTTPError (приведение типов)
		httpError, ok := err.(*echo.HTTPError)
		if !ok {
			httpError = &echo.HTTPError{
				Code: http.StatusInternalServerError,
				Message: map[string]string{"unknown": err.Error()},
			}
		}
		// если пришла *echo.HTTPError ошибка со строкой в качестве httpError.Message
		stringErrorMessage, ok := (httpError.Message).(string)
		if ok {
			httpError.Message = map[string]string{"unknown": stringErrorMessage}
		}

		// поле timestamp
		strTime := time.Now().Format(settings.TimeFmt)
		// поле path
		requestPath := context.Path()

		errMessage := map[string]interface{}{
			"status": "error",
			"statusCode": httpError.Code,
			"path": requestPath,
			"timestamp": strTime,
			"errors": httpError.Message,
		}

		// отправка ответа
		var respErr error
		if !context.Response().Committed {
			// если метод HEAD
			if context.Request().Method == http.MethodHead {
				respErr = context.NoContent(httpError.Code)
			} else {
				respErr = context.JSON(httpError.Code, errMessage)
			}

			if respErr != nil {
				context.Echo().Logger.Error(respErr)
			}
		}
	}
}
