package urls

import (
	echo "github.com/labstack/echo/v4"

	coreMiddlewares "github.com/echo-tokyo/TaskManager/core/middlewares"
	"github.com/echo-tokyo/TaskManager/test_app/handlers"
)


func RouterGroup(group *echo.Group) {
	group.GET("/test", handlers.Test)
	group.GET("/auth", handlers.Auth, coreMiddlewares.ValidateJWTMiddleware, coreMiddlewares.TokenIsAccessMiddleware)
	group.GET("/refresh-test", handlers.RefreshTest, coreMiddlewares.ValidateJWTMiddleware, coreMiddlewares.TokenIsRefreshMiddleware)
}
