package controllers

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
	"gitlab.com/vipindasvg/ticket-management/common"
	"gitlab.com/vipindasvg/ticket-management/models"
	"gitlab.com/vipindasvg/ticket-management/storage"
)

func CreateUser(c echo.Context) error {
	curs := storage.GetCursor()
	u := new(models.User)
	if err := c.Bind(u); err != nil {
		common.Log.WithField("handler", "create-user").WithField("issue", "request").
			Errorln("can not bind the request body into provided type:", err)
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	fmt.Println("Bind(u) :", u)
	created, err := curs.CreateUser(u)
	if err != nil {
		common.Log.WithField("handler", "create-user").WithField("issue", "cursor").
			Errorln("can not create user record in the database:", err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusCreated, created)
}

func Login(c echo.Context) error {
	curs := storage.GetCursor()
	user := new(models.User)
	if err := c.Bind(user); err != nil {
		common.Log.WithField("handler", "user-login").WithField("issue", "request").Warnln("can not binds the request body into provided type:", err)
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	var err error
	// authenticate the user
	user, err = curs.Login(user.Email, user.Password)
	if err != nil {
		if err.Error() == "record not found" {
			common.Log.WithField("handler", "user-login").WithField("issue", "cursor").Warnln("authenticate the user - record not found:", err)
			return echo.NewHTTPError(http.StatusNotFound, err.Error())
		}
		return echo.NewHTTPError(http.StatusUnauthorized, "Invalid login credentials")
	}

	// if login successful
	// generate JWT token
	fmt.Println("user.IsAdmin.Bool in login: ", user.IsAdmin.Bool)
	token, err := common.GenerateJWT(user.Id, user.IsAdmin.Bool, user.UserName)
	if err != nil {
		common.Log.WithField("handler", "user-login").WithField("issue", "request").
			Errorln("can not generate JWT token:", err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	fmt.Println("Login JWT token: ", token)
	dto := struct {
		User  models.User
		Token string
	}{*user, token}

	return c.JSON(http.StatusOK, dto)
}

//To test user requests
func SampleUserRequest(c echo.Context) error {
	return c.JSON(http.StatusOK, nil)
}

//Add tickets
func CreateTicket(c echo.Context) error {
	curs := storage.GetCursor()
	ticket := new(models.Ticket)

	if err := c.Bind(ticket); err != nil {
		common.Log.WithField("handler", "create-user").WithField("issue", "request").
			Errorln("can not bind the request body into provided type:", err)
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	fmt.Println("Bind(ticket) :", ticket)
	id, err := common.ParseUID(c.Request().URL.Path)
	//get id from link
	fmt.Println("c.Request().URL.Path :", id)
	//assign the id to ticket
	ticket.U_id = int(id)
	created, err := curs.CreateTicket(ticket)
	if err != nil {
		common.Log.WithField("handler", "create-user").WithField("issue", "cursor").
			Errorln("can not create ticket in the database:", err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusCreated, created)

}

//List all users.
func ListUser(c echo.Context) error {
	curs := storage.GetCursor()

	psql := fmt.Sprintf("SELECT * FROM users")

	data, err := curs.UserList(psql)
	if err != nil {
		common.Log.WithField("handler", "Get-user").WithField("issue", "cursor").
			Errorln("can not get:", err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, data)
}

//List all tickets
func ListTickets(c echo.Context) error {
	curs := storage.GetCursor()

	psql := fmt.Sprintf("SELECT * FROM tickets")

	data, err := curs.TicketList(psql)
	if err != nil {
		common.Log.WithField("handler", "Get-user").WithField("issue", "cursor").
			Errorln("can not get:", err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, data)
}

//List all tickets created by user.
func ListUserTickets(c echo.Context) error {
	curs := storage.GetCursor()

	id, err := common.ParseUID(c.Request().URL.Path)
	Id := int(id)
	fmt.Println("id :", Id)
	if err != nil {
		return err
	}

	psql := fmt.Sprintf("SELECT * FROM tickets WHERE u_id = %d;", int(Id))

	data, err := curs.ListUserTickets(psql)
	if err != nil {
		common.Log.WithField("handler", "Get-user").WithField("issue", "cursor").
			Errorln("can not get:", err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, data)
}

//get single ticket by id
func UserTicketDetail(c echo.Context) error {
	curs := storage.GetCursor()
	//ParseTwoID is a function in common package.
	//It takes url string as a argument and returns uid, tid and error.
	uid, tid, err := common.ParseTwoID(c.Request().URL.Path)
	if err != nil {
		return err
	}

	psql := fmt.Sprintf("SELECT * FROM tickets WHERE ticket_id = %d AND u_id = %d;", tid, uid)

	data, err := curs.UserTicketDetail(psql)
	if err != nil {
		common.Log.WithField("handler", "Get-user").WithField("issue", "cursor").
			Errorln("cannot get:", err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, data)
}

//Update ticket
func UpdateUserTicket(c echo.Context) error {
	curs := storage.GetCursor()
	ticket := new(models.Ticket)
	if err := c.Bind(ticket); err != nil {
		return err
	}
	uid, tid, err := common.ParseTwoID(c.Request().URL.Path)
	if err != nil {
		return err
	}
	err = curs.UpdateUserTicket(ticket, tid, uid)
	if err != nil {
		common.Log.WithField("handler", "Update-tickets").WithField("issue", "cursor").
			Errorln("cannot update ticket record in the database:", err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	common.Log.Debugf("Successfully Updated...")
	updated := map[string]string{
		"message": "successfully updated...",
	}
	return c.JSON(http.StatusCreated, updated)

}
