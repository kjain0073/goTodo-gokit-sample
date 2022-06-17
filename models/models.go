package models

import (
	"context"
	"time"

	"gopkg.in/mgo.v2/bson"
)

type (
	TodoModel struct {
		ID        bson.ObjectId `bson:"_id,omitempty"`
		Title     string        `bson:"title"`
		Completed bool          `bson:"completed"`
		CreatedAt time.Time     `bson:"createAt"`
	}
	Todo struct {
		ID        string    `json:"id"`
		Title     string    `json:"title"`
		Completed bool      `json:"completed"`
		CreatedAt time.Time `json:"created_at"`
	}
	Repository interface {
		CreateTodo(ctx context.Context, todo TodoModel) error
		GetTodos(ctx context.Context) ([]TodoModel, error)
		DeleteTodo(ctx context.Context, id string) error
		UpdateTodo(ctx context.Context, id string, title string, completed bool) error
	}
)
