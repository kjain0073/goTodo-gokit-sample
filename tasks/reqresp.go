package tasks

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	// "github.com/thedevsaddam/renderer"

	"github.com/kjain0073/go-Todo/models"
)

// var rnd *renderer.Render

type (
	CreateTodoRequest struct {
		Title string `json:"title"`
	}
	CreateTodoResponse struct {
		Ok string `json:"ok"`
	}

	GetTodosRequest struct {
	}

	GetTodosResponse struct {
		TodoList []models.Todo
	}

	DeleteTodoRequest struct {
		Id string `json:"id"`
	}
	DeleteTodoResponse struct {
		Ok string `json:"ok"`
	}

	UpdateTodoRequest struct {
		Id        string `json:"id"`
		Title     string `json:"title"`
		Completed bool   `json:"completed"`
	}
	UpdateTodoResponse struct {
		Ok string `json:"ok"`
	}
)

func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}

// func encodeGetResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
// 	err := rnd.Template(w, http.StatusOK, []string{"static/home.tpl"}, nil)
// 	if err != nil {
// 		var logger log.Logger
// 		level.Error(logger).Log("exit", err)
// 		os.Exit(-1)
// 	}
// 	return json.NewEncoder(w).Encode(response)
// }

func decodeCreateTodoReq(ctx context.Context, r *http.Request) (interface{}, error) {
	var req CreateTodoRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}
	return req, nil
}

func decodeGetTodosReq(ctx context.Context, r *http.Request) (interface{}, error) {
	var req GetTodosRequest
	return req, nil

}

func decodeUpdateTodoReq(ctx context.Context, r *http.Request) (interface{}, error) {
	var req UpdateTodoRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}
	return req, nil
}

func decodeDeleteTodoReq(ctx context.Context, r *http.Request) (interface{}, error) {
	var req DeleteTodoRequest
	vars := mux.Vars(r)

	req = DeleteTodoRequest{
		Id: vars["id"],
	}
	return req, nil
}
