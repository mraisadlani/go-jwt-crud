package controller

import (
	"encoding/json"
	"errors"
	"github.com/vanilla/go-jwt-crud/api/common"
	"github.com/vanilla/go-jwt-crud/api/entities"
	"github.com/vanilla/go-jwt-crud/api/payload"
	"github.com/vanilla/go-jwt-crud/api/repository"
	"github.com/vanilla/go-jwt-crud/api/repository/impl"
	"github.com/vanilla/go-jwt-crud/api/security"
	"net/http"
)

func DoSign(w http.ResponseWriter, r *http.Request) {
	db, err := common.Connect()
	defer db.Close()

	if err != nil {
		payload.ErrorResponse(w, 500, err)
		return
	}

	var account entities.Account
	err = json.NewDecoder(r.Body).Decode(&account)

	if err != nil {
		payload.ErrorResponse(w, 422, err)
		return
	}

	repo := impl.NewUserRepositoryImpl(db)
	func (userRepository repository.UserRepository) {
		user, err := userRepository.Login(account.Username.String)

		if err != nil {
			payload.ErrorResponse(w, 422, err)
			return
		}

		password := account.Password.String
		hashPassword := user.Password.String

		if len(user.Username.String) > 0 {
			hash := security.VerifyPassword(hashPassword, password)

			if hash == nil {
				token, err := security.GenerateToken(uint64(user.ID.Int64))

				if err != nil {
					payload.ErrorResponse(w, 422, err)
					return
				}

				payload.MessageToken(w, "Login Successfully", token, user, 200)
			} else {
				payload.ErrorResponse(w, 500, errors.New("Verify password is not match"))
			}
		} else {
			payload.ErrorResponse(w, 400, errors.New("User not found"))
			return
		}
	}(repo)
}