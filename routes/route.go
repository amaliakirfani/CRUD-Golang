package routes

import (
	"AttendanceApi/controllers"

	"github.com/gorilla/mux"
)

func Routes() *mux.Router {
	router := mux.NewRouter()

	// router.HandleFunc("/todolist", controllers.GetToDoList).Methods("GET")

	// router.HandleFunc("/", controllers.TodoListView).Methods("GET")
	// router.HandleFunc("/todolist/add", controllers.AddToDoList).Methods("POST")
	// router.HandleFunc("/todolist/delete/{id}", controllers.DeleteToDoList).Methods("DELETE")
	// router.HandleFunc("/todolist/edit/{id}", controllers.EditToDoList).Methods("GET")

	router.HandleFunc("/attendance", controllers.Attendance).Methods("PUT")

	return router
}
