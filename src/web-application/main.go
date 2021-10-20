package main

import (
	"go-practice-playground/web-application/echo"
	"go-practice-playground/web-application/gorilla"
	"go-practice-playground/web-application/mysql"
	nonframework "go-practice-playground/web-application/non-framework"
)

func main() {
	// use native golang HTTP server
	nonframework.Service.Run()
	// use gorilla tool libraries to support build web application
	gorilla.Service.Run()
	// use labstack/echo to build web application
	echo.Service.Run()

	// test mysql driver, should execute docker-compose first to build mysql docker container.
	mysql.Run()
}
