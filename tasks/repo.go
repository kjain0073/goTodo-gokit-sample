package tasks

import (
	"context"
	"errors"

	"github.com/go-kit/log"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"github.com/kjain0073/go-Todo/models"
)

const (
	HostName       string = "mongodb://localhost:27017"
	DbName         string = "demo_todo"
	CollectionName string = "todo"
	Port           string = ":9000"
)

type repo struct {
	db     *mgo.Database
	logger log.Logger
}

func NewRepo(db *mgo.Database, logger log.Logger) models.Repository {
	return &repo{
		db:     db,
		logger: log.With(logger, "repo", "mongoDB"),
	}
}

func (repo *repo) GetTodos(ctx context.Context) ([]models.TodoModel, error) {
	todos := []models.TodoModel{}

	if err := repo.db.C(CollectionName).Find(bson.M{}).All(&todos); err != nil {
		return nil, errors.New("unable to Fetch Todos")
	}

	return todos, nil
}

func (repo *repo) CreateTodo(ctx context.Context, todo models.TodoModel) error {

	if todo.Title == "" {
		return errors.New("title is required")
	}

	if err := repo.db.C(CollectionName).Insert(&todo); err != nil {
		return err
	}

	return nil
}

func (repo *repo) DeleteTodo(ctx context.Context, id string) error {
	if !bson.IsObjectIdHex(id) {
		return errors.New("invalid Task Id")
	}

	if err := repo.db.C(CollectionName).RemoveId(bson.ObjectIdHex(id)); err != nil {
		return errors.New("failed to delete Todo")
	}

	return nil
}

func (repo *repo) UpdateTodo(ctx context.Context, id string, title string, completed bool) error {

	if !bson.IsObjectIdHex(id) {
		return errors.New("id is invalid")
	}

	if title == "" {
		return errors.New("title field is required")
	}

	if err := repo.db.C(CollectionName).
		Update(
			bson.M{"_id": bson.ObjectIdHex(id)},
			bson.M{"title": title, "completed": completed},
		); err != nil {
		return errors.New("failed to update Todo")
	}
	return nil
}
