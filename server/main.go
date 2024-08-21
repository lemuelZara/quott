package main

import "github.com/lemuelZara/server/internal/database"

func main() {
	_, err := database.NewDatabase()
	if err != nil {
		panic(err)
	}
}
