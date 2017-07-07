package country

import (
	"orderedmap"
	"testing"
)

type NumberWithUnitsCase struct {
	s                  string
	expectedValue      float64
	expectedUnits      string
	expectedNoteExists bool
	expectedNote       string
	expectedError      error
}

var numberWithUnitsCases = []NumberWithUnitsCase{
	// simplest possible test
	NumberWithUnitsCase{
		s:             "1 m",
		expectedValue: 1,
		expectedUnits: "m",
		expectedError: nil,
	},
	// with decimal place
	NumberWithUnitsCase{
		s:             "1.1 m",
		expectedValue: 1.1,
		expectedUnits: "m",
		expectedError: nil,
	},
	// with thousand
	NumberWithUnitsCase{
		s:             "1.1 thousand m",
		expectedValue: 1100,
		expectedUnits: "m",
		expectedError: nil,
	},
	// with million
	NumberWithUnitsCase{
		s:             "1.1 million m",
		expectedValue: 1100000,
		expectedUnits: "m",
		expectedError: nil,
	},
	// with billion
	NumberWithUnitsCase{
		s:             "1.1 billion m",
		expectedValue: 1100000000,
		expectedUnits: "m",
		expectedError: nil,
	},
	// with trillion
	NumberWithUnitsCase{
		s:             "1.1 trillion m",
		expectedValue: 1100000000000,
		expectedUnits: "m",
		expectedError: nil,
	},
	// with commas
	NumberWithUnitsCase{
		s:             "1,234 m",
		expectedValue: 1234,
		expectedUnits: "m",
		expectedError: nil,
	},
	// with commas and magnitude
	NumberWithUnitsCase{
		s:             "1,234 million m",
		expectedValue: 1234000000,
		expectedUnits: "m",
		expectedError: nil,
	},
	// with units with spaces
	NumberWithUnitsCase{
		s:             "1,234 sq km",
		expectedValue: 1234,
		expectedUnits: "sq km",
		expectedError: nil,
	},
	// with leading, trailing and duplicate space
	NumberWithUnitsCase{
		s:             "\t1,234   sq km\n",
		expectedValue: 1234,
		expectedUnits: "sq km",
		expectedError: nil,
	},
	// with no units
	NumberWithUnitsCase{
		s:             "1",
		expectedValue: 1,
		expectedUnits: "",
		expectedError: nil,
	},
	// with magnitude but no units
	NumberWithUnitsCase{
		s:             "1 million",
		expectedValue: 1000000,
		expectedUnits: "",
		expectedError: nil,
	},
	// with no number
	NumberWithUnitsCase{
		s:             "a",
		expectedValue: 0,
		expectedUnits: "",
		expectedError: StringToNumberWithUnitsErr,
	},
	// with trailing data in parenthesis
	NumberWithUnitsCase{
		s:                  "1,234 km (landlocked)",
		expectedValue:      1234,
		expectedUnits:      "km",
		expectedNoteExists: true,
		expectedNote:       "landlocked",
		expectedError:      nil,
	},
	// with additional trailing data
	NumberWithUnitsCase{
		s:                  "200 nm or to the edge of the continental margin",
		expectedValue:      200,
		expectedUnits:      "nm",
		expectedNoteExists: true,
		expectedNote:       "or to the edge of the continental margin",
		expectedError:      nil,
	},
	// with number as NA
	NumberWithUnitsCase{
		s:                  "NA",
		expectedValue:      0,
		expectedUnits:      "",
		expectedNoteExists: false,
		expectedNote:       "",
		expectedError:      StringIsNaErr,
	},
	// with number in note
	NumberWithUnitsCase{
		s:                  "14 million sq km (280,000 sq km ice-free, 13.72 million sq km ice-covered) (est.)",
		expectedValue:      14000000,
		expectedUnits:      "sq km",
		expectedNoteExists: true,
		expectedNote:       "280,000 sq km ice-free, 13.72 million sq km ice-covered est.",
		expectedError:      nil,
	},
}

func TestStringToNumberWithUnits(t *testing.T) {
	for _, c := range numberWithUnitsCases {
		o, e := stringToNumberWithUnits(c.s)
		// error
		if e != c.expectedError {
			t.Error("Number With Units error", c.s, e)
		}
		// value
		v, _ := o.Get("value")
		if v.(float64) != c.expectedValue {
			t.Error("Number With Units value", c.s, v)
		}
		// units
		u, _ := o.Get("units")
		if c.expectedUnits != "" {
			if u.(string) != c.expectedUnits {
				t.Error("Number With Units units", c.s, u)
			}
		}
		// note
		n, exists := o.Get("note")
		if !c.expectedNoteExists && exists {
			t.Error("Number With Units note exists", c.s, n)
		}
		if c.expectedNoteExists && n != c.expectedNote {
			t.Error("Number With Units note", c.s, n)
		}
	}
}

