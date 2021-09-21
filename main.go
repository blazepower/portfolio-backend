package main

import (
	"fmt"
	Handlers2 "github.com/portfolio/src/Handlers"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("http server listening on port 1000")
	handleRequests()
}

func handleRequests() {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/skills", Handlers2.GetAllSkills)
	router.HandleFunc("/projects", Handlers2.GetAllProjectDetails)
	router.HandleFunc("/education", Handlers2.GetAllEducation)
	router.HandleFunc("/experience", Handlers2.GetAllExperience)

	err := http.ListenAndServe(":1000", router)
	if err != nil {
		panic(err)
	}
}
