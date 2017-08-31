package country

import (
	"errors"
	"github.com/PuerkitoBio/goquery"
	"net/url"
	"orderedmap"
	"os"
	"path"
	"strconv"
	"strings"
	"time"
)

const VERSION = "0.0.1-beta"

var NoValueErr = errors.New("No value")

type Page struct {
	filelocation string
	dom          *goquery.Document
	ParsedData   *orderedmap.OrderedMap
	NameKey      string
	HasData      bool
}

func NewPage(f string) (Page, error) {
	p := Page{
		filelocation: f,
		ParsedData:   orderedmap.New(),
	}
	// read the dom
	r, err := os.Open(p.filelocation)
	if err != nil {
		return p, err
	}
	p.dom, err = goquery.NewDocumentFromReader(r)
	if err != nil {
		return p, err
	}
	// get values from the dom
	err = p.parse()
	return p, err
}

// Reads the country list from the select element on the page
func (p Page) CountryList() ([]string, error) {
	return countryListFromDom(p.dom)
}

func (p *Page) parse() error {
	// get metadata
	metaData := orderedmap.New()
	date := p.dateStrFromFilename()
	url := p.urlFromFilename()
	yearlySummaryUrl := p.yearlySummaryUrl()
	metaData.Set("date", date)
	metaData.Set("source", url)
	metaData.Set("nearby_dates", yearlySummaryUrl)
	// get the page data
	pageData := orderedmap.New()
	// Set the name
	name, err := p.countryName()
	if err != nil {
		return err
	}
	pageData.Set("name", name)
	p.NameKey = stringToJsonKey(name)
	// Parse each section of the page
	tryAddingData(pageData, "introduction", p.introduction)
	tryAddingData(pageData, "geography", p.geography)
	tryAddingData(pageData, "people", p.people)
	tryAddingData(pageData, "government", p.government)
	tryAddingData(pageData, "economy", p.economy)
	tryAddingData(pageData, "energy", p.energy)
	tryAddingData(pageData, "communications", p.communications)
	tryAddingData(pageData, "transportation", p.transportation)
	tryAddingData(pageData, "military_and_security", p.militaryAndSecurity)
	tryAddingData(pageData, "transnational_issues", p.transnationalIssues)
	d := orderedmap.New()
	d.Set("data", pageData)
	d.Set("metadata", metaData)
	p.ParsedData = d
	p.HasData = len(pageData.Keys()) > 0
	return nil
}

func (p *Page) dateStrFromFilename() string {
	dir, _ := path.Split(p.filelocation)
	dateDir := path.Base(dir)
	return dateDir
}

func (p *Page) urlFromFilename() string {
	_, filename := path.Split(p.filelocation)
	url, err := url.QueryUnescape(filename)
	if err != nil {
		return p.filelocation
	}
	return url
}

func (p *Page) yearlySummaryUrl() string {
	_, filename := path.Split(p.filelocation)
	url, err := url.QueryUnescape(filename)
	if err != nil {
		return p.filelocation
	}
	if len(url) < 42 {
		return p.filelocation
	}
	url = url[0:36] + "000000*" + url[42:len(url)]
	return url
}

func (p *Page) countryName() (string, error) {
	return countryNameFromDom(p.dom)
}

func tryAddingData(d *orderedmap.OrderedMap, key string, valueFn func() (interface{}, error)) {
	value, err := valueFn()
	if err != nil {
		return
	}
	d.Set(key, value)
}

func (p *Page) tryAddingDataForSelector(d *orderedmap.OrderedMap, key, selector string, valueFn func(string) (interface{}, error)) {
	valueStr, err := textForFieldKey(p.dom, selector)
	valueStr = strings.Replace(valueStr, "\t", " ", -1)
	if err != nil {
		return
	}
	value, err := valueFn(valueStr)
	if err != nil {
		return
	}
	d.Set(key, value)
}

func (p *Page) introduction() (interface{}, error) {
	introData := orderedmap.New()
	p.tryAddingDataForSelector(introData, "background", "2028", introBackground)
	p.tryAddingDataForSelector(introData, "preliminary_statement", "2192", preliminaryStatement)
	if len(introData.Keys()) == 0 {
		return introData, NoValueErr
	}
	return introData, nil
}

func (p *Page) geography() (interface{}, error) {
	geoData := orderedmap.New()
	p.tryAddingDataForSelector(geoData, "overview", "2203", geographicOverview)
	p.tryAddingDataForSelector(geoData, "location", "2144", geographyLocation)
	p.tryAddingDataForSelector(geoData, "geographic_coordinates", "2011", geographicCoordinates)
	p.tryAddingDataForSelector(geoData, "map_references", "2145", mapReferences)
	tryAddingData(geoData, "area", p.geographyArea)
	p.tryAddingDataForSelector(geoData, "land_boundaries", "2096", landBoundaries)
	p.tryAddingDataForSelector(geoData, "coastline", "2060", coastline)
	p.tryAddingDataForSelector(geoData, "maritime_claims", "2106", maritimeClaims)
	p.tryAddingDataForSelector(geoData, "climate", "2059", climate)
	p.tryAddingDataForSelector(geoData, "terrain", "2125", terrain)
	p.tryAddingDataForSelector(geoData, "elevation", "2020", elevation)
	p.tryAddingDataForSelector(geoData, "natural_resources", "2111", naturalResources)
	p.tryAddingDataForSelector(geoData, "land_use", "2097", landUse)
	p.tryAddingDataForSelector(geoData, "irrigated_land", "2146", irrigatedLand)
	p.tryAddingDataForSelector(geoData, "total_renewable_water_sources", "2201", totalRenewableWaterSources)
	p.tryAddingDataForSelector(geoData, "freshwater_withdrawal", "2202", freshwaterWithdrawal)
	p.tryAddingDataForSelector(geoData, "population_distribution", "2266", populationDistribution)
	p.tryAddingDataForSelector(geoData, "natural_hazards", "2021", naturalHazards)
	tryAddingData(geoData, "environment", p.environment)
	p.tryAddingDataForSelector(geoData, "note", "2113", geographyNote)
	if len(geoData.Keys()) == 0 {
		return geoData, NoValueErr
	}
	return geoData, nil
}

func (p *Page) people() (interface{}, error) {
	peopleData := orderedmap.New()
	p.tryAddingDataForSelector(peopleData, "population", "2119", population)
	p.tryAddingDataForSelector(peopleData, "nationality", "2110", nationality)
	p.tryAddingDataForSelector(peopleData, "ethnic_groups", "2075", ethnicGroups)
	p.tryAddingDataForSelector(peopleData, "languages", "2098", languages)
	p.tryAddingDataForSelector(peopleData, "religions", "2122", religions)
	p.tryAddingDataForSelector(peopleData, "demographic_profile", "2257", demographicProfile)
	p.tryAddingDataForSelector(peopleData, "age_structure", "2010", ageStructure)
	p.tryAddingDataForSelector(peopleData, "dependency_ratios", "2261", dependencyRatios)
	p.tryAddingDataForSelector(peopleData, "median_age", "2177", medianAge)
	p.tryAddingDataForSelector(peopleData, "population_growth_rate", "2002", populationGrowthRate)
	p.tryAddingDataForSelector(peopleData, "birth_rate", "2054", birthRate)
	p.tryAddingDataForSelector(peopleData, "death_rate", "2066", deathRate)
	p.tryAddingDataForSelector(peopleData, "net_migration_rate", "2112", netMigrationRate)
	p.tryAddingDataForSelector(peopleData, "population_distribution", "2267", populationDistribution)
	p.tryAddingDataForSelector(peopleData, "urbanization", "2212", urbanization)
	p.tryAddingDataForSelector(peopleData, "major_urban_areas", "2219", majorUrbanAreas)
	p.tryAddingDataForSelector(peopleData, "sex_ratio", "2018", sexRatio)
	p.tryAddingDataForSelector(peopleData, "mothers_mean_age_at_first_birth", "2256", mothersMeanAgeAtFirstBirth)
	p.tryAddingDataForSelector(peopleData, "maternal_mortality_rate", "2223", maternalMortalityRate)
	p.tryAddingDataForSelector(peopleData, "infant_mortality_rate", "2091", infantMortalityRate)
	p.tryAddingDataForSelector(peopleData, "life_expectancy_at_birth", "2102", lifeExpectancyAtBirth)
	p.tryAddingDataForSelector(peopleData, "total_fertility_rate", "2127", totalFertilityRate)
	p.tryAddingDataForSelector(peopleData, "contraceptive_prevalence_rate", "2258", contraceptivePrevalenceRate)
	p.tryAddingDataForSelector(peopleData, "health_expenditures", "2225", healthExpenditures)
	p.tryAddingDataForSelector(peopleData, "physicians_density", "2226", physiciansDensity)
	p.tryAddingDataForSelector(peopleData, "hospital_bed_density", "2227", hospitalBedDensity)
	p.tryAddingDataForSelector(peopleData, "drinking_water_source", "2216", drinkingWaterSource)
	p.tryAddingDataForSelector(peopleData, "sanitation_facility_access", "2217", sanitationFacilityAccess)
	tryAddingData(peopleData, "hiv_aids", p.hivAids)
	p.tryAddingDataForSelector(peopleData, "major_infectious_diseases", "2193", majorInfectiousDiseases)
	p.tryAddingDataForSelector(peopleData, "adult_obesity", "2228", obesityAdultPrevalenceRate)
	p.tryAddingDataForSelector(peopleData, "underweight_children", "2224", childrenUnderFiveYearsUnderweight)
	p.tryAddingDataForSelector(peopleData, "education_expenditures", "2206", educationExpenditures)
	p.tryAddingDataForSelector(peopleData, "literacy", "2103", literacy)
	p.tryAddingDataForSelector(peopleData, "school_life_expectancy", "2205", schoolLifeExpectancy)
	p.tryAddingDataForSelector(peopleData, "child_labor", "2255", childLabor)
	p.tryAddingDataForSelector(peopleData, "youth_unemployment", "2229", youthUnemployment)
	p.tryAddingDataForSelector(peopleData, "note", "2022", peopleNote)
	if len(peopleData.Keys()) == 0 {
		return peopleData, NoValueErr
	}
	return peopleData, nil
}

func (p *Page) government() (interface{}, error) {
	governmentData := orderedmap.New()
	p.tryAddingDataForSelector(governmentData, "country_name", "2142", countryName)
	p.tryAddingDataForSelector(governmentData, "union_name", "2189", unionName)
	p.tryAddingDataForSelector(governmentData, "political_structure", "2190", politicalStructure)
	p.tryAddingDataForSelector(governmentData, "government_type", "2128", governmentType)
	p.tryAddingDataForSelector(governmentData, "capital", "2057", capital)
	p.tryAddingDataForSelector(governmentData, "member_states", "2191", memberStates)
	p.tryAddingDataForSelector(governmentData, "administrative_divisions", "2051", administrativeDivisions)
	p.tryAddingDataForSelector(governmentData, "dependent_areas", "2068", dependentAreas)
	p.tryAddingDataForSelector(governmentData, "independence", "2088", independence)
	p.tryAddingDataForSelector(governmentData, "national_holidays", "2109", nationalHoliday)
	p.tryAddingDataForSelector(governmentData, "constitution", "2063", constitution)
	p.tryAddingDataForSelector(governmentData, "legal_system", "2100", legalSystem)
	p.tryAddingDataForSelector(governmentData, "international_law_organization_participation", "2220", internationalLaw)
	p.tryAddingDataForSelector(governmentData, "citizenship", "2263", citizenship)
	p.tryAddingDataForSelector(governmentData, "suffrage", "2123", suffrage)
	p.tryAddingDataForSelector(governmentData, "executive_branch", "2077", executiveBranch)
	p.tryAddingDataForSelector(governmentData, "legislative_branch", "2101", legislativeBranch)
	p.tryAddingDataForSelector(governmentData, "judicial_branch", "2094", judicialBranch)
	p.tryAddingDataForSelector(governmentData, "political_parties_and_leaders", "2118", politicalPartiesAndLeaders)
	p.tryAddingDataForSelector(governmentData, "political_pressure_groups_and_leaders", "2115", politicalPressureGroupsAndLeaders)
	p.tryAddingDataForSelector(governmentData, "international_organization_participation", "2107", internationalOrganizationParticipation)
	tryAddingData(governmentData, "diplomatic_representation", p.diplomaticRepresentation)
	p.tryAddingDataForSelector(governmentData, "flag_description", "2081", flagDescription)
	p.tryAddingDataForSelector(governmentData, "national_symbol", "2230", nationalSymbol)
	tryAddingData(governmentData, "national_anthem", p.nationalAnthem)
	p.tryAddingDataForSelector(governmentData, "note", "2140", governmentNote)
	if len(governmentData.Keys()) == 0 {
		return governmentData, NoValueErr
	}
	return governmentData, nil
}

