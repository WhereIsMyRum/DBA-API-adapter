package models

import (
	"regexp"
	"strconv"
	"strings"

	"github.com/gocolly/colly"
)

type GeneralModel interface {
	Append(*colly.HTMLElement)
}

type CommonData struct {
	Id    string `json:"id"`
	Title string `json:"title"`
	Date  string `json:"date"`
	Price int    `json:"price"`
	URL   string `json:"url"`
}

func populateCommonData(e *colly.HTMLElement) CommonData {
	c := new(CommonData)

	c.Title = extractTitle(e)
	c.Date = extractDate(e)
	c.Price = extractPrice(e)
	c.URL = extractURL(e)
	c.Id = extractID(c.URL)

	return *c
}

func extractID(URL string) string {
	re := regexp.MustCompile(`/(?P<id>id-[0-9]+)`)
	return re.FindAllStringSubmatch(URL, -1)[0][1]
}

func extractTitle(e *colly.HTMLElement) string {
	title := e.ChildText("div.expandable-box")
	sliceLength := 56
	elipsis := "..."

	if len(title) < 56 {
		sliceLength = len(title)
		elipsis = ""
	}
	return title[:sliceLength] + elipsis
}

func extractDate(e *colly.HTMLElement) string {
	return e.ChildText("td[title='Dato']")
}

func extractPrice(e *colly.HTMLElement) int {
	replacer := strings.NewReplacer(".", "", ",", "")
	price, err := strconv.Atoi(replacer.Replace(strings.TrimSuffix(e.ChildText("td[title='Pris']"), " kr.")))
	return handleIntConversionErrors(price, err)
}

func extractMileage(e *colly.HTMLElement) int {
	replacer := strings.NewReplacer(".", "", ",", "")
	mileage, err := strconv.Atoi(replacer.Replace(e.ChildText("td[title='Km']")))
	return handleIntConversionErrors(mileage, err)
}

func extractYear(e *colly.HTMLElement) int {
	year, err := strconv.Atoi(e.ChildText("td[title='Modelår']"))
	return handleIntConversionErrors(year, err)
}

func extractFeet(e *colly.HTMLElement) int {
	feet, err := strconv.Atoi(e.ChildText("td[title='Fod']"))
	return handleIntConversionErrors(feet, err)
}

func handleIntConversionErrors(val int, err error) int {
	if err != nil {
		val = 0
	}
	return val
}

func extractURL(e *colly.HTMLElement) string {
	return e.ChildAttr("a", "href")
}
