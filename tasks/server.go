package tasks

import (
	"context"
	"net/http"

	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
)

func NewHTTPServer(ctx context.Context, endpoints Endpoints) http.Handler {
	r := mux.NewRouter()
	r.Use(commonMiddleware)

	r.Methods("POST").Path("/").Handler(httptransport.NewServer(
		endpoints.CreateTodo,
		decodeCreateTodoReq,
		encodeResponse,
	))

	r.Methods("GET").Path("/").Handler(httptransport.NewServer(
		endpoints.GetTodos,
		decodeGetTodosReq,
		encodeResponse,
	))

	r.Methods("DELETE").Path("/{id}").Handler(httptransport.NewServer(
		endpoints.DeleteTodo,
		decodeDeleteTodoReq,
		encodeResponse,
	))

	r.Methods("PUT").Path("/{id}").Handler(httptransport.NewServer(
		endpoints.UpdateTodo,
		decodeUpdateTodoReq,
		encodeResponse,
	))
	return r

}

func commonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}
