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
	"github.com/PolarSoft-Technologies/Go-JJC/src/validations"
	"github.com/gofiber/fiber/v2"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
	"encoding/json")

//UserController for user related operations
func UserController(app *fiber.App, db *gorm.DB) { //receiving the fiber object passed as app and the db instance

	/* NOTE: For the purpose of this tutorial, I'll be retrieving data sent
	*** by the client as query parameters. If you wish to use
	*** the body of the request or any other means, kindly review
	*** the users_controller.go file in the 'onlyrouting' module
	*** which explains how to retrieve
	*** data sent by the client in different ways.
	*** correct users in the users table in the database are:
	*** 1. Email: test@test.com, Password: 111111
	*** 2. Email: test1@test1.com, Password: 111111
	**/

	//handle sign-in operations here. Uses the GET HTTP method
	app.Get(config.GetAPIBase()+"user/signIn",
		func(c *fiber.Ctx) error {

			/* Here, our endpoint on the client side would be like this
			*** "{{BASE_API_URL}}/user/signIn?email=test@test.com&password=111111";
			**/

			//retrieving email & password from the client
			userEmail := c.Query("email")
			userPassword := c.Query("password")

			//create a variable of type validations.User to hold fields 
			//from the client and result of our db query
			user := validations.User{Email: userEmail, Password: userPassword}

			//sanitize the fields gotten from the client
			user.Prepare()

			//we check if any required field(s) value is/are missing and
			//that fields are valid
			err := user.IsValid("signin")

			//throw error if there's any invalid field value(s)
			if err != nil {
				return config.RESPONSE(c, 500, false,
					err.Error(), nil)
			}

			/*check if user with the supplied email exists and populate the result
			*** of the query into the user variable created above
			**/
			res := db.Find(&user, "email = ?", userEmail)
			//more queries can be found here https://gorm.io/docs/query.html

			if res.Error != nil { //user does not exist
				return config.RESPONSE(c, 500, false,
					config.MatchFailed, nil)
			}

			//validate password and return error if it is not correct
			err = config.VerifyPassword(user.Password, userPassword)
			if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
				return config.RESPONSE(c, 500, false,
					config.MatchFailed, nil)
			}

			//generate JWT token if the password is correct
			token, _ := config.CreateToken(user.ID)
			/* In production you need to handle error(s) thrown by the 
			*** second parameter, _. i.e change _ to err and handle
			*** the error thrown. do so everywhere it occurs in 
			*** order to have a production grade code
			*** refer to the introduction module if you do not understand what
			*** the underscore _ is used for.
			**/

			user.Token = token

			/*remove the password value from the response to be returned 
			** to the client
			*/
			user.Password = ""

			u,_ := json.Marshal(user)

			s := string(u) 

			//signin successful
			return config.RESPONSE(c, 200, true, config.SignInSuccess, &s)

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
