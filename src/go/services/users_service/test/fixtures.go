package test

import (
	"time"

	usermodels "truenorth/services/users_service/models"
)

var RandSalt = "SOME RAND SALT"

var TheGoat = &usermodels.User{
	Username:  "messi",
	Password:  "YeVYjDfUo8B9U7j990C1ar8zONs7i6a4L8sq11cuSb8=",
	Status:    usermodels.StatusActive,
	CreatedAt: time.Now(),
	UpdatedAt: time.Now(),
}
