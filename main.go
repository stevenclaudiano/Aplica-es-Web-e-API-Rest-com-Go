package main

import (
	"net/http"
	"text/template"
)

type Produto struct {
	Nome       string
	Descricao  string
	Preco      float64
	Quantidade int
}

var temp = template.Must(template.ParseGlob("templates/*.html"))

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":8000", nil)

}

func index(w http.ResponseWriter, r *http.Request) {
	produtos := []Produto{
		{"Notebook", "Notebook Dell", 3500.00, 10},
		{"Mouse", "Mouse sem fio", 80.00, 50},
	}
	temp.ExecuteTemplate(w, "index", produtos)
}
