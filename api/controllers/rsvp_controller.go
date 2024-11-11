package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/asimsharf/gatherhub/internal/models"
	"github.com/asimsharf/gatherhub/internal/services"
	"github.com/gorilla/mux"
)

type RsvpController struct {
	Service *services.RsvpService
}

func NewRsvpController(service *services.RsvpService) *RsvpController {
	return &RsvpController{Service: service}
}

func (c *RsvpController) CreateRsvp(w http.ResponseWriter, r *http.Request) {
	var item models.Rsvp
	if err := json.NewDecoder(r.Body).Decode(&item); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := c.Service.CreateRsvp(&item); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(item)
}

func (c *RsvpController) GetRsvp(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	item, err := c.Service.GetRsvpByID(id)
	if err != nil {
		http.Error(w, "Item not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(item)
}

func (c *RsvpController) GetAllRsvps(w http.ResponseWriter, r *http.Request) {
	items, err := c.Service.GetAllRsvps()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(items)
}
