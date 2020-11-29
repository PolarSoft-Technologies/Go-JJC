package config

/** This package contains app wide re-usable functions **/

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

//Config func to get env value from key
func Config(key string) string {
	// load .env file
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Print("Error loading .env file")
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

/** AppMode: used for switching between modes supported in
*** this project
**/
func AppMode() string {
	/**create an array of modes that are supported in this project
	*** the current mode determines the routes and packages
	*** that would be used
	**/
	modes := [20]string{"onlyrouting", "with_mysql",
		"with_mongodb", "with_mariadb", "with_firebase",
		"with_graphql", "with_redis", "with_firebase"}

	return modes[0] //return the current working mode. i.e onlyrouting
}
