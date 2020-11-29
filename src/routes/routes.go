package router

import (
	onlyrouting "github.com/PolarSoft-Technologies/Go-JJC/src/onlyrouting/controllers"
	"github.com/gofiber/fiber/v2"
)

// SetupRoutes func
func SetupRoutes(app *fiber.App) { /** We are receiving the Fiber app from main.go file as a param **/

	/** Setup app wide routing to any controller created in the
	*** controllers folder and passing app as a param tore
	**/
	onlyrouting.Controller(app)

}
