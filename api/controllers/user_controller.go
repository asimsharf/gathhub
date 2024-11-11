package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/asimsharf/gatherhub/internal/models"
	"github.com/asimsharf/gatherhub/internal/services"
	"github.com/gorilla/mux"
)

type UserController struct {
	Service *services.UserService
}

func NewUserController(service *services.UserService) *UserController {
	return &UserController{Service: service}
}

func (c *UserController) CreateUser(w http.ResponseWriter, r *http.Request) {
	var item models.User
	if err := json.NewDecoder(r.Body).Decode(&item); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := c.Service.CreateUser(&item); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(item)
}

func (c *UserController) GetUser(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	item, err := c.Service.GetUserByID(id)
	if err != nil {
		http.Error(w, "Item not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(item)
}

func (c *UserController) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	items, err := c.Service.GetAllUsers()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(items)
}
