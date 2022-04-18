package routers

import (
	"github.com/labstack/echo"
	"gitlab.com/vipindasvg/ticket-management/common/middleware"
	"gitlab.com/vipindasvg/ticket-management/controllers"
)

func SetUserRoutes(e *echo.Echo) *echo.Echo {
	e.POST(versionpref+"/users", controllers.CreateUser, middleware.AdminRBAC)
	e.POST(versionpref+"/login", controllers.Login)
	e.POST(versionpref+"/test"+ "/users/:uid", controllers.SampleUserRequest, middleware.UserRBAC)
	e.POST(versionpref+"/initial-user", controllers.CreateUser)
	return e
}