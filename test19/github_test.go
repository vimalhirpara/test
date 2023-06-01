package github

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGitHubCallSuccess(t *testing.T) {

	// build our response JSON
	jsonResponse := `[{
			"full_name": "mock-repo"
		}]`

	// create a new reader with that JSON
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		_, _ = w.Write([]byte(jsonResponse))
	}))
	defer server.Close()

	ghm := GitHubManager{
		BaseUrl: server.URL,
	}

	result, err := ghm.GetRepos("v1/clb/virtual-account-created")
	if err != nil {
		t.Error("TestGitHubCallSuccess failed.")
		return
	}

	assert.True(t, len(result) > 0)
	assert.Equal(t, result[0]["full_name"], "mock-repo")
	fmt.Println(result)
}

func TestGitHubCallFail(t *testing.T) {
	// create a new reader with that JSON
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(404)
	}))
	defer server.Close()

	ghm := GitHubManager{
		BaseUrl: server.URL,
	}

	_, err := ghm.GetRepos("v1/clb/virtual-account-created")
	if err == nil {
		t.Error("TestGitHubCallFail failed.")
		return
	}

	assert.Equal(t, "EOF", err.Error())
}
