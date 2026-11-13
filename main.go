package main

import "fmt"

func main() {
	fmt.Println("Hello, GO!")

	fileName := "report.pdf"
	fileType := "pdf"
	fileSize := "2048"
	fileAvailable := true
	fileOwner := "Kartik Gurve"
	fileFolder := "Documents"

	fmt.Println("---File Information---")

	fmt.Println("File Name : ", fileName)
	fmt.Println("File Type : ", fileType)
	fmt.Println("File Size : ", fileSize)
	fmt.Println("File Availability : ", fileAvailable)
	fmt.Println("File Owner : ", fileOwner)
	fmt.Println("File Folder : ", fileFolder)
}
