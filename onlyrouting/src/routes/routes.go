package router

import (
	controllers "github.com/PolarSoft-Technologies/Go-JJC/src/controllers"
	"github.com/gofiber/fiber/v2"
)

// SetupRoutes func
func SetupRoutes(app *fiber.App) { /** We are receiving the Fiber app from main.go file as a param **/

	/** Setup all controller routes in this project here
	**/
	controllers.UserController(app)
	controllers.ProductController(app)

}
