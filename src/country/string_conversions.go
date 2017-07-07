package country

import (
	"errors"
	"orderedmap"
	"regexp"
	"strcase"
	"strconv"
	"strings"
	"time"
)

var atLeastOneSpaceRe = regexp.MustCompile(`\s+`)
var startsWithSpaceOrParenthesis = regexp.MustCompile(`^[\s\(\)]+`)
var endsWithSpaceOrParenthesis = regexp.MustCompile(`[\s\(\)]+$`)
var ageStructureSplitterRe = regexp.MustCompile(`[:%\(\)/]+`)
var commaOrSemicolon = regexp.MustCompile(`[,;]+`)

var StringToMapErr = errors.New("String could not be converted to map")
var StringToMapOfNumbersErr = errors.New("String could not be converted to map of numbers")
var StringToNumberWithUnitsErr = errors.New("String could not be converted to number with units")
var StringToPercentageErr = errors.New("String could not be converted to percentage")
var StringToGPSErr = errors.New("String could not be converted to GPS")
var StringToPlaceAndNumberWithUnitsErr = errors.New("String could not be converted to place and number with units")
var StringIsNaErr = errors.New("String is NA")

type listConditions struct {
	keepAnds   bool
	splitChars string
}

func stringToJsonKey(s string) string {
	s = strings.Replace(s, ",", "", -1)
	s = strings.Replace(s, "(", "", -1)
	s = strings.Replace(s, ")", "", -1)
	s = strings.Replace(s, "-", "_", -1)
	s = strings.Replace(s, "/", "_", -1)
	s = strings.Replace(s, "\\", "", -1)
	return strcase.ToSnake(s)
}

func trimSpaceAndParenthesis(s string) string {
	trimmed := startsWithSpaceOrParenthesis.ReplaceAllString(s, "")
	trimmed = endsWithSpaceOrParenthesis.ReplaceAllString(trimmed, "")
	// remove any unclosed or unopened parenthesis
	firstOpen := strings.Index(trimmed, "(")
	firstClose := strings.Index(trimmed, ")")
	lastOpen := strings.LastIndex(trimmed, "(")
	lastClose := strings.LastIndex(trimmed, ")")
	// remove unclosed open parenthesis
	if lastOpen > lastClose && lastOpen > -1 && lastClose > -1 {
		trimmed = trimmed[0:lastOpen] + trimmed[lastOpen+1:len(trimmed)]
	}
	// remove unopened closed parenthesis
	if firstClose < firstOpen && firstClose > -1 && firstOpen > -1 {
		trimmed = trimmed[0:firstClose] + trimmed[firstClose+1:len(trimmed)]
	}
	return trimmed
}

func stringToMap(s string) (*orderedmap.OrderedMap, error) {
	o := orderedmap.New()
	lines := strings.Split(s, "\n")
	hasFirstKey := false
	currentKey := ""
	currentValue := ""
	for _, line := range lines {
		lineBits := strings.Split(line, ":")
		lineHasKey := len(lineBits) >= 2                  // must have a :
		lineHasKey = lineHasKey && len(lineBits[0]) < 100 // first bit (ie key) must be less than 100 chars
		// handle line without key
		if !lineHasKey {
			// handle multiline value
			if hasFirstKey {
				currentValue = currentValue + "\n" + line
			}
			continue
		}
		// handle previous key value
		if hasFirstKey {
			value := strings.TrimSpace(currentValue)
			if len(value) > 0 {
				o.Set(currentKey, value)
			}
			currentKey = ""
			currentValue = ""
		}
		// start building this key value pair
		hasFirstKey = true
		currentKey = stringToJsonKey(strings.TrimSpace(lineBits[0]))
		thisLineValue := strings.TrimSpace(strings.Join(lineBits[1:len(lineBits)], ":"))
		currentValue = strings.TrimSpace(currentValue + thisLineValue)
	}
	if hasFirstKey {
		value := strings.TrimSpace(currentValue)
		if len(value) > 0 {
			o.Set(currentKey, value)
		}
	}
	if len(o.Keys()) == 0 {
		return o, StringToMapErr
	}
	return o, nil
}

func stringToMapOfNumbersWithUnits(s string) (*orderedmap.OrderedMap, error) {
	s, date, hasDate := stringWithoutDate(s)
	o, err := stringToMap(s)
	if err != nil {
		return o, StringToMapOfNumbersErr
	}
	keys := o.Keys()
	keysToDelete := []string{}
	for _, key := range keys {
		valueInterface, _ := o.Get(key)
		valueStr := valueInterface.(string)
		// handle exceptional case for global rank
		if key == "country_comparison_to_the_world" {
			valueInt, err := strconv.Atoi(valueStr)
			if err != nil {
				continue
			}
			o.Set("global_rank", valueInt)
			keysToDelete = append(keysToDelete, key)
			continue
		}
		// handle note
		if key == "note" {
			continue
		}
		// handle all other keys
		valueNum, err := stringToNumberWithUnits(valueStr)
		if err == StringIsNaErr {
			keysToDelete = append(keysToDelete, key)
		} else if err == nil {
			o.Set(key, valueNum)
		}
	}
	// remove invalid values
	for _, key := range keysToDelete {
		o.Delete(key)
	}
	if hasDate {
		o.Set("date", date)
	}
	return o, nil
}

