package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

/** AppMiddleWares func for the app. You can add all
*** the middlewares needed in this project here.
*** Middlewares are a set of rules/functions governing
*** the use of a particular endpoint in a given api.
*** For example, an authentication middleware may check
*** if a user is authorised (based on their access/bearer token) 
*** before they can hit/access the
*** requested route. A cors middleware could be used to 
*** accept requests only from a given url on the client side;
*** for example, if https://google.com is only allowed,
*** requests made from other urls and even from Postman
*** might be rejected/prevented from accessing the api route!
*** etc.
**/
func AppMiddleWares(app *fiber.App) {
	//We use cors in order to allow cross-origin resource sharing
	app.Use(cors.New())

	/** Look up Fiber, https://github.com/gofiber/fiber, for all the
	*** possible middlewares you can make use of
	**/

	//do something here if the route match any request
	app.Use(func(c *fiber.Ctx) error {
		return c.Next()
	})

	//do something here if the route match any request starting with /api
	app.Use("/api", func(c *fiber.Ctx) error {
		return c.Next()
	})

	// Attach multiple handlers 
	app.Use("/api",func(c *fiber.Ctx) error {
		c.Set("X-Custom-Header", "some_random_number_here. e.g random.String(32)")
			return c.Next()
		}, func(c *fiber.Ctx) error {
			return c.Next()
	})
}
