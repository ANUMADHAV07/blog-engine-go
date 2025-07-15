package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/ANUMADHAV07/blog-engine-go.git/internal/app"
	"github.com/ANUMADHAV07/blog-engine-go.git/internal/routes"
)

func main() {

	content := []byte(`___
title: Test Post
date: 2024-01-15
tags: ["go", "api"]
___

# Hello World
This is a test post.`)

	app, err := app.NewApplication()

	post, err := app.Parser.Parse(content, "test.md")

	if err != nil {
		fmt.Printf("Error parsing post: %v", err)
	}

	fmt.Printf("   Title: %s\n", post.Title)
	fmt.Printf("   Slug: %s\n", post.Slug)
	fmt.Printf("   Date: %s\n", post.Date.Format("2006-01-02"))
	fmt.Printf("   Tags: %v\n", post.Tags)
	fmt.Printf("   HTML: %s\n", post.HTMLContent)

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
