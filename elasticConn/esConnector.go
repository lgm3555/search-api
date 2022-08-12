package elasticConn

import (
	"fmt"
	"time"

	esclient "github.com/olivere/elastic/v7"

	"github.com/search-api/common/config"
)

var EsClient *esclient.Client
var Err error

func InitEsClient() {

	url, username, password := config.GetElasticUser()

	// es conncet
	EsClient, Err = esclient.NewClient(
		esclient.SetURL(url),
		esclient.SetBasicAuth(username, password),
		esclient.SetHealthcheckInterval(10*time.Second),
		esclient.SetSniff(false),
		esclient.SetMaxRetries(5))
	if Err != nil {
		// Handle error
		fmt.Println(Err)
	}
}
