package file

import (
	"io/ioutil"
	"os"
)

func WriteToFile(path string, result string) {
	f, err := os.Create(path)
	HandleErrors(err)

	f.WriteString(result)
}

func ReadFromFile(path string) []byte {
	data, err := ioutil.ReadFile(path)
	HandleErrors(err)

	return data
}

func HandleErrors(err error) {
	if err != nil {
		panic(err)
	}
}
