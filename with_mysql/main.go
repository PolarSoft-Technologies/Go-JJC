package main

/** NOTE: you can willingly change both PolarSoft-Technologies and Go-JJC
*** to whatever name you want to use but make sure to be consistent
*** with your naming conventions! For e.g replace every occurence of
*** PolarSoft-Technologies in every file in this project with the name
*** you want to use. The Go-JJC part is nothing but your project's name;
*** i.e simply rename your project and be consitent with the name
*** everywhere it occurs in this project!
 **/

import (
	"github.com/PolarSoft-Technologies/Go-JJC/src/config"
	"github.com/PolarSoft-Technologies/Go-JJC/src/middleware"
	router "github.com/PolarSoft-Technologies/Go-JJC/src/routes"
	"github.com/gofiber/fiber/v2"
)

// entry point of this program
func main() {

	// Connect to database
	db := config.DB()

	// call the New() method - used to instantiate a new Fiber App
	app := fiber.New()

	// Setup app wide middlewares
	middleware.AppMiddleWares(app)

	//setup app wide routes. Passing it the db instance ensures a single 
	//db connection is used throughout the execution of the main function
	router.SetupRoutes(app,db)

	//close MySQL connection when the main function terminates
	defer db.Close()

	// start app and listen on port specified in the with_mysql.env file
	err := app.Listen(":" + config.Config("SERVER_PORT"))
	if err != nil {
		panic(err)
	}

}
