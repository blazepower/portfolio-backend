package main

import (
	"fmt"
	"github.com/blazepower/portfolio-backend/Handlers"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("http server listening on port 8080")
	handleRequests()
}

func handleRequests() {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/skills", Handlers.GetAllSkills)
	router.HandleFunc("/projects", Handlers.GetAllProjectDetails)
	router.HandleFunc("/education", Handlers.GetAllEducation)
	router.HandleFunc("/experience", Handlers.GetAllExperience)

	err := http.ListenAndServe(":8080", router)
	if err != nil {
		panic(err)
	}
}
