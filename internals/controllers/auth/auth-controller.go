// handler.go
package auth

import (
	"net/http"

	"github.com/jitendra/pastPal/internals/config"
	"github.com/jitendra/pastPal/internals/models/auth"
	"github.com/jitendra/pastPal/utils"
)

var userRepository auth.UserRepository

func init() {
	userRepository = &auth.MongoUserRepository{Collection: config.MongoConnect()}
}

func AddNewUserHandler(w http.ResponseWriter, r *http.Request) {
	user := &auth.UserInfoSignUp{}

	if err := utils.ParseBody(r, user); err != nil {
		http.Error(w, "Invalid request body: "+err.Error(), http.StatusBadRequest)
		return
	}

	if err := userRepository.AddUser(user); err != nil {
		http.Error(w, "Failed to add user: "+err.Error(), http.StatusInternalServerError)
		return
	}

	utils.ParseResponseInJson(w, user)
}
