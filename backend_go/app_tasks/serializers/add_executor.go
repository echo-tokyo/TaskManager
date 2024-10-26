package serializers

import (
	"gorm.io/gorm"
	echo "github.com/labstack/echo/v4"
	validate "github.com/gobuffalo/validate/v3"

	"github.com/echo-tokyo/TaskManager/app_tasks/errors"
    
    "github.com/echo-tokyo/TaskManager/core/db/models"
	coreValidator "github.com/echo-tokyo/TaskManager/core/validator"
)


// структура для входных данных добавления нового исполнителя задачи
type AddExecutorIn struct {
	TaskID int `json:"taskId" myvalid:"required" validate:"required"`
}

// базовая валидация полей по тегам
func (self *AddExecutorIn) IsValid(errors *validate.Errors) {
	coreValidator.BaseValidator(self, errors)
}

// более глубокая валидация с возвратом ошибок валидации
func (self *AddExecutorIn) Validate() error {
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
	return nil
}


// получение записи задачи из БД по её ID
func (self *AddExecutorIn) AddExecutorInGetTaskEntity(dbConnect *gorm.DB, taskEntity *models.Task) error {	
	// получение задачи из БД с таким ID
	findResult := dbConnect.Where("id = ?", self.TaskID).First(&taskEntity)
	if err := findResult.Error; err != nil {
		return errors.TaskFoundError
	}
	return nil
}
