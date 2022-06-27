package main

import (
	"log"
	"os"
	"path"

	"runtime"

	"flag"

	"github.com/JulianTeschner/githubCli/handler"
	"github.com/joho/godotenv"
)

func main() {
	loadEnv()

	githubUsername := flag.String("user", os.Getenv("DEFAULT_USER"), "Github Username")
	repoTemplate := flag.String("template", os.Getenv("DEFAULT_Template"), "Template to use")
	projectName := flag.String("name", "", "Name of the project")
	repoDescription := flag.String("description", "This is a auto-generated repository", "Description of the project")
	method := flag.String("method", "", "Method to use")
	clone := flag.Bool("clone", false, "Clone the new Repo to the current directory")

	flag.Parse()

	switch *method {
	case "create":
		handler.CreateEmptyRepo(repoTemplate, githubUsername, projectName, repoDescription)
	case "delete":
		handler.DeleteRepo(githubUsername, projectName)
	default:
		log.Println("No method specified")
	}

	if *clone {
		handler.CloneRepoHandler(githubUsername, projectName)
	}

}

func loadEnv() {
	if _, filename, _, ok := runtime.Caller(1); !ok {
		log.Fatal("No caller information")
	} else {
		filepath := path.Join(path.Dir(filename), ".env")
		err := godotenv.Load(filepath)
		if err != nil {
			log.Println("No .env file found in cmd Folder... Trying to use default values")
		}
	}
}
