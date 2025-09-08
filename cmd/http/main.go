package main

import (
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"
	"strings"

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

	slog.SetDefault(newLogger(os.Stderr, config.GetLevelLog()))
	slog.Info(fmt.Sprintf("Server running on port %s", config.ServerPort))

	mux := http.NewServeMux()

	staticHandler := http.FileServer(http.Dir("views/static/"))
	mux.Handle("/static/", http.StripPrefix("/static/", staticHandler))

	noteHandler := handlers.NewNoteHandler()
	mux.HandleFunc("/", noteHandler.NoteList)
	mux.HandleFunc("/note/view", noteHandler.NoteView)
	mux.HandleFunc("/note/new", noteHandler.NoteNew)
	mux.HandleFunc("/note/create", noteHandler.NoteCreate)

	http.ListenAndServe(fmt.Sprintf(":%s", config.ServerPort), mux)
}

func (c Config) GetLevelLog() slog.Level {
	switch strings.ToLower(c.LevelLog) {
	case "debug":
		return slog.LevelDebug
	case "info":
		return slog.LevelInfo
	case "warn":
		return slog.LevelWarn
	case "error":
		return slog.LevelError
	default:
		return slog.LevelInfo
	}
}
