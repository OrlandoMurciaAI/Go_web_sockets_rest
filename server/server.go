package server

import (
	"context"
	"errors"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	database "platzi.com/go/rest-ws-go/database"
	repository "platzi.com/go/rest-ws-go/repository"
)

type Config struct {
	Port        string
	JWTSecret   string
	DatabaseUrl string
}

type Server interface {
	Config() *Config // Retorna una configuración
}

// Gracias a la interface el broker comienza a funcionar
// como un servidor
type Broker struct {
	config *Config
	router *mux.Router
}

// Reciever function que crea un metodo
func (b *Broker) Config() *Config {
	return b.config
}

// Este metodo le permite a nuestr broker levantarse
func (b *Broker) Start(binder func(s Server, r *mux.Router)) {
	b.router = mux.NewRouter()
	binder(b, b.router)
	repo, err := database.NewPostgresRepository(b.config.DatabaseUrl)

	if err != nil {
		log.Fatal(err)
	}

	repository.SetRepository(repo)

	log.Println("Starting server on port", b.Config().Port)
	if err := http.ListenAndServe(b.config.Port, b.router); err != nil {
		log.Fatal("ListenAndServer: ", err) // Se sale del programa despues de mostrar un mensaje
	}
}

// El contexto sirve para encontrar posibles problemas
// dentro de nuestra aplicación
func NewServer(ctx context.Context, config *Config) (*Broker, error) {
	if config.Port == "" {
		return nil, errors.New("Port is required")
	}

	if config.JWTSecret == "" {
		return nil, errors.New("JWTSecret is required")
	}

	if config.DatabaseUrl == "" {
		return nil, errors.New("DatabaseUrl is required")
	}

	broker := &Broker{
		config: config,
		router: mux.NewRouter(),
	}
	return broker, nil
}