type StringToMapCase struct {
	s              string
	expectedKeys   []string
	expectedValues []string
	expectedError  error
}

var stringToMapCases = []StringToMapCase{
	// simplest possible test
	StringToMapCase{
		s: "a: 1\nb: 2",
		expectedKeys: []string{
			"a",
			"b",
		},
		expectedValues: []string{
			"1",
			"2",
		},
		expectedError: nil,
	},
	// keys are snake case
	StringToMapCase{
		s: "this is the key: value",
		expectedKeys: []string{
			"this_is_the_key",
		},
		expectedValues: []string{
			"value",
		},
		expectedError: nil,
	},
	// Keys with no value are ignored
	StringToMapCase{
		s: "a:\nb: 2",
		expectedKeys: []string{
			"b",
		},
		expectedValues: []string{
			"2",
		},
		expectedError: nil,
	},
	StringToMapCase{
		s: senegalAirCarriersMapStr,
		expectedKeys: []string{
			"annual_passenger_traffic_on_registered_air_carriers",
			"annual_freight_traffic_on_registered_air_carriers",
		},
		expectedValues: []string{
			"115,355",
			"3,095,523 mt-km",
		},
		expectedError: nil,
	},
}

const senegalAirCarriersMapStr = `
number of registered air carriers: 
inventory of registered aircraft operated by air carriers: 
annual passenger traffic on registered air carriers: 115,355
annual freight traffic on registered air carriers: 3,095,523 mt-km
`

func TestStringToMap(t *testing.T) {
	for testIndex, c := range stringToMapCases {
		o, e := stringToMap(c.s)
		// error
		if e != c.expectedError {
			t.Error("stringToMap error, testIndex: ", testIndex)
		}
		// keys and values
		keys := o.Keys()
		if len(keys) != len(c.expectedKeys) {
			t.Error("stringToMap keys length, testIndex: ", testIndex)
			continue
		}
		for i, key := range keys {
			// key
			if key != c.expectedKeys[i] {
				t.Error("stringToMap key, testIndex: ", testIndex)
			}
			// value
			v, _ := o.Get(key)
			if v.(string) != c.expectedValues[i] {
				t.Error("stringToMap value, testIndex: ", testIndex)
			}
		}
	}
}

func TestStringToGPS(t *testing.T) {
	o, e := stringToGPS("12 34 N, 123 01 E")
	if e != nil {
		t.Error("stringToGPS error")
	}
	keys := o.Keys()
	if len(keys) != 2 {
		t.Error("stringToGPS top level keys")
	}
	latitudeI, _ := o.Get("latitude")
	latitude := latitudeI.(*orderedmap.OrderedMap)
	d, _ := latitude.Get("degrees")
	if d.(int) != 12 {
		t.Error("stringToGPS latitude degrees")
	}
	m, _ := latitude.Get("minutes")
	if m.(int) != 34 {
		t.Error("stringToGPS latitude minutes")
	}
	h, _ := latitude.Get("hemisphere")
	if h.(string) != "N" {
		t.Error("stringToGPS latitude hemisphere")
	}
	longitudeI, _ := o.Get("longitude")
	longitude := longitudeI.(*orderedmap.OrderedMap)
	d, _ = longitude.Get("degrees")
	if d.(int) != 123 {
		t.Error("stringToGPS longitude degrees")
	}
	m, _ = longitude.Get("minutes")
	if m.(int) != 1 {
		t.Error("stringToGPS longitude minutes")
	}
	h, _ = longitude.Get("hemisphere")
	if h.(string) != "E" {
		t.Error("stringToGPS longitude hemisphere")
	}
}

type StringToListCase struct {
	s              string
	lc             listConditions
	expectedValues []string
	expectedError  error
}

var stringToListCases = []StringToListCase{
	// simplest possible test
	StringToListCase{
		s:  "a, b, c",
		lc: listConditions{},
		expectedValues: []string{
			"a",
			"b",
			"c",
		},
		expectedError: nil,
	},
	// blank string
	StringToListCase{
		s:              "",
		lc:             listConditions{},
		expectedValues: []string{},
		expectedError:  nil,
	},
	// separates on ' and '
	StringToListCase{
		s:  "a, b and c",
		lc: listConditions{},
		expectedValues: []string{
			"a",
			"b",
			"c",
		},
		expectedError: nil,
	},
	// separates on ',' without trailing space
	StringToListCase{
		s:  "a,b,c",
		lc: listConditions{},
		expectedValues: []string{
			"a",
			"b",
			"c",
		},
		expectedError: nil,
	},
	// excess whitespace
	StringToListCase{
		s:  " a,   b   and    c",
		lc: listConditions{},
		expectedValues: []string{
			"a",
			"b",
			"c",
		},
		expectedError: nil,
	},
	// multiple commas
	StringToListCase{
		s:  " a,,  ,   b  ,,,,    c",
		lc: listConditions{},
		expectedValues: []string{
			"a",
			"b",
			"c",
		},
		expectedError: nil,
	},
	// oxford comma
	StringToListCase{
		s:  "a, b, and c",
		lc: listConditions{},
		expectedValues: []string{
			"a",
			"b",
			"c",
		},
		expectedError: nil,
	},
	// keeping ands
	StringToListCase{
		s: "this and that and the other",
		lc: listConditions{
			keepAnds: true,
		},
		expectedValues: []string{
			"this and that and the other",
		},
		expectedError: nil,
	},
}

