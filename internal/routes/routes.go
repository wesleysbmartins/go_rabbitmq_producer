package routes

import (
	"fmt"
	"go_rabbitmq_producer/internal/controllers"
	"net/http"

	"github.com/gorilla/mux"
)

func healthCheck(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, i am RabbitMQ Producer Server!")
}

func Routes(router *mux.Router) {

	router.HandleFunc("/", healthCheck).Methods("GET")

	router.HandleFunc("/sale", controllers.SaleController).Methods("POST")
}
