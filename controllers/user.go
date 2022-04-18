package controllers

import (
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
	token, err := common.GenerateJWT(user.Id, user.IsAdmin.Bool)
	if err != nil {
		common.Log.WithField("handler", "user-login").WithField("issue", "request").
			Errorln("can not generate JWT token:", err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	dto := struct {
		User  models.User
		Token string
	}{*user, token}

	return c.JSON(http.StatusOK, dto)
}

//To test user requests
func SampleUserRequest(c echo.Context) error{
	return c.JSON(http.StatusOK, nil)
}