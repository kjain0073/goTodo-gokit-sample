package tasks

import (
	"context"
	"time"

	"github.com/go-kit/log"
	"github.com/go-kit/log/level"
	"github.com/gofrs/uuid"
	"github.com/kjain0073/go-Todo/models"
	"gopkg.in/mgo.v2/bson"
)

type service struct {
	repository models.Repository
	logger     log.Logger
}

func NewService(rep models.Repository, logger log.Logger) Service {
	return &service{
		repository: rep,
		logger:     logger,
	}
}

func (s service) CreateTodo(ctx context.Context, title string) (string, error) {
	logger := log.With(s.logger, "method", "CreateTodo")

	uuid, _ := uuid.NewV4()
	id := uuid.String()
	todo := models.TodoModel{
		ID:        bson.NewObjectId(),
		Title:     title,
		Completed: false,
		CreatedAt: time.Now(),
	}

	if err := s.repository.CreateTodo(ctx, todo); err != nil {
		level.Error(logger).Log("err", err)
		return "", err
	}

	logger.Log("create todo", id)

	return "Success", nil
}

func (s service) GetTodos(ctx context.Context) ([]models.Todo, error) {
	logger := log.With(s.logger, "method", "GetTodo")

	todos, err := s.repository.GetTodos(ctx)

	if err != nil {
		level.Error(logger).Log("err", err)
		return nil, err
	}

	todoList := []models.Todo{}

	for _, t := range todos {
		todoList = append(todoList, models.Todo{
			ID:        t.ID.Hex(),
			Title:     t.Title,
			Completed: t.Completed,
			CreatedAt: t.CreatedAt,
		})
	}

	logger.Log("Get Todo")

	return todoList, nil
}

func (s service) DeleteTodo(ctx context.Context, id string) (string, error) {
	logger := log.With(s.logger, "method", "DeleteTodo")

	if err := s.repository.DeleteTodo(ctx, id); err != nil {
		level.Error(logger).Log("err", err)
		return "", err
	}

	logger.Log("Delete Todo", id)

	return "Success", nil
}

func (s service) UpdateTodo(ctx context.Context, id string, title string, completed bool) (string, error) {
	logger := log.With(s.logger, "method", "UpdateTodo")

	if err := s.repository.UpdateTodo(ctx, id, title, completed); err != nil {
		level.Error(logger).Log("err", err)
		return "", err
	}

	logger.Log("Update todo", id)

	return "Success", nil
}
