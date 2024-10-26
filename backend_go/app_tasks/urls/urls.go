package urls

import (
	echo "github.com/labstack/echo/v4"

	coreMiddlewares "github.com/echo-tokyo/TaskManager/core/middlewares"
	"github.com/echo-tokyo/TaskManager/app_tasks/handlers"
)


func RouterGroup(group *echo.Group) {
	group.POST("/add-executor", handlers.AddExecutor, coreMiddlewares.ValidateJWTMiddleware, coreMiddlewares.TokenIsAccessMiddleware)
	group.GET("/date-filter", handlers.DateFilter, coreMiddlewares.ValidateJWTMiddleware, coreMiddlewares.TokenIsAccessMiddleware)
}
