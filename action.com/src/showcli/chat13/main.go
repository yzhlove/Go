package main

import "os"

//copy file

func main() {

	source, err := os.Open("./text.txt")
	if err != nil {
		panic(err)
	}

	distance, err := os.Create("backup.txt")
	if err != nil {
		panic(err)
	}

}

func readFile(file *os.File, length int) ([]byte, error) {

	return nil, nil
}
