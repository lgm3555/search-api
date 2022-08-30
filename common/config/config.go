package config

func GetElasticUser() (url, username, password string) {
	url = "http://localhost:9200"
	username = "elastic"
	password = "changeme"
	//username = os.Getenv("ELASTIC_USERNAME")
	//password = os.Getenv("ELASTIC_PASSWORD")
	return
}
