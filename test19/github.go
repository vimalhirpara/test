package github

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type GitHubManager struct {
	BaseUrl string
	Client  http.Client
}

// GetRepos takes a username and retreives
func (ghm *GitHubManager) GetRepos(username string) ([]map[string]interface{}, error) {
	url := fmt.Sprintf("%s/%s", ghm.BaseUrl, username)

	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	response, err := ghm.Client.Do(request)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	m := []map[string]interface{}{}
	err = json.NewDecoder(response.Body).Decode(&m)
	if err != nil {
		return nil, err
	}

	return m, nil
}
