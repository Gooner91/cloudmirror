package google_drive

import(
	"log"
	"fmt"
	"context"
	"os"
	"net/http"
	"encoding/json"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/drive/v3"
	"golang.org/x/oauth2"
)

func GoogleDriveOauthClient() *http.Client{
	credentials, err := os.ReadFile("credentials.json")
	if err != nil {
		log.Fatalf("Unable to read client secret file: %v", err)
	}

	config, err := google.ConfigFromJSON(credentials, drive.DriveMetadataReadonlyScope)

	if err != nil {
		log.Fatalf("Unable to parse client secret file to config: %v", err)
	}

	client := getClient(config)

	return client
}

func getClient(config *oauth2.Config) *http.Client {
	tokenFile := "token.json"
	token, err := tokenFromFile(tokenFile)
	if err != nil {
		token = tokenFromWeb(config)
		saveToken(tokenFile, token)
	}
	return config.Client(context.Background(), token)
}

func tokenFromFile(file string) (*oauth2.Token, error){
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}

	defer f.Close()
	token := &oauth2.Token{}
	err = json.NewDecoder(f).Decode(token)
	return token, err
}

func tokenFromWeb(config *oauth2.Config) *oauth2.Token {
	authURL := config.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
	fmt.Printf("Go to the following link in your browser then type the "+
		"authorization code: \n%v\n", authURL)

	var authCode string

	if _, err := fmt.Scan(&authCode); err != nil {
		log.Fatalf("Unable to read authorization code %v", err)
	}

	token, err := config.Exchange(context.TODO(), authCode)

	if err != nil {
		log.Fatalf("Unable to retrieve token from web %v", err)
	}

	return token
}

func saveToken(file string, token *oauth2.Token) {
	fmt.Printf("Saving credential file to: %s\n", file)
	f, err := os.OpenFile(file, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
		if err != nil {
			log.Fatalf("Unable to cache oauth token: %v", err)
		}

		defer f.Close()
		json.NewEncoder(f).Encode(token)
}