func stringToMapOfNumbers(s string) (*orderedmap.OrderedMap, error) {
	s, date, hasDate := stringWithoutDate(s)
	o, err := stringToMap(s)
	if err != nil {
		return o, StringToMapOfNumbersErr
	}
	keys := o.Keys()
	keysToDelete := []string{}
	for _, key := range keys {
		// handle exceptional case for global rank
		if key == "country_comparison_to_the_world" {
			valueStr, _ := o.Get(key)
			valueInt, err := strconv.Atoi(valueStr.(string))
			if err != nil {
				continue
			}
			o.Set("global_rank", valueInt)
			o.Delete(key)
			continue
		}
		if key == "note" {
			continue
		}
		// handle all other keys
		valueStr, _ := o.Get(key)
		valueNum, err := stringToNumber(valueStr.(string))
		if err == StringIsNaErr {
			keysToDelete = append(keysToDelete, key)
		} else if err == nil {
			o.Set(key, valueNum)
		}
	}
	// remove invalid values
	for _, key := range keysToDelete {
		o.Delete(key)
	}
	if hasDate {
		o.Set("date", date)
	}
	return o, nil
}

func stringToNumberWithUnits(s string) (*orderedmap.OrderedMap, error) {
	o := orderedmap.New()
	o.Set("value", 0.0)
	o.Set("units", "")
	sTrimmed := strings.TrimSpace(s)
	// check string for invalid values
	if sTrimmed == "NA" {
		return o, StringIsNaErr
	}
	// convert percentages into units
	sTrimmed = strings.Replace(sTrimmed, "%", " %", -1)
	// split number into bits
	bits := atLeastOneSpaceRe.Split(sTrimmed, -1)
	// get the number component
	var value float64
	foundValue := false
	startedNote := false
	magnitude := 1.0
	units := ""
	note := ""
	for _, bit := range bits {
		noCommas := strings.Replace(bit, ",", "", -1)
		possibleValue, err := strconv.ParseFloat(noCommas, 64)
		if err == nil && !foundValue {
			value = possibleValue
			foundValue = true
			units = ""
		} else if bit == "thousand" && !startedNote {
			magnitude = 1e3
			units = ""
		} else if bit == "million" && !startedNote {
			magnitude = 1e6
			units = ""
		} else if bit == "billion" && !startedNote {
			magnitude = 1e9
			units = ""
		} else if bit == "trillion" && !startedNote {
			magnitude = 1e12
			units = ""
		} else if isUnitsStr(bit) && !startedNote {
			units = units + " " + bit
		} else if foundValue {
			note = note + " " + bit
			startedNote = true
		}
	}
	if !foundValue {
		return o, StringToNumberWithUnitsErr
	}
	// tidy up the units and note
	units = strings.TrimSpace(units)
	note = trimSpaceAndParenthesis(note)
	// convert the number using the magnitude
	value = value * magnitude
	// bundle into a single map
	o.Set("value", value)
	if units != "" {
		o.Set("units", units)
	} else {
		o.Delete("units")
	}
	if len(note) > 0 {
		o.Set("note", note)
	}
	return o, nil
}

func stringToPercentage(s string) (*orderedmap.OrderedMap, error) {
	o := orderedmap.New()
	o.Set("value", 0.0)
	o.Set("units", "")
	s, date, hasDate := stringWithoutDate(s)
	sTrimmed := strings.TrimSpace(s)
	// check string for invalid values
	if sTrimmed == "NA" {
		return o, StringIsNaErr
	}
	// split number into bits
	bits := strings.SplitN(s, "%", 2)
	// check for no value
	if len(bits) == 0 {
		return o, NoValueErr
	}
	// get the value
	numStr := bits[0]
	noCommas := strings.Replace(numStr, ",", "", -1)
	value, err := strconv.ParseFloat(noCommas, 64)
	if err != nil {
		return o, StringToPercentageErr
	}
	// get the note
	note := ""
	if len(bits) > 1 {
		note = strings.Join(bits[1:len(bits)], " ")
		note = trimSpaceAndParenthesis(note)
		if len(note) > 5 && note[0:5] == "note:" {
			note = note[5:len(note)]
			note = trimSpaceAndParenthesis(note)
		}
	}
	// bundle into a single map
	o.Set("value", value)
	o.Set("units", "%")
	if len(note) > 0 {
		o.Set("note", note)
	}
	if hasDate {
		o.Set("date", date)
	}
	return o, nil
}

