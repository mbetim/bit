package bitbucket

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
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

func GetRepoAndWorkspaceNameFromCurrentDir() (string, string, error) {
	file, err := os.Open(".git/config")
	if err != nil {
		return "", "", err
	}
	defer file.Close()

	var repoName string
	var workspaceName string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		if strings.Contains(line, "url =") && strings.Contains(line, "bitbucket.org") {
			parts := strings.Split(line, "=")

			if len(parts) != 2 {
				return "", "", fmt.Errorf("unexpected format of .git/config")
			}

			url := strings.TrimSpace(parts[1])
			repoName, workspaceName = extractRepoAndWorkspaceNameFromUrl(url)
		}
	}

	if err := scanner.Err(); err != nil {
		return "", "", err
	}

	if strings.TrimSpace(repoName) == "" {
		return "", "", fmt.Errorf("repository URL not found")
	}

	return repoName, workspaceName, nil
}

func extractRepoAndWorkspaceNameFromUrl(url string) (string, string) {
	parts := strings.Split(url, "/")
	if len(parts) >= 2 {
		repo := parts[len(parts)-1]
		workspace := parts[len(parts)-2]
		return strings.TrimSuffix(repo, ".git"), workspace
	}

	return "", ""
}
