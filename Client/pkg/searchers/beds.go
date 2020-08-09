package searchers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"client.com/pkg/models"
)

const BED_URL = "http://scraper.docker/api/til-boligen/sengemoebler-og-udstyr/dobbeltsenge/type-dobbeltseng/?bredde=(120-140)&laengde=(200-200)&soegfra=2860&radius=4"
const BED_FILE_PATH = "./data-bed.json"

func LookForBeds() {
	newResult, previousResults, newMap, previousMap := initializeVariables()

	c1 := make(chan []byte, 1)
	c2 := make(chan string, 2)

	resp := fetch(BED_URL)
	defer resp.Body.Close()

	body := getResponseBody(resp)

}

func initializeVariables() (*[]models.Basic, *[]models.Basic, *map[string]models.Basic, *map[string]models.Basic) {
	var newResults, previousResults []models.Basic
	newMap := make(map[string]models.Basic)
	previousMap := make(map[string]models.Basic)

	return &newResults, &previousResults, &newMap, &previousMap
}

func handleErrors(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func fetch(URL string) *http.Response {
	resp, err := http.Get(URL)
	handleErrors(err)

	return resp
}

func getResponseBody(resp *http.Response) []byte {
	body, err := ioutil.ReadAll(resp.Body)
	handleErrors(err)

	return body
}

func readPreviousResults(filepath string, c chan []byte) {
	var (
		previousResults []models.Basic
		previousMap     = make(map[string]models.Basic)
	)

	jsonString := readFromFile(filepath)
	json.Unmarshal(data, &previousResults)
	convertToMap(previousResults, &previousMap)

	c <- jsonString
	c <- previousResults
	c <- previousMap
}

func convertToMap(objects []models.Basic, objectsMap *map[string]models.Basic) {
	for _, obj := range objects {
		(*objectsMap)[obj.Id] = obj
	}
}

func readFromFile(path string) []byte {
	data, err := ioutil.ReadFile(path)
	handleErrors(err)

	return data
}
