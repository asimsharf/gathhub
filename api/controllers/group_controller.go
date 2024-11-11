package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/asimsharf/gatherhub/internal/models"
	"github.com/asimsharf/gatherhub/internal/services"
	"github.com/gorilla/mux"
)

type GroupController struct {
	Service *services.GroupService
}

func NewGroupController(service *services.GroupService) *GroupController {
	return &GroupController{Service: service}
}

func (c *GroupController) CreateGroup(w http.ResponseWriter, r *http.Request) {
	var item models.Group
	if err := json.NewDecoder(r.Body).Decode(&item); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := c.Service.CreateGroup(&item); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(item)
}

func (c *GroupController) GetGroup(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	item, err := c.Service.GetGroupByID(id)
	if err != nil {
		http.Error(w, "Item not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(item)
}

func (c *GroupController) GetAllGroups(w http.ResponseWriter, r *http.Request) {
	items, err := c.Service.GetAllGroups()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(items)
}
