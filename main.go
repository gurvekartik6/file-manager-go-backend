package main

import "fmt"

func main() {
	fmt.Println("Hello, GO!")

	files := []string{
		"report.pdf",
		"document.pdf",
		"img1.jpg",
		"img2.jpg",
		"img3.png",
		"img4.png",
		"song1.mp3",
		"video1.mp4",
	}

	for _, fileName := range files {

		var category string

		switch {
		case len(fileName) >= 4:

			category = getCategory(fileName[len(fileName)-4:])

		default:
			category = "others"
		}

		fmt.Println("File Information : ")
		fmt.Println("File Name : ", fileName)
		fmt.Println("Category : ", category)
	}
}

func getCategory(extension string) string {

	categories := map[string]string{
		".pdf": "Documents",
		".txt": "Documents",
		".jpg": "Images",
		".png": "Images",
		".mp3": "Songs",
		".mp4": "Videos",
	}

	category, exists := categories[extension]

	if !exists {
		return "Others"
	}

	return category
}
