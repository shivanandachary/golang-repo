package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	var name string
	var age string
	var gender string
	fmt.Println("Enter your Name:")
	_, err := fmt.Scanln(&name)
	if err != nil {
		return
	}

	fmt.Println("Enter your Age:")
	_, err = fmt.Scanln(&age)
	if err != nil {
		return
	}

	fmt.Println("Enter your Gender :")
	_, err = fmt.Scanln(&gender)
	if err != nil {
		return
	}

	var filename string = "./sample.txt"

	var content = " Name : " + name + "\n Age : " + age + "\n Gender: " + gender

	writeFile(filename, content)
}

func writeFile(filename string, content string) {
	var file, err = os.Create(filename)
	if err != nil {
		panic(err)
	}
	length, err := io.WriteString(file, content)

	if err != nil {
		panic(err)
	}

	fmt.Println("File is written successfully :", length)
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)
}