func TestStringToList(t *testing.T) {
	for testIndex, c := range stringToListCases {
		list, err := stringToList(c.s, c.lc)
		if err != c.expectedError {
			t.Error("stringToList error, testIndex: ", testIndex)
			continue
		}
		if len(list) != len(c.expectedValues) {
			t.Error("stringToList length, testIndex: ", testIndex)
			continue
		}
		// values
		for i, value := range list {
			if value != c.expectedValues[i] {
				t.Error("stringToList value, testIndex: ", testIndex)
			}
		}
	}
}

type StringWithoutDateCase struct {
	s               string
	expectedS       string
	expectedDate    string
	expectedHasDate bool
}

var stringWithoutDateCases = []StringWithoutDateCase{
	// simplest possible test
	StringWithoutDateCase{
		s:               "Testing (2010)",
		expectedS:       "Testing",
		expectedDate:    "2010",
		expectedHasDate: true,
	},
	// with est.
	StringWithoutDateCase{
		s:               "Testing (2010 est.)",
		expectedS:       "Testing",
		expectedDate:    "2010",
		expectedHasDate: true,
	},
	// multiline
	StringWithoutDateCase{
		s:               "Test\nTest\nTest (2010 est.)",
		expectedS:       "Test\nTest\nTest",
		expectedDate:    "2010",
		expectedHasDate: true,
	},
}

func TestStringWithoutDate(t *testing.T) {
	for testIndex, c := range stringWithoutDateCases {
		s, d, h := stringWithoutDate(c.s)
		// new s
		if s != c.expectedS {
			t.Error("stringWithoutDate value, testIndex: ", testIndex)
		}
		// year
		if d != c.expectedDate {
			t.Error("stringWithoutDate year, testIndex: ", testIndex)
		}
		// has year
		if h != c.expectedHasDate {
			t.Error("stringWithoutDate hasDate, testIndex: ", testIndex)
		}
	}
}

type StringToNumberWithUnitsAndDateCase struct {
	s             string
	expectedNum   float64
	expectedUnits string
	expectedDate  string
	expectedError error
}

var stringToNumberWithUnitsAndDateCases = []StringToNumberWithUnitsAndDateCase{
	// simplest possible test
	StringToNumberWithUnitsAndDateCase{
		s:             "500.3 sq km (2010)",
		expectedNum:   500.3,
		expectedUnits: "sq km",
		expectedDate:  "2010",
		expectedError: nil,
	},
}

func TestStringToNumberWithUnitsAndDate(t *testing.T) {
	for testIndex, c := range stringToNumberWithUnitsAndDateCases {
		num, err := stringToNumberWithUnitsAndDate(c.s)
		// error
		if err != c.expectedError {
			t.Error("stringToNumberWithUnitsAndDate error, testIndex: ", testIndex)
		}
		// num
		n, _ := num.Get("value")
		if n != c.expectedNum {
			t.Error("stringToNumberWithUnitsAndDate num, testIndex: ", testIndex)
		}
		// units
		u, _ := num.Get("units")
		if u != c.expectedUnits {
			t.Error("stringToNumberWithUnitsAndDate units, testIndex: ", testIndex)
		}
		// date
		d, _ := num.Get("date")
		if d != c.expectedDate {
			t.Error("stringToNumberWithUnitsAndDate date, testIndex: ", testIndex)
		}
	}
}

type StringToPercentageListCase struct {
	s                       string
	expectedKeys            []string
	expectedNames           []string
	expectedValues          []float64
	expectedNotes           []string
	expectedBreakdownNames  []string
	expectedBreakdownValues []float64
	expectedDate            string
	expectedError           error
}

