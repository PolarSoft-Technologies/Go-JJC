package onlyrouting //declare package name here

/** This package is used for handling/controlling user related requests **/

import (
	"github.com/gofiber/fiber/v2"
)

//Controller holder for all only_routing functionalities
func Controller(app *fiber.App) {
	/** Include all other controllers in only_routing and passing them
	*** the app as a parameter
	**/
	UserController(app)
	ProductController(app)
}
