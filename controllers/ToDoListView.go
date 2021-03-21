package controllers

import (
	"net/http"
	"path"
	"text/template"
)

func TodoListView(res http.ResponseWriter, req *http.Request) {
	var filepath = path.Join("views", "ToDoList.html")
	var tmpl, err = template.ParseFiles(filepath)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}

	var data = map[string]interface{}{
		"title": "Todo List",
	}

	err = tmpl.Execute(res, data)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
	}
}
