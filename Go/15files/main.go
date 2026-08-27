package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("Files!")
	content := "this is content here for the file"

	file, err := os.Create("firstfile.txt")
	checknillerr(err)

	length, err := io.WriteString(file, content)
	checknillerr(err)

	fmt.Printf("Wrote %d characters to file.\n", length)
	file.Close()

	readFile("firstfile.txt")

}

func readFile(filename string) {
	databytes, err := ioutil.ReadFile(filename)
	checknillerr(err)

	fmt.Println("File read successfully.", string(databytes))
}

func checknillerr(err error) {
	if err != nil {
		panic(err)
	}
}