func (p *Page) economy() (interface{}, error) {
	economyData := orderedmap.New()
	p.tryAddingDataForSelector(economyData, "overview", "2116", economyOverview)
	tryAddingData(economyData, "gdp", p.gdp)
	p.tryAddingDataForSelector(economyData, "gross_national_saving", "2260", grossNationalSaving)
	p.tryAddingDataForSelector(economyData, "agriculture_products", "2052", agricultureProducts)
	p.tryAddingDataForSelector(economyData, "industries", "2090", industries)
	p.tryAddingDataForSelector(economyData, "industrial_production_growth_rate", "2089", industrialProductionGrowthRate)
	tryAddingData(economyData, "labor_force", p.laborForce)
	p.tryAddingDataForSelector(economyData, "unemployment_rate", "2129", unemploymentRate)
	p.tryAddingDataForSelector(economyData, "population_below_poverty_line", "2046", populationBelowPovertyLine)
	p.tryAddingDataForSelector(economyData, "household_income_by_percentage_share", "2047", householdIncomeByPercentageShare)
	p.tryAddingDataForSelector(economyData, "distribution_of_family_income", "2172", distributionOfFamilyIncome)
	p.tryAddingDataForSelector(economyData, "investment_gross_fixed", "2185", investmentGrossFixed)
	p.tryAddingDataForSelector(economyData, "budget", "2056", budget)
	p.tryAddingDataForSelector(economyData, "taxes_and_other_revenues", "2221", taxesAndOtherRevenues)
	p.tryAddingDataForSelector(economyData, "budget_surplus_or_deficit", "2222", budgetSurplusOrDeficit)
	p.tryAddingDataForSelector(economyData, "public_debt", "2186", publicDebt)
	p.tryAddingDataForSelector(economyData, "fiscal_year", "2080", fiscalYear)
	p.tryAddingDataForSelector(economyData, "inflation_rate", "2092", inflationRate)
	p.tryAddingDataForSelector(economyData, "central_bank_discount_rate", "2207", centralBankDiscountRate)
	p.tryAddingDataForSelector(economyData, "commercial_bank_prime_lending_rate", "2208", commercialBankPrimeLendingRate)
	p.tryAddingDataForSelector(economyData, "stock_of_money", "2209", stockOfMoney)
	p.tryAddingDataForSelector(economyData, "stock_of_quasi_money", "2210", stockOfQuasiMoney)
	p.tryAddingDataForSelector(economyData, "stock_of_narrow_money", "2214", stockOfNarrowMoney)
	p.tryAddingDataForSelector(economyData, "stock_of_broad_money", "2215", stockOfBroadMoney)
	p.tryAddingDataForSelector(economyData, "stock_of_domestic_credit", "2211", stockOfDomesticCredit)
	p.tryAddingDataForSelector(economyData, "market_value_of_publicly_traded_shares", "2200", marketValueOfPubliclyTradedShares)
	p.tryAddingDataForSelector(economyData, "current_account_balance", "2187", currentAccountBalance)
	tryAddingData(economyData, "exports", p.exports)
	tryAddingData(economyData, "imports", p.imports)
	p.tryAddingDataForSelector(economyData, "reserves_of_foreign_exchange_and_gold", "2188", reservesOfForeignExchangeAndGold)
	p.tryAddingDataForSelector(economyData, "external_debt", "2079", externalDebt)
	tryAddingData(economyData, "stock_of_direct_foreign_investment", p.stockOfDirectForeignInvestment)
	p.tryAddingDataForSelector(economyData, "exchange_rates", "2076", exchangeRates)
	//p.tryAddingDataForSelector(economyData, "economy_of_the_area_administered_by_turkish_cypriots", "2204", economyOfTurkishCypriots)
	if len(economyData.Keys()) == 0 {
		return economyData, NoValueErr
	}
	return economyData, nil
}

func (p *Page) energy() (interface{}, error) {
	energyData := orderedmap.New()
	tryAddingData(energyData, "electricity", p.electricity)
	tryAddingData(energyData, "crude_oil", p.crudeOil)
	tryAddingData(energyData, "refined_petroleum_products", p.refinedPetroleumProducts)
	tryAddingData(energyData, "natural_gas", p.naturalGas)
	p.tryAddingDataForSelector(energyData, "carbon_dioxide_emissions_from_consumption_of_energy", "2254", carbonDioxideEmissions)
	if len(energyData.Keys()) == 0 {
		return energyData, NoValueErr
	}
	return energyData, nil
}

func (p *Page) communications() (interface{}, error) {
	commsData := orderedmap.New()
	tryAddingData(commsData, "telephones", p.telephones)
	p.tryAddingDataForSelector(commsData, "broadcast_media", "2213", broadcastMedia)
	p.tryAddingDataForSelector(commsData, "radio_broadcast_stations", "2013", radioBroacastStations)
	p.tryAddingDataForSelector(commsData, "television_broadcast_stations", "2015", televisionBroacastStations)
	tryAddingData(commsData, "internet", p.internet)
	p.tryAddingDataForSelector(commsData, "note", "2138", communicationsNote)
	if len(commsData.Keys()) == 0 {
		return commsData, NoValueErr
	}
	return commsData, nil
}

func (p *Page) transportation() (interface{}, error) {
	transportData := orderedmap.New()
	tryAddingData(transportData, "air_transport", p.airTransport)
	p.tryAddingDataForSelector(transportData, "pipelines", "2117", pipelines)
	p.tryAddingDataForSelector(transportData, "railways", "2121", railways)
	p.tryAddingDataForSelector(transportData, "roadways", "2085", roadways)
	p.tryAddingDataForSelector(transportData, "waterways", "2093", waterways)
	p.tryAddingDataForSelector(transportData, "merchant_marine", "2108", merchantMarine)
	p.tryAddingDataForSelector(transportData, "ports_and_terminals", "2120", portsAndTerminals)
	p.tryAddingDataForSelector(transportData, "shipyards_and_ship_building", "2231", shipyardsAndShipBuilding)
	p.tryAddingDataForSelector(transportData, "note", "2008", transportNote)
	if len(transportData.Keys()) == 0 {
		return transportData, NoValueErr
	}
	return transportData, nil
}

func (p *Page) militaryAndSecurity() (interface{}, error) {
	militaryData := orderedmap.New()
	p.tryAddingDataForSelector(militaryData, "expenditures", "2034", militaryExpenditures)
	p.tryAddingDataForSelector(militaryData, "branches", "2055", militaryBranches)
	tryAddingData(militaryData, "manpower", p.militaryManpower)
	p.tryAddingDataForSelector(militaryData, "service_age_and_obligation", "2024", militaryServiceAgeAndObligation)
	p.tryAddingDataForSelector(militaryData, "terrorist_groups", "2265", terroristGroups)
	p.tryAddingDataForSelector(militaryData, "note", "2137", militaryNote)
	if len(militaryData.Keys()) == 0 {
		return militaryData, NoValueErr
	}
	return militaryData, nil
}

func (p *Page) transnationalIssues() (interface{}, error) {
	issuesData := orderedmap.New()
	p.tryAddingDataForSelector(issuesData, "disputes", "2070", disputes)
	p.tryAddingDataForSelector(issuesData, "refugees_and_iternally_displaced_persons", "2194", refugees)
	p.tryAddingDataForSelector(issuesData, "trafficking_in_persons", "2196", traffickingInPersons)
	p.tryAddingDataForSelector(issuesData, "illicit_drugs", "2086", illicitDrugs)
	if len(issuesData.Keys()) == 0 {
		return issuesData, NoValueErr
	}
	return issuesData, nil
}

func introBackground(value string) (interface{}, error) {
	return value, nil
}

func preliminaryStatement(value string) (interface{}, error) {
	return value, nil
}

func geographicOverview(value string) (interface{}, error) {
	return value, nil
}

func geographyLocation(value string) (interface{}, error) {
	// handle edge cases
	// see france
	value = convertToFranceValue(value)
	return value, nil
}

func geographicCoordinates(value string) (interface{}, error) {
	// handle edge cases
	// see france
	value = convertToFranceValue(value)
	return stringToGPS(value)
}

func mapReferences(value string) (interface{}, error) {
	// handle edge cases
	// see france
	value = convertToFranceValue(value)
	return value, nil
}

func (p *Page) geographyArea() (interface{}, error) {
	// Combination of two fields - areas and comparitive
	// areas
	areas, err := textForFieldKey(p.dom, "2147")
	if err != nil {
		return areas, err
	}
	areas = strings.Replace(areas, "water: NEGL", "water: 0 sq km", -1)
	m, err := stringToMapOfNumbersWithUnits(areas)
	if err != nil {
		return areas, err
	}
	// handle edge cases
	// see france
	if strings.Index(areas, "metropolitan France") > -1 {
		m.Delete("note")
		totalMap, ok := m.Get("total")
		if ok {
			totalStr, ok := totalMap.(*orderedmap.OrderedMap).Get("note")
			if ok {
				total, err := stringToNumberWithUnits(totalStr.(string))
				if err == nil {
					total.Delete("note")
					m.Set("total", total)
				}
			}
		}
		landMap, ok := m.Get("land")
		if ok {
			landStr, ok := landMap.(*orderedmap.OrderedMap).Get("note")
			if ok {
				land, err := stringToNumberWithUnits(landStr.(string))
				if err == nil {
					land.Delete("note")
					m.Set("land", land)
				}
			}
		}
		waterMap, ok := m.Get("water")
		if ok {
			waterStr, ok := waterMap.(*orderedmap.OrderedMap).Get("note")
			if ok {
				water, err := stringToNumberWithUnits(waterStr.(string))
				if err == nil {
					water.Delete("note")
					m.Set("water", water)
				}
			}
		}
		m.Set("note", "metropolitan France")
	}
	// comparative
	comparative, err := textForFieldKey(p.dom, "2023")
	if err != nil {
		return comparative, err
	}
	comparative, _ = firstLine(comparative)
	m.Set("comparative", comparative)
	return m, nil
}

func landBoundaries(value string) (interface{}, error) {
	// fix edge cases
	// see france
	if strings.Index(value, "metropolitan France") > -1 {
		value = strings.Replace(value, "metropolitan France - total:", "total:", -1)
		otherCountriesStart := strings.Index(value, "French Guiana - total")
		if otherCountriesStart > -1 {
			value = value[0:otherCountriesStart]
		}
	}
	// might be a map
	// or might be just a number
	// so try map conversion first
	// if it fails, treat it as a number
	boundaryMap, err := stringToMap(value)
	if err == nil {
		// format values more cleanly
		keys := boundaryMap.Keys()
		keysToRemove := []string{}
		for _, key := range keys {
			if key == "total" {
				boundaryTotalStr, _ := boundaryMap.Get(key)
				boundaryTotalNum, err := stringToNumberWithUnits(boundaryTotalStr.(string))
				if err != nil {
					return boundaryMap, err
				}
				boundaryMap.Set(key, boundaryTotalNum)
			} else if len(key) > 16 && key[0:17] == "border_countries_" {
				borderCountriesStr, _ := boundaryMap.Get(key)
				keysToRemove = append(keysToRemove, key)
				borderCountries, err := borderCountriesStringToSlice(borderCountriesStr.(string))
				if err != nil {
					return boundaryMap, err
				}
				boundaryMap.Set("border_countries", borderCountries)
			} else if len(key) > 16 && key[0:17] == "regional_borders_" {
				regionalBordersStr, _ := boundaryMap.Get(key)
				keysToRemove = append(keysToRemove, key)
				regionalBorders, err := borderCountriesStringToSlice(regionalBordersStr.(string))
				if err != nil {
					return boundaryMap, err
				}
				boundaryMap.Set("regional_borders", regionalBorders)
			}
		}
		for _, key := range keysToRemove {
			boundaryMap.Delete(key)
		}
	} else {
		boundaryTotal, err := stringToNumberWithUnits(value)
		if err != nil {
			return boundaryTotal, err
		}
		boundaryMap = orderedmap.New()
		boundaryMap.Set("total", boundaryTotal)
	}
	return boundaryMap, nil
}

func coastline(value string) (interface{}, error) {
	// fix edge cases
	// see Saint Helena, Ascension and Tristan Da Cunha
	if strings.Index(value, "Saint Helena:") > -1 {
		return value, NoValueErr
	}
	// see france
	value = convertToFranceValue(value)
	first, others := firstLine(value)
	o, err := stringToNumberWithUnits(first)
	if err != nil {
		return o, err
	}
	m, err := stringToMap(others)
	if err == nil {
		keys := m.Keys()
		for _, key := range keys {
			v, _ := m.Get(key)
			o.Set(key, v)
		}
	}
	return o, nil
}

func maritimeClaims(value string) (interface{}, error) {
	// handle unusual cases
	// see the world
	if strings.Index(value, "most countries make the following claims") > -1 {
		return value, nil
	}
	value = strings.Replace(value, "0-m ", "0 m ", -1)
	return stringToMapOfNumbersWithUnits(value)
}

func climate(value string) (interface{}, error) {
	// fix edge cases
	// see france
	value = convertToFranceValue(value)
	return value, nil
}

func terrain(value string) (interface{}, error) {
	// fix edge cases
	// see france
	value = convertToFranceValue(value)
	return value, nil
}

func elevation(value string) (interface{}, error) {
	// fix nested keys
	value = strings.Replace(value, "elevation extremes: ", "", -1)
	// convert to map.
	// can't use
	// stringToMapOfNumbersWithUnits
	// since some values are
	// stringToPlaceAndNumberWithUnits
	elevationMap, err := stringToMap(value)
	if err != nil {
		return elevationMap, err
	}
	keys := elevationMap.Keys()
	keysToDelete := []string{}
	for _, key := range keys {
		// mean elevation
		value, _ := elevationMap.Get(key)
		if key == "mean_elevation" || key == "mean_depth" {
			meanElNum, err := stringToNumberWithUnits(value.(string))
			if err == nil {
				elevationMap.Set(key, meanElNum)
			} else if err == StringIsNaErr {
				keysToDelete = append(keysToDelete, key)
			}
		} else if key == "highest_point" || key == "lowest_point" {
			// lowest and highest point
			v, err := stringToPlaceAndNumberWithUnits(value.(string), "name", "elevation")
			if err != nil {
				continue
			}
			elevationMap.Set(key, v)
		} else if key == "top_ten_highest_mountains_(measured_from_sea_level)" || key == "highest_point_on_each_continent" || key == "lowest_point_on_each_continent" {
			keysToDelete = append(keysToDelete, key)
		}
	}
	// delete invalid values
	for _, key := range keysToDelete {
		elevationMap.Delete(key)
	}
	return elevationMap, err
}

