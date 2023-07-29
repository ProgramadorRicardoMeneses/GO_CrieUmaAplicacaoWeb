package controllers

import (
	"html/template"
	"net/http"

	"github.com/ProgramadorRicardoMeneses/GO_CrieUmaAplicacaoWeb/models"
)

var tmpl = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	todosOsProdutos := models.BuscaTodosOsProdutos()

	tmpl.ExecuteTemplate(w, "Index", todosOsProdutos)
}

func New(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "New", nil)
}

func Insert(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		nomeProduto := r.FormValue("nome")
		descricaoProduto := r.FormValue("descricao")
		precoProduto := r.FormValue("preco")
		quantidadeProduto := r.FormValue("quantidade")
		models.CriaNovoProduto(nomeProduto, descricaoProduto, precoProduto, quantidadeProduto)
	}
	http.Redirect(w, r, "/", 301)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id") // Pega o id do
	models.DetetaProduto(id)
	http.Redirect(w, r, "/", 301)
}

func Edit(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	produto := models.EditaProduto(id)
	tmpl.ExecuteTemplate(w, "Edit", produto)
}

func Update(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		id := r.FormValue("id")
		nomeProduto := r.FormValue("nome")
		descricaoProduto := r.FormValue("descricao")
		precoProduto := r.FormValue("preco")
		quantidadeProduto := r.FormValue("quantidade")
		models.AlteraProduto(id, nomeProduto, descricaoProduto, precoProduto, quantidadeProduto)
	}
	http.Redirect(w, r, "/", 301)
}
