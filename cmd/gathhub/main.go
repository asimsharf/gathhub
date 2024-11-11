package main

import (
	"log"
	"net/http"

	"github.com/asimsharf/gatherhub/api/controllers"
	"github.com/asimsharf/gatherhub/internal/repositories"
	"github.com/asimsharf/gatherhub/internal/services"
	"github.com/asimsharf/gatherhub/pkg/db"
	"github.com/gorilla/mux"
)

func main() {
	db, port := db.ConnectMySQL()

	// Set up repositories
	userRepo := repositories.NewUserRepository(db)
	eventRepo := repositories.NewEventRepository(db)
	groupRepo := repositories.NewGroupRepository(db)
	// rsvpRepo := repositories.NewRSVPRepository(db)
	notificationRepo := repositories.NewNotificationRepository(db)

	// Set up services
	userService := services.NewUserService(userRepo)
	eventService := services.NewEventService(eventRepo)
	groupService := services.NewGroupService(groupRepo)
	// rsvpService := services.NewRSVPService(rsvpRepo)
	notificationService := services.NewNotificationService(notificationRepo)

	// Set up controllers
	userController := controllers.NewUserController(userService)
	eventController := controllers.NewEventController(eventService)
	groupController := controllers.NewGroupController(groupService)
	// rsvpController := controllers.NewRSVPController(rsvpService)
	notificationController := controllers.NewNotificationController(notificationService)

	// Set up the router and endpoints
	router := mux.NewRouter()

	// User Routes
	router.HandleFunc("/api/users", userController.CreateUser).Methods("POST")
	router.HandleFunc("/api/users/{id}", userController.GetUser).Methods("GET")
	router.HandleFunc("/api/users", userController.GetAllUsers).Methods("GET")
	// router.HandleFunc("/api/users/{id}", userController.UpdateUser).Methods("PUT")
	// router.HandleFunc("/api/users/{id}", userController.DeleteUser).Methods("DELETE")

	// Event Routes
	router.HandleFunc("/api/events", eventController.CreateEvent).Methods("POST")
	router.HandleFunc("/api/events/{id}", eventController.GetEvent).Methods("GET")
	router.HandleFunc("/api/events", eventController.GetAllEvents).Methods("GET")
	// router.HandleFunc("/api/events/{id}", eventController.UpdateEvent).Methods("PUT")
	// router.HandleFunc("/api/events/{id}", eventController.DeleteEvent).Methods("DELETE")

	// Group Routes
	router.HandleFunc("/api/groups", groupController.CreateGroup).Methods("POST")
	router.HandleFunc("/api/groups/{id}", groupController.GetGroup).Methods("GET")
	router.HandleFunc("/api/groups", groupController.GetAllGroups).Methods("GET")
	// router.HandleFunc("/api/groups/{id}", groupController.UpdateGroup).Methods("PUT")
	// router.HandleFunc("/api/groups/{id}", groupController.DeleteGroup).Methods("DELETE")

	// RSVP Routes
	// router.HandleFunc("/api/rsvps", rsvpController.CreateRSVP).Methods("POST")
	// router.HandleFunc("/api/rsvps/{id}", rsvpController.GetRSVP).Methods("GET")
	// router.HandleFunc("/api/rsvps", rsvpController.GetAllRSVPs).Methods("GET")
	// router.HandleFunc("/api/rsvps/{id}", rsvpController.UpdateRSVP).Methods("PUT")
	// router.HandleFunc("/api/rsvps/{id}", rsvpController.DeleteRSVP).Methods("DELETE")

	// Notification Routes
	router.HandleFunc("/api/notifications", notificationController.CreateNotification).Methods("POST")
	router.HandleFunc("/api/notifications/{id}", notificationController.GetNotification).Methods("GET")
	router.HandleFunc("/api/notifications", notificationController.GetAllNotifications).Methods("GET")
	// router.HandleFunc("/api/notifications/{id}", notificationController.UpdateNotification).Methods("PUT")
	// router.HandleFunc("/api/notifications/{id}", notificationController.DeleteNotification).Methods("DELETE")

	// Start the server
	if port == "" {
		port = "3000"
	}

	log.Printf("Server started at :%s", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