func isUnitsStr(s string) bool {
	isUnits := false
	isUnits = isUnits || s == "km"
	isUnits = isUnits || s == "sq"
	isUnits = isUnits || s == "nm"
	isUnits = isUnits || s == "m"
	isUnits = isUnits || s == "years"
	isUnits = isUnits || s == "year"
	isUnits = isUnits || s == "deaths_per_1000_live_births"
	isUnits = isUnits || s == "%"
	isUnits = isUnits || s == "USD"
	isUnits = isUnits || s == "kWh"
	isUnits = isUnits || s == "cu"
	return isUnits
}

func stringToGPS(s string) (*orderedmap.OrderedMap, error) {
	o := orderedmap.New()
	sTrimmed := strings.TrimSpace(s)
	bits := atLeastOneSpaceRe.Split(sTrimmed, -1)
	if len(bits) != 6 {
		return o, StringToGPSErr
	}
	// latitude
	latitude := orderedmap.New()
	// latitude degrees
	latD, err := strconv.Atoi(bits[0])
	if err != nil {
		return o, StringToGPSErr
	}
	latitude.Set("degrees", latD)
	// latitude minutes
	latM, err := strconv.Atoi(bits[1])
	if err != nil {
		return o, StringToGPSErr
	}
	latitude.Set("minutes", latM)
	// latitude hemisphere
	latH := strings.Replace(bits[2], ",", "", -1)
	latitude.Set("hemisphere", latH)
	// longitude
	longitude := orderedmap.New()
	// longitude degrees
	longD, err := strconv.Atoi(bits[3])
	if err != nil {
		return o, StringToGPSErr
	}
	longitude.Set("degrees", longD)
	// longitude minutes
	longM, err := strconv.Atoi(bits[4])
	if err != nil {
		return o, StringToGPSErr
	}
	longitude.Set("minutes", longM)
	// longitude hemisphere
	longH := strings.Replace(bits[5], ",", "", -1)
	longitude.Set("hemisphere", longH)
	// save in top level map
	o.Set("latitude", latitude)
	o.Set("longitude", longitude)
	return o, nil
}

func firstLine(s string) (string, string) {
	sTrimmed := strings.TrimSpace(s)
	lines := strings.SplitN(sTrimmed, "\n", 2)
	for len(lines) < 2 {
		lines = append(lines, "")
	}
	return lines[0], lines[1]
}

func stringToPlaceAndNumberWithUnits(s, placekey, numberkey string) (*orderedmap.OrderedMap, error) {
	o := orderedmap.New()
	sTrimmed := strings.TrimSpace(s)
	bits := atLeastOneSpaceRe.Split(sTrimmed, -1)
	if len(bits) < 3 {
		return o, StringToPlaceAndNumberWithUnitsErr
	}
	// find index of number
	indexOfNumber := 0
	for i, bit := range bits {
		bit = strings.Replace(bit, ",", "", -1)
		_, err := strconv.ParseFloat(bit, 64)
		if err == nil {
			indexOfNumber = i
			break
		}
	}
	nameStrs := bits[0:indexOfNumber]
	name := strings.Join(nameStrs, " ")
	numberStrs := bits[indexOfNumber:len(bits)]
	numberStr := strings.Join(numberStrs, " ")
	number, err := stringToNumberWithUnits(numberStr)
	if err != nil {
		return o, err
	}
	o.Set(placekey, name)
	o.Set(numberkey, number)
	// move note from number to top-level map if present
	note, exists := number.Get("note")
	if exists {
		o.Set("note", note)
		number.Delete("note")
	}
	return o, nil
}

func borderCountriesStringToSlice(s string) ([]*orderedmap.OrderedMap, error) {
	c := []*orderedmap.OrderedMap{}
	countryStrs := strings.Split(s, ", ")
	for _, countryStr := range countryStrs {
		country, err := stringToPlaceAndNumberWithUnits(countryStr, "country", "border_length")
		if err != nil {
			return c, err
		}
		c = append(c, country)
	}
	return c, nil
}

func stringToList(s string, lc listConditions) ([]string, error) {
	if s == "none" {
		return []string{}, nil
	}
	if !lc.keepAnds {
		s = strings.Replace(s, " and ", ", ", -1)
	}
	splitterRe := regexp.MustCompile("[,]+")
	var err error
	if len(lc.splitChars) > 0 {
		// TODO consider [ or ] in splitChars
		splitterRe, err = regexp.Compile("[" + lc.splitChars + "]+")
		if err != nil {
			return []string{}, err
		}
	}
	list := splitterRe.Split(s, -1)
	cleanList := []string{}
	for _, item := range list {
		cleanItem := strings.TrimSpace(item)
		if len(cleanItem) > 0 {
			cleanList = append(cleanList, cleanItem)
		}
	}
	return cleanList, nil
}

