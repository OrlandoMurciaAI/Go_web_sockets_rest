package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"platzi.com/go/rest-ws-go/handlers"
	"platzi.com/go/rest-ws-go/middleware"
	"platzi.com/go/rest-ws-go/server"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file %v\n", err)
	}

	PORT := os.Getenv("PORT")
	JWT_SECRET := os.Getenv("JWT_SECRET")
	DATABASE_URL := os.Getenv("DATABASE_URL")

	s, err := server.NewServer(context.Background(), &server.Config{
		Port:        PORT,
		JWTSecret:   JWT_SECRET,
		DatabaseUrl: DATABASE_URL,
	})

	if err != nil {
		log.Fatalf("Error creating server %v\n", err)
	}

	s.Start(BindRoutes)
}

func BindRoutes(s server.Server, r *mux.Router) {
	api := r.PathPrefix("/api/v1").Subrouter() // Creamos un subrouter  para que el middelware lo use en vez del principal
	// Todas las rutas que lleven api antes seran protegidas
	api.Use(middleware.CheckAuthMiddleware(s)) // para cada una de las ruta utilice el middleware

	r.HandleFunc("/", handlers.HomeHandler(s)).Methods(http.MethodGet)

	r.HandleFunc("/api/v1/signup", handlers.SignUpHandler(s)).Methods(http.MethodPost)
	r.HandleFunc("/api/v1/login", handlers.LoginHandler(s)).Methods(http.MethodPost)
	api.HandleFunc("/me", handlers.MeHandler(s)).Methods(http.MethodGet)

	api.HandleFunc("/posts", handlers.InsertPostHandler(s)).Methods(http.MethodPost)
	api.HandleFunc("/posts/{postId}", handlers.DeletePostByIdHandler(s)).Methods(http.MethodDelete)
	api.HandleFunc("/posts/{postId}", handlers.UpdatePostByIdHandler(s)).Methods(http.MethodPut)
	r.HandleFunc("/api/v1/posts/{postId}", handlers.GetPostByIdHandler(s)).Methods(http.MethodGet)
	r.HandleFunc("/api/v1/posts", handlers.ListPostHandler(s)).Methods(http.MethodGet)

	r.HandleFunc("/ws", s.Hub().HandleWebSocket) // hacemos que el websocket sea de manera publica
}
