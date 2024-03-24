package cmd

import (
	"log"
	"net/http"
	"text/template"
)

func Home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		ErrorHandler(w, http.StatusNotFound)
		return
	}

	if r.Method != http.MethodGet {
		ErrorHandler(w, http.StatusMethodNotAllowed)
		return
	}

	tmpl, err := template.ParseFiles(ListPage, TemplatePage)
	if err != nil {
		log.Println(err.Error())
		ErrorHandler(w, http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, ListTasksAll)
	if err != nil {
		log.Println(err.Error())
		ErrorHandler(w, http.StatusInternalServerError)
		return
	}
}
