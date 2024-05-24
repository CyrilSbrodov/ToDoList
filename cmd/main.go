package main

import (
	"fmt"
	"github.com/CyrilSbrodov/ToDoList/internal/app"
)

func main() {
	srv := app.NewServerApp()
	srv.Run()
	fmt.Println("done")
}
