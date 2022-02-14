package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {

	//user input
	var file string
	fmt.Printf("What is the file name?")
	fmt.Scan(&file)

	//open file
	text, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Println(err)
	}
	openedFile := string(text)
	lines := strings.Split(openedFile, "\n")

	fmt.Println(lines)

	//calls countWords function
	for index, element := range countWords(openedFile) {
		fmt.Println(index, "=", element)
	}

}

func countWords(lines string) map[string]int {

	newMap := make(map[string]int)
	inputFile := strings.Fields(lines)
	for _, word := range inputFile {
		_, num := newMap[word]
		if num {
			newMap[word] += 1
		} else {
			newMap[word] = 1
		}
	}
	return newMap
}

func reportResults(openedFile string) {

	for index, element := range countWords(openedFile) {
		fmt.Println(index, "=", element)
	}
}
