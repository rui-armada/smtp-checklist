package api

import (
	"log"
	"net/http"
	"text/template"

	"github.com/gorilla/mux"
)

var (
	tpl *template.Template
)

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}

func SetupAPI() {
	/*----------- API routes ------------*/
	router := mux.NewRouter()

	router.HandleFunc("/", index)
	router.HandleFunc("/process", sendEmail)

	log.Fatal(http.ListenAndServe(":8080", router))
}

func index(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "index.gohtml", nil)
}
