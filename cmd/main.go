package main

import (
	"log"
	"path"

	"runtime"

	"github.com/JulianTeschner/githubCli/handler"
	"github.com/joho/godotenv"
)

func main() {
	loadEnv()

	// handler.GetUser()
	handler.CreateEmptyRepo("EmptyTemplate", "JulianTeschner", "githubCli", "This is a auto-generated repository")
	// handler.DeleteRepo("JulianTeschner", "test")
}

func loadEnv() {
	if _, filename, _, ok := runtime.Caller(1); !ok {
		log.Fatal("No caller information")
	} else {
		filepath := path.Join(path.Dir(filename), ".env")
		err := godotenv.Load(filepath)
		if err != nil {
			log.Fatal("No .env file found in cmd Folder")
		}
	}
}
