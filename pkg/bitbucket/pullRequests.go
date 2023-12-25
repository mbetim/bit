package bitbucket

import (
	"fmt"
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

	resp, err := MakeHttpRequest("GET", BaseURL+"/repositories/"+workspace+"/"+repo+"/pullrequests", nil, &prs)
	if err != nil {
		log.Fatal(err)
	}

	if resp.StatusCode == http.StatusNotFound {
		return nil, fmt.Errorf("repo %v not found in workspace %v", repo, workspace)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error: status code getting prs is %d", resp.StatusCode)
	}

	return prs.Values, nil
}