func naturalResources(value string) (interface{}, error) {
	o := orderedmap.New()
	// fix edge cases
	// see france
	value = convertToFranceValue(value)
	// list is only first line
	first, others := firstLine(value)
	lc := listConditions{
		splitChars: ",()",
	}
	// fix some inconsistencies
	// see the world
	if strings.Index(first, "pose serious long-term problems") > -1 {
		return "", NoValueErr
	}
	// see antarctica
	first = strings.Replace(first, "none presently exploited;", "none presently exploited,", -1)
	first = strings.Replace(first, "have been taken by commercial fisheries", "", -1)
	// see antigua and barbuda
	first = strings.Replace(first, "NEGL;", "NEGL,", -1)
	// see United States Pacific Island Wildlife Refuges
	first = strings.Replace(first, "terrestrial and aquatic wildlife", "terrestrial wildlife, aquatic wildlife", -1)
	// see china
	first = strings.Replace(first, "(world's largest)", "", -1)
	// see georgia
	first = strings.Replace(first, "minor coal and oil deposits;", "minor coal deposits, minor oil deposits,", -1)
	first = strings.Replace(first, "coastal climate and soils allow for important tea and citrus growth", "tea, citrus", -1)
	// see greenland
	first = strings.Replace(first, "possible oil and gas", "possible oil, possible gas", -1)
	// see guam
	first = strings.Replace(first, "aquatic wildlife (supporting tourism)", "aquatic wildlife supporting tourism", -1)
	first = strings.Replace(first, "fishing (largely undeveloped)", "fishing; largely undeveloped", -1)
	// see india
	first = strings.Replace(first, "(fourth-largest reserves in the world)", "", -1)
	// see kiribati
	first = strings.Replace(first, "(production discontinued in 1979)", "; production discontinued in 1979", -1)
	// see kyrgyzstan
	first = strings.Replace(first, "hydropower; gold", "hydropower, gold", -1)
	first = strings.Replace(first, "rare earth metals;", "rare earth metals,", -1)
	first = strings.Replace(first, "locally exploitable coal, oil, and natural gas;", "coal, oil, natural gas,", -1)
	first = strings.Replace(first, "other deposits of nepheline", "nepheline", -1)
	// see luxembourg
	first = strings.Replace(first, " (no longer exploited)", "; no longer exploited", -1)
	// see malawi
	first = strings.Replace(first, "unexploited deposits of uranium, coal, and bauxite", "uranium; unexploited, coal; unexploited, bauxite; unexploited", -1)
	// see mali
	if others == "note: bauxite, iron ore, manganese, tin, and copper deposits are known but not exploited" {
		first = first + ", bauxite; unexploited, iron ore; unexploited, manganese; unexploited, tin; unexploited, copper deposits; unexploited"
		others = ""
	}
	// see palau
	first = strings.Replace(first, " (especially gold)", ", gold", -1)
	// see pitcairn islands
	first = strings.Replace(first, "miro trees (used for handicrafts)", "miro trees; used for handicrafts", -1)
	// see puerto rico
	first = strings.Replace(first, "some copper and nickel;", "copper, nickel,", -1)
	first = strings.Replace(first, "potential for onshore and offshore oil", "onshore oil; unexploited, offshore oil; unexploited", -1)
	// see russia
	first = strings.Replace(first, "wide natural resource base including major deposits of ", "", -1)
	// see saint barthelemy
	first = strings.Replace(first, "few natural resources; ", "few natural resources, ", -1)
	// see slovakia
	first = strings.Replace(first, "manganese ore; salt; arable land", "manganese ore, salt, arable land", -1)
	// see somalia
	first = strings.Replace(first, "largely unexploited reserves of iron ore", "iron ore; largely unexploited", -1)
	// see southern ocean
	first = strings.Replace(first, "probable large oil and gas fields on the continental margin;", "oil, gas,", -1)
	first = strings.Replace(first, "fresh water as icebergs;", "fresh water as icebergs,", -1)
	first = strings.Replace(first, "squid, whales, and seals - none exploited;", "squid; unexploited, whales; unexploited, seals; unexploited,", -1)
	// see sudan
	first = strings.Replace(first, "petroleum; small reserves of iron ore,", "petroleum, iron ore,", -1)
	first = strings.Replace(first, "gold; hydropower", "gold, hydropower", -1)
	// see suriname
	first = strings.Replace(first, "small amounts of ", "", -1)
	// see swaziland
	first = strings.Replace(first, "small gold and diamond deposits", "gold, diamond", -1)
	// see taiwan
	first = strings.Replace(first, "small deposits of ", "", -1)
	// see yemen
	first = strings.Replace(first, "copper; fertile soil in west", "copper, fertile soil in west", -1)
	list, err := stringToList(first, lc)
	if err != nil {
		return o, err
	}
	if len(list) > 0 {
		o.Set("resources", list)
	}
	// other values, such as note
	others = strings.TrimSpace(others)
	if len(others) > 0 {
		m, err := stringToMap(others)
		if err != nil {
			return o, err
		}
		keys := m.Keys()
		for _, k := range keys {
			v, _ := m.Get(k)
			o.Set(k, v)
		}
	}
	if len(o.Keys()) == 0 {
		return o, NoValueErr
	}
	return o, nil
}

func landUse(value string) (interface{}, error) {
	// sanitize the string
	value = strings.Replace(value, "; ", "\n", -1)
	value = strings.Replace(value, "agricultural land:", "agricultural land total: ", -1)
	value = strings.Replace(value, "arable land ", "agricultural land arable land: ", -1)
	value = strings.Replace(value, "permanent crops ", "agricultural land permanent crops: ", -1)
	value = strings.Replace(value, "permanent pasture ", "agricultural land permanent pasture: ", -1)
	return stringToPercentageMap(value, "by_sector")
}

func irrigatedLand(value string) (interface{}, error) {
	// fix edge cases
	// see france
	value = convertToFranceValue(value)
	return stringToNumberWithUnitsAndDate(value)
}

func totalRenewableWaterSources(value string) (interface{}, error) {
	return stringToNumberWithUnitsAndDate(value)
}

func freshwaterWithdrawal(value string) (interface{}, error) {
	return stringToMap(value)
}

func populationDistribution(value string) (interface{}, error) {
	return value, nil
}

func naturalHazards(value string) (interface{}, error) {
	// fix edge cases
	// see france
	value = convertToFranceValue(value)
	// first line is a list
	firstLine, otherLines := firstLine(value)
	lc := listConditions{
		keepAnds:   true,
		splitChars: ";",
	}
	list, err := stringToList(firstLine, lc)
	if err != nil {
		return list, err
	}
	// include type with each hazard
	hazards := []*orderedmap.OrderedMap{}
	for _, hazard := range list {
		h := orderedmap.New()
		h.Set("description", hazard)
		h.Set("type", "hazard")
		hazards = append(hazards, h)
	}
	// other hazards are listed as a map, eg volcanism
	hazardMap, _ := stringToMap(otherLines)
	// combine map values into the list
	keys := hazardMap.Keys()
	for _, key := range keys {
		value, _ := hazardMap.Get(key)
		h := orderedmap.New()
		h.Set("description", value.(string))
		h.Set("type", key)
		hazards = append(hazards, h)
	}
	return hazards, nil
}

func (p *Page) environment() (interface{}, error) {
	// combination of two field keys - current issues and international agreements
	o := orderedmap.New()
	// current issues
	currentIssues, err := textForFieldKey(p.dom, "2032")
	if err == nil {
		// see The World
		a := "pollution (air, water, acid rain, toxic substances)"
		b := "air pollution, water pollution, acid rain, toxic substances"
		currentIssues = strings.Replace(currentIssues, a, b, -1)
		a = "vegetation (overgrazing, deforestation, desertification)"
		b = "vegetation, overgrazing, deforestation, desertification"
		currentIssues = strings.Replace(currentIssues, a, b, -1)
		lc := listConditions{
			keepAnds:   true,
			splitChars: ",;",
		}
		issuesList, err := stringToList(currentIssues, lc)
		if err == nil {
			o.Set("current_issues", issuesList)
		}
	}
	// international agreements
	internationalAgreements, err := textForFieldKey(p.dom, "2033")
	if err == nil {
		agreementsMap, err := stringToMap(internationalAgreements)
		if err == nil {
			keys := agreementsMap.Keys()
			lc := listConditions{
				keepAnds: true,
			}
			for _, key := range keys {
				value, _ := agreementsMap.Get(key)
				if key == "signed_but_not_ratified" && value == "none of the selected agreements" {
					agreementsMap.Set(key, []string{})
					continue
				}
				list, err := stringToList(value.(string), lc)
				if err != nil {
					continue
				}
				agreementsMap.Set(key, list)
			}
			o.Set("international_agreements", agreementsMap)
		}
	}
	keys := o.Keys()
	if len(keys) == 0 {
		return o, NoValueErr
	}
	return o, nil
}

func geographyNote(value string) (interface{}, error) {
	return value, nil
}

func population(value string) (interface{}, error) {
	return stringToNumberWithGlobalRankAndDate(value, "total")
}

func nationality(value string) (interface{}, error) {
	return stringToMap(value)
}

func ethnicGroups(value string) (interface{}, error) {
	return stringToPercentageList(value, "ethnicity")
}

func languages(value string) (interface{}, error) {
	return stringToPercentageList(value, "language")
}

func religions(value string) (interface{}, error) {
	value = strings.Replace(value, "%, 2", "% (2", -1) // date
	// fix typos
	// see australia
	value = strings.Replace(value, "Baptist, ", "Baptist ", -1)
	o, err := stringToPercentageList(value, "religion")
	if err != nil {
		return o, err
	}
	// remove religious affiliation
	o.Delete("religious_affiliation")
	return o, nil
}

func demographicProfile(value string) (interface{}, error) {
	return value, nil
}

func ageStructure(value string) (interface{}, error) {
	value, date, hasDate := stringWithoutDate(value)
	lines := strings.Split(value, "\n")
	o := orderedmap.New()
	for _, line := range lines {
		key := ""
		if strings.Index(line, "0-14 years:") == 0 {
			key = "0_to_14"
		} else if strings.Index(line, "15-24 years:") == 0 {
			key = "15_to_24"
		} else if strings.Index(line, "25-54 years:") == 0 {
			key = "25_to_54"
		} else if strings.Index(line, "55-64 years:") == 0 {
			key = "55_to_64"
		} else if strings.Index(line, "65 years and over:") == 0 {
			key = "65_and_over"
		}
		if key == "" {
			continue
		}
		ages, err := stringToAgeStructureMap(line)
		if err != nil {
			continue
		}
		if len(ages.Keys()) > 0 {
			o.Set(key, ages)
		}
	}
	if hasDate {
		o.Set("date", date)
	}
	return o, nil
}

func dependencyRatios(value string) (interface{}, error) {
	value, date, hasDate := stringWithoutDate(value)
	o, err := stringToPercentageMap(value, "ratios")
	if err != nil {
		return o, err
	}
	if hasDate {
		o.Set("date", date)
	}
	return o, nil
}

func medianAge(value string) (interface{}, error) {
	value, date, hasDate := stringWithoutDate(value)
	o, err := stringToMapOfNumbersWithUnits(value)
	if err != nil {
		return o, err
	}
	if hasDate {
		o.Set("date", date)
	}
	return o, nil
}

func populationGrowthRate(value string) (interface{}, error) {
	return stringToNumberWithGlobalRankAndDate(value, "growth_rate")
}

func birthRate(value string) (interface{}, error) {
	return stringToNumberWithGlobalRankAndDate(value, "births_per_1000_population")
}

func deathRate(value string) (interface{}, error) {
	return stringToNumberWithGlobalRankAndDate(value, "deaths_per_1000_population")
}

func netMigrationRate(value string) (interface{}, error) {
	return stringToNumberWithGlobalRankAndDate(value, "migrants_per_1000_population")
}

func urbanization(value string) (interface{}, error) {
	o, err := stringToMap(value)
	if err != nil {
		return o, err
	}
	for _, k := range o.Keys() {
		vInterface, _ := o.Get(k)
		v := vInterface.(string)
		v, date, hasDate := stringWithoutDate(v)
		vBits := strings.Split(v, "%")
		if len(vBits) == 2 {
			percent, err := strconv.ParseFloat(vBits[0], 64)
			if err != nil {
				continue
			}
			percentMap := orderedmap.New()
			percentMap.Set("value", percent)
			percentMap.Set("units", "%")
			if hasDate {
				percentMap.Set("date", date)
			}
			o.Set(k, percentMap)
		} else if k == "ten_largest_urban_agglomerations" {
			citiesMap := orderedmap.New()
			cities := []*orderedmap.OrderedMap{}
			bits := strings.Split(v, "; ")
			for _, bit := range bits {
				bit, ps := removeParenthesis(bit)
				if len(ps) != 1 {
					continue
				}
				country := ps[0]
				parts := strings.Split(bit, " - ")
				if len(parts) != 2 {
					continue
				}
				city := parts[0]
				population, err := stringToNumber(parts[1])
				if err != nil {
					continue
				}
				c := orderedmap.New()
				c.Set("city", city)
				c.Set("country", country)
				c.Set("population", population)
				cities = append(cities, c)
			}
			if len(cities) > 0 {
				citiesMap.Set("by_population", cities)
				if hasDate {
					citiesMap.Set("date", date)
				}
				o.Set(k, citiesMap)
			}
		}
	}
	if len(o.Keys()) == 0 {
		return o, NoValueErr
	}
	return o, nil
}

func majorUrbanAreas(value string) (interface{}, error) {
	// parse date
	s, date, hasDate := stringWithoutDate(value)
	// parse list of places and populations
	o := orderedmap.New()
	places := []*orderedmap.OrderedMap{}
	bits := strings.Split(s, ";")
	for _, bit := range bits {
		placeMap := orderedmap.New()
		s, ps := removeParenthesis(bit)
		// get place name
		placeStr := ""
		placeStrEndIndex := 0
		placeBits := strings.Split(s, " ")
		for i, placeBit := range placeBits {
			if !startsWithNumber(placeBit) {
				placeStr = placeStr + " " + placeBit
				placeStrEndIndex = i
			} else {
				break
			}
		}
		placeStr = strings.TrimSpace(placeStr)
		placeStr = strings.ToLower(placeStr)
		placeStr = strings.Title(placeStr)
		if len(placeStr) == 0 {
			continue
		}
		// get place population
		hasPopulation := false
		populationStr := strings.Join(placeBits[placeStrEndIndex:len(placeBits)], " ")
		population, err := stringToNumber(populationStr)
		if err == nil {
			hasPopulation = true
		}
		// get capital / note
		notes := []string{}
		isCapital := false
		for _, p := range ps {
			if p == "capital" {
				isCapital = true
			} else {
				notes = append(notes, p)
			}
		}
		note := strings.Join(notes, "; ")
		// store the parsed data
		placeMap.Set("place", placeStr)
		if hasPopulation {
			placeMap.Set("population", population)
		}
		if isCapital {
			placeMap.Set("is_capital", isCapital)
		}
		if len(note) > 0 {
			placeMap.Set("note", note)
		}
		places = append(places, placeMap)
	}
	if len(places) == 0 {
		return o, NoValueErr
	}
	o.Set("places", places)
	// set date
	if hasDate {
		o.Set("date", date)
	}
	return o, nil
}

