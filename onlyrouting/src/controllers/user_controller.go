package controllers //declare package name here

/** This package is used for handling/controlling user related requests **/

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
	"github.com/gofiber/fiber/v2"
)

/** create a struct model for holding the data in the body of the request.
*** this is helpful when you need to keep track of what's needed
*** by this endpoint. Otherwise you could blindly retrieve the data
*** in the body of the request using an interface instead.
*** Field names should start with an uppercase letter
**/
type BodyData struct {
	Email    string `json:"email" xml:"email" form:"email"`
	Password string `json:"password" xml:"password" form:"password"`
}

//UserController for user related operations
func UserController(app *fiber.App) {//receiving the fiber object passed as app

	//handle sign-in operations here. Uses the Post HTTP method
	app.Post(config.GetAPIBase()+"user/signIn",
		func(c *fiber.Ctx) error {
			/** The Fiber context c contains the HTTP request and response.
			*** It has methods for the request query string, parameters, body,
			*** HTTP headers, and so on. You can read up more at
			*** https://docs.gofiber.io/
			*** some examples are illustrated below
			**/

			/** 1. Retrieving a query string value in the request
			*** Suppose our endpoint on the client side is
			*** "user/signIn?email=test@test.com&password=123456";
			*** we retrieve the query string values of email and password
			*** from the Query function of the Fiber context, which takes the
			*** name of the query string as a parameter
			**/
			emailQueryString := c.Query("email")
			passwordQueryString := c.Query("password")

			/** 2. Retrieving a parameter value in the request
			*** Suppose our endpoint is "user/signIn/:email/:password",
			*** i.e on the client side we have
			*** "user/signIn/test@test.com/123456";
			*** we retrieve the parameters, email and password, from
			*** the Params function of the Fiber context, which takes the
			*** name of the parameter as a parameter
			**/
			emailQueryParam := c.Params("email")
			passwordQueryParam := c.Params("password")

			/** 3. Retrieving data from the body in the request
			*** Suppose our endpoint is "user/signIn",
			*** we retrieve the body of the request; from which
			*** we can access the individual data in the request's body
			**/

			//create a new instance of the BodyData struct as bodyData
			bodyData := new(BodyData)
			//populate bodyData with the data in the body of the request.
			if err := c.BodyParser(bodyData); err != nil {
				return err
			}

			//create a variable holding all response in this endpoint
			result := map[string]string{
				"emailQueryString":    emailQueryString,
				"passwordQueryString": passwordQueryString,
				"emailQueryParam":     emailQueryParam,
				"passwordQueryParam":  passwordQueryParam,
				"reqBodyEmail":        bodyData.Email,
				"reqBodyPassword":     bodyData.Password,
			}

			//return a response to the client in JSON format
			return c.JSON(&fiber.Map{
				"success": true,
				"message": "Signed In Successfully",
				"data":    result,
			})
		})

	//handle sign-up operations here. Uses the Post HTTP method
	app.Post(config.GetAPIBase()+"user/signUp", func(c *fiber.Ctx) error {
		//return a response to the client in JSON format
		return c.JSON(&fiber.Map{
			"success": true,
			"message": "Signed Up Successfully",
			"data":    "Some data from sign-up endpoint",
		})
	})

	//handle fetch-users operations here. Uses the Get HTTP method
	app.Get(config.GetAPIBase()+"user/fetchAll", func(c *fiber.Ctx) error {
		//return a response to the client in JSON format
		return c.JSON(&fiber.Map{
			"success": true,
			"message": "Fetched All Users Successfully",
			"data":    "Some data from fetch-users endpoint",
		})
	})

	/** ... you can add more functions here for handling the
	** specific api requests depending on your needs.
	**/

}