func stringToPercentageMap(s, datakey string) (*orderedmap.OrderedMap, error) {
	// remove the date
	s, date, hasDate := stringWithoutDate(s)
	// convert entire string to a map
	raw, err := stringToMap(s)
	if err != nil {
		return raw, err
	}
	// prepare maps to keep filtered data
	data := orderedmap.New()
	metadata := orderedmap.New()
	final := orderedmap.New()
	// filter data to either the data, or metadata
	keys := raw.Keys()
	for _, key := range keys {
		value, _ := raw.Get(key)
		num, err := stringToPercentage(value.(string))
		if err == nil {
			data.Set(key, num)
		} else {
			metadata.Set(key, value)
		}
	}
	// add data to the root map
	dataKeys := data.Keys()
	if len(dataKeys) > 0 {
		final.Set(datakey, data)
	}
	// add any metadata (to retain data as the first key)
	metadataKeys := metadata.Keys()
	for _, key := range metadataKeys {
		value, _ := metadata.Get(key)
		final.Set(key, value)
	}
	// if date is there, add it to the root map
	if hasDate {
		final.Set("date", date)
	}
	return final, nil
}

func stringWithoutDate(s string) (string, string, bool) {
	_, ps := removeParenthesis(s)
	dateStr := ""
	sNoDate := s
	for _, p := range ps {
		pTidy := strings.Replace(p, " est.", "", -1)
		pTidy = strings.Replace(pTidy, " es", "", -1)
		bits := strings.Split(pTidy, " ")
		if len(bits) == 1 {
			if len(pTidy) == 7 && pTidy[4] == '/' {
				pTidy = pTidy[0:4]
			}
			t, err := time.Parse("2006", pTidy)
			if err != nil {
				continue
			}
			dateStr = t.Format("2006")
			sNoDate = strings.Replace(sNoDate, "("+p+")", "", -1)
			break
		} else if len(bits) == 2 {
			t, err := time.Parse("January 2006", pTidy)
			if err != nil {
				continue
			}
			dateStr = t.Format("2006-01-02")
			sNoDate = strings.Replace(sNoDate, "("+p+")", "", -1)
			break
		} else if len(bits) == 3 {
			t, err := time.Parse("2 January 2006", pTidy)
			if err != nil {
				continue
			}
			dateStr = t.Format("2006-01-02")
			sNoDate = strings.Replace(sNoDate, "("+p+")", "", -1)
			break
		}
	}
	hasDate := dateStr != ""
	sNoDate = strings.TrimSpace(sNoDate)
	return sNoDate, dateStr, hasDate
}

func removeParenthesis(s string) (string, []string) {
	ps := []string{}
	sNoPs := ""
	p := ""
	depth := 0
	for _, r := range s {
		if r == ')' {
			depth = depth - 1
			if depth > 0 {
				p = p + string(r)
			} else {
				ps = append(ps, p)
				p = ""
			}
		} else if depth > 0 {
			p = p + string(r)
		} else if r == '(' {
			if depth > 0 {
				p = p + string(r)
			}
			depth = depth + 1
		} else {
			sNoPs = sNoPs + string(r)
		}
	}
	return sNoPs, ps
}

func convertToFranceValue(s string) string {
	if strings.Index(s, "metropolitan France: ") == -1 {
		return s
	}
	o, err := stringToMap(s)
	if err != nil {
		return s
	}
	v, ok := o.Get("metropolitan_france")
	if !ok {
		return s
	}
	return v.(string)
}

func stringToNumberWithUnitsAndDate(s string) (*orderedmap.OrderedMap, error) {
	s, date, hasDate := stringWithoutDate(s)
	num, err := stringToNumberWithUnits(s)
	if err != nil {
		return num, err
	}
	if hasDate {
		num.Set("date", date)
	}
	return num, nil
}

