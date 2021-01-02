package router

/** NOTE: you can willingly change both PolarSoft-Technologies and Go-JJC
*** to whatever name you want to use but make sure to be consistent
*** with your naming conventions! For e.g replace every occurence of
*** PolarSoft-Technologies in every file in this project with the name
*** you want to use. The Go-JJC part is nothing but your project's name;
*** i.e simply rename your project and be consitent with the name
*** everywhere it occurs in this project!
**/

import (
	controllers "github.com/PolarSoft-Technologies/Go-JJC/src/controllers"
	"github.com/gofiber/fiber/v2"
)

// SetupRoutes func
func SetupRoutes(app *fiber.App) { /** We are receiving the Fiber app from main.go file as a param **/

	/** Setup all controller routes in this project here
	**/
	controllers.UserController(app)

}
