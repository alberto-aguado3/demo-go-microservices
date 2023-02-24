package http

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type HttpClient struct {
	client *http.Client
}

func NewClient() *HttpClient {
	return &HttpClient{
		client: http.DefaultClient,
	}
}

func (h *HttpClient) Get(url string, responseBody interface{}) error {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return err
	}

	req.Header.Add("someHeader", "valueForHeader")

	resp, err := h.client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	if err = json.Unmarshal(respBytes, &responseBody); err != nil {
		return err
	}
	return nil
}
