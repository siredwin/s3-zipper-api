package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"log"
	_"net/http"
	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"net/http"
)

func main() {
	e := echo.New()
	e.Debug = true
	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.Secure())
	//e.Use(middleware.Gzip()) // DONT USE GZIP WHEN STREAMING FILES// BAD THINGS HAPPEN
	e.Use(session.Middleware(sessions.NewCookieStore([]byte("secret"))))
	//CORS
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
	}))
	// #*************** // WEB  ROUTES *****************# //
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, Welcome to S3zipper API testing! \n")
	})
	/********** JWT  API  CALLERS******************************/
	e.Match([]string{"GET", "POST"}, "/zipstart", s3JwtStart)
	e.Match([]string{"GET", "POST"}, "/zipstate", s3JwtState)
	e.Match([]string{"GET", "POST"}, "/zipresult", s3JwtResult)

	e.Match([]string{"GET", "POST", "PUT"}, "/token1", getJWTToken)
	e.Match([]string{"GET", "POST", "PUT"}, "/token2", getJWTToken2) // same as the other

	// #*************** START SERVER *****************# //
	log.Printf("Start Go HTTP Server")
	e.Logger.Fatal(e.Start(HOST_IP))

}
