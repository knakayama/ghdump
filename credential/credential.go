package credential

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/bgentry/speakeasy"
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
		SetCredential()
	}

	return parseCredential()
}

func SetCredential() {
	var user, pass string
	var err error

	if _, err := os.Stat(credentialPath); os.IsNotExist(err) == false {
		return
	}

	fmt.Print("Enter your github user name: ")
	fmt.Scanln(&user)

	pass, err = speakeasy.Ask("Enter your github password: ")
	utils.Dieif(err)

	token, err := getGitHubToken(user, pass)
	utils.Dieif(err)

	byt, _ := json.MarshalIndent(map[string]interface{}{
		"user":               user,
		"oauth_access_token": token,
	}, "", "  ")

	err = ioutil.WriteFile(credentialPath, byt, 0600)
	utils.Dieif(err)

	return
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

func getGitHubToken(user, pass string) (*oauth2.Token, error) {
	hostname, _ := os.Hostname()
	clientID := "ghdump for " + os.Getenv("USER") + "@" + hostname

	outhConf := &oauth2.Config{
		ClientID:     clientID,
		ClientSecret: "test",
		Scopes:       []string{"repo"},
		Endpoint:     githuboauth.Endpoint,
	}
	return oauthConf.PasswordCredentialsToken(oauth2.NoContext, user, pass)
}
