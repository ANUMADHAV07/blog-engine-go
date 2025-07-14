package app

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/ANUMADHAV07/blog-engine-go.git/internal/api"
)

type Application struct {
	Logger  *log.Logger
	Hanlder *api.Handler
}

func NewApplication() (*Application, error) {

	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	handler := api.NewHandler()

	app := Application{
		Logger:  logger,
		Hanlder: handler,
	}

	return &app, nil
}

func (app *Application) HealthCheck(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Status is available")
}
