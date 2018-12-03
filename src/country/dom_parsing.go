package country

import (
	"errors"
	"github.com/PuerkitoBio/goquery"
	"strings"
)

var IncorrectNumberOfSelects = errors.New("Number of select elements != 1")
var IncorrectNumberOfFieldKeyLinks = errors.New("Number of fieldkey links != 1")
var IncorrectNumberOfAudioTags = errors.New("Number of audio tags != 1")
var NoCountryNameError = errors.New("Country Name not found in DOM")
var NoValueError = errors.New("No value found in DOM")
var NoSrcAttribute = errors.New("No value found for src attribute")

var FilenameBlacklist = []string{
	"fs.html", // French Southern and Antarctic Lands
	"um.html", // United States Pacific Island Wildlife Refuges
	"fq.html", // Baker Island
}

func countryListFromDom(doc *goquery.Document) ([]string, error) {
	l := []string{}
	// get the select element
	selects := doc.Find("select")
	// check there's only one (there should be only one)
	if selects.Length() < 1 {
		return l, IncorrectNumberOfSelects
	}
	// Extract the filenames for each country
	selects.Each(func(i int, s *goquery.Selection) {
		s.Find("option").Each(func(j int, o *goquery.Selection) {
			value, exists := o.Attr("value")
			if exists && value != "" {
				valueBits := strings.Split(value, "/")
				filename := valueBits[len(valueBits)-1]
				filename = strings.ToLower(filename)
				if !endsWith(filename, ".html") && len(filename) == 2 {
					filename = filename + ".html"
				}
				if !fileIsBlacklisted(filename) && endsWith(filename, ".html") {
					l = append(l, filename)
				}
			}
		})
	})
	return l, nil
}

func fileIsBlacklisted(f string) bool {
	for _, filename := range FilenameBlacklist {
		if f == filename {
			return true
		}
	}
	return false
}

func textForSelector(doc *goquery.Document, selector Selector) (string, error) {
	// Find the heading node for this fieldkey
	selectorStr := "a[href*='fieldkey=" + selector.FieldKey + "']"
	links := doc.Find(selectorStr)
	if links.Length() < 1 {
		selectorStr = "a[href*='fields/" + selector.FieldKey + ".html']"
		links = doc.Find(selectorStr)
		if links.Length() < 1 {
			return "", IncorrectNumberOfFieldKeyLinks
		}
	}
	linkParent := links.First().ParentsFiltered("[class$='_light']").First()
	// Gather text from next nodes until heading is reached
	textNodes := linkParent.NextUntil("[class$='_light']")
	if linkParent.Length() < 1 {
		textNodes = links.First().Parent()
	}
	s := ""
	textNodes.Each(func(i int, node *goquery.Selection) {
		s = s + "\n" + node.Text()
	})
	s = strings.TrimSpace(s)
	return s, nil
}

func countryNameFromDom(doc *goquery.Document) (string, error) {
	name := doc.Find(".countryName").First().Text()
	name = strings.TrimSpace(name)
	if len(name) == 0 {
		name = doc.Find(".region").First().Text()
		name = strings.TrimSpace(name)
		if len(name) == 0 {
			name = doc.Find("font[face='arial']").Eq(1).Text()
			name = strings.TrimSpace(name)
			if len(name) == 0 {
				return name, NoCountryNameError
			}
		}
	}
	name = strings.ToLower(name)
	name = strings.Title(name)
	return name, nil
}

func nationalAnthemMp3FromDom(doc *goquery.Document) (string, error) {
	root := "https://www.cia.gov/library/publications/the-world-factbook/"
	audio := doc.Find("audio")
	if audio.Length() != 1 {
		return "", IncorrectNumberOfAudioTags
	}
	url, exists := audio.First().Attr("src")
	if !exists {
		return "", NoSrcAttribute
	}
	url = root + strings.Replace(url, "../", "", -1)
	return url, nil
}
