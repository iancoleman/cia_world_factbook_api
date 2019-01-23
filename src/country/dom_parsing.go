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
	// Prepare response
	s := ""
	// Find the heading node for this fieldkey
	selectorStr := "a[href*='fieldkey=" + selector.FieldKey + "']"
	links := doc.Find(selectorStr)
	if links.Length() < 1 {
		selectorStr = "a[href*='fields/" + selector.FieldKey + ".html']"
		links = doc.Find(selectorStr)
	}
	// if fieldkey has found a result, use it to get the text for this field
	if links.Length() == 1 {
		linkParent := links.First().ParentsFiltered("[class$='_light']").First()
		// Gather text from next nodes until heading is reached
		textNodes := linkParent.NextUntil("[class$='_light']")
		if linkParent.Length() < 1 {
			textNodes = links.First().Parent()
		}
		textNodes.Each(func(i int, node *goquery.Selection) {
			s = s + "\n" + node.Text()
		})
	}
	// if fieldkey has no result, try using id
	if s == "" {
		selectorStr = "#field-anchor-" + selector.Id
		links = doc.Find(selectorStr)
		if links.Length() < 1 {
			return "", IncorrectNumberOfFieldKeyLinks
		}
		rootLink := links.First()
		textNodes := rootLink.Next().Children()
		textNodes.Each(func(i int, node *goquery.Selection) {
			// get text for this node
			nodeText := node.Text()
			childText := ""
			shouldUseChildText := false
			// check if children mean we need to use a more complex extraction
			// eg <p> children should have \n separators
			childTextNodes := node.Children()
			childTextNodes.Each(func(i int, childNode *goquery.Selection) {
				childText = childText + " " + strings.TrimSpace(childNode.Text())
				if childNode.Is("p") {
					childText = childText + "\n"
					shouldUseChildText = true
				}
			})
			if shouldUseChildText {
				s = s + " " + childText
			} else {
				s = s + " " + nodeText
			}
		})
		// fix global rank value to be on same line as key
		s = globalRankSpaces.ReplaceAllString(s, "country comparison to the world: ")
		// fix newlines and spaces before parenthesis
		s = spacesBeforeParenthesis.ReplaceAllString(s, " (")
		// fix multiple newlines
		s = multipleNewlines.ReplaceAllString(s, "\n")
		// fix multiple spaces after colon
		s = spacesAfterColon.ReplaceAllString(s, ": ")
	}
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
