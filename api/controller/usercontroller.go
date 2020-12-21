package controller

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/vanilla/go-jwt-crud/api/common"
	"github.com/vanilla/go-jwt-crud/api/entities"
	"github.com/vanilla/go-jwt-crud/api/payload"
	"github.com/vanilla/go-jwt-crud/api/repository"
	"github.com/vanilla/go-jwt-crud/api/repository/impl"
	"github.com/vanilla/go-jwt-crud/api/security"
	"net/http"
	"strconv"
)

func GetAllUser(w http.ResponseWriter, r *http.Request) {
	db, err := common.Connect()
	defer db.Close()

	if err != nil {
		payload.ErrorResponse(w, 500, err)
		return
	}

	repo := impl.NewUserRepositoryImpl(db)
	func (userRepository repository.UserRepository) {
		user, err := userRepository.FindAll()

		if err != nil {
			payload.ErrorResponse(w, 422, err)
			return
		}

		payload.MessageResponse(w, "Get users successfully", user, 200)
	}(repo)
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	db, err := common.Connect()
	defer db.Close()

	if err != nil {
		payload.ErrorResponse(w, 500, err)
		return
	}

	vars := mux.Vars(r)
	uid, err := strconv.ParseInt(vars["id"], 10, 64)

	if err != nil {
		payload.ErrorResponse(w, 500, err)
		return
	}

	repo := impl.NewUserRepositoryImpl(db)
	func (userRepository repository.UserRepository) {
		user, err := userRepository.FindById(uint64(uid))

		if err != nil {
			payload.ErrorResponse(w, 422, err)
			return
		}

		payload.MessageResponse(w, "Get detail user successfully", user, 200)
	}(repo)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	db, err := common.Connect()
	defer db.Close()

	if err != nil {
		payload.ErrorResponse(w, 500, err)
		return
	}

	var user entities.User
	err = json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		payload.ErrorResponse(w, 422, err)
		return
	}

	encrypt, err := security.HashPassword(user.Password.String)

	if err != nil {
		payload.ErrorResponse(w, 500, err)
		return
	}

	user.Password.String = encrypt
	repo := impl.NewUserRepositoryImpl(db)
	func (userRepository repository.UserRepository) {
		getuser, err := userRepository.Save(user)

		if err != nil {
			payload.ErrorResponse(w, 422, err)
			return
		}

		payload.MessageResponse(w, "create user successfully", getuser, 201)
	}(repo)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	db, err := common.Connect()
	defer db.Close()

	if err != nil {
		payload.ErrorResponse(w, 500, err)
		return
	}

	var user entities.User
	err = json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		payload.ErrorResponse(w, 422, err)
		return
	}

	encrypt, err := security.HashPassword(user.Password.String)

	if err != nil {
		payload.ErrorResponse(w, 500, err)
		return
	}

	vars := mux.Vars(r)
	uid, err := strconv.ParseInt(vars["id"], 10, 64)

	if err != nil {
		payload.ErrorResponse(w, 500, err)
		return
	}

	user.Password.String = encrypt
	repo := impl.NewUserRepositoryImpl(db)
	func (userRepository repository.UserRepository) {
		getuser, err := userRepository.Update(uint64(uid), user)

		if err != nil {
			payload.ErrorResponse(w, 422, err)
			return
		}

		payload.MessageResponse(w, "update user successfully", getuser, 200)
	}(repo)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	db, err := common.Connect()
	defer db.Close()

	if err != nil {
		payload.ErrorResponse(w, 500, err)
		return
	}

	vars := mux.Vars(r)
	uid, err := strconv.ParseInt(vars["id"], 10, 64)

	if err != nil {
		payload.ErrorResponse(w, 500, err)
		return
	}

	repo := impl.NewUserRepositoryImpl(db)
	func (userRepository repository.UserRepository) {
		user, err := userRepository.Delete(uint64(uid))

		if err != nil {
			payload.ErrorResponse(w, 422, err)
			return
		}

		payload.MessageResponse(w, "delete user successfully", user, 200)
	}(repo)
}
