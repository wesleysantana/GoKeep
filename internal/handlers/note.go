package handlers

import (
	"errors"
	"fmt"
	"html/template"
	"net/http"

	"github.com/wesleysantana/GoKeep/internal/apperror"
)

type noteHander struct{}

func NewNoteHandler() *noteHander {
	return &noteHander{}
}

func (h *noteHander) NoteList(w http.ResponseWriter, r *http.Request) error {
	if r.URL.Path != "/" {
		return errors.New("página não encontrada")
	}
	files := []string{
		"views/templates/base.html",
		"views/templates/pages/home.html",
	}
	t, err := template.ParseFiles(files...)
	if err != nil {
		return errors.New("aconteceu um erro ao executar essa página")
	}
	return t.ExecuteTemplate(w, "base", nil)
}

func (h *noteHander) NoteView(w http.ResponseWriter, r *http.Request) error {
	id := r.URL.Query().Get("id")
	if id == "" {
		return apperror.WithStatus(errors.New("o ID da nota é obrigatória"),
			http.StatusBadRequest)
	}
	files := []string{
		"views/templates/base.html",
		"views/templates/pages/note-view.html",
	}
	t, err := template.ParseFiles(files...)
	if err != nil {
		return errors.New("aconteceu um erro ao executar essa página")
	}
	return t.ExecuteTemplate(w, "base", id)
}

func (h *noteHander) NoteNew(w http.ResponseWriter, r *http.Request) error {
	files := []string{
		"views/templates/base.html",
		"views/templates/pages/note-new.html",
	}
	t, err := template.ParseFiles(files...)
	if err != nil {
		return errors.New("aconteceu um erro ao executar essa página")
	}
	return t.ExecuteTemplate(w, "base", nil)
}

func (h *noteHander) NoteCreate(w http.ResponseWriter, r *http.Request) error {
	if r.Method != http.MethodPost {

		w.Header().Set("Allow", http.MethodPost)

		return errors.New("método não permitido")
	}
	fmt.Fprint(w, "Criando uma nova nota...")
	return nil
}
