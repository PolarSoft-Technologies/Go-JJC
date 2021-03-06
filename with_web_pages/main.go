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
	"github.com/gofiber/fiber/v2"
)

// entry point of this program
func main() {

	// call the New() method - used to instantiate a new Fiber App
	app := fiber.New()

	// Setup app wide middlewares
	middleware.AppMiddleWares(app)

	//serve our files in the public directory
	app.Static("/", "./public")
	/** you can also serve files from multiple directories
	*** Simply do this if that's what you want:
	app.Static("/","./public") //serve from public directory
	app.Static("/", "./files") //serve from files directory etc

	*** checkout the guide here https://docs.gofiber.io/api/app#static
	**/

	// start app and listen on port specified in the with_web_pages.env file
	err := app.Listen(":" + config.Config("SERVER_PORT"))
	if err != nil {
		panic(err)
	}

}
