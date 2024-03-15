package route

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

type Todo struct {
	ID   int
	Task string
}

var todos []Todo

func Router(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Add("Content-Type", "application/json")
	rw.Header().Add("Access-Control-Allow-Origin", "*")

	query := r.URL.Query()
	id, _ := strconv.Atoi(query.Get("id"))

	switch r.Method {
	case "GET":
		rw.WriteHeader(http.StatusOK)
		json.NewEncoder(rw).Encode(todos)
	case "POST":
		var data Todo
		bodyBytes, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Printf("Error reading request body: %v", err)
			rw.WriteHeader(http.StatusBadRequest)
			return
		}
		err = json.Unmarshal(bodyBytes, &data)
		if err != nil {
			log.Printf("Error decoding JSON: %v", err)
			rw.WriteHeader(http.StatusBadRequest)
			return
		}

		// Check for duplicate ID
		for _, todo := range todos {
			if todo.ID == data.ID {
				rw.WriteHeader(http.StatusConflict)
				rw.Write([]byte(`{"message": "Duplicate ID found"}`))
				return
			}
		}

		// Append new todo
		todos = append(todos, data)

		rw.WriteHeader(http.StatusCreated)
		json.NewEncoder(rw).Encode(data)
	case "DELETE":
		for index, todo := range todos {
			if todo.ID == id {
				todos = append(todos[:index], todos[index+1:]...)
				rw.WriteHeader(http.StatusOK)
				rw.Write([]byte(`{"message": "Success to delete"}`))
			}
		}
	case "PUT":
		for index, todo := range todos {
			if todo.ID == id {
				json.NewDecoder(r.Body).Decode(&todo)
				todos[index].ID = todo.ID
				todos[index].Task = todo.Task
				rw.WriteHeader(http.StatusOK)
				rw.Write([]byte(`{"message": "Success to update"}`))
			}
		}
	}

	log.Printf("%v", todos)
}
