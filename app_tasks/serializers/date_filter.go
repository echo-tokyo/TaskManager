package serializers

import (
	"fmt"
	"time"

	echo "github.com/labstack/echo/v4"
	validate "github.com/gobuffalo/validate/v3"

    "github.com/echo-tokyo/TaskManager/core/db/models"
	"github.com/echo-tokyo/TaskManager/app_tasks/errors"

	coreValidator "github.com/echo-tokyo/TaskManager/core/validator"
)


// структура для входных данных добавления нового исполнителя задачи
type DateFilterIn struct {
	DateFrom	string `json:"date_from" myvalid:"required" validate:"required"`
	DateTo 		string `json:"date_to" myvalid:"required" validate:"required"`
}


// базовая валидация полей по тегам
func (self *DateFilterIn) IsValid(errors *validate.Errors) {
	coreValidator.BaseValidator(self, errors)
}

// более глубокая валидация с возвратом ошибок валидации
func (self *DateFilterIn) Validate() error {
	// базовая валидация полей по тегам
	var validateErrors *validate.Errors = validate.Validate(self)

	if len(validateErrors.Errors) > 0 {
		// словарь для ошибок
		errMap := make(map[string]string, len(validateErrors.Errors))

		for key, value := range validateErrors.Errors {
			errMap[key] = value[0]
		}
		// возвращаем *echo.HTTPError
		httpError := echo.NewHTTPError(400, errMap)
		return httpError
	}

	// перевод дат из string в time.Time
	parsedDateFrom, err := time.Parse("2006-01-02", self.DateFrom)
	if err != nil {
		return echo.NewHTTPError(400, map[string]string{"date_from": err.Error()})
	}
	parsedDateTo, err := time.Parse("2006-01-02", self.DateTo)
	if err != nil {
		return echo.NewHTTPError(400, map[string]string{"date_from": err.Error()})
	}

	if parsedDateFrom.After(parsedDateTo) {
		return errors.DateValidateError
	}

	// делаем верхний порог включительным в диапазоне фильтра
	(*self).DateTo = (parsedDateTo.Add(time.Hour * 24)).Format("2006-01-02")

	return nil
}


// структура с отфильтрованными задачами по статусам
type DateFilteredTasksByStatus struct {
	BacklogTasks []models.Task
	ProccessingTasks []models.Task
	FinishedTasks []models.Task
}


// получение структуры с разделёнными задачами по статусам
func (self *DateFilteredTasksByStatus) Get(tasks []models.Task) {
	for i, task := range tasks {
		fmt.Println("task", i, "->", task)
		fmt.Println("task.Status:", task.Status)
		fmt.Printf("task.Status: %T\n", task.Status)

		switch (task.Status) {
			case "backlog":
				(*self).BacklogTasks = append((*self).BacklogTasks, task)
			case "proccessing":
				(*self).ProccessingTasks = append((*self).ProccessingTasks, task)
			case "finished":
				(*self).FinishedTasks = append((*self).FinishedTasks, task)
		}
	}
}
