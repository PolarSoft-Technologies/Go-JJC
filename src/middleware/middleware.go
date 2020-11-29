package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

//AppMiddleWares func for the app
func AppMiddleWares(app *fiber.App) {
	//We use cors in order to allow cross-origin resource sharing
	app.Use(cors.New())

	/** Look up Fiber, https://github.com/gofiber/fiber, for all the
	*** possible middlewares you can make use of
	**/
}
