package handlers

import (
	"net/http"

	echo "github.com/labstack/echo/v4"

    "github.com/echo-tokyo/TaskManager/app_tasks/serializers"

    "github.com/echo-tokyo/TaskManager/core/db/models"
    coreDB "github.com/echo-tokyo/TaskManager/core/db"
    coreServices "github.com/echo-tokyo/TaskManager/core/services"
    coreErrors "github.com/echo-tokyo/TaskManager/core/errors"
)


// Input:
// (JSON body)
// {
//   "taskId": 3
// }
// (Header: Bearer jwt.access.token)

// Output: {
//   "username": "petr",
//   "fio": "Петров Пётр петрович",
//   "created_at": "2024-10-26T08:47:42.436476+03:00",
//   "job_title": "Руководитель"
// }

// добавление исполнителя к задаче
func AddExecutor(context echo.Context) error {
	var dataIn serializers.AddExecutorIn
	var err error
	var user models.User
	var task models.Task

	dbConnect, err := coreDB.GetConnection()
	if err != nil {
		return coreErrors.DBConnectError
	}

	// получение юзера из запроса
	if err = coreServices.GetUserFromRequest(context, &user); err != nil {
		return err
	}

	// парсинг JSON-body
	if err = context.Bind(&dataIn); err != nil {
		return err
	}
	// валидация полученной структуры
	if err := dataIn.Validate(); err != nil {
		return err
	}

	// получение записи задачи
	if err = dataIn.AddExecutorInGetTaskEntity(dbConnect, &task); err != nil {
		return err
	}

	err = dbConnect.Model(&user).Association("Tasks").Append(&task)
	if err != nil {
		return err
	}

	return context.JSON(http.StatusOK, user)
}
