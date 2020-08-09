package controllers

import (
	"net/http"
	"regexp"
	"strings"

	"dba-scraper.com/pkg/models"
	"dba-scraper.com/pkg/scraper"
	"github.com/gin-gonic/gin"
)

const BASE_URL = "https://www.dba.dk/"

func Fetch(c *gin.Context) {
	data := getModelBasedOnURL(c)
	URL := GetFullRequestURLFromContext(c)
	err := scraper.Scrap(URL, &data)

	if err != nil {
		c.JSON(http.StatusAccepted, gin.H{"a-status": err.Error(), "data": data})
		return
	}

	c.JSON(http.StatusOK, gin.H{"a-status": "ok", "data": data})
}

func GetFullRequestURLFromContext(c *gin.Context) map[string]string {
	URL := make(map[string]string)

	URL["base"] = BASE_URL
	URL["path"] = strings.TrimLeft(c.Request.URL.Path, "/api")
	URL["query"] = createQueryString(c)

	return URL
}

func createQueryString(c *gin.Context) string {
	queryString := ""
	queryParams := c.Request.URL.Query()

	for key, value := range queryParams {
		if len(queryString) == 0 {
			queryString = "?" + key + "=" + value[0]
		} else {
			queryString = queryString + "&" + key + "=" + value[0]
		}
	}

	return queryString

}

func getModelBasedOnURL(c *gin.Context) models.GeneralModel {
	re := regexp.MustCompile(`/api/(?P<type>[a-z-_]+)`)
	model := re.FindAllStringSubmatch(c.Request.URL.Path, -1)[0][1]

	switch model {
	case "biler":
		return &models.Cars{}
	case "camping":
		return &models.Campings{}
	case "baade":
		return &models.Boats{}
	default:
		return &models.BasicCollection{}
	}
}
