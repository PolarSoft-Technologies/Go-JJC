package controllers //declare package name here

/** This package is used for handling/controlling product related requests **/

import (
	"github.com/PolarSoft-Technologies/Go-JJC/src/config"
	"github.com/gofiber/fiber/v2"
)

//ProductController for product related operations
func ProductController(app *fiber.App) {

	//handle add-product operations here. Uses the Post HTTP method
	app.Post(config.GetAPIBase()+"products/add", func(c *fiber.Ctx) error {
		//return a response to the client in JSON format
		return c.JSON(&fiber.Map{
			"success": "true",
			"message": "Product Added Successfully",
			"data":    "Some data from add-product endpoint",
		})
	})

	//handle update-product operations here. Uses the Post HTTP method
	app.Post(config.GetAPIBase()+"products/update", func(c *fiber.Ctx) error {
		//return a response to the client in JSON format
		return c.JSON(&fiber.Map{
			"success": "true",
			"message": "Product Updated Successfully",
			"data":    "Some data from update-product endpoint",
		})
	})

	//handle fetch-products operations here. Uses the Get HTTP method
	app.Get(config.GetAPIBase()+"products/fetchAll", func(c *fiber.Ctx) error {
		//return a response to the client in JSON format
		return c.JSON(&fiber.Map{
			"success": "true",
			"message": "Fetched All Products Successfully",
			"data":    "Some data from fetch-products endpoint",
		})
	})

	/** ... you can add more functions here for handling the
	** specific api requests depending on your needs.
	**/

}
