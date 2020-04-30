package controllers

import (
	"encoding/json"
	"net/http"
	"regexp"
	"strconv"

	"github.com/imayavgi/gows/internal/pkg/models"
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
	if request.URL.Path == "/users" {
		switch request.Method {
		case http.MethodGet:
			uc.getAll(response, request)
		case http.MethodPost:
			uc.post(response, request)
		default:
			response.WriteHeader(http.StatusNotImplemented)
		}
	} else {
		matches := uc.userIDPattern.FindStringSubmatch(request.URL.Path)
		if len(matches) == 0 {
			response.WriteHeader(http.StatusNotFound)
		}
		id, err := strconv.Atoi(matches[1])
		if err != nil {
			response.WriteHeader(http.StatusNotFound)
		}
		switch request.Method {
		case http.MethodGet:
			uc.get(id, response)
		case http.MethodPut:
			uc.put(id, response, request)
		case http.MethodDelete:
			uc.delete(id, response)
		default:
			response.WriteHeader(http.StatusNotImplemented)
		}
	}
}

func (uc *userController) getAll(response http.ResponseWriter, request *http.Request) {
	encodeResponseAsJSON(models.GetUsers(), response)
}

func (uc *userController) get(id int, response http.ResponseWriter) {
	u, err := models.GetUserByID(id)

	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		return
	}

	encodeResponseAsJSON(u, response)
}

func (uc *userController) post(response http.ResponseWriter, request *http.Request) {
	u, err := uc.parseRequest(request)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte("Could not parse User object"))
		return
	}

	u, err = models.AddUser(u)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(err.Error()))
		return
	}

	encodeResponseAsJSON(u, response)
}

func (uc *userController) put(id int, response http.ResponseWriter, request *http.Request) {
	u, err := uc.parseRequest(request)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte("Could not parse User object"))
		return
	}

	if id != u.ID {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte("ID passed did match ID in URL"))
		return
	}

	u, err = models.UpdateUser(u)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(err.Error()))
		return
	}

	encodeResponseAsJSON(u, response)
}

func (uc *userController) delete(id int, response http.ResponseWriter) {
	err := models.RemoveUserByID(id)

	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(err.Error()))
		return
	}

	response.WriteHeader(http.StatusOK)
}

func (uc *userController) parseRequest(request *http.Request) (models.User, error) {
	dec := json.NewDecoder(request.Body)
	var u models.User
	err := dec.Decode(&u)
	if err != nil {
		return models.User{}, err
	}
	return u, nil
}
