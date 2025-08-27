package google_drive

import(
	"google.golang.org/api/drive/v3"
	"log"
	"fmt"
)

type GoogleDriveClient struct {
	Service *drive.Service
}



func (client *GoogleDriveClient) ListFiles(pageSize int64) {
	r, err := client.Service.Files.List().PageSize(10).Fields("nextPageToken, files(id, name)").Do()
	if err != nil {
		log.Fatalf("Unable to retrieve files: %v", err)
	}
	fmt.Println("Files:")

	if len(r.Files) == 0 {
		fmt.Println("No files found.")
	} else {
		for _, i := range r.Files {
			fmt.Printf("%s (%s)\n", i.Name, i.Id)
		}
	}
}
