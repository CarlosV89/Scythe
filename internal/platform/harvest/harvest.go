package harvest

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

type Harvester interface {
	Request(
		requestType string,
		requestBody map[string]string,
		params map[string]string,
		endpoint string,
	) []byte
}

type harvester struct {
	c  http.Client // client
	t  string      // Authorization token
	id string      // Harvest account id
}

func NewHarvester(t string, id string) Harvester {
	return &harvester{http.Client{}, t, id}
}

func (h *harvester) Request(
	requestType string,
	requestBody map[string]string,
	params map[string]string,
	endpoint string,
) []byte {
	url := fmt.Sprintf("https://api.harvestapp.com/v2/%s", endpoint)
	req, _ := http.NewRequest(requestType, url, nil)
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", h.t))
	req.Header.Add("Harvest-Account-Id", fmt.Sprintf("%v", h.id))
	resp, _ := h.c.Do(req)
	body, _ := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	return body
}
