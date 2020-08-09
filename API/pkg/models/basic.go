package models

import "github.com/gocolly/colly"

type BasicCollection struct {
	Common []CommonData `json:"general"`
}

func (m *BasicCollection) Append(e *colly.HTMLElement) {
	m.Common = append(m.Common, populateCommonData(e))
}
