package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/asimsharf/gatherhub/internal/models"
	"github.com/asimsharf/gatherhub/internal/services"
	"github.com/gorilla/mux"
)

type EventController struct {
	Service *services.EventService
}

func NewEventController(service *services.EventService) *EventController {
	return &EventController{Service: service}
}

func (c *EventController) CreateEvent(w http.ResponseWriter, r *http.Request) {
	var item models.Event
	if err := json.NewDecoder(r.Body).Decode(&item); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := c.Service.CreateEvent(&item); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(item)
}

func (c *EventController) GetEvent(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	item, err := c.Service.GetEventByID(id)
	if err != nil {
		http.Error(w, "Item not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(item)
}

func (c *EventController) GetAllEvents(w http.ResponseWriter, r *http.Request) {
	items, err := c.Service.GetAllEvents()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(items)
}
