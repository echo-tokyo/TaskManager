package urls

import (
	echo "github.com/labstack/echo/v4"

	coreMiddlewares "github.com/echo-tokyo/TaskManager/core/middlewares"
	"github.com/echo-tokyo/TaskManager/app_tasks/handlers"
)


func RouterGroup(group *echo.Group) {
	group.GET("/all", handlers.Tasks)
	group.GET("/executor", handlers.UserTasks, coreMiddlewares.ValidateJWTMiddleware, coreMiddlewares.TokenIsAccessMiddleware)
}
