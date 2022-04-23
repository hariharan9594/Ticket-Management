package routers

import (
	"github.com/labstack/echo"
	"gitlab.com/vipindasvg/ticket-management/common/middleware"
	"gitlab.com/vipindasvg/ticket-management/controllers"
)

func SetUserRoutes(e *echo.Echo) *echo.Echo {
	e.POST(versionpref+"/users", controllers.CreateUser, middleware.AdminRBAC)
	e.POST(versionpref+"/initial-user", controllers.CreateUser)
	e.POST(versionpref+"/login", controllers.Login)
	e.POST(versionpref+"/test"+"/users/:uid", controllers.SampleUserRequest, middleware.UserRBAC)
	e.POST(versionpref+"/ticket/:id", controllers.CreateTicket, middleware.UserRBAC)
	//Get method
	e.GET(versionpref+"/admin/listUser", controllers.ListUser, middleware.AdminRBAC)
	e.GET(versionpref+"/admin/listTickets", controllers.ListTickets, middleware.AdminRBAC)
	e.GET(versionpref+"/User/listTickets/:id", controllers.ListUserTickets, middleware.UserRBAC)
	e.GET(versionpref+"/User/singleTicket/:id", controllers.UserTicketDetail, middleware.UserRBAC)
	//Update method
	e.PUT(versionpref+"/User/updateTicket/:tid", controllers.UpdateUserTicket, middleware.UserRBAC)

	return e
}
