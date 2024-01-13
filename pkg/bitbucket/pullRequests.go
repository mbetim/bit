package bitbucket

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
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

func GetPullRequestById(workspace string, repo string, prId int) (PullRequest, error) {
	var pr PullRequest

	resp, err := MakeHttpRequest("GET", BaseURL+"/repositories/"+workspace+"/"+repo+"/pullrequests/"+strconv.Itoa(prId), nil, &pr)
	if err != nil {
		return pr, err
	}

	if resp.StatusCode == http.StatusNotFound {
		return pr, fmt.Errorf("pull request %d from %v/%v not found", prId, workspace, repo)
	}

	if resp.StatusCode != http.StatusOK {
		return pr, fmt.Errorf("error: status code getting pr is %d", resp.StatusCode)
	}

	return pr, nil
}
