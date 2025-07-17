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
	Manager *blog.Manager
}

func NewApplication() (*Application, error) {

	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	contentDir := "./content"

	parser := blog.NewParser()
	manager := blog.NewManager(contentDir)
	handler := api.NewHandler(manager)

	app := Application{
		Logger:  logger,
		Handler: handler,
		Parser:  parser,
		Manager: manager,
	}

	return &app, nil
}

func (app *Application) HealthCheck(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Status is available")
}
