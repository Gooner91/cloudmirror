package google_drive

import(
	"context"
	"google.golang.org/api/drive/v3"
	"google.golang.org/api/option"
)

type GoogleDriveClient struct {
	Service *drive.Service
}

func NewGoogleDriveClient() (*GoogleDriveClient, error) {
	context := context.Background()
	oauthClient := GoogleDriveOauthClient()

	service, error := drive.NewService(context, option.WithHTTPClient(oauthClient))
	if error != nil {
		return nil, error
	}

	return &GoogleDriveClient{Service: service}, nil
}

func (client *GoogleDriveClient) ListFiles(pageSize int64) ([]*drive.File, error) {
	response, error := client.Service.Files.List().PageSize(pageSize).Do()
	if error != nil {
		return nil, error
	}

	return response.Files, nil
}
