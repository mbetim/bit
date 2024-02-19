package bitbucket

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os/exec"
	"strings"

	"github.com/mbetim/bit/pkg/config"
)

const (
	BaseURL = "https://api.bitbucket.org/2.0"
)

func addTokenToRequest(method string, url string, body io.Reader) (*http.Request, error) {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return req, err
	}

	token, err := config.GetToken()
	if err != nil {
		return req, err
	}

	req.Header.Add("Authorization", "Basic "+token)

	return req, err
}

func MakeHttpRequest(method string, url string, body io.Reader, response interface{}) (*http.Response, error) {
	client := &http.Client{}

	req, err := addTokenToRequest(method, url, body)
	if err != nil {
		return nil, err
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusUnauthorized {
		log.Fatalf("Invalid credentials")
	}

	if resp.StatusCode != http.StatusOK {
		return resp, nil
	}

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return resp, err
	}

	if err := json.Unmarshal(respBody, &response); err != nil {
		return resp, err
	}

	return resp, nil
}

func GetRepoAndWorkspaceNameFromCli() (string, string, error) {
	commandOutput, err := exec.Command("bash", "-c", "git remote -v | awk '{print $2}'").Output()
	if err != nil {
		return "", "", err
	}

	url := strings.Split(string(commandOutput), "\n")
	repoName, workspaceName := extractRepoAndWorkspaceNameFromRemote(url[0])

	return repoName, workspaceName, nil
}

func extractRepoAndWorkspaceNameFromRemote(url string) (string, string) {
	if strings.HasPrefix(url, "https") {
		return extractRepoAndWorkspaceNameFromHttps(url)
	}

	return extractRepoAndWorkspaceNameFromSsh(url)
}

func extractRepoAndWorkspaceNameFromHttps(url string) (string, string) {
	urlParts := strings.Split(url, "/")

	repo := urlParts[len(urlParts)-1]
	workspace := urlParts[len(urlParts)-2]

	return strings.TrimSuffix(repo, ".git"), workspace
}

func extractRepoAndWorkspaceNameFromSsh(url string) (string, string) {
	workspaceRepoString := strings.Split(url, ":")[1]
	workspaceRepoParts := strings.Split(workspaceRepoString, "/")

	workspace := workspaceRepoParts[0]
	repo := strings.TrimSuffix(workspaceRepoParts[1], ".git")

	return repo, workspace
}
