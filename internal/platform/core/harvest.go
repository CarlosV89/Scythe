package harvest

import "net/http"

type HarvestConfig struct {
	Id    string `yaml: "AccountID"`
	Token string `yaml: "Token"`
}

func request(
	requestType string,
	body struct,
	params struct,
	endpoint string
) {
	client := &http.Client{}
	url := Sprintf("https://api.harvestapp.com/v2/%s", endpoint)
	req, _ := http.NewRequest(requestType, url, nil)
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", viper.Get("Harvest.Token")))
	req.Header.Add("Harvest-Account-Id", viper.Get("Harvest.AccountId"))
	res, _ := client.Do(req)
}
