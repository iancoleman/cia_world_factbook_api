package country

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"orderedmap"
	"path"
	"scraper"
	"strings"
	"time"
)

var NoPagesForCountryError = errors.New("No pages for country")
var NoPagesForTimeError = errors.New("No pages for this country and time")

var countryCache = map[string]Country{}

type Country struct {
	filename string
	files    []scraper.CountryFile
	pages    map[time.Time]Page
	jsons    map[time.Time]*orderedmap.OrderedMap
}

func ForFilename(f string) Country {
	// try to fetch country from cache
	c, isCached := countryCache[f]
	if isCached {
		return c
	}
	// create new country if not in cache
	c = Country{
		filename: f,
		pages:    map[time.Time]Page{},
		jsons:    map[time.Time]*orderedmap.OrderedMap{},
	}
	// get the files and dates for this country from the scraper
	c.files, _ = scraper.AllFilesForCountry(c.filename)
	// add new country to cache
	countryCache[f] = c
	return c
}

// Returns the most recent page before the specified time
func (c Country) PageForDate(t time.Time) (Page, error) {
	p := Page{}
	// check if no pages
	if len(c.files) == 0 {
		return p, NoPagesForCountryError
	}
	// check if no pages before the specified time
	defaultFile, _ := getEarliestFile(c.files)
	if defaultFile.ScrapedDate.Time.After(t) {
		return p, NoPagesForTimeError
	}
	// get the latest file which is before the specified time
	latestFile := defaultFile
	for _, f := range c.files {
		if f.ScrapedDate.Time.Before(t) || f.ScrapedDate.Time.Equal(t) {
			if f.ScrapedDate.Time.After(latestFile.ScrapedDate.Time) {
				latestFile = f
			}
		}
	}
	// check if this page has been cached
	pageDate := latestFile.ScrapedDate
	page, cachedPageExists := c.pages[pageDate.Time]
	// if the page is not cached, parse it and cache it
	if !cachedPageExists {
		filelocation := path.Join(pageDate.DirStr, latestFile.Filename)
		var err error
		page, err = NewPage(filelocation)
		if err != nil {
			return page, err
		}
		c.pages[pageDate.Time] = page
	}
	return page, nil
}

// Returns the most recent json before the specified time
func (c Country) JsonForDate(t time.Time, countryHtmlRoot, countryJsonRoot string) (*orderedmap.OrderedMap, string, error) {
	o := orderedmap.New()
	namekey := ""
	// check if no json
	if len(c.files) == 0 {
		return o, namekey, NoPagesForCountryError
	}
	// check if no pages before the specified time
	defaultFile, _ := getEarliestFile(c.files)
	if defaultFile.ScrapedDate.Time.After(t) {
		return o, namekey, NoPagesForTimeError
	}
	// get the latest file which is before the specified time
	latestFile := defaultFile
	for _, f := range c.files {
		if f.ScrapedDate.Time.Before(t) || f.ScrapedDate.Time.Equal(t) {
			if f.ScrapedDate.Time.After(latestFile.ScrapedDate.Time) {
				latestFile = f
			}
		}
	}
	// check if this page has been cached
	jsonDate := latestFile.ScrapedDate
	o, cachedJsonExists := c.jsons[jsonDate.Time]
	// if the json is not cached, parse it and cache it
	if !cachedJsonExists {
		jsonDateDir := strings.Replace(jsonDate.DirStr, countryHtmlRoot, countryJsonRoot, 1)
		filelocation := path.Join(jsonDateDir, latestFile.Filename+".json")
		var err error
		jsonBytes, err := ioutil.ReadFile(filelocation)
		if err != nil {
			return o, namekey, err
		}
		o = orderedmap.New()
		err = json.Unmarshal(jsonBytes, o)
		if err != nil {
			return o, namekey, err
		}
		// cache parsed result
		c.jsons[jsonDate.Time] = o
	}
	// get name key
	dataInterface, exists := o.Get("data")
	if !exists {
		return o, namekey, NoValueErr
	}
	data := dataInterface.(orderedmap.OrderedMap)
	name, exists := data.Get("name")
	if !exists {
		return o, namekey, NoValueErr
	}
	namekey = stringToJsonKey(name.(string))
	return o, namekey, nil
}

func (c Country) ClearCacheAfter(t time.Time) {
	// pages
	pageDatesToClear := []time.Time{}
	for d, _ := range c.pages {
		if d.After(t) {
			pageDatesToClear = append(pageDatesToClear, d)
		}
	}
	for _, d := range pageDatesToClear {
		delete(c.pages, d)
	}
	// jsons
	jsonDatesToClear := []time.Time{}
	for d, _ := range c.jsons {
		if d.After(t) {
			jsonDatesToClear = append(jsonDatesToClear, d)
		}
	}
	for _, d := range jsonDatesToClear {
		delete(c.jsons, d)
	}
}

func getEarliestFile(fs []scraper.CountryFile) (scraper.CountryFile, error) {
	e := scraper.CountryFile{}
	if len(fs) == 0 {
		return e, NoPagesForCountryError
	}
	e = fs[0]
	for _, f := range fs {
		if f.ScrapedDate.Time.Before(e.ScrapedDate.Time) {
			e = f
		}
	}
	return e, nil
}
