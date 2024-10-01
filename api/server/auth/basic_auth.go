package auth

import (
	errors "github.com/go-openapi/errors"
	models "github.com/openclarity/kubeclarity/api/server/models"
)

var (
	userDb map[string]string
)

func init() {
	// emulates the loading of a local users database
	userDb = map[string]string{
		"tung": "pass",
		"vic":  "word",
	}
}

func IsLoggedIn(user, pass string) (interface{}, error) {
	password, ok := userDb[user]
	if !ok || pass != password {
		return nil, errors.New(401, "Unauthorized: not a registered user")
	}

	return models.Principal{
		Name: user,
	}, nil
}
