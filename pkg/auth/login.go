package auth

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/AlecAivazis/survey/v2"
	"github.com/tensorleap/helm-charts/pkg/server"
	"github.com/tensorleap/leap-cli/pkg/api"
	"github.com/tensorleap/leap-cli/pkg/log"
)

func AskForUrl(defaultUrl string) (string, error) {
	if defaultUrl == "" {
		defaultUrl = fmt.Sprintf("http://localhost:%v", server.DefaultHttpPort)
	}
	prompt := survey.Input{
		Message: "Enter the API URL",
		Default: defaultUrl,
	}
	url := defaultUrl
	err := survey.AskOne(&prompt, &url)
	return url, err
}

func AskIfUseUserAndPassword() (bool, error) {
	useLogin := false
	prompt := &survey.Confirm{
		Message: "Login with username and password?",
		Default: true,
	}
	err := survey.AskOne(prompt, &useLogin)
	return useLogin, err
}

func AskForApiKey() (string, error) {
	apiKeyPrompt := survey.Password{
		Message: "Enter API key",
	}
	apiKey := ""
	err := survey.AskOne(&apiKeyPrompt, &apiKey)
	return apiKey, err
}

func AskForUserNameAndPassword(userName, password string) (string, string, error) {
	if userName == "" {
		userNamePrompt := survey.Input{
			Message: "Enter your username",
		}
		err := survey.AskOne(&userNamePrompt, &userName)
		if err != nil {
			return "", "", err
		}
	}
	if password == "" {
		passwordPrompt := survey.Password{
			Message: "Enter your password",
		}
		err := survey.AskOne(&passwordPrompt, &password)
		if err != nil {
			return "", "", err
		}
	}
	return userName, password, nil
}

func LoginAndGetAuthToken(apiUrl, username, password string) (string, error) {

	baseUrl := strings.TrimSuffix(apiUrl, API_PATH)
	clientID := "tensorleap-client"
	realm := "tensorleap"
	keycloakURL := fmt.Sprintf("%s/auth", baseUrl)
	token, err := loginToKeycloak(clientID, username, password, realm, keycloakURL)
	if err != nil {
		return "", err
	}
	jwtHeader := fmt.Sprintf("KBearer %s", token.AccessToken)
	getApiTokenUrl := fmt.Sprintf("%s/auth/keygen", apiUrl)

	req, err := http.NewRequest("POST", getApiTokenUrl, nil)
	if err != nil {
		return "", err
	}
	req.Header.Set("authorization", jwtHeader)

	log.Info("Getting API token...")

	resp, err := http.DefaultClient.Do(req)
	if err = api.CheckRes(resp, err); err != nil {
		return "", err
	}
	// the response is a json with key apiToken
	type ApiKeyResponse struct {
		ApiKey string `json:"apiKey"`
	}
	var apiKey ApiKeyResponse
	err = json.NewDecoder(resp.Body).Decode(&apiKey)
	if err != nil {
		return "", err
	}
	return apiKey.ApiKey, nil
}

type TokenResponse struct {
	AccessToken  string `json:"access_token"`
	TokenType    string `json:"token_type"`
	ExpiresIn    int    `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
}

func loginToKeycloak(clientID, username, password, realm, keycloakURL string) (*TokenResponse, error) {
	data := url.Values{}
	data.Set("client_id", clientID)
	data.Set("grant_type", "password")
	data.Set("username", username)
	data.Set("password", password)

	tokenURL := fmt.Sprintf("%s/realms/%s/protocol/openid-connect/token", keycloakURL, realm)

	resp, err := http.Post(tokenURL, "application/x-www-form-urlencoded", strings.NewReader(data.Encode()))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to login: %s", resp.Status)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var tokenResponse TokenResponse
	if err := json.Unmarshal(body, &tokenResponse); err != nil {
		return nil, err
	}

	return &tokenResponse, nil
}
