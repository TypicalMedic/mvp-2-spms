package clouddrive

import (
	"log"
	googleapi "mvp-2-spms/integrations/google-api"
	"strings"

	"google.golang.org/api/drive/v3"
	"google.golang.org/api/option"
)

const DAYS_PERIOD = 7
const HOURS_IN_DAY = 24
const EVENT_DURATION_HOURS = 1

type googleDriveApi struct {
	api *drive.Service
}

func InitDriveApi(googleAPI googleapi.GoogleAPI) googleDriveApi {
	api, err := drive.NewService(googleAPI.Context, option.WithHTTPClient(googleAPI.Client))
	c := googleDriveApi{api}
	if err != nil {
		log.Fatalf("Unable to retrieve Drive client: %v", err)
	}
	return c
}

func (d *googleDriveApi) CreateFolder(folderName string, parentFolder string) (*drive.File, error) {
	fileMetadata := &drive.File{
		Name:     folderName,
		MimeType: "application/vnd.google-apps.folder",
		Parents:  []string{parentFolder},
	}
	file, err := d.api.Files.Create(fileMetadata).Fields("id", "webViewLink").Do()
	if err == nil {
		return file, nil
	}
	return nil, err
}

func (d *googleDriveApi) AddTextFileToFolder(fileName string, fileText string, parentFolderId string) (*drive.File, error) {
	fileMetadata := &drive.File{
		Name:     fileName,
		MimeType: "application/vnd.google-apps.document", // google document type
		Parents:  []string{parentFolderId},
	}

	r := strings.NewReader(fileText)
	file, err := d.api.Files.Create(fileMetadata).Media(r).Fields("id").Do()
	if err == nil {
		return file, nil
	}
	return nil, err
}
