package manager

import (
	"strings"

	"github.com/olivere/elastic/v7"
)

// EsProdSearch sort
func EsProdSearchSort(sort string) (fieldSorts elastic.Sorter) {

	barArray := strings.Split(sort, "|")

	if len(barArray) == 2 {
		fieldSort := ""

		if barArray[0] == "1" {
			fieldSort = "price"
		} else if barArray[0] == "2" {
			fieldSort = "registerDate"
		} else {
			fieldSort = "price"
		}

		if barArray[1] == "1" {
			fieldSorts = elastic.NewFieldSort(fieldSort).Asc()
		} else if barArray[1] == "2" {
			fieldSorts = elastic.NewFieldSort(fieldSort).Desc()
		} else {
			fieldSorts = elastic.NewFieldSort(fieldSort).Asc()
		}
	}

	return
}
