package scraper

import (
	"errors"
	"fmt"
	"log"
	"strconv"

	"dba-scraper.com/pkg/models"
	"github.com/gocolly/colly"
)

const PAGE_LIMIT = 20

var pageNumber = 1

func Scrap(URL map[string]string, modelArray *models.GeneralModel) error {
	var collyInstance = colly.NewCollector()

	collyInstance.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	collyInstance.OnError(func(_ *colly.Response, err error) {
		log.Println("Something went wrong:", err)
	})

	collyInstance.OnHTML("tr.dbaListing", func(e *colly.HTMLElement) {
		(*modelArray).Append(e)
	})

	collyInstance.OnHTML("a[data-ga-lbl='paging-next']", func(element *colly.HTMLElement) {
		if checkForLongRunningQueries(pageNumber) {
			return
		}
		pageNumber = pageNumber + 1

		collyInstance.Visit(element.Request.AbsoluteURL(element.Attr("href")))
	})

	collyInstance.OnHTML("span[data-ga-lbl='paging-next']", func(element *colly.HTMLElement) {
		pageNumber = pageNumber + 1
		dynamicURL := URL["base"] + URL["path"] + "side-" + strconv.Itoa(pageNumber) + "/" + URL["query"]

		if checkForLongRunningQueries(pageNumber) {
			return
		}

		collyInstance.Visit(dynamicURL)
	})

	collyInstance.Visit(URL["base"] + URL["path"] + URL["query"])
	collyInstance.Wait()

	if pageNumber > 20 {
		pageNumber = 1
		return errors.New("Query is too long. Please modify your query")
	}

	pageNumber = 1
	return nil
}

func checkForLongRunningQueries(pageNumber int) bool {
	if pageNumber > 20 {
		return true
	}
	return false
}
