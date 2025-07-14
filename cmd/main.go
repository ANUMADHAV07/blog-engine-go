package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/ANUMADHAV07/blog-engine-go.git/internal/app"
	"github.com/ANUMADHAV07/blog-engine-go.git/internal/routes"
)

func main() {

	app, err := app.NewApplication()

	if err != nil {
		fmt.Println("err", err)
		panic(err)
	}

	r := routes.SetupRoute(app)

	server := &http.Server{
		Addr:         ":8080",
		Handler:      r,
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}
	app.Logger.Printf("we are running our app %d\n", 8080)

	err = server.ListenAndServe()

	if err != nil {
		fmt.Println("err", err)
		app.Logger.Fatal()
	}

}
