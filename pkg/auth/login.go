package auth

import (
	"context"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/AlecAivazis/survey/v2"
	"github.com/tensorleap/helm-charts/pkg/server"
	"github.com/tensorleap/leap-cli/pkg/api"
	"github.com/tensorleap/leap-cli/pkg/local"
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

func AskIfUseLogin() (bool, error) {
	useLogin := false
	prompt := &survey.Confirm{
		Message: "Login with username and password?",
		Default: true,
	}
	err := survey.AskOne(prompt, &useLogin)
	return useLogin, err
}

func AskIfOpenBrowser() (bool, error) {
	openBrowser := false
	prompt := &survey.Confirm{
		Message: "Open browser for authentication?",
		Default: true,
	}
	err := survey.AskOne(prompt, &openBrowser)
	return openBrowser, err
}

// the response is a json with key apiToken
type ApiKeyResponse struct {
	ApiKey string `json:"apiKey"`
}

func LoginAndGetAuthTokenWithBrowser(ctx context.Context, apiUrl string) (string, error) {
	baseUrl := api.ChangeToUIUrl(apiUrl)
	codeVerifier, err := generateCodeVerifier()
	if err != nil {
		return "", fmt.Errorf("failed to generate code verifier: %w", err)
	}
	codeChallenge := generateCodeChallenge(codeVerifier)

	url := fmt.Sprintf("%s/authorization-request?code_challenge=%s", baseUrl, codeChallenge)
	err = local.OpenLink(url)
	if err != nil {
		log.Warnf("Failed to open browser: %v", err)
	}
	log.Info("Please open the following link in your browser (if not opened automatically):")
	log.Info(url)
	var apiKey string

	err = api.WaitForCondition(ctx, "Waiting for authentication...", func() (bool, error) {
		url := fmt.Sprintf("%s/auth/getApiKeyByCode", apiUrl)
		type GetApiKeyByCode struct {
			Code string `json:"codeVerifier"`
		}

		body, err := json.Marshal(GetApiKeyByCode{Code: codeVerifier})
		if err != nil {
			return false, err
		}

		req, err := http.NewRequest("POST", url, strings.NewReader(string(body)))
		if err != nil {
			return false, err
		}
		req.Header.Set("Content-Type", "application/json")
		res, err := http.DefaultClient.Do(req)
		if err != nil {
			return false, err
		}

		if res.StatusCode == http.StatusOK {
			var r ApiKeyResponse
			err = json.NewDecoder(res.Body).Decode(&r)
			if err != nil {
				return false, err
			}
			apiKey = r.ApiKey
			return true, nil

		}
		return false, nil

	}, time.Second*3, time.Minute*5)

	if err != nil {
		return "", fmt.Errorf("failed to get API key: %w", err)
	}
	if apiKey == "" {
		return "", fmt.Errorf("authentication failed")
	}
	return apiKey, nil
}

// Generate a code verifier: a cryptographically random string
func generateCodeVerifier() (string, error) {
	verifier := make([]byte, 32)
	if _, err := rand.Read(verifier); err != nil {
		return "", err
	}
	// Base64-url encode without padding
	return base64.RawURLEncoding.EncodeToString(verifier), nil
}

// Generate the code challenge by hashing the verifier with SHA-256 and base64-url encoding
func generateCodeChallenge(verifier string) string {
	hash := sha256.Sum256([]byte(verifier))
	// Base64-url encode the hash without padding
	return base64.RawURLEncoding.EncodeToString(hash[:])
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

	baseUrl := api.ChangeToUIUrl(apiUrl)
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
