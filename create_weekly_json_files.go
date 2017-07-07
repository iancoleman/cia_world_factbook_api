package main

import (
	"country"
	"encoding/json"
	"io/ioutil"
	"logger"
	"orderedmap"
	"os"
	"path"
	"scraper"
	"time"
)

var countryHtmlRoot = ""
var countryJsonRoot = ""
var weeklyJsonRoot = ""

// Combines data for every country for every Monday
func main() {
	// read config
	configBytes, err := ioutil.ReadFile("config.json")
	if err != nil {
		logger.Stderr("Error reading config.json")
		logger.Stderr(err)
		return
	}
	// parse config
	config := map[string]string{}
	err = json.Unmarshal(configBytes, &config)
	var exists bool
	countryHtmlRoot, exists = config["country_html_root"]
	if !exists {
		logger.Stderr("Missing config value: country_html_root")
	}
	countryJsonRoot, exists = config["country_json_root"]
	if !exists {
		logger.Stderr("Missing config value: country_json_root")
	}
	weeklyJsonRoot, exists = config["weekly_json_root"]
	if !exists {
		logger.Stderr("Missing config value: weekly_json_root")
	}
	// get first date
	firstDate, err := scraper.FirstDate(countryHtmlRoot)
	if err != nil {
		logger.Stderr("Error getting first date")
		logger.Stderr(err)
		return
	}
	logger.Stdout("Earliest date found is", firstDate.Format("2006-01-02"))
	// find prior Monday
	now := time.Now().UTC()
	mondayToParse := mondayBefore(now)
	// parse every Monday back to the first date
	for mondayToParse.After(firstDate) || mondayToParse.Equal(firstDate) {
		logger.StdoutInline("Parsing ", mondayToParse.Format("2006-01-02"), " ")
		start := time.Now()
		parseForDate(mondayToParse)
		end := time.Now()
		logger.Stdout("took", end.Sub(start))
		mondayToParse = mondayToParse.Add(-7 * 24 * time.Hour)
	}
	logger.Stdout("Complete")
}

func parseForDate(d time.Time) {
	// get world page for this date
	w := country.ForFilename("xx.html")
	wp, err := w.PageForDate(d)
	if err != nil {
		logger.Stderr("Error getting page for world on date", d)
		logger.Stderr(err)
		return
	}
	// get country list for world on this date
	countryFilenames, err := wp.CountryList()
	if err != nil {
		logger.Stderr("Error getting country list for world on date", d)
		logger.Stderr(err)
		return
	}
	// prepare the data container to hold the parsed result
	countries := orderedmap.New()
	// iterate over countries and get json
	for _, f := range countryFilenames {
		c := country.ForFilename(f)
		if err != nil {
			logger.Stderr("Error getting country for", f, "on date", d)
			logger.Stderr(err)
			continue
		}
		// clear cache for any times older than this
		c.ClearCacheAfter(d)
		// get json for this country
		cj, namekey, err := c.JsonForDate(d, countryHtmlRoot, countryJsonRoot)
		if err != nil {
			logger.Stderr("Error getting json for", f, "on date", d)
			logger.Stderr(err)
			continue
		}
		// save the values for this country to the final result
		countries.Set(namekey, cj)
	}
	// prepare metadata
	metadata := orderedmap.New()
	metadata.Set("date", d.Format("2006-01-02"))
	metadata.Set("parser_version", country.VERSION)
	metadata.Set("parsed_time", time.Now().UTC().Format("2006-01-02 15:04:05 MST"))
	// save the parsed data
	parsed := orderedmap.New()
	parsed.Set("countries", countries)
	parsed.Set("metadata", metadata)
	content, err := json.MarshalIndent(parsed, "", "  ")
	if err != nil {
		logger.Stderr("Error marshalling parsed result to json")
		logger.Stderr(err)
		return
	}
	err = os.MkdirAll(weeklyJsonRoot, 0777)
	if err != nil {
		logger.Stderr("Error creating weeklyJsonRoot to store parsed result")
		logger.Stderr(err)
		return
	}
	filename := d.Format("2006-01-02") + "_factbook.json"
	filelocation := path.Join(weeklyJsonRoot, filename)
	err = ioutil.WriteFile(filelocation, []byte(content), 0664)
	if err != nil {
		logger.Stderr("Error saving parsed result")
		logger.Stderr(filelocation)
		logger.Stderr(err)
		return
	}
}

func mondayBefore(date time.Time) time.Time {
	daysDifference := (int(date.Weekday()-time.Monday) + 7) % 7
	if daysDifference == 0 {
		daysDifference = 7
	}
	priorDate := date.Add(time.Duration(daysDifference*-24) * time.Hour)
	return priorDate
}
