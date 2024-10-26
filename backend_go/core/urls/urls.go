package urls

import (
	echo "github.com/labstack/echo/v4"

	tasksUrls "github.com/echo-tokyo/TaskManager/app_tasks/urls"
)


// подгрузка urls каждого микроприложения и их общая настройка
func InitUrlRouters(echoApp *echo.Echo) {
	apiTasksGroup := echoApp.Group("/api/v2/tasks")
	tasksUrls.RouterGroup(apiTasksGroup)
}
