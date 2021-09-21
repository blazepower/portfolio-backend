package Handlers

import (
	"encoding/base64"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type Detail struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

type ProjectDetail struct {
	Title   string   `json:"title"`
	Img     string   `json:"img"`
	Url     string   `json:"url"`
	Details []Detail `json:"details"`
}

type ProjectDetails struct {
	Title   string   `json:"title"`
	Img     string   `json:"img"`
	Url     string   `json:"url"`
	Details []Detail `json:"details"`
}

var (
	projectDetail  []ProjectDetail
	projectDetails []ProjectDetails
)

func GetAllProjectDetails(w http.ResponseWriter, _ *http.Request) {
	if os.Getenv("dev") == "true" {
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	}

	getJsonData("details", &projectDetail)
	projectDetails = make([]ProjectDetails, len(projectDetail))

	for i, detail := range projectDetail {
		projectDetails[i].Title = detail.Title
		projectDetails[i].Url = detail.Url
		projectDetails[i].Details = detail.Details
		imageBytes, _ := ioutil.ReadFile(detail.Img)
		projectDetails[i].Img = base64.StdEncoding.EncodeToString(imageBytes)
	}

	err := json.NewEncoder(w).Encode(projectDetails)
	if err != nil {
		log.Fatal(err)
	}
}
