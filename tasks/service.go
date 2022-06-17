package tasks

import (
	"context"

	"github.com/kjain0073/go-Todo/models"
)

type Service interface {
	CreateTodo(ctx context.Context, Title string) (string, error)
	GetTodos(ctx context.Context) ([]models.Todo, error)
	DeleteTodo(ctx context.Context, Id string) (string, error)
	UpdateTodo(ctx context.Context, Id string, Title string, Completed bool) (string, error)
}
