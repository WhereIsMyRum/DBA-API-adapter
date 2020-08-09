package searchers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"reflect"
	"strings"

	"client.com/pkg/file"
	"client.com/pkg/models"
)

func LookFor(URL string, FilePath string) {
	newResults, previousResults, newMap, previousMap := initializeVariables()

	c1 := make(chan []byte, 1)
	c2 := make(chan string, 2)

	resp := fetch(URL)
	defer resp.Body.Close()

	body := getResponseBody(resp)

	go readPreviousResults(FilePath, c1, previousResults, previousMap)
	go readNewResults(string(body), c2, newResults, newMap)

	<-c1
	result := <-c2

	if !reflect.DeepEqual(newResults, previousResults) {
		newItems := findNew(*newMap, *previousMap)
		if newItems != nil {
			sendNotification(newItems)
		}
	}

	if result != "" {
		file.WriteToFile(FilePath, result)
	}

}

func initializeVariables() (*[]models.Basic, *[]models.Basic, *map[string]models.Basic, *map[string]models.Basic) {
	var newResults, previousResults []models.Basic
	newMap := make(map[string]models.Basic)
	previousMap := make(map[string]models.Basic)

	return &newResults, &previousResults, &newMap, &previousMap
}

func fetch(URL string) *http.Response {
	resp, err := http.Get(URL)
	file.HandleErrors(err)

	return resp
}

func getResponseBody(resp *http.Response) []byte {
	body, err := ioutil.ReadAll(resp.Body)
	file.HandleErrors(err)

	return body
}

func readPreviousResults(filepath string, c chan []byte, previousResults *[]models.Basic, previousMap *map[string]models.Basic) {
	jsonString := file.ReadFromFile(filepath)
	json.Unmarshal(jsonString, previousResults)
	convertToMap(*previousResults, previousMap)

	c <- jsonString
}

func readNewResults(body string, c chan string, newResults *[]models.Basic, newMap *map[string]models.Basic) {
	replacer := strings.NewReplacer("{\"a-status\":\"ok\",\"data\":{\"general\":", "")
	result := replacer.Replace(string(body))
	result = result[:len(result)-2]
	json.Unmarshal([]byte(result), &newResults)
	convertToMap(*newResults, newMap)

	c <- result
}

func convertToMap(objects []models.Basic, objectsMap *map[string]models.Basic) {
	for _, obj := range objects {
		(*objectsMap)[obj.Id] = obj
	}
}

func findNew(result map[string]models.Basic, historic map[string]models.Basic) []models.Basic {
	var newItems []models.Basic

	for k, val := range result {
		if _, ok := historic[k]; !ok {
			newItems = append(newItems, val)
		}
	}

	return newItems
}

func sendNotification(items []models.Basic) {
	email := models.Email{}

	email.CreateEmail(&items)
	email.Send()
}
