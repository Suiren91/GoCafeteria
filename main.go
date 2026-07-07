package main

import (
	"log"

	"github.com/Suiren91/go-cafeteria/internal/handler"
)

func main() {
	r := handler.NewRouter()

	if err := r.Run(); err != nil {
		log.Fatal(err)
	}
}
