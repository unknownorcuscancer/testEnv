package main

import (
	"fmt"
	"os"
	"time"
)

type Vars struct {
	name      string
	character string
}

func LoadVars() (*Vars, error) {
	// err := godotenv.Load()
	// if err != nil {
	// 	log.Fatal("Error loading .env file")
	// }
	vars := &Vars{
		name:      os.Getenv("NAME"),
		character: os.Getenv("CHARACTER"),
	}

	return vars, nil
}

func main() {
	vars, err := LoadVars()

	if err != nil {
		os.Exit(1)
	}

	fmt.Printf("Hello,"+" %s"+"!", vars.name)
	fmt.Printf("\nFav Character:"+" %s", vars.character)

	time.Sleep(20 * time.Minute)

}
