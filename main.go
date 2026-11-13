package main

import "fmt"

func main() {
	fmt.Println("Hello, GO!")

	type FileRecord struct {
		fileName      string
		fileSize      int
		fileOwner     string
		fileType      string
		fileOrganized bool
	}

	file1 := FileRecord{
		fileName:  "Kartik_resume.pdf",
		fileSize:  2048,
		fileOwner: "Kartik Gurve",
	}

	fmt.Println(file1)
}
