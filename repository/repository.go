package repository

import (
	"context"

	"platzi.com/go/rest-ws-go/models"
)

type Repository interface {
	InsertUser(ctx context.Context, user *models.User) error
	GetUserById(ctx context.Context, id string) (*models.User, error)
	GetUserByEmail(ctx context.Context, email string) (*models.User, error)
	InsertPost(ctx context.Context, post *models.Post) error
	GetPostById(ctx context.Context, id string) (*models.Post, error)
	DeletePost(ctx context.Context, id string, userId string) error
	UpdatePost(ctx context.Context, post *models.Post, userId string) error
	Close() error
}

var implementation Repository

func SetRepository(repository Repository) {
	implementation = repository
}

func InsertUser(ctx context.Context, user *models.User) error {
	return implementation.InsertUser(ctx, user)
}

func GetUserById(ctx context.Context, id string) (*models.User, error) {
	return implementation.GetUserById(ctx, id)
}

func GetUserByEmail(ctx context.Context, email string) (*models.User, error) {
	return implementation.GetUserByEmail(ctx, email)
}

func InsertPost(ctx context.Context, post *models.Post) error {
	return implementation.InsertPost(ctx, post)
}

func GetPostById(ctx context.Context, id string) (*models.Post, error) {
	return implementation.GetPostById(ctx, id)
}

func DeletePost(ctx context.Context, id string, userId string) error {
	return implementation.DeletePost(ctx, id, userId)
}

func UpdatePost(ctx context.Context, post *models.Post, userId string) error {
	return implementation.UpdatePost(ctx, post, userId)
}

// Abstracciones > concretas
// Principios de codigos solid donde toda la implementación se hace de
// forma modular buscando que sea de la forma mas general posible

// esta implementacion de user repository lo que hace es definir lo que haran
// estas funciones en tiempo de ejecucion y con esto no lo hago de forma
// tan implicita en el código sino que se realiza al vuelo.
