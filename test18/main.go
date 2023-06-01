package api

import (
	"io/ioutil"
	"net/http"
)

type API struct {
	Client  *http.Client
	baseURL string
}

func (api *API) DoStuff() ([]byte, error) {
	resp, err := api.Client.Get(api.baseURL + "/VirtualAccountV2")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	// handling error and doing stuff with body that needs to be unit tested
	return body, err
}
