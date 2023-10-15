package rest

import (
	"errors"
	"net/http"
	"strings"

	"github.com/kelas.work/project-go-restful-api/internal/model"
)

func GetSessionData(r *http.Request) (model.UserSession, error) {
	authString := r.Header.Get("Authorization")
	splitSting := strings.Split(authString, " ")
	if len(splitSting) != 2 {
		return model.UserSession{}, errors.New("unauthorized")
	}

	accessString := splitSting[1]

	return model.UserSession{
		JWTToken: accessString,
	}, nil
}