func sexRatio(value string) (interface{}, error) {
	// parse date
	s, date, hasDate := stringWithoutDate(value)
	// parse ratios
	ratios := orderedmap.New()
	m, err := stringToMap(s)
	if err != nil {
		return m, err
	}
	var total float64
	hasTotal := false
	keys := m.Keys()
	units := "males/female"
	for _, k := range keys {
		v, _ := m.Get(k)
		vStr := strings.TrimSpace(v.(string))
		bits := strings.Split(vStr, " ")
		if len(bits) == 0 {
			continue
		}
		value, err := stringToNumber(bits[0])
		if err != nil {
			continue
		}
		if k == "total_population" {
			total = value
			hasTotal = true
		} else {
			k = strings.Replace(k, "0_14", "0_to_14", -1)
			k = strings.Replace(k, "15_24", "15_to_24", -1)
			k = strings.Replace(k, "25_54", "25_to_54", -1)
			k = strings.Replace(k, "55_64", "55_to_64", -1)
			valueWithUnits := orderedmap.New()
			valueWithUnits.Set("value", value)
			valueWithUnits.Set("units", units)
			ratios.Set(k, valueWithUnits)
		}
	}
	// set ratios
	if len(ratios.Keys()) == 0 {
		return ratios, NoValueErr
	}
	o := orderedmap.New()
	o.Set("by_age", ratios)
	if hasTotal {
		totalWithUnits := orderedmap.New()
		totalWithUnits.Set("value", total)
		totalWithUnits.Set("units", units)
		o.Set("total_population", totalWithUnits)
	}
	// set date
	if hasDate {
		o.Set("date", date)
	}
	return o, nil
}

func mothersMeanAgeAtFirstBirth(value string) (interface{}, error) {
	// parse date
	s, date, hasDate := stringWithoutDate(value)
	// parse age
	s = strings.TrimSpace(s)
	age, err := stringToNumber(s)
	if err != nil {
		return age, err
	}
	o := orderedmap.New()
	o.Set("age", age)
	if hasDate {
		o.Set("date", date)
	}
	return o, nil
}

func maternalMortalityRate(value string) (interface{}, error) {
	return stringToNumberWithGlobalRankAndDate(value, "deaths_per_100k_live_births")
}

func infantMortalityRate(value string) (interface{}, error) {
	value = strings.Replace(value, "deaths/1,000 live births", "deaths_per_1000_live_births", -1)
	return stringToMapOfNumbersWithUnits(value)
}

func lifeExpectancyAtBirth(value string) (interface{}, error) {
	return stringToMapOfNumbersWithUnits(value)
}

func totalFertilityRate(value string) (interface{}, error) {
	return stringToNumberWithGlobalRankAndDate(value, "children_born_per_woman")
}

func contraceptivePrevalenceRate(value string) (interface{}, error) {
	return stringToPercentage(value)
}

func healthExpenditures(value string) (interface{}, error) {
	value = strings.Replace(value, "% of GDP", " percent_of_gdp", -1)
	return stringToNumberWithGlobalRankAndDate(value, "percent_of_gdp")
}

func physiciansDensity(value string) (interface{}, error) {
	value = strings.Replace(value, "physicians/1,000 population", " physicians_per_1000_population", -1)
	return stringToNumberWithGlobalRankAndDate(value, "physicians_per_1000_population")
}

func hospitalBedDensity(value string) (interface{}, error) {
	value = strings.Replace(value, "beds/1,000 population", " beds_per_1000_population", -1)
	return stringToNumberWithGlobalRankAndDate(value, "beds_per_1000_population")
}

func drinkingWaterSource(value string) (interface{}, error) {
	return stringToImprovedUnimprovedList(value)
}

func sanitationFacilityAccess(value string) (interface{}, error) {
	return stringToImprovedUnimprovedList(value)
}

func (p *Page) hivAids() (interface{}, error) {
	// contains three field keys
	o := orderedmap.New()
	// adult prevalence rate
	rateStr, err := textForFieldKey(p.dom, "2155")
	if err == nil {
		rate, err := stringToNumberWithGlobalRankAndDate(rateStr, "percent_of_adults")
		if err == nil {
			o.Set("adult_prevalence_rate", rate)
		}
	}
	// people living with hiv aids
	livingStr, err := textForFieldKey(p.dom, "2156")
	if err == nil {
		living, err := stringToNumberWithGlobalRankAndDate(livingStr, "total")
		if err == nil {
			o.Set("people_living_with_hiv_aids", living)
		}
	}
	// deaths
	deathsStr, err := textForFieldKey(p.dom, "2157")
	if err == nil {
		deaths, err := stringToNumberWithGlobalRankAndDate(deathsStr, "total")
		if err == nil {
			o.Set("deaths", deaths)
		}
	}
	// check for no values
	keys := o.Keys()
	if len(keys) == 0 {
		return o, NoValueErr
	}
	return o, nil
}

func majorInfectiousDiseases(value string) (interface{}, error) {
	value = strings.Replace(value, "bacterial and protozoal diarrhea", "bacterial diarrhea, protozoal diarrhea", -1)
	value = strings.Replace(value, " disease:", " diseases:", -1)
	value = strings.Replace(value, "hepatitis A and E", "hepatitis A, hepatitis E", -1)
	value, date, hasDate := stringWithoutDate(value)
	o, err := stringToMap(value)
	if err != nil {
		return o, err
	}
	keys := o.Keys()
	for _, key := range keys {
		vInterface, _ := o.Get(key)
		v := vInterface.(string)
		if key == "food_or_waterborne_diseases" ||
			key == "vectorborne_diseases" ||
			key == "soil_contact_diseases" ||
			key == "respiratory_diseases" ||
			key == "water_contact_diseases" ||
			key == "animal_contact_diseases" {
			l, err := stringToList(v, listConditions{})
			if err == nil {
				o.Set(key, l)
			}
		}
	}
	if hasDate {
		o.Set("date", date)
	}
	return o, nil
}

func obesityAdultPrevalenceRate(value string) (interface{}, error) {
	return stringToNumberWithGlobalRankAndDate(value, "percent_of_adults")
}

func childrenUnderFiveYearsUnderweight(value string) (interface{}, error) {
	return stringToNumberWithGlobalRankAndDate(value, "percent_of_children_under_the_age_of_five")
}

func educationExpenditures(value string) (interface{}, error) {
	value = strings.Replace(value, "% of GDP", " percent_of_gdp", -1)
	return stringToNumberWithGlobalRankAndDate(value, "percent_of_gdp")
}

func literacy(value string) (interface{}, error) {
	value, date, hasDate := stringWithoutDate(value)
	o, err := stringToMap(value)
	if err != nil {
		return o, err
	}
	keys := o.Keys()
	for _, key := range keys {
		vInterface, _ := o.Get(key)
		v := vInterface.(string)
		if key == "total_population" ||
			key == "male" ||
			key == "female" {
			num, err := stringToNumberWithUnits(v)
			if err != nil {
				continue
			}
			o.Set(key, num)
		}
	}
	if hasDate {
		o.Set("date", date)
	}
	return o, nil
}

func schoolLifeExpectancy(value string) (interface{}, error) {
	return stringToMapOfNumbersWithUnits(value)
}

func childLabor(value string) (interface{}, error) {
	value, date, hasDate := stringWithoutDate(value)
	o, err := stringToMap(value)
	if err != nil {
		return o, err
	}
	keys := o.Keys()
	for _, key := range keys {
		vInterface, _ := o.Get(key)
		v := vInterface.(string)
		if key == "total_number" {
			num, err := stringToNumber(v)
			if err != nil {
				continue
			}
			o.Set(key, num)
		}
		if key == "percentage" {
			num, err := stringToNumberWithUnits(v)
			if err != nil {
				continue
			}
			o.Set(key, num)
		}
	}
	if hasDate {
		o.Set("date", date)
	}
	return o, nil
}

func youthUnemployment(value string) (interface{}, error) {
	return stringToMapOfNumbersWithUnits(value)
}

func peopleNote(value string) (interface{}, error) {
	return value, nil
}

func countryName(value string) (interface{}, error) {
	return stringToMap(value)
}

func unionName(value string) (interface{}, error) {
	return stringToMap(value)
}

func politicalStructure(value string) (interface{}, error) {
	return value, nil
}

func governmentType(value string) (interface{}, error) {
	return value, nil
}

func capital(value string) (interface{}, error) {
	// see Samoa
	value = strings.Replace(value, "during Standard Time)\n+1hr", "during Standard Time)\ndaylight saving time: +1hr", -1)
	m, err := stringToMap(value)
	if err != nil {
		return m, err
	}
	// geographic coordinates
	gpsStr, ok := m.Get("geographic_coordinates")
	if ok {
		gps, err := stringToGPS(gpsStr.(string))
		if err == nil {
			m.Set("geographic_coordinates", gps)
		}
	}
	// time difference
	timeDiffInterface, ok := m.Get("time_difference")
	if ok {
		timeDiffStr := timeDiffInterface.(string)
		timeDiffStr, ps := removeParenthesis(timeDiffStr)
		note := strings.Join(ps, "; ")
		timeDiffStr = strings.Replace(timeDiffStr, "UTC", "", -1)
		timeDiffStr = strings.TrimSpace(timeDiffStr)
		timeDiff, err := strconv.ParseFloat(timeDiffStr, 64)
		if err == nil {
			o := orderedmap.New()
			o.Set("timezone", timeDiff)
			o.Set("note", note)
			m.Set("time_difference", o)
		}
	}
	return m, nil
}

func memberStates(value string) (interface{}, error) {
	return value, nil
}

func administrativeDivisions(value string) (interface{}, error) {
	// Moldova and Taiwan
	isKeyBased := strings.Index(value, "counties:") > -1 ||
		strings.Index(value, "raions:") > -1 ||
		strings.Index(value, "divisions:") > -1
	if isKeyBased {
		o, err := stringToMap(value)
		if err != nil {
			return o, err
		}
		keys := o.Keys()
		names := []*orderedmap.OrderedMap{}
		for _, key := range keys {
			if key == "note" {
				continue
			}
			tidyType := strings.Replace(key, "_", " ", -1)
			singularType := singularize(tidyType)
			namesStr, _ := o.Get(key)
			namesBits := strings.Split(namesStr.(string), ", ")
			for _, nameBit := range namesBits {
				nameBit, _ := removeParenthesis(nameBit)
				n := orderedmap.New()
				n.Set("name", nameBit)
				n.Set("type", singularType)
				names = append(names, n)
			}
		}
		if len(names) == 0 {
			return names, NoValueErr
		}
		return names, nil
	}
	// Falkland Islands
	if strings.Index(value, "none") == 0 {
		return value, NoValueErr
	}
	// si.html
	value = strings.Replace(value, ") Ajdovscina", "); Ajdovscina", -1)
	// mv.html
	value = strings.Replace(value, "the capital city", "1 capital city", -1)
	// mu.html
	value = strings.Replace(value, ") Ad Dakhiliyah", "); Ad Dakhiliyah", -1)
	// ma.html Madagascar 2014-02-09 19:14:19
	value = strings.Replace(value, "Hi Trent,", "", -1)
	firstline, _ := firstLine(value)
	// Bosnia and Herzegovina
	firstline = strings.Replace(firstline, " - ", "; ", 1)
	// Turkmenistan
	firstline = strings.Replace(firstline, "*: ", "*; ", 1)
	// All other countries
	bits := strings.Split(firstline, ";")
	if len(bits) != 2 {
		return bits, NoValueErr
	}
	// get types
	typesStr, _ := removeParenthesis(bits[0])
	typesStr = strings.Replace(typesStr, ", ", " ", -1) // see Gambia, The
	types := strings.Split(typesStr, " ")
	typeIndex := -1
	typeNames := []string{}
	// build primary and secondary and tertiary type
	for _, typeStr := range types {
		if startsWithNumber(typeStr) {
			typeIndex = typeIndex + 1
		} else if len(typeStr) > 0 {
			if typeStr == "and" {
				continue
			}
			if len(typeNames) == typeIndex {
				typeNames = append(typeNames, "")
			}
			typeNames[typeIndex] = typeNames[typeIndex] + " " + typeStr
		}
	}
	// tidy up type names
	for i, typeName := range typeNames {
		typeName = strings.Replace(typeName, "*", "", -1)
		typeName = strings.TrimSpace(typeName)
		typeName = singularize(typeName)
		typeNames[i] = typeName
	}
	// get names
	names := strings.Split(strings.TrimSpace(bits[1]), ", ")
	// add type to names
	namesWithTypes := []*orderedmap.OrderedMap{}
	for _, name := range names {
		n := orderedmap.New()
		// get type name
		typeNameIndex := strings.Count(name, "*")
		typeName := ""
		typeName = typeNames[typeNameIndex]
		// set name
		name = strings.Replace(name, "*", "", -1)
		n.Set("name", name)
		// set type
		n.Set("type", typeName)
		namesWithTypes = append(namesWithTypes, n)
	}
	if len(namesWithTypes) == 0 {
		return namesWithTypes, NoValueErr
	}
	return namesWithTypes, nil
}

func dependentAreas(value string) (interface{}, error) {
	firstline, otherlines := firstLine(value)
	o := orderedmap.New()
	// areas
	areas := strings.Split(firstline, ",")
	for i, area := range areas {
		areas[i] = strings.TrimSpace(area)
	}
	if len(areas) == 0 {
		return areas, NoValueErr
	}
	o.Set("areas", areas)
	// note
	n, _ := stringToMap(otherlines)
	keys := n.Keys()
	for _, key := range keys {
		v, _ := n.Get(key)
		o.Set(key, v)
	}
	return o, nil
}

func independence(value string) (interface{}, error) {
	bits := strings.Split(value, "; ")
	if len(bits) < 1 {
		return bits, NoValueErr
	}
	s := bits[0]
	s, ps := removeParenthesis(s)
	o := orderedmap.New()
	// date
	sBits := strings.Split(s, " ")
	if len(sBits) > 2 {
		possibleDate := strings.Join(sBits[0:3], " ")
		t, err := time.Parse("2 January 2006", possibleDate)
		if err != nil {
			return t, NoValueErr
		}
		dateStr := t.Format("2006-01-02")
		o.Set("date", dateStr)
	}
	// note
	if len(ps) > 0 {
		if len(bits) > 1 {
			ps = append(ps, bits[1:len(bits)]...)
		}
		note := strings.Join(ps, "; ")
		o.Set("note", note)
	}
	return o, nil
}