func stringToNumber(s string) (float64, error) {
	clean := s
	clean = strings.TrimSpace(clean)
	// remove range, use first value only
	rangeBits := strings.Split(clean, "%")
	if len(rangeBits) > 1 {
		clean = rangeBits[0]
	}
	if strings.Index(clean, "-") > 0 {
		rangeBits = strings.Split(clean, "-")
		clean = rangeBits[0]
	}
	// clean it up
	clean = strings.Replace(clean, "$", "", -1)
	clean = strings.Replace(clean, "<", "", -1)
	clean = strings.Replace(clean, ">", "", -1)
	clean = strings.TrimSpace(clean)
	bits := atLeastOneSpaceRe.Split(clean, -1)
	if len(bits) == 0 {
		return 0, NoValueErr
	}
	// split number into bits
	// get the number component
	var value float64
	foundValue := false
	magnitude := 1.0
	for _, bit := range bits {
		noCommas := strings.Replace(bit, ",", "", -1)
		possibleValue, err := strconv.ParseFloat(noCommas, 64)
		if err == nil && !foundValue {
			value = possibleValue
			foundValue = true
		} else if bit == "thousand" {
			magnitude = 1e3
		} else if bit == "million" {
			magnitude = 1e6
		} else if bit == "billion" {
			magnitude = 1e9
		} else if bit == "trillion" {
			magnitude = 1e12
		}
	}
	if !foundValue {
		return value, StringToNumberWithUnitsErr
	}
	// convert the number using the magnitude
	value = value * magnitude
	return value, nil
}

func stringToNumberWithGlobalRankAndDate(s, numberKey string) (*orderedmap.OrderedMap, error) {
	o := orderedmap.New()
	// get date
	s, date, hasDate := stringWithoutDate(s)
	// get number
	firstLine, otherLines := firstLine(s)
	number, err := stringToNumber(firstLine)
	if err != nil {
		return o, err
	}
	o.Set(numberKey, number)
	// get global rank
	lines := strings.Split(otherLines, "\n")
	for _, line := range lines {
		if len(line) > 31 && line[0:31] == "country comparison to the world" {
			rankStr := strings.Replace(line, "country comparison to the world:", "", -1)
			rankStr = strings.TrimSpace(rankStr)
			rankInt, err := strconv.Atoi(rankStr)
			if err == nil {
				o.Set("global_rank", rankInt)
			}
		}
	}
	// set date
	if hasDate {
		o.Set("date", date)
	}
	return o, nil
}

func stringToListOfAnnualValues(s, units string) (*orderedmap.OrderedMap, error) {
	o := orderedmap.New()
	annualValues := []*orderedmap.OrderedMap{}
	lines := strings.Split(s, "\n")
	globalRank := 0.0
	hasGlobalRank := false
	notes := []string{}
	for _, line := range lines {
		isNumberLine := false
		isNumberLine = isNumberLine || strings.Index(line, "$") == 0
		isNumberLine = isNumberLine || strings.Index(line, "-$") == 0
		isNumberLine = isNumberLine || strings.Index(line, "%") > -1
		isNumberLine = isNumberLine || startsWithNumber(line)
		if isNumberLine {
			line, date, hasDate := stringWithoutDate(line)
			v, err := stringToNumber(line)
			if err != nil {
				continue
			}
			annualValue := orderedmap.New()
			annualValue.Set("value", v)
			annualValue.Set("units", units)
			if hasDate {
				annualValue.Set("date", date)
			}
			annualValues = append(annualValues, annualValue)
		} else if startsWith(line, "note") {
			m, err := stringToMap(line)
			if err != nil {
				continue
			}
			keys := m.Keys()
			for _, k := range keys {
				v, _ := m.Get(k)
				notes = append(notes, v.(string))
			}
		} else if startsWith(line, "country comparison to the world") {
			m, err := stringToMap(line)
			if err != nil {
				continue
			}
			keys := m.Keys()
			for _, k := range keys {
				v, _ := m.Get(k)
				vFloat, err := stringToNumber(v.(string))
				if err != nil {
					continue
				}
				globalRank = vFloat
				hasGlobalRank = true
			}
		} else {
			notes = append(notes, line)
		}
	}
	if len(annualValues) > 0 {
		o.Set("annual_values", annualValues)
	}
	if hasGlobalRank {
		o.Set("global_rank", globalRank)
	}
	if len(notes) > 0 {
		note := strings.Join(notes, "; ")
		o.Set("note", note)
	}
	keys := o.Keys()
	if len(keys) == 0 {
		return o, NoValueErr
	}
	return o, nil
}

