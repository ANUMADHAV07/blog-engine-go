package app

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/ANUMADHAV07/blog-engine-go.git/internal/api"
	"github.com/ANUMADHAV07/blog-engine-go.git/internal/blog"
)

type Application struct {
	Logger  *log.Logger
	Handler *api.Handler
	Parser  *blog.Parser
}

func NewApplication() (*Application, error) {

	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	handler := api.NewHandler()
	parser := blog.NewParser()

	app := Application{
		Logger:  logger,
		Handler: handler,
		Parser:  parser,
	}

	return &app, nil
}

func (app *Application) HealthCheck(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Status is available")
}
