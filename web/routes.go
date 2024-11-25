package main

import (
	"log"
	"net/http"
	"runtime"

	"os"
	"path/filepath"

	"github.com/go-chi/chi/v5"
)

func (app *application) routes() http.Handler {
	mux := chi.NewRouter()

	mux.Get("/virtual-terminal", app.VirtualTerminal)
	mux.Post("/payment-succeeded", app.PaymentSucceeded)

	mux.Get("/charge-once", app.ChargeOnce)

	_, currentFile, _, _ := runtime.Caller(0)
	projectRoot := filepath.Join(filepath.Dir(currentFile), "../..")

	err := os.Chdir(projectRoot)
	if err != nil {
		log.Fatal("Error changing working directory: ", err)
	}

	fileServer := http.FileServer(http.Dir("./cmd/internal/static"))
	mux.Handle("/static/*", http.StripPrefix("/static/", fileServer))

	return mux
}