// subsets are saved in key 'breakdown'
// eg Catholic 25.3% (Roman Catholic 25.1%, other Catholic 0.2%)
// {
//   "name": "Catholic",
//   "percent": 25.3,
//   "breakdown": [
//     {
//       "name": "Roman Catholic",
//       "percent": 25.1
//     },
//     {
//       "name": "other Catholic",
//       "percent": 0.2
//     }
//   ]
// }
//
// ranges use the lower bound
// eg Muslim 99.7% (Sunni 84.7 - 89.7%, Shia 10 - 15%)
// {
//   "name": "Muslim",
//   "percent": 99.7,
//   "breakdown": [
//     {
//       "name": "Sunni",
//       "percent": 84.7
//     },
//     {
//       "name": "Shia",
//       "percent": 10
//     }
//   ]
// }
func stringToPercentageList(s, key string) (*orderedmap.OrderedMap, error) {
	o := orderedmap.New()
	// get date
	s, date, hasDate := stringWithoutDate(s)
	// get list
	firstLine, otherLines := firstLine(s)
	bits := splitIgnoringParenthesis(firstLine, ',')
	list := []*orderedmap.OrderedMap{}
	for _, bit := range bits {
		bitNoPs, ps := removeParenthesis(bit)
		bitNoPs = strings.TrimSpace(bitNoPs)
		bit = strings.Replace(bit, " - ", "-", -1)
		percent := 0.0
		hasPercent := false
		var err error
		name := ""
		parts := strings.Split(bitNoPs, " ")
		for _, part := range parts {
			if startsWithNumber(part) {
				percentStr := strings.Replace(part, "%", "", -1)
				percent, err = stringToNumber(percentStr)
				if err == nil {
					hasPercent = true
				} else {
					name = name + " " + part
				}
			} else {
				name = name + " " + part
			}
		}
		name = strings.TrimSpace(name)
		m := orderedmap.New()
		if len(name) > 0 {
			m.Set("name", name)
		}
		if hasPercent {
			m.Set("percent", percent)
		}
		// handle subsets
		if len(ps) > 0 {
			breakdown := []*orderedmap.OrderedMap{}
			notes := []string{}
			for _, p := range ps {
				if strings.Index(p, "%") > -1 {
					breakdownBits := strings.Split(p, ",")
					for _, breakdownBit := range breakdownBits {
						breakdownBit = strings.TrimSpace(breakdownBit)
						breakdownBit = strings.Replace(breakdownBit, " - ", "-", -1)
						bPercent := 0.0
						bHasPercent := false
						bName := ""
						bParts := strings.Split(breakdownBit, " ")
						for _, bPart := range bParts {
							if startsWithNumber(bPart) {
								bPercent, err = stringToNumber(bPart)
								if err == nil {
									bHasPercent = true
								} else {
									bName = bName + " " + bPart
								}
							} else {
								bName = bName + " " + bPart
							}
						}
						bm := orderedmap.New()
						bName = strings.TrimSpace(bName)
						if len(bName) > 0 {
							bm.Set("name", bName)
						}
						if bHasPercent {
							bm.Set("percent", bPercent)
						}
						breakdown = append(breakdown, bm)
					}
				} else {
					notes = append(notes, p)
				}
			}
			if len(breakdown) > 0 {
				m.Set("breakdown", breakdown)
			}
			if len(notes) > 0 {
				note := strings.Join(notes, "; ")
				m.Set("note", note)
			}
		}
		if len(m.Keys()) > 0 {
			list = append(list, m)
		}
	}
	if len(list) > 0 {
		o.Set(key, list)
	}
	// get note and other extra data
	if len(otherLines) > 0 {
		extraData, err := stringToMap(otherLines)
		if err == nil {
			keys := extraData.Keys()
			for _, k := range keys {
				v, _ := extraData.Get(k)
				o.Set(k, v)
			}
		}
	}
	// set date
	if hasDate {
		o.Set("date", date)
	}
	if len(o.Keys()) == 0 {
		return o, NoValueErr
	}
	return o, nil
}

func stringToAgeStructureMap(s string) (*orderedmap.OrderedMap, error) {
	o := orderedmap.New()
	bits := ageStructureSplitterRe.Split(s, -1)
	if len(bits) == 6 {
		// percent
		percentStr := strings.TrimSpace(bits[1])
		percent, err := strconv.ParseFloat(percentStr, 64)
		if err == nil {
			o.Set("percent", percent)
		}
		// male
		maleStr := strings.Replace(bits[3], "male ", "", -1)
		maleStr = strings.Replace(maleStr, ",", "", -1)
		male, err := strconv.Atoi(maleStr)
		if err == nil {
			o.Set("males", male)
		}
		// female
		femaleStr := strings.Replace(bits[4], "female ", "", -1)
		femaleStr = strings.Replace(femaleStr, ",", "", -1)
		female, err := strconv.Atoi(femaleStr)
		if err == nil {
			o.Set("females", female)
		}
	}
	return o, nil
}

func startsWithNumber(s string) bool {
	s = strings.TrimSpace(s)
	if len(s) > 0 && s[0] == '<' {
		s = s[1:len(s)]
	}
	if len(s) > 0 && s[0] == '>' {
		s = s[1:len(s)]
	}
	if len(s) > 0 && s[0] == '-' {
		s = s[1:len(s)]
	}
	s = strings.TrimSpace(s)
	if len(s) == 0 {
		return false
	}
	return strings.IndexAny(string(s[0]), "0123456789") > -1
}

func startsWith(s, t string) bool {
	// tests if s starts with t
	return len(s) >= len(t) && s[0:len(t)] == t
}

