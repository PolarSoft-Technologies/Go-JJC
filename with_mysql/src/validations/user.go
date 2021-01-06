package validations

import (
	"errors"
	"html"
	"strings"

	"github.com/PolarSoft-Technologies/Go-JJC/src/config"
	"github.com/PolarSoft-Technologies/Go-JJC/src/models"
	"github.com/badoux/checkmail"
)

//User alias of the User model in models.User
type User models.User

//BeforeCreate initiates the hashing of the user's password
func (u *User) BeforeCreate() error {
	hashedPassword, err := config.Hash(u.Password)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)
	return nil
}

//Prepare sanitizes the values gotten from the client
func (u *User) Prepare() {
	u.ID = 0
	u.FirstName = html.EscapeString(strings.TrimSpace(u.FirstName))
	u.LastName = html.EscapeString(strings.TrimSpace(u.LastName))
	u.Gender = html.EscapeString(strings.TrimSpace(u.Gender))
	u.Email = html.EscapeString(strings.TrimSpace(u.Email))
	u.Password = html.EscapeString(strings.TrimSpace(u.Password))
	u.Country = html.EscapeString(strings.TrimSpace(u.Country))
}

//IsValid checks if the supplied data from the client are valid/acceptable
func (u *User) IsValid(action string) error {
	switch strings.ToLower(action) {
	case "update":
		if u.FirstName == "" {
			return errors.New("FirstName is required")
		}
		if u.LastName == "" {
			return errors.New("LastName is required")
		}
		if u.ID == 0 {
			return errors.New("User ID is required")
		}

	case "signin":
		if u.Password == "" {
			return errors.New("Password is required")
		}
		if u.Email == "" {
			return errors.New("Email is required")
		}
		if err := checkmail.ValidateFormat(u.Email); err != nil {
			return errors.New("Email is invalid")
		}

	default:
		if u.FirstName == "" {
			return errors.New("FirstName is required")
		}
		if u.Password == "" {
			return errors.New("Password is required")
		}
		if u.Email == "" {
			return errors.New("Email is required")
		}
		if err := checkmail.ValidateFormat(u.Email); err != nil {
			return errors.New("Email is invalid")
		}
	}

	return nil
}
