package main

import (
	"fmt"
	"github.com/Gooner91/cloudmirror/internal/google_drive"
)

func main(){

  service, error := google_drive.NewGoogleDriveClient()

	if error != nil {
		fmt.Println("No files found.")
	}

	service.ListFiles(10)
	// if len(r.Files) == 0 {
	// 	fmt.Println("No files found.")
	// } else {
	// 	for _, i := range r.Files {
	// 		fmt.Printf("%s (%s)\n", i.Name, i.Id)
	// 	}
	// }
}
