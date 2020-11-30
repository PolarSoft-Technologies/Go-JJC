package config

/** This package contains app wide re-usable functions **/

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

//Config func to get env value from key
func Config(key string) string {
	// load onlyrouting.env file
	err := godotenv.Load("onlyrouting.env")
	if err != nil {
		fmt.Print("Error loading onlyrouting.env file")
	}
	return os.Getenv(key)

}

/** GetAPIBase: returns the api base key/path in the .env file.
*** Doing this here is very important just in case
*** you need to change your api key/path later in the future.
*** Changing it once is better than changing it for every single
*** controllers you create.
**/
func GetAPIBase() string {

	return Config("API_KEYWORD")

}
