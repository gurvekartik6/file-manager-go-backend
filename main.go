package main

import "fmt"

func main() {
	fmt.Println("Hello, GO!")

	fileName := "report"
	fileType := ".zip"
	fileSize := "2048"
	fileAvailable := true
	fileOwner := "Kartik Gurve"
	// fileFolder := "Documents"

	fmt.Println("---File Information---")

	fmt.Println("File Name : ", fileName)
	fmt.Println("File Type : ", fileType)
	fmt.Println("File Size : ", fileSize)
	fmt.Println("File Availability : ", fileAvailable)
	fmt.Println("File Owner : ", fileOwner)
	// fmt.Println("File Folder : ", fileFolder)

	switch fileType {
	case ".jpg", ".jpeg", ".png":
		fmt.Println("Images Folder")
	case ".mp3", ".mp2", ".mp1":
		fmt.Println("Songs Folder")
	case ".pdf", ".txt", ".docx":
		fmt.Println("Documents Folder")
	case ".mp4":
		fmt.Println("Videos Folder")
	default:
		fmt.Println("Other")
	}
}
