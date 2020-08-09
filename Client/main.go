package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"reflect"
	"strings"
	"time"

	"client.com/pkg/models"
)

const PATH = "./data.json"

func main() {
	for {
		var resultParsed, historicParsed []models.Basic
		resultMap := make(map[string]models.Basic)
		historicMap := make(map[string]models.Basic)

		resp, err := http.Get("http://scraper.docker/api/til-boligen/sengemoebler-og-udstyr/dobbeltsenge/type-dobbeltseng/?bredde=(120-140)&laengde=(200-200)&soegfra=2860&radius=4")
		c1 := make(chan []byte, 1)
		c2 := make(chan string, 2)

		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)

		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		go func() {
			historic := readFromFile(PATH)
			json.Unmarshal(historic, &historicParsed)
			convertToMap(historicParsed, &historicMap)
			c1 <- historic
		}()

		go func() {
			replacer := strings.NewReplacer("{\"a-status\":\"ok\",\"data\":{\"general\":", "")
			result := replacer.Replace(string(body))
			result = result[:len(result)-2]
			json.Unmarshal([]byte(result), &resultParsed)
			convertToMap(resultParsed, &resultMap)

			c2 <- result
		}()

		<-c1
		result := <-c2

		//writeToFile(PATH, result)
		if !reflect.DeepEqual(resultParsed, historicParsed) {
			newItems := findNew(resultMap, historicMap)
			sendNotification(newItems)
		}

		if result != "" {
			writeToFile(PATH, result)
		}

		time.Sleep(600 * time.Second)
	}

}

func sendNotification(items []models.Basic) {
	email := models.Email{}

	email.CreateEmail(&items)
	email.Send()
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

func writeToFile(path string, result string) {
	f, err := os.Create(path)
	checkErrors(err)

	f.WriteString(result)
}

func readFromFile(path string) []byte {
	data, err := ioutil.ReadFile(path)
	checkErrors(err)

	return data
}

func checkErrors(err error) {
	if err != nil {
		panic(err)
	}
}

func convertToMap(objects []models.Basic, objectsMap *map[string]models.Basic) {
	for _, obj := range objects {
		(*objectsMap)[obj.Id] = obj
	}
}
