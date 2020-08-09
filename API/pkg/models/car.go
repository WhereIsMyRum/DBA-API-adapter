package models

import (
	"github.com/gocolly/colly"
)

type Car struct {
	Common  CommonData `json:"general"`
	Mileage int        `json:"mileage"`
	Year    int        `json:"year"`
}

type Cars struct {
	Car []Car `json:"Cars"`
}

func (c *Cars) Append(e *colly.HTMLElement) {
	c.Car = append(c.Car, c.parseHTML(e))
}

func (c *Cars) parseHTML(e *colly.HTMLElement) Car {
	car := new(Car)

	car.Common = populateCommonData(e)
	car.Mileage = extractMileage(e)
	car.Year = extractYear(e)

	return *car
}
