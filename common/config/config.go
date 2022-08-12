package config

import (
	"os"
)

func GetElasticUser() (url, username, password string) {
	url = "http://localhost:9200"
	username = os.Getenv("ELASTIC_SEARCH_USERNAME")
	password = os.Getenv("ELASTIC_SEARCH_PASSWORD")
	return
}
