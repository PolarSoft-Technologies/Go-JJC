package models

//User model
type User struct {
	ID uint64
	FirstName string `gorm:"size:64" json:"first_name"`
	LastName string `gorm:"size:64" json:"last_name"`
	Gender string `gorm:"size:64" json:"gender"`
	Email string `gorm:"size:128;not null;unique" json:"email"`
	Password string `gorm:"size:64;not null;" json:"password"`
	Country string `gorm:"size:64" json:"country"`
	Token string `json:"token"`
  }