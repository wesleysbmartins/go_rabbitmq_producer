package server

import (
	"fmt"
	configs "go_rabbitmq_producer/config"
	"go_rabbitmq_producer/internal/routes"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

type Server struct {
	Port           int      `yaml:"port"`
	AllowedOrigins []string `yaml:"allowed-origins"`
}

type IServer interface {
	Run()
}

func (s *Server) Run() {
	router := mux.NewRouter()
	router.Use()
	routes.Routes(router)
	http.Handle("/", router)

	credentials := Server{}
	config := configs.Config{}
	config.Load("server", &credentials)

	corsHandler := handlers.CORS(handlers.AllowedOrigins(credentials.AllowedOrigins))(router)

	fmt.Println(fmt.Printf("SERVER LISTENNING ON PORT: %v", credentials.Port))
	panic(http.ListenAndServe(fmt.Sprintf(":%v", credentials.Port), corsHandler))
}
