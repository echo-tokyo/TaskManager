package urls

import (
	echo "github.com/labstack/echo/v4"

	testUrls "github.com/echo-tokyo/TaskManager/test_app/urls"
)


// подгрузка urls каждого микроприложения и их общая настройка
func InitUrlRouters(echoApp *echo.Echo) {
	apiTestGroup := echoApp.Group("/api/test")
	testUrls.RouterGroup(apiTestGroup)
}
