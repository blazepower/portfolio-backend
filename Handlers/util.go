package Handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

func getJsonData(file string, output interface{})  {
	dat, err := os.ReadFile(fmt.Sprintf("data/%s.json", file))
	if err != nil {
		dat = make([]byte, 0)
		log.Println(err)
	}

	err = json.Unmarshal(dat, &output)
	if err != nil {
		log.Println(err)
		output = make([]interface{}, 0)
	}
}