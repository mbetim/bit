package bitbucket

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/mbetim/bit/pkg/config"
)

const (
	BaseURL = "https://api.bitbucket.org/2.0"
)

func MakeHttpRequest(method string, url string, body io.Reader) (*http.Request, error) {
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

func GetRepoNameFromCurrentDir() (string, error) {
	file, err := os.Open(".git/config")
	if err != nil {
		return "", err
	}
	defer file.Close()

	var repoName string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		if strings.Contains(line, "url =") && strings.Contains(line, "bitbucket.org") {
			parts := strings.Split(line, "=")

			if len(parts) != 2 {
				return "", fmt.Errorf("unexpected format of .git/config")
			}

			url := strings.TrimSpace(parts[1])
			repoName = extractReponameFromUrl(url)
		}
	}

	if err := scanner.Err(); err != nil {
		return "", err
	}

	if strings.TrimSpace(repoName) == "" {
		return "", fmt.Errorf("repository URL not found")
	}

	return repoName, nil
}

func extractReponameFromUrl(url string) string {
	parts := strings.Split(url, "/")
	if len(parts) >= 2 {
		repo := parts[len(parts)-1]
		return strings.TrimSuffix(repo, ".git")
	}

	return ""
}
