package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"github.com/search-api/common/utils"
	elasticConn "github.com/search-api/elasticConn"
	"github.com/search-api/manager"
	"github.com/search-api/model"

	"github.com/olivere/elastic/v7"

	"github.com/search-api/response"
)

var logger = utils.Logger

func EsProdSearch(data map[string]string) (result response.Result, err error) {

	// bool
	query := elastic.NewBoolQuery()

	// must - match
	for _, subKeyword := range strings.Split(data["keyword"], " ") {
		matchQuery := elastic.NewMatchQuery("productName", subKeyword)
		query.Must(matchQuery)
	}

	// filter - range
	filter1 := elastic.NewRangeQuery("price").From(data["minPrice"]).To(data["maxPrice"])
	query.Filter(filter1)

	// sort
	fieldSort := manager.EsProdSearchSort(data["sort"])

	// from, size
	from, _ := strconv.Atoi(data["from"])
	size, _ := strconv.Atoi(data["size"])

	src, err := query.Source()
	if err != nil {
		fmt.Println(err)
	}

	q, err := json.MarshalIndent(src, "", "  ")
	if err == nil {
		fmt.Println("\"from\" : ", from)
		fmt.Println(",\" size\" : ", size)
		fmt.Println(",\" query\" : ", string(q))

		if fieldSort != nil {
			sortParam, _ := fieldSort.Source()
			sort, _ := json.MarshalIndent(sortParam, "", "  ")
			fmt.Println(",\"sort\" : ", "["+string(sort)+"]")
		}
	}

	cleint := elasticConn.EsClient

	res, err := cleint.Search().
		Index("prod").           // index 명
		Query(query).            // query
		From(from).Size(size).   // start - limit
		SortBy(fieldSort).       // sort
		Pretty(true).            // pretty print request and response JSON
		Do(context.Background()) // execute

	if err == nil {

		searchList := make([]interface{}, 0)

		for _, element := range res.Hits.Hits {
			searchEsList := &model.ProdSearchEs{}
			err := json.Unmarshal(element.Source, &searchEsList)

			if err != nil {
				logger.SetPrefix(utils.Error)
				logger.Println(err.Error())
			}

			searchList = append(searchList, searchEsList)
		}
		result.SearchList = searchList
	}

	return
}