func endsWith(s, t string) bool {
	// tests if s ends with t
	start := len(s) - len(t)
	end := len(s)
	return len(s) >= len(t) && s[start:end] == t
}

func stringToImprovedUnimprovedList(s string) (*orderedmap.OrderedMap, error) {
	o := orderedmap.New()
	improved := orderedmap.New()
	unimproved := orderedmap.New()
	// get date
	s, date, hasDate := stringWithoutDate(s)
	// get lines
	lines := strings.Split(s, "\n")
	topKey := ""
	for _, line := range lines {
		// check for top key
		if startsWith(line, "improved") {
			topKey = "improved"
			continue
		}
		if startsWith(line, "unimproved") {
			topKey = "unimproved"
			continue
		}
		// get second key
		line = strings.Replace(line, "% of population", "", -1)
		bits := strings.Split(line, ": ")
		if len(bits) != 2 {
			continue
		}
		secondKey := bits[0]
		// get percentage
		p, err := stringToNumber(bits[1])
		if err != nil {
			continue
		}
		percentage := orderedmap.New()
		percentage.Set("value", p)
		percentage.Set("units", "percent of population")
		// set values
		if topKey == "improved" {
			improved.Set(secondKey, percentage)
		} else if topKey == "unimproved" {
			unimproved.Set(secondKey, percentage)
		}
	}
	// set improved
	keys := improved.Keys()
	if len(keys) > 0 {
		o.Set("improved", improved)
	}
	// set unimproved
	keys = unimproved.Keys()
	if len(keys) > 0 {
		o.Set("unimproved", unimproved)
	}
	// check for no value
	keys = o.Keys()
	if len(keys) == 0 {
		return o, NoValueErr
	}
	// set date
	if hasDate {
		o.Set("date", date)
	}
	return o, nil
}

// See https://github.com/bep/inflect
// for more comprehensive ruleset
var singulars = map[string]string{
	"administrative areas":                 "administrative area",
	"administrative districts":             "administrative district",
	"areas":                                "area",
	"autonomous districts":                 "autonomous district",
	"autonomous regions":                   "autonomous region",
	"cantons":                              "canton",
	"cities":                               "city",
	"cities with provincial status":        "city with provincial status",
	"communes":                             "commune",
	"counties":                             "country",
	"departments":                          "department",
	"dependencies":                         "dependency",
	"districts":                            "district",
	"divisions":                            "division",
	"emirates":                             "emirate",
	"ethnically based states":              "ethnically based state",
	"first-order administrative divisions": "first-order administrative division",
	"governorates":                         "governate",
	"indigenous territories":               "indigenous territory",
	"islands":                              "island",
	"island councils":                      "island council",
	"island divisions":                     "island division",
	"island groups":                        "island group",
	"localities":                           "locality",
	"municipalities":                       "municipality",
	"oblasts":                              "oblast",
	"parishes":                             "parish",
	"prefectures":                          "prefecture",
	"provinces":                            "province",
	"quarters":                             "quarter",
	"raions":                               "raion",
	"regions":                              "region",
	"regions administrative":               "region administrative",
	"self-governing administrations":       "self-governing administration",
	"states":            "state",
	"territories":       "territory",
	"town councils":     "town council",
	"union territories": "union territory",
	"zones":             "zone",
}

func singularize(s string) string {
	v, ok := singulars[s]
	if !ok {
		return s
	}
	return v
}

func startsWithCapitalLetter(s string) bool {
	if len(s) == 0 {
		return false
	}
	firstChar := string(s[0])
	firstCharLower := strings.ToLower(firstChar)
	return firstCharLower != firstChar
}

func splitByCommaOrSemicolon(s string) []string {
	return commaOrSemicolon.Split(s, -1)
}

func stringToDiplomat(s string) (*orderedmap.OrderedMap, error) {
	// May be opportunity for further parsing...?
	s = strings.Replace(s, "FAX:", "fax:", -1)
	return stringToMap(s)
}

func splitIgnoringParenthesis(s string, sep rune) []string {
	bits := []string{}
	currentBit := ""
	depth := 0
	for _, c := range s {
		if c == '(' {
			depth = depth + 1
			currentBit = currentBit + string(c)
		} else if c == ')' {
			depth = depth - 1
			currentBit = currentBit + string(c)
		} else if c == sep && depth == 0 {
			bits = append(bits, currentBit)
			currentBit = ""
		} else {
			currentBit = currentBit + string(c)
		}
	}
	bits = append(bits, currentBit)
	return bits
}

