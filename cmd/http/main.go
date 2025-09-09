package main

import (
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"

	"github.com/wesleysantana/GoKeep/internal/handlers"
	configloader "github.com/wesleysantana/config-loader"
)

type Config struct {
	ServerPort string `env:"SERVER_PORT,5000"`
	DBPassword string `env:"DB_PASSWORD,required"`
	LevelLog   string `env:"LEVEL_LOG,info"`
}

func main() {

	config := Config{}
	if err := configloader.Load(&config); err != nil {
		log.Fatal(err)
	}

	logger := handlers.Logger{}
	slog.SetDefault(logger.NewLogger(os.Stderr, logger.GetLevelLog(config.LevelLog)))
	slog.Info(fmt.Sprintf("Server running on port %s", config.ServerPort))

	mux := http.NewServeMux()

	staticHandler := http.FileServer(http.Dir("views/static/"))
	mux.Handle("/static/", http.StripPrefix("/static/", staticHandler))

	noteHandler := handlers.NewNoteHandler()
	mux.Handle("/", handlers.HandlerWithError(noteHandler.NoteList))
	mux.Handle("/note/view", handlers.HandlerWithError(noteHandler.NoteView))
	mux.Handle("/note/new", handlers.HandlerWithError(noteHandler.NoteNew))
	mux.Handle("/note/create", handlers.HandlerWithError(noteHandler.NoteCreate))

	if err := http.ListenAndServe(
		fmt.Sprintf(":%s", config.ServerPort), mux); err != nil {
		panic(err)
	}
}
