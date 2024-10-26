package errors

import (
	echo "github.com/labstack/echo/v4"
)


// ошибка нахождения записи задачи
var TaskFoundError *echo.HTTPError = echo.NewHTTPError(400, map[string]string{"db_not_found": "Task with given ID was not found"})
