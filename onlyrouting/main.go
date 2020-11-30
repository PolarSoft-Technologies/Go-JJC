package main

import (
	"github.com/PolarSoft-Technologies/Go-JJC/src/config"
	"github.com/PolarSoft-Technologies/Go-JJC/src/middleware"
	router "github.com/PolarSoft-Technologies/Go-JJC/src/routes"
	"github.com/gofiber/fiber/v2"
)

// entry point to this program
func main() {

	// call the New() method - used to instantiate a new Fiber App
	app := fiber.New()

	// Setup app wide middlewares
	middleware.AppMiddleWares(app)

	//setup app wide routes
	router.SetupRoutes(app)

	// start app and listen on port specified in the onlyrouting.env file
	err := app.Listen(":" + config.Config("SERVER_PORT"))
	if err != nil {
		panic(err)
	}

}
