package scraper

import (
	"errors"
	"io/ioutil"
	"path"
	"sort"
	"time"
)

const dateFormat = "2006-01-02"
const dateUtcFormat = "2006-01-02 MST"

var noValidDateErr = errors.New("No valid date found")
var noFilesFound = errors.New("No files found")

var dates []ScrapedDate
var datesHaveBeenParsed = false

var filesForCountries = map[string][]CountryFile{}

type CountryFile struct {
	ScrapedDate ScrapedDate
	Filename    string
}

type ScrapedDate struct {
	DirStr  string
	TimeStr string
	Time    time.Time
}

type ByTime []ScrapedDate

func (t ByTime) Len() int           { return len(t) }
func (t ByTime) Swap(i, j int)      { t[i], t[j] = t[j], t[i] }
func (t ByTime) Less(i, j int) bool { return t[i].Time.Before(t[j].Time) }

func FirstDate(countryHtmlRoot string) (time.Time, error) {
	firstDate := time.Now().UTC()
	if !datesHaveBeenParsed {
		parseDatesFromDirs(countryHtmlRoot)
	}
	if len(dates) == 0 {
		return firstDate, noValidDateErr
	}
	return dates[0].Time, nil
}

func AllFilesForCountry(filename string) ([]CountryFile, error) {
	countryFiles, exists := filesForCountries[filename]
	if !exists {
		return countryFiles, noFilesFound
	}
	return countryFiles, nil
}

func parseDatesFromDirs(countryHtmlRoot string) {
	// list the dates
	// from stored format of ./pages/YYYY-MM-DD/ENCODED_URL
	files, err := ioutil.ReadDir(countryHtmlRoot)
	if err != nil {
		return
	}
	// get dates as strings
	for _, f := range files {
		// get date as string
		dirStr := f.Name()
		// convert to time.Time
		dirStrUtc := dirStr + " UTC"
		dirDate, err := time.Parse(dateUtcFormat, dirStrUtc)
		if err != nil {
			continue
		}
		// convert to ScrapedDate type
		fullDirStr := path.Join(countryHtmlRoot, dirStr)
		date := ScrapedDate{
			DirStr:  fullDirStr,
			TimeStr: dirStr,
			Time:    dirDate,
		}
		// add it to the list of dates
		dates = append(dates, date)
	}
	// sort dates with earliest first
	sort.Sort(ByTime(dates))
	datesHaveBeenParsed = true
	// cache dates for each country
	for _, date := range dates {
		fileinfos, err := ioutil.ReadDir(date.DirStr)
		if err != nil {
			continue
		}
		for _, fileinfo := range fileinfos {
			name := fileinfo.Name() // eg https%3A%2F...geos%2Fxx.html
			if len(name) < 7 {
				continue
			}
			filename := name[len(name)-7 : len(name)] // eg xx.html
			_, exists := filesForCountries[filename]
			if !exists {
				filesForCountries[filename] = []CountryFile{}
			}
			countryFile := CountryFile{
				ScrapedDate: date,
				Filename:    name,
			}
			countryFiles := filesForCountries[filename]
			countryFiles = append(countryFiles, countryFile)
			filesForCountries[filename] = countryFiles
		}
	}
}