func stringsInParenthesis(s string) []string {
	bits := []string{}
	currentBit := ""
	depth := 0
	for _, c := range s {
		if c == '(' {
			depth = depth + 1
			if depth == 1 {
				continue
			}
		} else if c == ')' {
			depth = depth - 1
			if depth == 0 {
				bits = append(bits, currentBit)
				currentBit = ""
				continue
			}
		}
		if depth > 0 {
			currentBit = currentBit + string(c)
		}
	}
	return bits
}

func stringToImportExportPartnerList(s string) (*orderedmap.OrderedMap, error) {
	s = strings.Replace(s, "% (", "%, (", -1)
	o := orderedmap.New()
	ps := []*orderedmap.OrderedMap{}
	s, date, hasDate := stringWithoutDate(s)
	partnerStrs := strings.Split(s, ", ")
	for _, partnerStr := range partnerStrs {
		bits := strings.Split(partnerStr, " ")
		if len(bits) < 2 {
			continue
		}
		percent, err := stringToNumber(bits[len(bits)-1])
		if err != nil {
			continue
		}
		name := strings.Join(bits[0:len(bits)-1], " ")
		p := orderedmap.New()
		p.Set("name", name)
		p.Set("percent", percent)
		ps = append(ps, p)
	}
	if len(ps) > 0 {
		o.Set("by_country", ps)
	}
	if hasDate {
		o.Set("date", date)
	}
	keys := o.Keys()
	if len(keys) == 0 {
		return o, NoValueErr
	}
	return o, nil
}

// eg 18,727 km 1.435-m gauge (650 km electrified)
func stringToRailLengthMap(s string) (*orderedmap.OrderedMap, error) {
	o := orderedmap.New()
	s, ps := removeParenthesis(s)
	bits := strings.Split(s, " ")
	if len(bits) > 1 {
		// length
		length, err := stringToNumber(bits[0])
		if err != nil {
			return o, err
		}
		o.Set("length", length)
		// electrified
		for _, p := range ps {
			pBits := strings.Split(p, " ")
			if len(pBits) > 2 && pBits[2] == "electrified" {
				length, err := stringToNumber(pBits[0])
				if err != nil {
					continue
				}
				o.Set("electrified", length)
			}
		}
		// units
		units := bits[1]
		o.Set("units", units)
	}
	if len(bits) > 3 {
		// guage size
		gauge := bits[2]
		o.Set("gauge", gauge)
	}
	return o, nil
}

// eg bulk carrier 8, cargo 7, liquefied gas 4, passenger 6, passenger/cargo 6
func stringToListOfCounts(s, nameKey string) ([]*orderedmap.OrderedMap, error) {
	maps := []*orderedmap.OrderedMap{}
	countStrs := strings.Split(s, ", ")
	for _, countStr := range countStrs {
		countStrBits := strings.Split(countStr, " ")
		if len(countStrBits) > 1 {
			cStr := countStrBits[len(countStrBits)-1]
			count, err := stringToNumber(cStr)
			if err != nil {
				return maps, err
			}
			name := strings.Join(countStrBits[0:len(countStrBits)-1], " ")
			m := orderedmap.New()
			m.Set(nameKey, name)
			m.Set("count", count)
			maps = append(maps, m)
		}
	}
	return maps, nil
}

// eg 17 (Canada 5, Germany 2, Singapore 2, South Africa 1, UK 5, US 2)
func stringToListOfCountsWithTotal(s, nameKey string) (*orderedmap.OrderedMap, error) {
	o := orderedmap.New()
	s, ps := removeParenthesis(s)
	total, err := stringToNumber(s)
	if err != nil {
		return o, err
	}
	// total
	o.Set("total", total)
	// others
	if len(ps) == 1 {
		m, err := stringToListOfCounts(ps[0], nameKey)
		if err != nil {
			return o, err
		}
		o.Set("by_"+nameKey, m)
	}
	return o, nil
}

// eg Dampier (iron ore), Dalrymple Bay (coal), Hay Point (coal)
func stringToListWithItemNotes(s, itemKey, noteKey string) ([]*orderedmap.OrderedMap, error) {
	o := []*orderedmap.OrderedMap{}
	bits := strings.Split(s, ", ")
	for _, bit := range bits {
		bit, ps := removeParenthesis(bit)
		item := orderedmap.New()
		// name
		name := strings.TrimSpace(bit)
		// try converting to number if possible
		nameNum, err := stringToNumber(name)
		if err == nil {
			item.Set(itemKey, nameNum)
		} else {
			item.Set(itemKey, name)
		}
		// note
		note := strings.Join(ps, "; ")
		if len(note) > 0 {
			// try converting to number if possible
			noteNum, err := stringToNumber(note)
			if err == nil {
				item.Set(noteKey, noteNum)
			} else {
				item.Set(noteKey, note)
			}
		}
		o = append(o, item)
	}
	if len(o) == 0 {
		return o, NoValueErr
	}
	return o, nil
}
