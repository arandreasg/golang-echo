package main

import (
	"os"
	"pegadaianempat/config"
	"pegadaianempat/controller"
	auth "pegadaianempat/middleware"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	echoSwagger "github.com/swaggo/echo-swagger"
)

// @title Apps Car Documentation API
// @version 1.0
// @description This is a sample service for managing cars
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email soberkoder@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:8080
// @BasePath /
func main() {

	// initialize database connection
	config.Connect()

	e := echo.New()

	e.Use(middleware.CSRFWithConfig(middleware.CSRFConfig{
		TokenLookup: "header:" + config.CSRFTokenHeader,
		ContextKey:  config.CSRFKey,
	}))

	e.GET("/index", controller.Index)
	e.POST("/sayhello", controller.SayHello)

	emm := e.Group("/Employee")

	emm.Use(auth.Authentication())
	emm.PUT("/", controller.UpdateEmployee)
	emm.DELETE("/:id", controller.DeleteEmployee)

	itm := e.Group("/item")

	itm.Use(auth.Authentication())
	itm.POST("/", controller.CreateItem)

	// e.GET("/", controller.HelloWorld)

	// //e.GET("/json", controller.JsonMap)

	// e.GET("/page1", controller.Page1)

	// //e.Any("/user", controller.User)

	// e.POST("/employee", controller.CreateEmployee)

	// e.PUT("/employee", controller.UpdateEmployee)

	// e.DELETE("/employee/:id", controller.DeleteEmployee)

	// route for swagger

	e.GET("/swagger/*", echoSwagger.WrapHandler)

	e.POST("/login", controller.UserLogin)
	e.POST("/register", controller.CreateEmployee)

	var PORT = os.Getenv("PORT")

	e.Logger.Fatal(e.Start(":" + PORT))
}
