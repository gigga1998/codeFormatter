package codeFormatter


import (
	"os"
	"log"
	"bufio"
	"strings"
)


func ReadFile(filename string) []string{
	output := []string{}
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}

	scaner := bufio.NewScanner(file)
	for scaner.Scan() {
		output = append(output, scaner.Text())
	}

	err = file.Close()
	if err != nil {
		log.Fatal((err))
	}

	if scaner.Err() != nil {
		log.Fatal(err)
	}
	return output
}


func RemoveSpaces(filename string) {
	listOfStrings := []string{}
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}

	scaner := bufio.NewScanner(file)
	for scaner.Scan() {
		new_string := strings.TrimRight(scaner.Text(), " ")
		listOfStrings = append(listOfStrings, new_string)
	}

	err = file.Close()
	if err != nil {
		log.Fatal((err))
	}

	if scaner.Err() != nil {
		log.Fatal(err)
	}
	WriteIntoFile(filename, listOfStrings)
}


func WriteIntoFile(filePath string, listOfStrings []string) {
    file, err := os.Create(filePath)
     
    if err != nil{
        log.Fatal(err)
    }
    defer file.Close()

	N := len(listOfStrings) - 1
	for i, text := range listOfStrings {
		if i != N {
			file.WriteString(text + "\n")
		} else {
			file.WriteString(text)
		}
	}
}
