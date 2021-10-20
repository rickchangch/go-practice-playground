package users

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

type (
	UserID struct {
		ID int `json:"id" param:"id" query:"id" form:"id" validate:"gt=0"`
	}

	User struct {
		ID    int    `json:"id" param:"id" form:"id"`
		Name  string `json:"name" query:"name" form:"name"`
		Email string `json:"email" query:"email" form:"email"`
	}
)

var (
	counter = 1
	data    = map[int]*User{
		1: {
			ID:    1,
			Name:  "test 1",
			Email: "test1@gmail.com",
		},
	}
)

func Index(c echo.Context) error {
	u := new(UserID)
	if err := c.Bind(u); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	users := []*User{}

	for _, v := range data {
		users = append(users, v)
	}
	return c.JSONPretty(http.StatusOK, echo.Map{
		"code": "0000",
		"data": users,
	}, "\t")
}

func View(c echo.Context) error {
	u := new(UserID)
	if err := c.Bind(u); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}
	if err := c.Validate(u); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	ret, ok := data[u.ID]
	if ok && ret != nil {
		return c.JSONPretty(
			http.StatusOK,
			echo.Map{
				"code": "0000",
				"data": ret,
			}, "\t")
	}
	return c.JSONPretty(
		http.StatusNotFound,
		echo.Map{
			"code": "0001",
			"msg":  fmt.Sprintf("user [%d] not found", u.ID),
		}, "\t")
}

func Create(c echo.Context) error {
	u := new(User)
	if err := c.Bind(u); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	counter += 1
	u.ID = counter
	data[u.ID] = u
	c.Response().Header().Set("Location", fmt.Sprintf("/api/v1/user/%d", u.ID))
	return c.String(http.StatusCreated, "")
}

func Update(c echo.Context) error {
	user := new(User)
	if err := c.Bind(user); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	data[user.ID] = user

	return c.JSONPretty(http.StatusOK, echo.Map{
		"code": "0000",
		"data": user,
	}, "\t")
}