var stringToPercentageListCases = []StringToPercentageListCase{
	// with notes
	StringToPercentageListCase{
		s: "Muslim (official; predominantly Sunni) 99%, other (includes Christian and Jewish) <1% (2012 est.)",
		expectedKeys: []string{
			"religion",
			"date",
		},
		expectedNames: []string{
			"Muslim",
			"other",
		},
		expectedValues: []float64{
			99,
			1,
		},
		expectedNotes: []string{
			"official; predominantly Sunni",
			"includes Christian and Jewish",
		},
		expectedBreakdownNames:  []string{},
		expectedBreakdownValues: []float64{},
		expectedDate:            "2012",
		expectedError:           nil,
	},
	// with breakdown
	StringToPercentageListCase{
		s: "Catholic 25.3% (Roman Catholic 25.1%, other Catholic 0.2%)",
		expectedKeys: []string{
			"religion",
		},
		expectedNames: []string{
			"Catholic",
		},
		expectedValues: []float64{
			25.3,
		},
		expectedNotes: []string{
			"",
		},
		expectedBreakdownNames: []string{
			"Roman Catholic",
			"other Catholic",
		},
		expectedBreakdownValues: []float64{
			25.1,
			0.2,
		},
		expectedDate:  "",
		expectedError: nil,
	},
	// with ranges
	StringToPercentageListCase{
		s: "Muslim 99.7% (Sunni 84.7 - 89.7%, Shia 10 - 15%)",
		expectedKeys: []string{
			"religion",
		},
		expectedNames: []string{
			"Muslim",
		},
		expectedValues: []float64{
			99.7,
		},
		expectedNotes: []string{
			"",
		},
		expectedBreakdownNames: []string{
			"Sunni",
			"Shia",
		},
		expectedBreakdownValues: []float64{
			84.7,
			10,
		},
		expectedDate:  "",
		expectedError: nil,
	},
}

func TestStringToPercentageList(t *testing.T) {
	for testIndex, c := range stringToPercentageListCases {
		o, err := stringToPercentageList(c.s, "religion")
		if err != c.expectedError {
			t.Error("stringToPercentageList len expected error, testIndex: ", testIndex)
		}
		// keys
		keys := o.Keys()
		if len(keys) != len(c.expectedKeys) {
			t.Error("stringToPercentageList len expected keys, testIndex: ", testIndex)
		}
		for i, key := range keys {
			if key != c.expectedKeys[i] {
				t.Error("stringToPercentageList expected key, testIndex: ", testIndex, key)
			}
		}
		// prepare for percent values
		percentsInterface, _ := o.Get("religion")
		percents := percentsInterface.([]*orderedmap.OrderedMap)
		// names
		if len(percents) != len(c.expectedNames) {
			t.Error("stringToPercentageList len expected names, testIndex: ", testIndex)
		}
		for i, percent := range percents {
			name, _ := percent.Get("name")
			if name.(string) != c.expectedNames[i] {
				t.Error("stringToPercentageList expected name, testIndex: ", testIndex, name)
			}
		}
		// values
		if len(percents) != len(c.expectedValues) {
			t.Error("stringToPercentageList len expected names, testIndex: ", testIndex)
		}
		for i, percent := range percents {
			value, _ := percent.Get("percent")
			if value.(float64) != c.expectedValues[i] {
				t.Error("stringToPercentageList expected value, testIndex: ", testIndex, value)
			}
		}
		// notes
		if len(percents) != len(c.expectedNotes) {
			t.Error("stringToPercentageList len expected notes, testIndex: ", testIndex)
		}
		for i, percent := range percents {
			note, _ := percent.Get("note")
			if c.expectedNotes[i] != "" && note.(string) != c.expectedNotes[i] {
				t.Error("stringToPercentageList expected note, testIndex: ", testIndex, note)
			}
		}
		// breakdown names
		if len(c.expectedBreakdownNames) > 0 {
			breakdownInterface, _ := percents[0].Get("breakdown")
			breakdown := breakdownInterface.([]*orderedmap.OrderedMap)
			if len(breakdown) != len(c.expectedBreakdownNames) {
				t.Error("stringToPercentageList len expected breakdown names, testIndex: ", testIndex)
			}
			for i, b := range breakdown {
				name, _ := b.Get("name")
				if name.(string) != c.expectedBreakdownNames[i] {
					t.Error("stringToPercentageList expected breakdown name, testIndex: ", testIndex, name)
				}
			}
		}
		// breakdown values
		if len(c.expectedBreakdownValues) > 0 {
			breakdownInterface, _ := percents[0].Get("breakdown")
			breakdown := breakdownInterface.([]*orderedmap.OrderedMap)
			if len(breakdown) != len(c.expectedBreakdownValues) {
				t.Error("stringToPercentageList len expected breakdown values, testIndex: ", testIndex)
			}
			for i, b := range breakdown {
				value, _ := b.Get("percent")
				if value.(float64) != c.expectedBreakdownValues[i] {
					t.Error("stringToPercentageList expected breakdown value, testIndex: ", testIndex, value)
				}
			}
		}
		// date
		date, _ := o.Get("date")
		if c.expectedDate != "" && date.(string) != c.expectedDate {
			t.Error("stringToPercentageList len expected date, testIndex: ", testIndex, date)
		}
	}
}
