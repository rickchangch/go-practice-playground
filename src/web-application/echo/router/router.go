package router

import (
	"go-practice-playground/web-application/echo/users"

	"github.com/labstack/echo/v4"
)

func AddRules(e *echo.Echo) {
	// 建立 共通 prefix 的 group router
	v1 := e.Group("api/v1")

	// curl -i 127.0.0.1:8080/api/v1/users
	v1.GET("/users", users.Index)

	// crul -i -X PUT -d "name=rick" 127.0.0.1:8080/api/v1/users/1
	v1.PUT("/users/:id", users.Update)

	// curl -i -X POST -d "name=rachel" -d "email=nueip@staff.com" 127.0.0.1:8080/api/v1/users
	v1.POST("/users", users.Create)

	// curl -i 127.0.0.1:8080/api/v1/users/2
	v1.GET("/users/:id", users.View)
}
