package google_drive

import(
	"context"
	"google.golang.org/api/drive/v3"
	"google.golang.org/api/option"
)

func GoogleDriveService() (*drive.Service, error) {
	ctx := context.Background()
	oauthClient := GoogleDriveOauthClient()

	service, err := drive.NewService(ctx, option.WithHTTPClient(oauthClient))
	if err != nil {
		return nil, err
	}

	return service, nil
}
