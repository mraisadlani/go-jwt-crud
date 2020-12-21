package controller

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/vanilla/go-jwt-crud/api/common"
	"github.com/vanilla/go-jwt-crud/api/entities"
	"github.com/vanilla/go-jwt-crud/api/payload"
	"github.com/vanilla/go-jwt-crud/api/repository"
	"github.com/vanilla/go-jwt-crud/api/repository/impl"
	"net/http"
	"strconv"
)

func GetAllProduct(w http.ResponseWriter, r *http.Request) {
	db, err := common.Connect()
	defer db.Close()

	if err != nil {
		payload.ErrorResponse(w, 500, err)
		return
	}

	repo := impl.NewProductRepositoryImpl(db)
	func (productRepository repository.ProductRepository) {
		product, err := productRepository.FindAll()

		if err != nil {
			payload.ErrorResponse(w, 422, err)
			return
		}

		payload.MessageResponse(w, "Get products successfully", product, 200)
	}(repo)
}

func GetProduct(w http.ResponseWriter, r *http.Request) {
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

	repo := impl.NewProductRepositoryImpl(db)
	func (productRepository repository.ProductRepository) {
		product, err := productRepository.FindById(uint64(uid))

		if err != nil {
			payload.ErrorResponse(w, 400, err)
			return
		}

		payload.MessageResponse(w, "Get detail product successfully", product, 200)
	}(repo)
}

func CreateProduct(w http.ResponseWriter, r *http.Request) {
	db, err := common.Connect()
	defer db.Close()

	if err != nil {
		payload.ErrorResponse(w, 500, err)
		return
	}

	var product entities.Product
	err = json.NewDecoder(r.Body).Decode(&product)

	if err != nil {
		payload.ErrorResponse(w, 422, err)
		return
	}

	repo := impl.NewProductRepositoryImpl(db)
	func (productRepository repository.ProductRepository) {
		product, err := productRepository.Save(product)

		if err != nil {
			payload.ErrorResponse(w, 422, err)
			return
		}

		payload.MessageResponse(w, "Create product successfully", product, 201)
	}(repo)
}

func UpdateProduct(w http.ResponseWriter, r *http.Request) {
	db, err := common.Connect()
	defer db.Close()

	if err != nil {
		payload.ErrorResponse(w, 500, err)
		return
	}

	vars := mux.Vars(r)
	uid, err := strconv.ParseInt(vars["id"], 10, 64)

	if err != nil {
		payload.ErrorResponse(w, 400, err)
		return
	}

	var product entities.Product
	err = json.NewDecoder(r.Body).Decode(&product)

	if err != nil {
		payload.ErrorResponse(w, 422, err)
		return
	}

	repo := impl.NewProductRepositoryImpl(db)
	func (productRepository repository.ProductRepository) {
		product, err := productRepository.Update(uint64(uid), product)

		if err != nil {
			payload.ErrorResponse(w, 422, err)
			return
		}

		payload.MessageResponse(w, "Update product successfully", product, 200)
	}(repo)
}

func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	db, err := common.Connect()
	defer db.Close()

	if err != nil {
		payload.ErrorResponse(w, 500, err)
		return
	}

	vars := mux.Vars(r)
	uid, err := strconv.ParseInt(vars["id"], 10, 64)

	if err != nil {
		payload.ErrorResponse(w, 400, err)
		return
	}

	repo := impl.NewProductRepositoryImpl(db)
	func (productRepository repository.ProductRepository) {
		product, err := productRepository.Delete(uint64(uid))

		if err != nil {
			payload.ErrorResponse(w, 422, err)
			return
		}

		payload.MessageResponse(w, "Delete product successfully", product, 200)
	}(repo)
}
