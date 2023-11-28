package main

import "github.com/CyrilSbrodov/ToDoList/internal/app"

func main() {
	srv := app.NewServerApp()
	srv.Run()
}
