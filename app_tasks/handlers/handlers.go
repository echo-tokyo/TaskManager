package handlers

import (
	"fmt"
	"net/http"

	echo "github.com/labstack/echo/v4"

    "github.com/echo-tokyo/TaskManager/core/db/models"
    
    coreDB "github.com/echo-tokyo/TaskManager/core/db"
    coreErrors "github.com/echo-tokyo/TaskManager/core/errors"
    // coreServices "github.com/echo-tokyo/TaskManager/core/services"
)


// эндпоинт получения всех задач
func Tasks(context echo.Context) error {
	dbConnect, err := coreDB.GetConnection()
	if err != nil {
		return coreErrors.DBConnectError
	}

	var board models.Board
	// подгружаем поле Tasks с зависимыми задачами
	findResult := dbConnect.Preload("Tasks").First(&board, 1)
	// если юзер с таким ID не найден
	if err = findResult.Error; err != nil {
		return err
	}

	fmt.Println("board:", board)

	return context.JSON(http.StatusOK, board)

}


// эндпоинт получения задач юзера (исполнитель - юзер)
func UserTasks(context echo.Context) error {
	// userID, err := coreServices.GetUserIDFromRequest(context)
	// if err != nil {
	// 	return err
	// }

	dbConnect, err := coreDB.GetConnection()
	if err != nil {
		return coreErrors.DBConnectError
	}

	var tasks []models.Task
	// подгружаем поле Executors с зависимыми исполнителями
	findResult := dbConnect.Preload("Executors").Find(&tasks) //, "id = ?", userID)
	// если юзер с таким ID не найден
	if err = findResult.Error; err != nil {
		return err
	}

	fmt.Println("tasks:", tasks)
	// fmt.Println("requestUser.Tasks:", requestUser.Tasks)

	return context.JSON(http.StatusOK, tasks)
}
