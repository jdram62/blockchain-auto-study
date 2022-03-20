package services

import (
	"encoding/json"
	"github.com/gocolly/colly"
	"io/ioutil"
	"os"
	"strings"
	"time"
)

// UpdateLocalContractInfo
//		Automates retrieval ABI get when incompatible contract query
///*
func UpdateLocalContractInfo() {
	file, err := os.Open("watchlist.csv")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	//csvReader := csv.NewReader(file)
	//csvData, err := csvReader.ReadAll()
	//chFin := make(chan string)
	//fetchInfo(csvData)
}

func fetchInfo(csvData [][]string) {
	endPoints := GetEndpoints()
	theURL := endPoints[1][0]
	c := colly.NewCollector(colly.AllowedDomains(endPoints[1][1]), colly.Async(true))
	c.Limit(&colly.LimitRule{
		DomainGlob:  endPoints[1][1] + "/*",
		Delay:       1 * time.Second,
		Parallelism: 2,
	})
	c.OnHTML("button", func(e *colly.HTMLElement) {
		if e.Attr("aria-label") == "Copy Contract ABI" {
			var jsonMap []map[string]interface{}
			abiPulled := e.Attr("data-clipboard-text")
			json.Unmarshal([]byte(abiPulled), &jsonMap)
			address := (strings.Split(e.Request.URL.Path, "/"))[2]
			jsonString, _ := json.Marshal(jsonMap[:])
			ioutil.WriteFile(address+".json", jsonString, os.ModePerm)
			os.Mkdir("./contracts/uni/", os.ModePerm)
			os.Rename(address+".json", "./info/"+address+".json")
		}
	})
	for _, row := range csvData {
		c.Visit(theURL + row[1] + "/contracts")
	}
	c.Wait()

}
