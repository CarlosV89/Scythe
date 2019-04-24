package harvest

type HarvestConfig struct {
	Id    string `yaml:"AccountID"`
	Token string `yaml: "Token"`
}

func writeConfigFile(id string, token string) {

}