func nationalHoliday(value string) (interface{}, error) {
	// see Bahrain
	value = strings.Replace(value, "; note - ", " ", -1)
	// see Boznia and Hergezovia
	if strings.Index(value, "Bosnia and Herzegovina") > -1 {
		value = strings.Replace(value, "\n", "; ", -1)
	}
	// see China
	if strings.Index(value, "China") > -1 {
		before := ", the anniversary of the founding of the People's Republic of China"
		after := " (the anniversary of the founding of the People's Republic of China)"
		value = strings.Replace(value, before, after, -1)
	}
	// see Croatia
	if strings.Index(value, "Croatian") > -1 {
		value = strings.Replace(value, ") and", "); ", -1)
		value = strings.Replace(value, "independence; following", "independence. Following", -1)
	}
	// See Curacao
	value = strings.Replace(value, "27 April 1967", "27 April (1967)", -1)
	// See Iraq
	value = strings.Replace(value, "July 14", "14 July", -1)
	// See Liechtenstein
	if strings.Index(value, "Assumption Day") > -1 {
		value = strings.Replace(value, ", 15 August, and", " and", -1)
	}
	// See Luxembourg
	if strings.Index(value, "Grand Duke Henri") > -1 {
		value = strings.Replace(value, ") ", "), ", -1)
	}
	// See Paraguay
	value = strings.Replace(value, " 1811 ", " (1811) ", -1)
	// See European Union
	if strings.Index(value, "Europe Day") > -1 {
		value = strings.Replace(value, "Day) ", "Day), ", -1)
	}
	// get holidays
	bits := strings.Split(value, "; ")
	holidays := []*orderedmap.OrderedMap{}
	for _, bit := range bits {
		h := orderedmap.New()
		s, ps := removeParenthesis(bit)
		sBits := strings.Split(s, ", ")
		if len(sBits) < 2 {
			continue
		}
		notes := []string{}
		name := strings.TrimSpace(sBits[0])
		dayStr := strings.TrimSpace(strings.Join(sBits[1:len(sBits)], ", "))
		dayBits := strings.Split(dayStr, " ")
		day := dayStr
		if startsWithNumber(dayStr) {
			if len(dayBits) > 1 {
				day = strings.Join(dayBits[0:2], " ")
			}
			if len(dayBits) > 2 {
				n := strings.Join(dayBits[2:len(dayBits)], " ")
				notes = append(notes, strings.TrimSpace(n))
			}
		}
		originalYear := ""
		for _, p := range ps {
			_, err := time.Parse("2006", p)
			if err != nil {
				notes = append(notes, p)
			} else {
				originalYear = p
			}
		}
		h.Set("name", name)
		h.Set("day", day)
		if len(originalYear) > 0 {
			h.Set("original_year", originalYear)
		}
		if len(notes) > 0 {
			h.Set("note", strings.Join(notes, "; "))
		}
		holidays = append(holidays, h)
	}
	if len(holidays) == 0 {
		return holidays, NoValueErr
	}
	return holidays, nil
}

func constitution(value string) (interface{}, error) {
	// see Armenia
	value = strings.Replace(value, "; note - ", "\nnote: ", -1)
	firstline, otherlines := firstLine(value)
	// first line
	o, err := stringToMap(firstline)
	if err != nil {
		o = orderedmap.New()
		o.Set("history", firstline)
	}
	// other lines
	b, _ := stringToMap(otherlines)
	keys := b.Keys()
	for _, key := range keys {
		v, _ := b.Get(key)
		o.Set(key, v)
	}
	return o, nil
}

func legalSystem(value string) (interface{}, error) {
	return value, nil
}

func internationalLaw(value string) (interface{}, error) {
	return strings.Split(value, "; "), nil
}

func citizenship(value string) (interface{}, error) {
	return stringToMap(value)
}

func suffrage(value string) (interface{}, error) {
	o := orderedmap.New()
	bits := strings.Split(value, " ")
	if len(bits) > 0 {
		age, err := stringToNumber(bits[0])
		if err == nil {
			o.Set("age", age)
		}
	}
	if strings.Index(value, "universal and compulsory") > -1 {
		o.Set("universal", true)
		o.Set("compulsory", true)
	} else if strings.Index(value, "universal") > -1 {
		o.Set("universal", true)
		o.Set("compulsory", false)
	}
	// note
	// ie anything not of the common form "X years of age; universal [and compulsory]"
	withoutAge := strings.Join(bits[1:len(bits)], " ")
	withoutAge = strings.Replace(withoutAge, "universal and compulsory", "", -1)
	withoutAge = strings.Replace(withoutAge, "universal", "", -1)
	withoutAge = strings.Replace(withoutAge, ";", "", -1)
	withoutAge = strings.Replace(withoutAge, ",", "", -1)
	withoutAge = strings.Replace(withoutAge, "years of age", "", -1)
	withoutAge = strings.TrimSpace(withoutAge)
	if len(withoutAge) > 0 {
		o.Set("note", value)
	}
	keys := o.Keys()
	if len(keys) == 0 {
		return o, NoValueErr
	}
	return o, nil
}

func executiveBranch(value string) (interface{}, error) {
	return stringToMap(value)
}

func legislativeBranch(value string) (interface{}, error) {
	return stringToMap(value)
}

func judicialBranch(value string) (interface{}, error) {
	return stringToMap(value)
}

