package googleapi

import (
	"context"
	"log"
	"net/http"
	"os"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

// токен пользователя надо хранить в бд или куках или где, и как работать с кучей пользователей со своими акичами??????????????
type GoogleAPI struct {
	Client  *http.Client
	Context context.Context
	config  *oauth2.Config
}

func InitGoogleAPI(scope ...string) GoogleAPI {
	ctx := context.Background()
	b, err := os.ReadFile("credentials.json")
	if err != nil {
		log.Fatalf("Unable to read client secret file: %v", err)
	} // If modifying these scopes, delete your previously saved token.json.
	config, err := google.ConfigFromJSON(b, scope...)
	if err != nil {
		log.Fatalf("Unable to parse client secret file to config: %v", err)
	}

	return GoogleAPI{Context: ctx, config: config}
}

func (g *GoogleAPI) GetAuthLink(redirectURI string, state string) string {
	// work with oauth state!!!
	g.config.RedirectURL = redirectURI
	authURL := g.config.AuthCodeURL(state, oauth2.AccessTypeOffline)
	return authURL
}

func (g *GoogleAPI) GetToken(authCode string) *oauth2.Token {
	tok, err := g.config.Exchange(context.TODO(), authCode)
	if err != nil {
		log.Fatalf("Unable to retrieve token from web: %v", err)
	}
	return tok
}

func (g *GoogleAPI) SetupClient(token *oauth2.Token) {
	g.Client = g.config.Client(context.Background(), token)
}
