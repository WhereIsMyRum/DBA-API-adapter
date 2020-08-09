package main

import (
	"time"

	"client.com/pkg/searchers"
)

const Bed_URL = "http://scraper.docker/api/til-boligen/sengemoebler-og-udstyr/dobbeltsenge/type-dobbeltseng/?bredde=(120-140)&laengde=(200-200)&soegfra=2860&radius=4"
const Bed_FilePath = "./data-bed.json"

const Bike_URL = "http://scraper.docker/api/cykler/cykler-og-cykelanhaengere/herrecykler/?pris=(-1000)&soegfra=2860&radius=4"
const Bike_FilePath = "./data-bike.json"

const Nightstand_URL = "http://scraper.docker/api/til-boligen/sengemoebler-og-udstyr/andre-sovevaerelsesmoebler-og-udstyr/produkt-natbord/?pris=(-200)&soegfra=2860&radius=4"
const Nightstand_FilePath = "./data-nightstand.json"

func main() {
	for {
		go searchers.LookFor(Bed_URL, Bed_FilePath)
		go searchers.LookFor(Bike_URL, Bike_FilePath)
		go searchers.LookFor(Nightstand_URL, Nightstand_FilePath)

		time.Sleep(600 * time.Second)
	}

}
