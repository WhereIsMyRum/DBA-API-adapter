package models

import (
	"github.com/gocolly/colly"
)

type Boat struct {
	Common CommonData `json:"general"`
	Feet   int        `json:"feet"`
	Year   int        `json:"year"`
}

type Boats struct {
	Boat []Boat `json:"Boats"`
}

func (b *Boats) Append(e *colly.HTMLElement) {
	b.Boat = append(b.Boat, b.parseHTML(e))
}

func (b *Boats) parseHTML(e *colly.HTMLElement) Boat {
	boat := new(Boat)

	boat.Common = populateCommonData(e)
	boat.Feet = extractFeet(e)
	boat.Year = extractYear(e)

	return *boat
}
