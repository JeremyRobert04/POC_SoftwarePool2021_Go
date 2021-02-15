package data

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func LineToCSV(line string) ([]string, error) {
	arr := strings.Split(line, ",")
	if len(arr) == 0 {
		return nil, fmt.Errorf("epmpty string")
	}
	return arr, nil
}

func ReadFile(path string) ([]string, error) {
	content, err := ioutil.ReadFile(path)
	str := string(content)
	if err != nil {
		log.Fatal(err)
		return strings.Split(str, "\n"), nil
	}
	return strings.Split(str, "\n"), nil
}