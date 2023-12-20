package bitbucket

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type Source struct {
	Branch struct {
		Name string `json:"name"`
	} `json:"branch"`
}

type PullRequest struct {
	Title  string `json:"title"`
	Source Source `json:"source"`
	Id     int    `json:"id"`
}

type PullRequestResponse struct {
	Values []PullRequest `json:"values"`
}

func GetPullRequestsFromRepo(workspace string, repo string) ([]PullRequest, error) {
	var prs PullRequestResponse
	client := &http.Client{}

	req, err := MakeHttpRequest("GET", BaseURL+"/repositories/"+workspace+"/"+repo+"/pullrequests", nil)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := client.Do(req)
	if err != nil {
		return prs.Values, err
	}
	defer resp.Body.Close()

	if resp.StatusCode == 401 {
		log.Fatalf("Invalid credentials")
	}

	if resp.StatusCode == http.StatusNotFound {
		log.Fatalf("Repo %v not found in workspace %v", repo, workspace)
	}

	if resp.StatusCode != http.StatusOK {
		return prs.Values, fmt.Errorf("error: status code getting prs is %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return prs.Values, err
	}

	if err := json.Unmarshal(body, &prs); err != nil {
		return prs.Values, err
	}

	return prs.Values, nil
}
