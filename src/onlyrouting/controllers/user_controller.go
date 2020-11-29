package onlyrouting //declare package name here

/** This package is used for handling/controlling user related requests **/

import (
	"github.com/PolarSoft-Technologies/Go-JJC/src/config"
	"github.com/gofiber/fiber/v2"
)

//UserController for user related operations
func UserController(app *fiber.App) {

	//handle sign-in operations here. Uses the Post HTTP method
	app.Post(config.GetAPIBase()+"user/signIn", func(c *fiber.Ctx) error {
		//return a response to the client in JSON format
		return c.JSON(&fiber.Map{
			"success": "true",
			"message": "Signed In Successfully",
			"data":    "Some data from sign-in endpoint",
		})
	})

	//handle sign-up operations here. Uses the Post HTTP method
	app.Post(config.GetAPIBase()+"user/signUp", func(c *fiber.Ctx) error {
		//return a response to the client in JSON format
		return c.JSON(&fiber.Map{
			"success": "true",
			"message": "Signed Up Successfully",
			"data":    "Some data from sign-up endpoint",
		})
	})

	//handle fetch-users operations here. Uses the Get HTTP method
	app.Get(config.GetAPIBase()+"user/fetchAll", func(c *fiber.Ctx) error {
		//return a response to the client in JSON format
		return c.JSON(&fiber.Map{
			"success": "true",
			"message": "Fetched All Users Successfully",
			"data":    "Some data from fetch-users endpoint",
		})
	})

	/** ... you can add more functions here for handling the
	** specific api requests depending on your needs.
	**/

}
