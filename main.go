package main

import (
	"github.com/Gooner91/cloudmirror/cmd"
	// "fmt"
	// "github.com/Gooner91/cloudmirror/internal/google_drive"
)

func main() {
	cmd.Execute()
	// service, error := google_drive.GoogleDriveService()

	// if error != nil {
	// 	fmt.Println("Error fetching third party service object.")
	// }

	// client := google_drive.GoogleDriveClient{Service: service}
	// client.ListFiles(10)
}
