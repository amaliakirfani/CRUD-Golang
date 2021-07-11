package controllers

import (
	"ToDoList/helpers"
	"ToDoList/infrastructures"
	"ToDoList/model"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

func GetToDoList(res http.ResponseWriter, req *http.Request) {

	var ToDoList []model.ToDoList

	pg := infrastructures.ToDoListConn()

	pg.Select("*").From("todolist").Load(&ToDoList)
	fmt.Println(ToDoList)

	helpers.Response(res, http.StatusOK, helpers.ResponsePattern{
		Code:    2200,
		Message: "success",
		Status:  "success",
		Data:    ToDoList,
	})
}

func AddToDoList(res http.ResponseWriter, req *http.Request) {

	var AddToDoList model.AddToDoList
	err := json.NewDecoder(req.Body).Decode(&AddToDoList)

	if err != nil {
		panic(err)
	}

	pg := infrastructures.ToDoListConn()

	rec, err := pg.InsertInto("todolist").
		Columns("name", "status").
		Record(&AddToDoList).
		Exec()

	if err != nil {
		panic(err)
	}
	fmt.Println(rec)

	helpers.Response(res, http.StatusOK, helpers.ResponsePattern{
		Code:    2200,
		Message: "success",
		Status:  "success",
		Data:    AddToDoList,
	})

}

func DeleteToDoList(res http.ResponseWriter, req *http.Request) {
	r := strings.Split(req.URL.Path, "/")
	Id := r[3]

	pg := infrastructures.ToDoListConn()

	_, err := pg.DeleteFrom("todolist").Where("id = ?", Id).Exec()

	if err != nil {
		panic(err)
	}

	helpers.Response(res, http.StatusOK, helpers.ResponsePattern{
		Code:    2200,
		Message: "success",
		Status:  "success",
		Data:    nil,
	})
}

func EditToDoList(res http.ResponseWriter, req *http.Request) {
	r := strings.Split(req.URL.Path, "/")
	Id := r[3]

	var ToDoList model.ToDoList

	pg := infrastructures.ToDoListConn()

	rec, err := pg.Select("*").From("todolist").Where("id = ?", Id).Load(&ToDoList)

	if err != nil {
		panic(err)
	}

	if rec == 0 {
		helpers.Response(res, http.StatusOK, helpers.ResponsePattern{
			Code:    2200,
			Message: "There is no data!",
			Status:  "success",
			Data:    nil,
		})

		return
	}

	helpers.Response(res, http.StatusOK, helpers.ResponsePattern{
		Code:    2200,
		Message: "success",
		Status:  "success",
		Data:    ToDoList,
	})
}
