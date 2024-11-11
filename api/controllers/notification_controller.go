package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/asimsharf/gatherhub/internal/models"
	"github.com/asimsharf/gatherhub/internal/services"
	"github.com/gorilla/mux"
)

type NotificationController struct {
	Service *services.NotificationService
}

func NewNotificationController(service *services.NotificationService) *NotificationController {
	return &NotificationController{Service: service}
}

func (c *NotificationController) CreateNotification(w http.ResponseWriter, r *http.Request) {
	var item models.Notification
	if err := json.NewDecoder(r.Body).Decode(&item); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := c.Service.CreateNotification(&item); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(item)
}

func (c *NotificationController) GetNotification(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	item, err := c.Service.GetNotificationByID(id)
	if err != nil {
		http.Error(w, "Item not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(item)
}

func (c *NotificationController) GetAllNotifications(w http.ResponseWriter, r *http.Request) {
	items, err := c.Service.GetAllNotifications()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(items)
}
