package harvest

import (
	"fmt"
	"net/http"

	"github.com/spf13/viper"
)

func request(
	requestType string,
	body map[string]string,
	params map[string]string,
	endpoint string,
) *http.Response {

	client := &http.Client{}
	url := fmt.Sprintf("https://api.harvestapp.com/v2/%s", endpoint)
	req, _ := http.NewRequest(requestType, url, nil)
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", viper.Get("Harvest.Token")))
	req.Header.Add("Harvest-Account-Id", fmt.Sprintf("%v", viper.Get("Harvest.AccountId")))
	res, _ := client.Do(req)
	return res
}
