package handlers

import (
	"net/http"

	echo "github.com/labstack/echo/v4"

    "github.com/echo-tokyo/TaskManager/app_tasks/serializers"

    "github.com/echo-tokyo/TaskManager/core/db/models"
    coreDB "github.com/echo-tokyo/TaskManager/core/db"
    coreErrors "github.com/echo-tokyo/TaskManager/core/errors"
)


// Input:
// {
//   "date_from": "2024-10-23",
//   "date_to": "2024-10-26"
// }

// Output:
// {
//  "backlog": [
// 	 {
//    "title": "task3",
//    "task_desc": "no",
//    "status": "processing",
//    "board_id": 1,
//    "created_at": "2024-10-23T00:00:00+03:00",
//    "executors": []
//   },
//   ...
//  ],
//  "proccessing": [...],
//  "finished": [...]
// }

// фильтр задач по времени с/по
func DateFilter(context echo.Context) error {
	var filteredTasksByStatus serializers.DateFilteredTasksByStatus
	var dataIn serializers.DateFilterIn
	var err error
	var filteredTasks []models.Task

	dbConnect, err := coreDB.GetConnection()
	if err != nil {
		return coreErrors.DBConnectError
	}

	// парсинг JSON-body
	if err = context.Bind(&dataIn); err != nil {
		return err
	}
	if err = dataIn.Validate(); err != nil {
		return err
	}

	// получение отфильтрованных задач из БД
	findResult := dbConnect.Preload("Executors").Where("created_at BETWEEN ? AND ?", dataIn.DateFrom, dataIn.DateTo).Find(&filteredTasks)
	if err := findResult.Error; err != nil {
		return err
	}

	// разделение задач по статусам
	filteredTasksByStatus.Get(filteredTasks)

	return context.JSON(http.StatusOK, map[string]interface{}{
		"backlog": filteredTasksByStatus.BacklogTasks,
		"proccessing": filteredTasksByStatus.ProccessingTasks,
		"finished": filteredTasksByStatus.FinishedTasks,
	})
}
