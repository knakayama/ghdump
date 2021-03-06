package credential

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/google/go-github/github"
	"github.com/knakayama/ghdump/utils"
	"golang.org/x/oauth2"
)

type Credential struct {
	User             string `json:"user"`
	OauthAccessToken string `json:"oauth_access_token"`
}

var credentialPath = filepath.Join(os.Getenv("HOME"), ".ghdump.json")

func GetCredential() (Credential, error) {

	if _, err := os.Stat(credentialPath); os.IsNotExist(err) {
		utils.Dieif(err)
	}

	return parseCredential()
}

func parseCredential() (Credential, error) {
	var credential Credential

	file, err := ioutil.ReadFile(credentialPath)
	utils.Dieif(err)
	err = json.Unmarshal(file, &credential)

	return credential, err
}

func GetGithubClient() (*github.Client, string) {
	credential, err := GetCredential()
	utils.Dieif(err)

	ts := oauth2.StaticTokenSource(
		&oauth2.Token{
			AccessToken: credential.OauthAccessToken,
		},
	)
	tc := oauth2.NewClient(oauth2.NoContext, ts)
	// TODO: error handling
	return github.NewClient(tc), credential.User
}
