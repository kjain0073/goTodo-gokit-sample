package middlewareTodo

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/go-chi/chi"
	"github.com/thedevsaddam/renderer"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"github.com/kjain0073/go-Todo/models"
)

var rnd *renderer.Render
var db *mgo.Database

const (
	hostName       string = "mongodb://localhost:27017"
	dbName         string = "demo_todo"
	collectionName string = "todo"
	Port           string = ":9000"
)

func init() {
	// to connect with DB
	rnd = renderer.New()
	sess, err := mgo.Dial(hostName)
	checkErr(err)
	sess.SetMode(mgo.Monotonic, true)
	db = sess.DB(dbName)
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	err := rnd.Template(w, http.StatusOK, []string{"static/home.tpl"}, nil)
	checkErr(err)
}

func FetchTodos(w http.ResponseWriter, r *http.Request) {
	todos := []models.TodoModel{}

	if err := db.C(collectionName).Find(bson.M{}).All(&todos); err != nil {
		rnd.JSON(w, http.StatusProcessing, renderer.M{
			"message": "Failed to fetch todo list",
		})
		return
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
	rnd.JSON(w, http.StatusOK, renderer.M{
		"data": todoList,
	})
}

func CreateTodo(w http.ResponseWriter, r *http.Request) {
	var t models.Todo

	if err := json.NewDecoder(r.Body).Decode(&t); err != nil {
		rnd.JSON(w, http.StatusProcessing, err)
		return
	}

	if t.Title == "" {
		rnd.JSON(w, http.StatusBadRequest, renderer.M{
			"message": "Title is required ",
		})
		return
	}

	td := models.TodoModel{
		ID:        bson.NewObjectId(),
		Title:     t.Title,
		Completed: false,
		CreatedAt: time.Now(),
	}

	if err := db.C(collectionName).Insert(&td); err != nil {
		rnd.JSON(w, http.StatusProcessing, renderer.M{
			"message": "Failed to save todo",
			"error":   err,
		})
		return
	}

	rnd.JSON(w, http.StatusCreated, renderer.M{
		"message": "Success! todo created",
		"todo_id": td.ID.Hex(),
	})
}

func DeleteTodo(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimSpace(chi.URLParam(r, "id"))
	if !bson.IsObjectIdHex(id) {
		rnd.JSON(w, http.StatusBadRequest, renderer.M{
			"message": "The id is invalid",
		})
		return
	}

	if err := db.C(collectionName).RemoveId(bson.ObjectIdHex(id)); err != nil {
		rnd.JSON(w, http.StatusProcessing, renderer.M{
			"message": "Failed to delete todo",
			"error":   err,
		})
		return
	}

	rnd.JSON(w, http.StatusOK, renderer.M{
		"message": "todo deleted successfully",
	})
}

func UpdateTodo(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimSpace(chi.URLParam(r, "id"))

	if !bson.IsObjectIdHex(id) {
		rnd.JSON(w, http.StatusBadRequest, renderer.M{
			"message": "The id is invalid",
		})
		return
	}

	var t models.Todo

	if err := json.NewDecoder(r.Body).Decode(&t); err != nil {
		rnd.JSON(w, http.StatusProcessing, err)
		return
	}

	if t.Title == "" {
		rnd.JSON(w, http.StatusBadRequest, renderer.M{
			"message": "Title field is required",
		})
		return
	}

	if err := db.C(collectionName).
		Update(
			bson.M{"_id": bson.ObjectIdHex(id)},
			bson.M{"title": t.Title, "completed": t.Completed},
		); err != nil {
		rnd.JSON(w, http.StatusProcessing, renderer.M{
			"message": "failed to update todo",
			"error":   err,
		})
		return
	}
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
