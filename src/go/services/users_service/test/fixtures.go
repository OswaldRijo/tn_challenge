package test

import (
	"time"

	"truenorth/services/users_service/models"
)

var RandSalt = "SOME RAND SALT"

var TheGoat = &models.User{
	Username:  "messi",
	Password:  "YeVYjDfUo8B9U7j990C1ar8zONs7i6a4L8sq11cuSb8=",
	Status:    models.StatusActive,
	CreatedAt: time.Now(),
	UpdatedAt: time.Now(),
}
