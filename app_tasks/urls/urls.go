package urls

import (
	echo "github.com/labstack/echo/v4"

	coreMiddlewares "github.com/echo-tokyo/TaskManager/core/middlewares"
	"github.com/echo-tokyo/TaskManager/app_tasks/handlers"
)


func RouterGroup(group *echo.Group) {
	// test
	group.GET("/all", handlers.Tasks)
	group.GET("/executor", handlers.UserTasks, coreMiddlewares.ValidateJWTMiddleware, coreMiddlewares.TokenIsAccessMiddleware)
	// prod
	group.POST("/add-executor", handlers.AddExecutor, coreMiddlewares.ValidateJWTMiddleware, coreMiddlewares.TokenIsAccessMiddleware)
	group.POST("/date-filter", handlers.DateFilter, coreMiddlewares.ValidateJWTMiddleware, coreMiddlewares.TokenIsAccessMiddleware)
}
