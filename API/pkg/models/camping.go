package models

import (
	"github.com/gocolly/colly"
)

type Camping struct {
	Common CommonData `json:"general"`
	Year   int        `json:"year"`
}

type Campings struct {
	Camping []Camping `json:"Campings"`
}

func (c *Campings) Append(e *colly.HTMLElement) {
	c.Camping = append(c.Camping, c.parseHTML(e))
}

func (c *Campings) parseHTML(e *colly.HTMLElement) Camping {
	camping := new(Camping)

	camping.Common = populateCommonData(e)
	camping.Year = extractYear(e)

	return *camping
}
