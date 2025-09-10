package main

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/wesleysantana/GoKeep/internal/handlers"
	configloader "github.com/wesleysantana/config-loader"
)

type Config struct {
	ServerPort string `env:"QNS_SERVER_PORT,5000"`
	DBConnURL  string `env:"QNS_DB_CONN_URL,required"`
	LevelLog   string `env:"QNS_LEVEL_LOG,info"`
	// MailHost     string `env:"QNS_MAIL_HOST,required"`
	// MailPort     string `env:"QNS_MAIL_PORT,required"`
	// MailUsername string `env:"QNS_MAIL_USERNAME,required"`
	// MailPassword string `env:"QNS_MAIL_PASSWORD,required"`
	// MailFrom     string `env:"QNS_MAIL_FROM,nao-responder@quick.com"`
	// CSRFKey      string `env:"QNS_CSRF_KEY,required"`
}

func main() {

	config := Config{}
	if err := configloader.Load(&config); err != nil {
		log.Fatal(err)
	}

	logger := handlers.Logger{}
	slog.SetDefault(logger.NewLogger(os.Stderr, logger.GetLevelLog(config.LevelLog)))
	slog.Info(fmt.Sprintf("Server running on port %s", config.ServerPort))

	dbpool, err := pgxpool.New(context.Background(), config.DBConnURL)
	if err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}
	slog.Info("Conexão com o banco aconteceu com sucesso")
	defer dbpool.Close()

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
