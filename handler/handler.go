package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"

	"github.com/JulianTeschner/githubCli/models"
)

var baseUrl string = "https://api.github.com"
var client *http.Client = &http.Client{}

func GetUser() {

	url := baseUrl + "/user"
	method := "GET"

	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		log.Println(err)
		return
	}
	req.Header.Add("Accept", "application/vnd.github.v3+json")
	req.Header.Add("Authorization", "token "+os.Getenv("ACCESS_TOKEN"))

	res, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println(string(body))
}

func CreateEmptyRepo(template *string, owner *string, name *string, description *string) {

	url := baseUrl + fmt.Sprintf("/repos/JulianTeschner/%s/generate", *template)
	method := "POST"

	repo := models.NewRepo(*owner, *name, *description)

	body, err := json.Marshal(repo)
	if err != nil {
		log.Println(err)
		return
	}

	data := []byte(body)
	req, err := http.NewRequest(method, url, bytes.NewBuffer(data))

	if err != nil {
		log.Println(err)
		return
	}
	req.Header.Add("Accept", "application/vnd.github.v3+json")
	req.Header.Add("Authorization", "token "+os.Getenv("ACCESS_TOKEN"))

	res, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return
	}
	defer res.Body.Close()

	_, err = ioutil.ReadAll(res.Body)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println("Repo was successfully created")
}

func DeleteRepo(owner *string, name *string) {

	url := baseUrl + fmt.Sprintf("/repos/%s/%s", *owner, *name)
	method := "DELETE"

	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		log.Println(err)
		return
	}
	req.Header.Add("Accept", "application/vnd.github.v3+json")
	req.Header.Add("Authorization", "token "+os.Getenv("ACCESS_TOKEN"))

	res, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return
	}
	defer res.Body.Close()

	_, err = ioutil.ReadAll(res.Body)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println("Repo successfully deleted")
}

func CloneRepoHandler(owner *string, name *string) {
	log.Println("Clone new Repo")

	s := fmt.Sprintf("git@github.com:%s/%s.git", *owner, *name)
	cmd := exec.Command("git", "clone", s)
	err := cmd.Run()
	errorHandler(err)
}

func errorHandler(err error) {
	if err != nil {
		log.Println(err)
	}
}