func politicalPartiesAndLeaders(value string) (interface{}, error) {
	value = strings.Replace(value, `“`, `"`, -1)
	value = strings.Replace(value, `”`, `"`, -1)
	// See Albania
	value = strings.Replace(value, "other parties:", "", -1)
	value = strings.Replace(value, "]:", "]", -1)
	// See Argentina
	value = strings.Replace(value, "numerous provincial parties", "Numerous provincial parties", -1)
	// See Austria
	value = strings.Replace(value, `"Team Stronach" [Frank STRONACH]`, "Team Stronach [Frank STRONACH]", -1)
	// See Belarus
	value = strings.Replace(value, "pro-government parties:", "", -1)
	value = strings.Replace(value, "opposition parties:", "", -1)
	// See Belgium
	value = strings.Replace(value, "Flemish parties:", "", -1)
	value = strings.Replace(value, "Francophone parties:", "", -1)
	value = strings.Replace(value, "other minor parties", "Other minor parties", -1)
	// See Burma
	value = strings.Replace(value, "numerous smaller parties", "Numerous smaller parties", -1)
	// See Cote D'Ivoire
	value = strings.Replace(value, "more than 144 smaller registered parties", "More than 144 smaller registered parties", -1)
	// See Cyprus
	value = strings.Replace(value, "area under government control:", "", -1)
	value = strings.Replace(value, "area administered by Turkish Cypriots:", "", -1)
	// See Czechia
	value = strings.Replace(value, "parties in parliament: ", "", -1)
	value = strings.Replace(value, "parties outside parliament:", "", -1)
	// See Egypt
	value = strings.Replace(value, "officially recognized: ", "", -1)
	// See Equatorial Guinea
	value = strings.Replace(value, "not officially registered parties: ", "", -1)
	// See Hong Kong
	value = strings.Replace(value, "parties:", "", -1)
	value = strings.Replace(value, "others:", "", -1)
	// See Italy
	value = strings.Replace(value, "Ruling left-center-right coalition:", "", -1)
	value = strings.Replace(value, "Center-right opposition:", "", -1)
	value = strings.Replace(value, "Other parties and parliamentary groups:", "", -1)
	value = strings.Replace(value, ", and ", ", ", -1)
	// See Jersey
	value = strings.Replace(value, "one registered party:", "", -1)
	// See North Korea
	value = strings.Replace(value, "major party:", "", -1)
	value = strings.Replace(value, "minor parties:", "", -1)
	// See Lebanon
	value = strings.Replace(value, "14 March Coalition:", "", -1)
	value = strings.Replace(value, "Hizballah-led bloc (formerly 8 March Coalition):", "", -1)
	value = strings.Replace(value, "Independent:", "", -1)
	// See Moldova
	value = strings.Replace(value, "represented in Parliament:", "", -1)
	value = strings.Replace(value, "not represented in Parliament, participated in recent elections (2014-2016):", "", -1)
	value = strings.Replace(value, `"Motherland" Party`, "Motherland Party", -1)
	value = strings.Replace(value, `"Right" Party`, "Right Party", -1)
	// See Sao Tome And Principe
	value = strings.Replace(value, "other small parties", "Other small parties", -1)
	// See Sierra Leone
	value = strings.Replace(value, "numerous other parties", "Numerous other parties", -1)
	// See Slovakia
	value = strings.Replace(value, "parties in the Parliament:", "", -1)
	value = strings.Replace(value, "selected parties outside the Parliament:", "", -1)
	// See Syria
	value = strings.Replace(value, "legal parties/alliances: ", "", -1)
	value = strings.Replace(value, "Kurdish parties (considered illegal): ", "", -1)
	value = strings.Replace(value, "other: ", "", -1)
	lines := strings.Split(value, "\n")
	parties := []*orderedmap.OrderedMap{}
	notes := []string{}
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if len(line) == 0 {
			continue
		}
		// get whole line note
		isParty := startsWithCapitalLetter(line) || startsWithNumber(line)
		isNote := len(line) > 6 && line[0:4] == "note"
		isNote = isNote || !isParty
		if isNote {
			note := line
			if len(line) > 6 && line[0:4] == "note" {
				note = strings.TrimSpace(line[6:len(line)])
			}
			notes = append(notes, note)
			continue
		}
		// prepare name and party notes
		name, ps := removeParenthesis(line)
		// get party notes
		hasNote := false
		partyNotes := []string{}
		if len(ps) > 0 {
			hasNote = true
			partyNotes = append(partyNotes, ps...)
		}
		notePrefixes := []string{
			"; note -",
			"; note:",
			"note:",
			"note -",
		}
		for _, notePrefix := range notePrefixes {
			noteIndex := strings.Index(name, notePrefix)
			if noteIndex > -1 {
				noteBits := strings.Split(name, notePrefix)
				if len(noteBits) == 2 {
					name = strings.TrimSpace(noteBits[0])
					partyNote := strings.TrimSpace(noteBits[1])
					partyNotes = append(partyNotes, partyNote)
					hasNote = true
				}
			}
		}
		// check if party has leaders
		hasLeaders := true
		startName := strings.Index(name, "[")
		if startName == -1 {
			hasLeaders = false
		}
		endName := strings.Index(name, "]")
		if endName == -1 {
			hasLeaders = false
		}
		// get leaders
		leaders := []string{}
		if hasLeaders {
			// convert leaders into slice
			leadersStr := name[startName+1 : endName]
			leaders = strings.Split(leadersStr, ", ")
			// remove leaders from name
			name = name[0:startName-1] + name[endName+1:len(name)]
		}
		// get alternative name
		hasAlternativeName := strings.Index(name, " or ") > -1
		alternativeName := ""
		if hasAlternativeName {
			nameBits := strings.Split(name, " or ")
			if len(nameBits) == 2 {
				name = strings.TrimSpace(nameBits[0])
				alternativeName = strings.TrimSpace(nameBits[1])
			}
		}
		// set values
		p := orderedmap.New()
		p.Set("name", name)
		if hasAlternativeName {
			p.Set("name_alternative", alternativeName)
		}
		if hasLeaders {
			p.Set("leaders", leaders)
		}
		if hasNote {
			partyNote := strings.Join(partyNotes, "; ")
			p.Set("note", partyNote)
		}
		parties = append(parties, p)
	}
	o := orderedmap.New()
	if len(parties) > 0 {
		o.Set("parties", parties)
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

func politicalPressureGroupsAndLeaders(value string) (interface{}, error) {
	// See Austria
	value = strings.Replace(value, ", and ", ", ", -1)
	lines := strings.Split(value, "\n")
	groups := []*orderedmap.OrderedMap{}
	notes := []string{}
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if len(line) == 0 {
			continue
		}
		// handle simple list of parties
		if len(line) > 8 && line[0:6] == "other:" {
			line = strings.Replace(line, "other:", "", -1)
			lineBits := splitByCommaOrSemicolon(line)
			for _, bit := range lineBits {
				name, notes := removeParenthesis(bit)
				if len(name) > 0 {
					p := orderedmap.New()
					p.Set("name", strings.TrimSpace(name))
					if len(notes) > 0 {
						p.Set("note", strings.Join(notes, "; "))
					}
					groups = append(groups, p)
				}
			}
			continue
		}
		// get whole line note
		isParty := startsWithCapitalLetter(line) || startsWithNumber(line)
		isNote := len(line) > 6 && line[0:4] == "note"
		isNote = isNote || !isParty
		if isNote {
			note := line
			if len(line) > 6 && line[0:4] == "note" {
				note = strings.TrimSpace(line[6:len(line)])
			}
			notes = append(notes, note)
			continue
		}
		// prepare name and party notes
		name, ps := removeParenthesis(line)
		// get party notes
		hasNote := false
		partyNotes := []string{}
		if len(ps) > 0 {
			hasNote = true
			partyNotes = append(partyNotes, ps...)
		}
		notePrefixes := []string{
			"; note -",
			"; note:",
			"note:",
			"note -",
		}
		for _, notePrefix := range notePrefixes {
			noteIndex := strings.Index(name, notePrefix)
			if noteIndex > -1 {
				noteBits := strings.Split(name, notePrefix)
				if len(noteBits) == 2 {
					name = strings.TrimSpace(noteBits[0])
					partyNote := strings.TrimSpace(noteBits[1])
					partyNotes = append(partyNotes, partyNote)
					hasNote = true
				}
			}
		}
		// check if party has leaders
		hasLeaders := true
		startName := strings.Index(name, "[")
		if startName == -1 {
			hasLeaders = false
		}
		endName := strings.Index(name, "]")
		if endName == -1 {
			hasLeaders = false
		}
		// get leaders
		leaders := []string{}
		if hasLeaders {
			// convert leaders into slice
			leadersStr := name[startName+1 : endName]
			leaders = strings.Split(leadersStr, ", ")
			// remove leaders from name
			name = name[0:startName-1] + name[endName+1:len(name)]
		}
		// get alternative name
		hasAlternativeName := strings.Index(name, " or ") > -1
		alternativeName := ""
		if hasAlternativeName {
			nameBits := strings.Split(name, " or ")
			if len(nameBits) == 2 {
				name = strings.TrimSpace(nameBits[0])
				alternativeName = strings.TrimSpace(nameBits[1])
			}
		}
		// set values
		p := orderedmap.New()
		p.Set("name", name)
		if hasAlternativeName {
			p.Set("name_alternative", alternativeName)
		}
		if hasLeaders {
			p.Set("leaders", leaders)
		}
		if hasNote {
			partyNote := strings.Join(partyNotes, "; ")
			p.Set("note", partyNote)
		}
		groups = append(groups, p)
	}
	o := orderedmap.New()
	if len(groups) > 0 {
		o.Set("pressure_groups", groups)
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

func internationalOrganizationParticipation(value string) (interface{}, error) {
	bits := strings.Split(value, ", ")
	groups := []*orderedmap.OrderedMap{}
	for _, bit := range bits {
		group := orderedmap.New()
		name, notes := removeParenthesis(bit)
		group.Set("organization", name)
		if len(notes) > 0 {
			group.Set("note", strings.Join(notes, "; "))
		}
		groups = append(groups, group)
	}
	if len(groups) == 0 {
		return groups, NoValueErr
	}
	return groups, nil
}

func (p *Page) diplomaticRepresentation() (interface{}, error) {
	// contains two field keys
	o := orderedmap.New()
	// in the US
	diplomatInUs, err := textForFieldKey(p.dom, "2149")
	if err == nil {
		person, err := stringToDiplomat(diplomatInUs)
		if err == nil {
			o.Set("in_united_states", person)
		}
	}
	// from the US
	diplomatFromUs, err := textForFieldKey(p.dom, "2007")
	if err == nil {
		person, err := stringToDiplomat(diplomatFromUs)
		if err == nil {
			o.Set("from_united_states", person)
		}
	}
	// check for no values
	keys := o.Keys()
	if len(keys) == 0 {
		return o, NoValueErr
	}
	return o, nil
}

func flagDescription(value string) (interface{}, error) {
	description, notes := firstLine(value)
	if len(description) > 4 && description[0:4] == "note" {
		notes = description
		description = ""
	}
	o := orderedmap.New()
	if len(description) > 0 {
		o.Set("description", description)
	}
	m, _ := stringToMap(notes)
	keys := m.Keys()
	for _, k := range keys {
		v, _ := m.Get(k)
		o.Set(k, v)
	}
	return o, nil
}

func nationalSymbol(value string) (interface{}, error) {
	categories := orderedmap.New()
	bits := splitIgnoringParenthesis(value, ';')
	rootKey := "symbols"
	childKey := "symbol"
	items := []*orderedmap.OrderedMap{}
	for _, bit := range bits {
		bit = strings.TrimSpace(bit)
		// change to national colors if required
		isColors := startsWith(bit, "national colors:")
		isColors = isColors || startsWith(bit, "union colors:")
		if isColors {
			if len(items) > 0 {
				categories.Set(rootKey, items)
				items = []*orderedmap.OrderedMap{}
			}
			rootKey = "colors"
			childKey = "color"
			bit = strings.Replace(bit, "national colors:", "", -1)
			bit = strings.Replace(bit, "union colors:", "", -1)
		}
		ps := stringsInParenthesis(bit)
		region := ""
		hasRegion := false
		if len(ps) == 1 {
			possibleRegionName := strings.Replace(ps[0], "in general", "", -1)
			if possibleRegionName == strings.Title(possibleRegionName) {
				region = ps[0]
				hasRegion = true
			}
		}
		bitsByComma := splitIgnoringParenthesis(bit, ',')
		for _, bitByComma := range bitsByComma {
			bitByComma = strings.TrimSpace(bitByComma)
			isForPrevNote := startsWith(bitByComma, "which ")                  // see India
			isForPrevNote = isForPrevNote || startsWith(bitByComma, "is the ") // see India
			if isForPrevNote {
				prevNote, exists := items[len(items)-1].Get("note")
				if exists {
					prevNote = prevNote.(string) + "; " + bitByComma
				} else {
					prevNote = bitByComma
				}
				items[len(items)-1].Set("note", prevNote)
				continue
			}
			isForPrevName := bitByComma == "five-pointed"                                        // see European Union
			isForPrevName = isForPrevName || bitByComma == "golden yellow stars on a blue field" // see European Union
			isForPrevName = isForPrevName || bitByComma == "yellow stars"                        // see Cabo Verde
			if isForPrevName {
				itemValue, _ := items[len(items)-1].Get(childKey)
				itemValue = itemValue.(string) + ", " + bitByComma
				items[len(items)-1].Set(childKey, itemValue)
				continue
			}
			item := orderedmap.New()
			bitByComma, ps = removeParenthesis(bitByComma)
			note := ""
			hasNote := false
			if len(ps) > 0 {
				note = strings.Join(ps, "; ")
				if note != region {
					hasNote = true
				}
			}
			item.Set(childKey, bitByComma)
			if hasNote {
				item.Set("note", note)
			}
			if hasRegion {
				item.Set("region", region)
			}
			items = append(items, item)
		}
	}
	if len(items) > 0 {
		categories.Set(rootKey, items)
	}
	keys := categories.Keys()
	if len(keys) == 0 {
		return categories, NoValueErr
	}
	return categories, nil
}

func (p *Page) nationalAnthem() (interface{}, error) {
	anthemStr, err := textForFieldKey(p.dom, "2218")
	if err != nil {
		return anthemStr, err
	}
	m, err := stringToMap(anthemStr)
	if err != nil {
		return m, err
	}
	url, err := nationalAnthemMp3FromDom(p.dom)
	if err == nil {
		m.Set("audio_url", url)
	}
	return m, nil
}

func governmentNote(value string) (interface{}, error) {
	return value, nil
}

func economyOverview(value string) (interface{}, error) {
	return value, nil
}

func (p *Page) gdp() (interface{}, error) {
	// includes six fields
	gdp := orderedmap.New()
	p.tryAddingDataForSelector(gdp, "purchasing_power_parity", "2001", gdpPpp)
	p.tryAddingDataForSelector(gdp, "official_exchange_rate", "2195", gdpOfficialExchangeRate)
	p.tryAddingDataForSelector(gdp, "real_growth_rate", "2003", gdpRealGrowthRate)
	p.tryAddingDataForSelector(gdp, "per_capita_purchasing_power_parity", "2004", gdpPerCapitaPpp)
	tryAddingData(gdp, "composition", p.gdpComposition)
	keys := gdp.Keys()
	if len(keys) == 0 {
		return gdp, NoValueErr
	}
	return gdp, nil
}

func gdpPpp(value string) (interface{}, error) {
	return stringToListOfAnnualValues(value, "USD")
}

func gdpOfficialExchangeRate(value string) (interface{}, error) {
	return stringToNumberWithGlobalRankAndDate(value, "USD")
}

func gdpRealGrowthRate(value string) (interface{}, error) {
	return stringToListOfAnnualValues(value, "%")
}

func gdpPerCapitaPpp(value string) (interface{}, error) {
	return stringToListOfAnnualValues(value, "USD")
}

func grossNationalSaving(value string) (interface{}, error) {
	return stringToListOfAnnualValues(value, "percent_of_gdp")
}

func (p *Page) gdpComposition() (interface{}, error) {
	composition := orderedmap.New()
	p.tryAddingDataForSelector(composition, "by_end_use", "2259", gdpCompositionByEndUse)
	p.tryAddingDataForSelector(composition, "by_sector_of_origin", "2012", gdpCompositionBySector)
	keys := composition.Keys()
	if len(keys) == 0 {
		return composition, NoValueErr
	}
	return composition, nil
}

func gdpCompositionByEndUse(value string) (interface{}, error) {
	// See American Samoa
	value = strings.Replace(value, "investment if", "investment in", -1)
	// See San Marino
	value = strings.Replace(value, "investments in", "investment in", -1)
	return stringToPercentageMap(value, "end_uses")
}

func gdpCompositionBySector(value string) (interface{}, error) {
	return stringToPercentageMap(value, "sectors")
}

func agricultureProducts(value string) (interface{}, error) {
	// see Japan
	value = strings.Replace(value, "/", ", ", -1)
	// see Cyprus
	value = strings.Replace(value, "\n", ", ", -1)
	// see Timor-Leste
	value = strings.Replace(value, " (", ", ", -1)
	value = strings.Replace(value, ")", "", -1)
	// see New Caledonia
	value = strings.Replace(value, "other ", "", -1)
	// see Armenia
	value = strings.Replace(value, "especially", "", -1)
	// see Pitcairn Islands
	value = strings.Replace(value, "wide variety of", "", -1)
	// see Anguilla
	value = strings.Replace(value, "small quantities of", "", -1)
	// see Saint Vincent
	value = strings.Replace(value, "small numbers of", "", -1)
	// see Western Sahara
	value = strings.Replace(value, "grown in the few oases", "", -1)
	value = strings.Replace(value, "kept by nomads", "", -1)
	// see Gabon
	value = strings.Replace(value, "a tropical softwood", "", -1)
	// see Anguilla
	value = strings.Replace(value, "raising", "", -1)
	// see Croatia
	value = strings.Replace(value, "for wine", "", -1)
	// see Indonesia
	value = strings.Replace(value, "and similar products", "", -1)
	value = strings.Replace(value, "its similar products", "", -1)
	// see Cyprus
	value = strings.Replace(value, "Agriculture - products:", "", -1)
	lc := listConditions{
		splitChars: ",;",
	}
	untidyProducts, err := stringToList(value, lc)
	if err != nil {
		return untidyProducts, err
	}
	products := []string{}
	notes := []string{}
	for _, untidyProduct := range untidyProducts {
		untidyProduct = strings.TrimSpace(untidyProduct)
		// find notes
		// See China
		isNote := startsWith(untidyProduct, "world leader")
		// See Macau
		isNote = isNote || startsWith(untidyProduct, "only ")
		isNote = isNote || startsWith(untidyProduct, "mainly ")
		if isNote {
			notes = append(notes, untidyProduct)
			continue
		}
		// find bits that can be ignored
		// see Macau
		canIgnore := untidyProduct == "mostly for crustaceans"
		canIgnore = canIgnore || untidyProduct == "is important"
		canIgnore = canIgnore || startsWith(untidyProduct, "some of the catch")
		if canIgnore {
			continue
		}
		bits := strings.Split(untidyProduct, " and ")
		for _, bit := range bits {
			// see Malaysia
			hyphenRemoved := strings.Split(bit, " - ")
			bit = hyphenRemoved[len(hyphenRemoved)-1]
			bit = strings.TrimSpace(bit)
			// see Monaco
			if len(bit) > 0 && bit != "none" {
				products = append(products, bit)
			}
		}
	}
	o := orderedmap.New()
	if len(products) > 0 {
		o.Set("products", products)
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

func industries(value string) (interface{}, error) {
	// see Spain
	value = strings.Replace(value, "(including", "including", -1)
	// see Hungary
	value = strings.Replace(value, "(especially", "especially", -1)
	// see Christmas Island
	value = strings.Replace(value, "near depletion", "", -1)
	// see Iran
	value = strings.Replace(value, "(particularly", "particularly", -1)
	// see Antigua And Barbuda
	value = strings.Replace(value, " (", ", ", -1)
	value = strings.Replace(value, ")", "", -1)
	// see Cyprus
	value = strings.Replace(value, "Industries:", "", -1)
	value = strings.Replace(value, "\n", ", ", -1)
	// see Thailand
	value = strings.Replace(value, "world's second-largest tungsten producer and third-largest tin producer", "tungsten, tin", -1)
	// see South Africa
	value = strings.Replace(value, "world's largest producer of", "", -1)
	// see West Bank
	value = strings.Replace(value, ", and ", ", ", -1)
	lc := listConditions{
		splitChars: ",;:",
	}
	industriesList, err := stringToList(value, lc)
	if err != nil {
		return industriesList, err
	}
	industries := []string{}
	notes := []string{}
	// see The World
	if startsWith(value, "dominated by the onrush") {
		notes = append(notes, value)
		industriesList = []string{}
	}
	for _, industry := range industriesList {
		industry = strings.TrimSpace(industry)
		// see Malaysia
		hyphenRemoved := strings.Split(industry, " - ")
		industry = hyphenRemoved[len(hyphenRemoved)-1]
		// see Mauritania
		if industry == "note" {
			continue
		}
		// find notes
		// see Mauritania
		isNote := startsWith(industry, "gypsum deposits have")
		// see China
		isNote = isNote || startsWith(industry, "world leader")
		// see European Union
		isNote = isNote || startsWith(industry, "among the world's largest")
		isNote = isNote || startsWith(industry, "the EU industrial")
		// see Japan
		isNote = isNote || startsWith(industry, "among world's largest")
		// see United States
		isNote = isNote || startsWith(industry, "highly diversified")
		isNote = isNote || startsWith(industry, "world leading")
		isNote = isNote || startsWith(industry, "high-technology innovator")
		isNote = isNote || startsWith(industry, "second-largest industrial output")
		if isNote {
			notes = append(notes, industry)
			continue
		}
		industries = append(industries, industry)
	}
	o := orderedmap.New()
	if len(industries) > 0 {
		o.Set("industries", industries)
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

func industrialProductionGrowthRate(value string) (interface{}, error) {
	return stringToNumberWithGlobalRankAndDate(value, "annual_percentage_increase")
}

func (p *Page) laborForce() (interface{}, error) {
	force := orderedmap.New()
	p.tryAddingDataForSelector(force, "total_size", "2095", laborForceTotal)
	p.tryAddingDataForSelector(force, "by_occupation", "2048", laborForceByOccupation)
	keys := force.Keys()
	if len(keys) == 0 {
		return force, NoValueErr
	}
	return force, nil
}

func laborForceTotal(value string) (interface{}, error) {
	return stringToNumberWithGlobalRankAndDate(value, "total_people")
}

func laborForceByOccupation(value string) (interface{}, error) {
	k := "transport and communications"
	value = strings.Replace(value, "transport and communication", k, -1)
	value = strings.Replace(value, "transportation and communication", k, -1)
	value = strings.Replace(value, "transportation and communications", k, -1)
	k = "wholesale and retail"
	value = strings.Replace(value, "wholesale and retail trade", k, -1)
	value = strings.Replace(value, "wholesale and retail trade, restaurants, and hotels", k, -1)
	value = strings.Replace(value, "wholesale and retail distribution", k, -1)
	m, err := stringToPercentageMap(value, "occupation")
	if err != nil {
		return m, err
	}
	// See Cyprus
	m.Delete("labor_force___by_occupation")
	return m, nil
}

func unemploymentRate(value string) (interface{}, error) {
	return stringToListOfAnnualValues(value, "%")
}

func populationBelowPovertyLine(value string) (interface{}, error) {
	return stringToPercentage(value)
}

func householdIncomeByPercentageShare(value string) (interface{}, error) {
	if strings.Index(value, "NA%") > -1 {
		return value, NoValueErr
	}
	value = strings.Replace(value, " 10%:", "_ten_percent:", -1)
	return stringToMapOfNumbersWithUnits(value)
}

func distributionOfFamilyIncome(value string) (interface{}, error) {
	return stringToListOfAnnualValues(value, "gini_index")
}

func investmentGrossFixed(value string) (interface{}, error) {
	return stringToListOfAnnualValues(value, "percent_of_gdp")
}

func budget(value string) (interface{}, error) {
	value = strings.Replace(value, "$", "", -1)
	value = strings.Replace(value, "illion", "illion USD", -1)
	return stringToMapOfNumbersWithUnits(value)
}

func taxesAndOtherRevenues(value string) (interface{}, error) {
	value = strings.Replace(value, "% of GDP", " percent_of_gdp", -1)
	return stringToNumberWithGlobalRankAndDate(value, "percent_of_gdp")
}

func budgetSurplusOrDeficit(value string) (interface{}, error) {
	value = strings.Replace(value, "% of GDP", " percent_of_gdp", -1)
	return stringToNumberWithGlobalRankAndDate(value, "percent_of_gdp")
}

func publicDebt(value string) (interface{}, error) {
	return stringToListOfAnnualValues(value, "percent_of_gdp")
}

func fiscalYear(value string) (interface{}, error) {
	first, others := firstLine(value)
	firstBits := strings.Split(first, " - ")
	o := orderedmap.New()
	if first == "calendar year" {
		o.Set("start", "1 January")
		o.Set("end", "31 December")
	} else if len(firstBits) == 2 {
		o.Set("start", firstBits[0])
		o.Set("end", firstBits[1])
	} else {
		return value, NoValueErr
	}
	if len(others) > 0 {
		m, err := stringToMap(others)
		if err == nil {
			keys := m.Keys()
			for _, key := range keys {
				v, _ := m.Get(key)
				o.Set(key, v)
			}
		} else {
			o.Set("note", others)
		}
	}
	return o, nil
}

func inflationRate(value string) (interface{}, error) {
	return stringToListOfAnnualValues(value, "%")
}

func centralBankDiscountRate(value string) (interface{}, error) {
	return stringToListOfAnnualValues(value, "%")
}

func commercialBankPrimeLendingRate(value string) (interface{}, error) {
	return stringToListOfAnnualValues(value, "%")
}

func stockOfMoney(value string) (interface{}, error) {
	return stringToListOfAnnualValues(value, "USD")
}

func stockOfQuasiMoney(value string) (interface{}, error) {
	return stringToListOfAnnualValues(value, "USD")
}

func stockOfNarrowMoney(value string) (interface{}, error) {
	return stringToListOfAnnualValues(value, "USD")
}

func stockOfBroadMoney(value string) (interface{}, error) {
	return stringToListOfAnnualValues(value, "USD")
}

func stockOfDomesticCredit(value string) (interface{}, error) {
	return stringToListOfAnnualValues(value, "USD")
}

func marketValueOfPubliclyTradedShares(value string) (interface{}, error) {
	return stringToListOfAnnualValues(value, "USD")
}

func currentAccountBalance(value string) (interface{}, error) {
	return stringToListOfAnnualValues(value, "USD")
}

func (p *Page) exports() (interface{}, error) {
	exports := orderedmap.New()
	p.tryAddingDataForSelector(exports, "total_value", "2078", importExportsTotalValue)
	p.tryAddingDataForSelector(exports, "commodities", "2049", importExportsCommodities)
	p.tryAddingDataForSelector(exports, "partners", "2050", importExportsPartners)
	keys := exports.Keys()
	if len(keys) == 0 {
		return exports, NoValueErr
	}
	return exports, nil
}

func (p *Page) imports() (interface{}, error) {
	exports := orderedmap.New()
	p.tryAddingDataForSelector(exports, "total_value", "2087", importExportsTotalValue)
	// import commodities may be same as list of exorts, detect that here
	value, err := textForFieldKey(p.dom, "2058")
	if err == nil {
		isSameAsExports := strings.Index(value, "see listing for exports") > -1
		if isSameAsExports {
			p.tryAddingDataForSelector(exports, "commodities", "2049", importExportsCommodities)
		} else {
			p.tryAddingDataForSelector(exports, "commodities", "2058", importExportsCommodities)
		}
	}
	p.tryAddingDataForSelector(exports, "partners", "2061", importExportsPartners)
	keys := exports.Keys()
	if len(keys) == 0 {
		return exports, NoValueErr
	}
	return exports, nil
}

func importExportsTotalValue(value string) (interface{}, error) {
	return stringToListOfAnnualValues(value, "USD")
}

func importExportsCommodities(value string) (interface{}, error) {
	// see The World
	value = strings.Replace(value, "\n", ", ", -1)
	value = strings.Replace(value, "the whole range of industrial and agricultural goods and services", "", -1)
	value = strings.Replace(value, "top ten - share of world trade: ", "", -1)
	value = strings.Replace(value, "including ", ", ", -1)
	value, date, hasDate := stringWithoutDate(value)
	commodities, err := stringToList(value, listConditions{
		splitChars: ",;",
	})
	if err != nil {
		return value, err
	}
	commodityList := []string{}
	notes := []string{}
	for _, commodity := range commodities {
		commodity = strings.TrimSpace(commodity)
		// see The World
		isNote := false
		isNote = isNote || startsWith(commodity, "the whole range")
		isNote = isNote || startsWith(commodity, "see listing for exports")
		if isNote {
			notes = append(notes, commodity)
		}
		shouldRemoveLastWord := false
		// see Azerbaijan
		shouldRemoveLastWord = shouldRemoveLastWord || endsWith(commodity, "%")
		// see The World
		shouldRemoveLastWord = shouldRemoveLastWord || endsWith(commodity, "and ")
		if shouldRemoveLastWord {
			commodityBits := strings.Split(commodity, " ")
			commodityBits = commodityBits[0 : len(commodityBits)-1]
			commodity = strings.Join(commodityBits, " ")
		}
		commodityList = append(commodityList, commodity)
	}
	o := orderedmap.New()
	if len(commodityList) > 0 {
		o.Set("by_commodity", commodityList)
	}
	if len(notes) > 0 {
		note := strings.Join(notes, "; ")
		o.Set("note", note)
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

func importExportsPartners(value string) (interface{}, error) {
	return stringToImportExportPartnerList(value)
}

func reservesOfForeignExchangeAndGold(value string) (interface{}, error) {
	return stringToListOfAnnualValues(value, "USD")
}

func externalDebt(value string) (interface{}, error) {
	return stringToListOfAnnualValues(value, "USD")
}

func (p *Page) stockOfDirectForeignInvestment() (interface{}, error) {
	stock := orderedmap.New()
	p.tryAddingDataForSelector(stock, "at_home", "2198", stockOfDirectForeignInvestmentAtHome)
	p.tryAddingDataForSelector(stock, "abroad", "2199", stockOfDirectForeignInvestmentAbroad)
	keys := stock.Keys()
	if len(keys) == 0 {
		return stock, NoValueErr
	}
	return stock, nil
}

func stockOfDirectForeignInvestmentAtHome(value string) (interface{}, error) {
	return stringToListOfAnnualValues(value, "USD")
}

func stockOfDirectForeignInvestmentAbroad(value string) (interface{}, error) {
	return stringToListOfAnnualValues(value, "USD")
}

func exchangeRates(value string) (interface{}, error) {
	value = strings.Replace(value, " -", "", -1)
	return stringToListOfAnnualValues(value, "USD")
}

func economyOfTurkishCypriots(value string) (interface{}, error) {
	return stringToMap(value)
}

func (p *Page) electricity() (interface{}, error) {
	electricity := orderedmap.New()
	p.tryAddingDataForSelector(electricity, "access", "2268", electricityAccess)
	p.tryAddingDataForSelector(electricity, "production", "2232", electricityTotalKwh)
	p.tryAddingDataForSelector(electricity, "consumption", "2233", electricityTotalKwh)
	p.tryAddingDataForSelector(electricity, "exports", "2234", electricityTotalKwh)
	p.tryAddingDataForSelector(electricity, "imports", "2235", electricityTotalKwh)
	p.tryAddingDataForSelector(electricity, "installed_generating_capacity", "2236", electricityTotalKw)
	tryAddingData(electricity, "by_source", p.electricityFrom)
	keys := electricity.Keys()
	if len(keys) == 0 {
		return electricity, NoValueErr
	}
	return electricity, nil
}

func electricityAccess(value string) (interface{}, error) {
	value = strings.Replace(value, "electrification - total population", "total electrification", -1)
	value = strings.Replace(value, "electrification - urban areas", "urban electrification", -1)
	value = strings.Replace(value, "electrification - rural areas", "rural electrification", -1)
	value, date, hasDate := stringWithoutDate(value)
	o, err := stringToMap(value)
	if err != nil {
		return o, err
	}
	keys := o.Keys()
	for _, k := range keys {
		vInterface, _ := o.Get(k)
		v := vInterface.(string)
		if k == "population_without_electricity" {
			vNum, err := stringToNumber(v)
			if err == nil {
				m := orderedmap.New()
				m.Set("value", vNum)
				m.Set("units", "people")
				o.Set(k, m)
			}
		} else {
			vNum, err := stringToNumberWithUnits(v)
			if err == nil {
				o.Set(k, vNum)
			}
		}
	}
	if hasDate {
		o.Set("date", date)
	}
	return o, nil
}

func electricityTotalKwh(value string) (interface{}, error) {
	return stringToNumberWithGlobalRankAndDate(value, "kWh")
}

func electricityTotalKw(value string) (interface{}, error) {
	return stringToNumberWithGlobalRankAndDate(value, "kW")
}

func (p *Page) electricityFrom() (interface{}, error) {
	from := orderedmap.New()
	p.tryAddingDataForSelector(from, "fossil_fuels", "2237", electricityPercent)
	p.tryAddingDataForSelector(from, "nuclear_fuels", "2239", electricityPercent)
	p.tryAddingDataForSelector(from, "hydroelectric_plants", "2238", electricityPercent)
	p.tryAddingDataForSelector(from, "other_renewable_sources", "2240", electricityPercent)
	keys := from.Keys()
	if len(keys) == 0 {
		return from, NoValueErr
	}
	return from, nil
}

func electricityPercent(value string) (interface{}, error) {
	value = strings.Replace(value, " of total installed capacity", "", -1)
	return stringToNumberWithGlobalRankAndDate(value, "percent")
}

func (p *Page) crudeOil() (interface{}, error) {
	oil := orderedmap.New()
	p.tryAddingDataForSelector(oil, "production", "2241", crudeOilBblPerDay)
	p.tryAddingDataForSelector(oil, "exports", "2242", crudeOilBblPerDay)
	p.tryAddingDataForSelector(oil, "imports", "2243", crudeOilBblPerDay)
	p.tryAddingDataForSelector(oil, "proved_reserves", "2244", crudeOilBbl)
	keys := oil.Keys()
	if len(keys) == 0 {
		return oil, NoValueErr
	}
	return oil, nil
}

func crudeOilBblPerDay(value string) (interface{}, error) {
	return stringToNumberWithGlobalRankAndDate(value, "bbl_per_day")
}

func crudeOilBbl(value string) (interface{}, error) {
	return stringToNumberWithGlobalRankAndDate(value, "bbl")
}

func (p *Page) refinedPetroleumProducts() (interface{}, error) {
	petrol := orderedmap.New()
	p.tryAddingDataForSelector(petrol, "production", "2245", crudeOilBblPerDay)
	p.tryAddingDataForSelector(petrol, "consumption", "2246", crudeOilBblPerDay)
	p.tryAddingDataForSelector(petrol, "exports", "2247", crudeOilBblPerDay)
	p.tryAddingDataForSelector(petrol, "imports", "2248", crudeOilBblPerDay)
	keys := petrol.Keys()
	if len(keys) == 0 {
		return petrol, NoValueErr
	}
	return petrol, nil
}

func (p *Page) naturalGas() (interface{}, error) {
	gas := orderedmap.New()
	p.tryAddingDataForSelector(gas, "production", "2249", naturalGasCuM)
	p.tryAddingDataForSelector(gas, "consumption", "2250", naturalGasCuM)
	p.tryAddingDataForSelector(gas, "exports", "2251", naturalGasCuM)
	p.tryAddingDataForSelector(gas, "imports", "2252", naturalGasCuM)
	p.tryAddingDataForSelector(gas, "proved_reserves", "2253", naturalGasCuM)
	keys := gas.Keys()
	if len(keys) == 0 {
		return gas, NoValueErr
	}
	return gas, nil
}

func naturalGasCuM(value string) (interface{}, error) {
	return stringToNumberWithGlobalRankAndDate(value, "cubic_metres")
}

func carbonDioxideEmissions(value string) (interface{}, error) {
	return stringToNumberWithGlobalRankAndDate(value, "megatonnes")
}

func (p *Page) telephones() (interface{}, error) {
	t := orderedmap.New()
	p.tryAddingDataForSelector(t, "fixed_lines", "2150", telephonesFixedLines)
	p.tryAddingDataForSelector(t, "mobile_cellular", "2151", telephonesMobileCellular)
	p.tryAddingDataForSelector(t, "system", "2124", telephoneSystem)
	keys := t.Keys()
	if len(keys) == 0 {
		return t, NoValueErr
	}
	return t, nil
}

func telephonesFixedLines(value string) (interface{}, error) {
	value = strings.Replace(value, "per 100", "per one hundred", -1)
	return stringToMapOfNumbers(value)
}

func telephonesMobileCellular(value string) (interface{}, error) {
	// See Guam
	value = strings.Replace(value, "inhabitatnts", "inhabitants", -1)
	// Use words instead of numbers for json keys
	value = strings.Replace(value, "per 100", "per one hundred", -1)
	return stringToMapOfNumbers(value)
}

func telephoneSystem(value string) (interface{}, error) {
	return stringToMap(value)
}

func broadcastMedia(value string) (interface{}, error) {
	return value, nil
}

func radioBroacastStations(value string) (interface{}, error) {
	return value, nil
}

func televisionBroacastStations(value string) (interface{}, error) {
	return value, nil
}

func (p *Page) internet() (interface{}, error) {
	i := orderedmap.New()
	p.tryAddingDataForSelector(i, "country_code", "2154", internetCountryCode)
	p.tryAddingDataForSelector(i, "hosts", "2184", internetHosts)
	p.tryAddingDataForSelector(i, "users", "2153", internetUsers)
	keys := i.Keys()
	if len(keys) == 0 {
		return i, NoValueErr
	}
	return i, nil
}

func internetCountryCode(value string) (interface{}, error) {
	return value, nil
}

func internetHosts(value string) (interface{}, error) {
	return stringToMapOfNumbers(value)
}

func internetUsers(value string) (interface{}, error) {
	return stringToMapOfNumbers(value)
}

func communicationsNote(value string) (interface{}, error) {
	return value, nil
}

func (p *Page) airTransport() (interface{}, error) {
	t := orderedmap.New()
	p.tryAddingDataForSelector(t, "national_system", "2269", nationalAirTransportSystem)
	p.tryAddingDataForSelector(t, "civil_aircraft_registration_country_code_prefix", "2270", civilAircraftRegistrationCountryCodePrefix)
	tryAddingData(t, "airports", p.airports)
	p.tryAddingDataForSelector(t, "heliports", "2019", heliports)
	keys := t.Keys()
	if len(keys) == 0 {
		return t, NoValueErr
	}
	return t, nil
}

func nationalAirTransportSystem(value string) (interface{}, error) {
	return stringToMapOfNumbers(value)
}

func civilAircraftRegistrationCountryCodePrefix(value string) (interface{}, error) {
	value, date, hasDate := stringWithoutDate(value)
	o := orderedmap.New()
	o.Set("prefix", strings.TrimSpace(value))
	if hasDate {
		o.Set("date", date)
	}
	return o, nil
}

func (p *Page) airports() (interface{}, error) {
	a := orderedmap.New()
	p.tryAddingDataForSelector(a, "total", "2053", airportsTotal)
	p.tryAddingDataForSelector(a, "paved", "2030", airportsRunways)
	p.tryAddingDataForSelector(a, "unpaved", "2031", airportsRunways)
	keys := a.Keys()
	if len(keys) == 0 {
		return a, NoValueErr
	}
	return a, nil
}

func airportsTotal(value string) (interface{}, error) {
	return stringToNumberWithGlobalRankAndDate(value, "airports")
}

func airportsRunways(value string) (interface{}, error) {
	value = strings.Replace(value, ",", "", -1)
	value = strings.Replace(value, " m:", " metres:", -1)
	return stringToMapOfNumbers(value)
}

func heliports(value string) (interface{}, error) {
	return stringToNumberWithGlobalRankAndDate(value, "total")
}

func pipelines(value string) (interface{}, error) {
	value, date, hasDate := stringWithoutDate(value)
	p := orderedmap.New()
	ps := []*orderedmap.OrderedMap{}
	bits := strings.Split(value, "; ")
	for _, bit := range bits {
		o := orderedmap.New()
		bit = strings.TrimSpace(bit)
		b := strings.Split(bit, " ")
		if len(b) < 3 {
			continue
		}
		// type
		pipeType := strings.Join(b[0:len(b)-2], " ")
		o.Set("type", pipeType)
		// length
		length, err := stringToNumber(b[len(b)-2])
		if err != nil {
			continue
		}
		o.Set("length", length)
		// units
		units := b[len(b)-1]
		o.Set("units", units)
		ps = append(ps, o)
	}
	if len(ps) > 0 {
		p.Set("by_type", ps)
	}
	if hasDate {
		p.Set("date", date)
	}
	keys := p.Keys()
	if len(keys) == 0 {
		return p, NoValueErr
	}
	return p, nil
}

func railways(value string) (interface{}, error) {
	value, date, hasDate := stringWithoutDate(value)
	o, err := stringToMap(value)
	if err != nil {
		return o, err
	}
	keys := o.Keys()
	for _, key := range keys {
		value, _ := o.Get(key)
		// global rank
		if key == "country_comparison_to_the_world" {
			vInt, err := stringToNumber(value.(string))
			if err != nil {
				continue
			}
			o.Set("global_rank", vInt)
			continue
		}
		// rail value
		newValue, err := stringToRailLengthMap(value.(string))
		if err != nil {
			continue
		}
		o.Set(key, newValue)
	}
	// remove duplicate global_rank
	o.Delete("country_comparison_to_the_world")
	// set date
	if hasDate {
		o.Set("date", date)
	}
	return o, nil
}

func roadways(value string) (interface{}, error) {
	value, date, hasDate := stringWithoutDate(value)
	o, err := stringToMapOfNumbersWithUnits(value)
	if err != nil {
		return o, err
	}
	// set date
	if hasDate {
		o.Set("date", date)
	}
	return o, nil
}

func waterways(value string) (interface{}, error) {
	// date
	value, date, hasDate := stringWithoutDate(value)
	// notes
	value, ps := removeParenthesis(value)
	// length
	first, others := firstLine(value)
	o, err := stringToNumberWithUnits(first)
	if err != nil {
		return o, err
	}
	// note
	if len(ps) > 0 {
		note := strings.Join(ps, "; ")
		o.Set("note", note)
	}
	// other values such as global rank
	m, err := stringToMap(others)
	if err == nil {
		keys := m.Keys()
		for _, key := range keys {
			v, _ := m.Get(key)
			if key == "country_comparison_to_the_world" {
				vInt, err := stringToNumber(v.(string))
				if err == nil {
					o.Set("global_rank", vInt)
				}
			} else {
				o.Set(key, v.(string))
			}
		}
	}
	// date
	if hasDate {
		o.Set("date", date)
	}
	keys := o.Keys()
	if len(keys) == 0 {
		return o, NoValueErr
	}
	return o, nil
}

func merchantMarine(value string) (interface{}, error) {
	// date
	value, date, hasDate := stringWithoutDate(value)
	o, err := stringToMap(value)
	if err != nil {
		return o, err
	}
	keys := o.Keys()
	for _, key := range keys {
		vInterface, _ := o.Get(key)
		v := vInterface.(string)
		if key == "country_comparison_to_the_world" {
			vInt, err := stringToNumber(v)
			if err == nil {
				o.Set("global_rank", vInt)
			}
		} else if key == "total" {
			vInt, err := stringToNumber(v)
			if err == nil {
				o.Set(key, vInt)
			}
		} else if key == "by_type" {
			m, err := stringToListOfCounts(v, "type")
			if err == nil {
				o.Set(key, m)
			}
		} else if key == "foreign_owned" {
			m, err := stringToListOfCountsWithTotal(v, "country")
			if err == nil {
				o.Set(key, m)
			}
		} else if key == "registered_in_other_countries" {
			m, err := stringToListOfCountsWithTotal(v, "country")
			if err == nil {
				o.Set(key, m)
			}
		}
	}
	// remove duplicate global rank
	o.Delete("country_comparison_to_the_world")
	// date
	if hasDate {
		o.Set("date", date)
	}
	keys = o.Keys()
	if len(keys) == 0 {
		return o, NoValueErr
	}
	return o, nil
}

func portsAndTerminals(value string) (interface{}, error) {
	value = strings.Replace(value, " (TEUs)", "", -1)
	value = strings.Replace(value, "terminal (", "terminal(s) (", -1)
	value = strings.Replace(value, ", India)", " India)", -1)
	value = strings.Replace(value, "river or lake", "river and lake", -1)
	value = strings.Replace(value, "LNG terminal", "liquid_natural_gas terminal", -1)
	// date
	value, date, hasDate := stringWithoutDate(value)
	o, err := stringToMap(value)
	if err != nil {
		return o, err
	}
	keys := o.Keys()
	lc := listConditions{
		splitChars: ",;",
	}
	for _, key := range keys {
		vInterface, _ := o.Get(key)
		v := vInterface.(string)
		if key == "major_seaports" {
			l, err := stringToList(v, lc)
			if err == nil {
				o.Set(key, l)
			}
		} else if key == "dry_bulk_cargo_ports" {
			m, err := stringToListWithItemNotes(v, "place", "type")
			if err == nil {
				o.Set(key, m)
			}
		} else if key == "container_ports" {
			m, err := stringToListWithItemNotes(v, "place", "twenty_foot_equivalent_units")
			if err == nil {
				o.Set(key, m)
			}
		} else if key == "liquid_natural_gas_terminals_import" {
			m, err := stringToList(v, lc)
			if err == nil {
				o.Set(key, m)
			}
		} else if key == "liquid_natural_gas_terminals_export" {
			m, err := stringToList(v, lc)
			if err == nil {
				o.Set(key, m)
			}
		} else if key == "oil_terminals" {
			m, err := stringToList(v, lc)
			if err == nil {
				o.Set(key, m)
			}
		} else if key == "river_and_lake_ports" {
			m, err := stringToList(v, lc)
			if err == nil {
				o.Set(key, m)
			}
		} else if key == "bulk_cargo_ports" {
			m, err := stringToList(v, lc)
			if err == nil {
				o.Set(key, m)
			}
		} else if key == "lake_ports" {
			m, err := stringToList(v, lc)
			if err == nil {
				o.Set(key, m)
			}
		} else if key == "river_ports" {
			m, err := stringToList(v, lc)
			if err == nil {
				o.Set(key, m)
			}
		} else if key == "cruise_ferry_ports" {
			m, err := stringToList(v, lc)
			if err == nil {
				o.Set(key, m)
			}
		} else if key == "cruise_ports" {
			m, err := stringToList(v, lc)
			if err == nil {
				o.Set(key, m)
			}
		}
	}
	// remove duplicate global rank
	o.Delete("country_comparison_to_the_world")
	// date
	if hasDate {
		o.Set("date", date)
	}
	keys = o.Keys()
	if len(keys) == 0 {
		return o, NoValueErr
	}
	return o, nil
}

func shipyardsAndShipBuilding(value string) (interface{}, error) {
	return stringToMapOfNumbers(value)
}

func transportNote(value string) (interface{}, error) {
	return value, nil
}

func militaryExpenditures(value string) (interface{}, error) {
	return stringToListOfAnnualValues(value, "percent_of_gdp")
}

func militaryBranches(value string) (interface{}, error) {
	o := orderedmap.New()
	value, date, hasDate := stringWithoutDate(value)
	l, err := stringToList(value, listConditions{
		splitChars: ",;:",
	})
	if err != nil {
		return o, err
	}
	if len(l) == 0 {
		return o, NoValueErr
	}
	o.Set("by_name", l)
	if hasDate {
		o.Set("date", date)
	}
	return o, nil
}

func (p *Page) militaryManpower() (interface{}, error) {
	manpower := orderedmap.New()
	p.tryAddingDataForSelector(manpower, "available_for_military_service", "2105", manpowerNumbers)
	p.tryAddingDataForSelector(manpower, "fit_for_military_service", "2025", manpowerNumbers)
	p.tryAddingDataForSelector(manpower, "reaching_militarily_significant_age_annually", "2026", manpowerNumbers)
	keys := manpower.Keys()
	if len(keys) == 0 {
		return manpower, NoValueErr
	}
	return manpower, nil
}

func manpowerNumbers(value string) (interface{}, error) {
	return stringToMapOfNumbers(value)
}

func militaryServiceAgeAndObligation(value string) (interface{}, error) {
	o := orderedmap.New()
	value, date, hasDate := stringWithoutDate(value)
	bits, _ := stringToList(value, listConditions{
		splitChars: " -",
	})
	if len(bits) > 0 {
		ageNum, err := stringToNumber(bits[0])
		if err == nil {
			o.Set("years_of_age", ageNum)
		}
	}
	o.Set("note", value)
	if hasDate {
		o.Set("date", date)
	}
	return o, nil
}

func terroristGroups(value string) (interface{}, error) {
	return value, nil
}

func militaryNote(value string) (interface{}, error) {
	return value, nil
}

func disputes(value string) (interface{}, error) {
	list := strings.Split(value, "; ")
	return list, nil
}

func refugees(value string) (interface{}, error) {
	value = strings.Replace(value, " (country of origin)", "", -1)
	value = strings.Replace(value, "IDPs", "internally_displaced_persons", -1)
	o, err := stringToMap(value)
	if err != nil {
		return o, err
	}
	keys := o.Keys()
	for _, key := range keys {
		vInterface, _ := o.Get(key)
		v := vInterface.(string)
		if key == "refugees" {
			v, date, hasDate := stringWithoutDate(v)
			byCountry, err := stringToListWithItemNotes(v, "people", "country_of_origin")
			if err != nil {
				continue
			}
			m := orderedmap.New()
			m.Set("by_country", byCountry)
			if hasDate {
				m.Set("date", date)
			}
			o.Set(key, m)
		} else if key == "internally_displaced_persons" {
			m := orderedmap.New()
			v, date, hasDate := stringWithoutDate(v)
			v, ps := removeParenthesis(v)
			v = strings.TrimSpace(v)
			vNum, err := stringToNumber(v)
			if err == nil {
				m.Set("people", vNum)
			} else {
				ps = append([]string{v}, ps...)
			}
			if len(ps) > 0 {
				note := strings.Join(ps, "; ")
				m.Set("note", note)
			}
			if hasDate {
				m.Set("date", date)
			}
			o.Set(key, m)
		} else if key == "stateless_persons" {
			m := orderedmap.New()
			v, date, hasDate := stringWithoutDate(v)
			vNum, err := stringToNumber(v)
			if err == nil {
				m.Set("people", vNum)
			} else {
				m.Set("note", v)
			}
			if hasDate {
				m.Set("date", date)
			}
			o.Set(key, m)
		}
	}
	if len(keys) == 0 {
		return o, NoValueErr
	}
	return o, nil
}

func traffickingInPersons(value string) (interface{}, error) {
	return stringToMap(value)
}

func illicitDrugs(value string) (interface{}, error) {
	if len(value) == 0 {
		return value, NoValueErr
	}
	o, err := stringToMap(value)
	keys := o.Keys()
	if err != nil || len(keys) == 0 {
		o = orderedmap.New()
		o.Set("note", value)
	}
	return o, nil
}
