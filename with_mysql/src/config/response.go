package config

/** This package contains app wide re-usable api response **/

import (
	"github.com/gofiber/fiber/v2"
)

//RESPONSE response
func RESPONSE(ctx *fiber.Ctx, statusCode int,
	succ bool, msg string, data *string) error {

	return ctx.Status(statusCode).JSON(&fiber.Map{
		"success": succ,
		"message": msg,
		"data":    data,
	})
}
