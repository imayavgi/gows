package controllers

import (
	"net/http"
	"regexp"
)

type userController struct {
	userIDPattern *regexp.Regexp
}

func newUserController() *userController {
	return &userController{
		userIDPattern: regexp.MustCompile((`^/users/(\d+)/?`)),
	}
}

func (uc userController) ServeHTTP(response http.ResponseWriter, request *http.Request) {
	response.Write([]byte(" Simple response from User Controller"))
}
