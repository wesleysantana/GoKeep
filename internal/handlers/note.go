package handlers

import (
	"fmt"
	"html/template"
	"net/http"
)

type noteHander struct{}

func NewNoteHandler() *noteHander {
	return &noteHander{}
}

func (h *noteHander) NoteList(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	files := []string{
		"views/templates/base.html",
		"views/templates/pages/home.html",
	}
	t, err := template.ParseFiles(files...)
	if err != nil {
		http.Error(w, "Aconteceu um erro ao executar essa página", http.StatusInternalServerError)
		return
	}
	t.ExecuteTemplate(w, "base", nil)
}

func (h *noteHander) NoteView(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "Nota não encontrada", http.StatusNotFound)
		return
	}
	files := []string{
		"views/templates/base.html",
		"views/templates/pages/note-view.html",
	}
	t, err := template.ParseFiles(files...)
	if err != nil {
		http.Error(w, "Aconteceu um erro ao executar essa página", http.StatusInternalServerError)
		return
	}
	t.ExecuteTemplate(w, "base", id)
}

func (h *noteHander) NoteNew(w http.ResponseWriter, r *http.Request) {
	files := []string{
		"views/templates/base.html",
		"views/templates/pages/note-new.html",
	}
	t, err := template.ParseFiles(files...)
	if err != nil {
		http.Error(w, "Aconteceu um erro ao executar essa página", http.StatusInternalServerError)
		return
	}
	t.ExecuteTemplate(w, "base", nil)
}

func (h *noteHander) NoteCreate(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {

		w.Header().Set("Allow", http.MethodPost)

		//rejeitar a requisição
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		return
	}
	fmt.Fprint(w, "Criando uma nova nota...")
}
