package country

import (
	"errors"
	"github.com/PuerkitoBio/goquery"
	"io/ioutil"
	"net/url"
	"orderedmap"
	"path"
	"strconv"
	"strings"
	"time"
)

const VERSION = "0.0.3-beta"

var NoValueErr = errors.New("No value")

type Page struct {
	filelocation string
	dom          *goquery.Document
	ParsedData   *orderedmap.OrderedMap
	NameKey      string
	HasData      bool
}

type Selector struct {
	FieldKey string
	Id       string
}

func NewPage(f string) (Page, error) {
	p := Page{
		filelocation: f,
		ParsedData:   orderedmap.New(),
	}
	// read the html file
	fileBytes, err := ioutil.ReadFile(p.filelocation)
	if err != nil {
		return p, err
	}
	// fix <br> tags before parsing to include newline as text
	fileString := string(fileBytes)
	fileString = strings.Replace(fileString, "<br>", "\n<br>", -1)
	// create new document
	p.dom, err = goquery.NewDocumentFromReader(strings.NewReader(fileString))
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
	tryAddingData(pageData, "terrorism", p.terrorism)
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

func (p *Page) tryAddingDataForSelector(d *orderedmap.OrderedMap, key string, selector Selector, valueFn func(string) (interface{}, error)) {
	valueStr, err := textForSelector(p.dom, selector)
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
	p.tryAddingDataForSelector(introData, "background", Selector{"2028", "introduction-background"}, introBackground)
	p.tryAddingDataForSelector(introData, "preliminary_statement", Selector{"2192", "introduction-preliminary-statement"}, preliminaryStatement)
	if len(introData.Keys()) == 0 {
		return introData, NoValueErr
	}
	return introData, nil
}

func (p *Page) geography() (interface{}, error) {
	geoData := orderedmap.New()
	p.tryAddingDataForSelector(geoData, "overview", Selector{"2203", "geography-geographic-overview"}, geographicOverview)
	p.tryAddingDataForSelector(geoData, "location", Selector{"2144", "geography-location"}, geographyLocation)
	p.tryAddingDataForSelector(geoData, "geographic_coordinates", Selector{"2011", "geography-geographic-coordinates"}, geographicCoordinates)
	p.tryAddingDataForSelector(geoData, "map_references", Selector{"2145", "geography-map-references"}, mapReferences)
	tryAddingData(geoData, "area", p.geographyArea)
	p.tryAddingDataForSelector(geoData, "land_boundaries", Selector{"2096", "geography-land-boundaries"}, landBoundaries)
	p.tryAddingDataForSelector(geoData, "coastline", Selector{"2060", "geography-coastline"}, coastline)
	p.tryAddingDataForSelector(geoData, "maritime_claims", Selector{"2106", "geography-maritime-claims"}, maritimeClaims)
	p.tryAddingDataForSelector(geoData, "climate", Selector{"2059", "geography-climate"}, climate)
	p.tryAddingDataForSelector(geoData, "terrain", Selector{"2125", "geography-terrain"}, terrain)
	p.tryAddingDataForSelector(geoData, "elevation", Selector{"2020", "geography-elevation"}, elevation)
	p.tryAddingDataForSelector(geoData, "natural_resources", Selector{"2111", "geography-natural-resources"}, naturalResources)
	p.tryAddingDataForSelector(geoData, "land_use", Selector{"2097", "geography-land-use"}, landUse)
	p.tryAddingDataForSelector(geoData, "irrigated_land", Selector{"2146", "geography-irrigated-land"}, irrigatedLand)
	p.tryAddingDataForSelector(geoData, "total_renewable_water_sources", Selector{"2201", ""}, totalRenewableWaterSources) // deprecated before id selectors came into use
	p.tryAddingDataForSelector(geoData, "freshwater_withdrawal", Selector{"2202", ""}, freshwaterWithdrawal)               // deprecated before id selectors came into use
	p.tryAddingDataForSelector(geoData, "population_distribution", Selector{"2266", "geography-population-distribution"}, populationDistribution)
	p.tryAddingDataForSelector(geoData, "natural_hazards", Selector{"2021", "geography-natural-hazards"}, naturalHazards)
	tryAddingData(geoData, "environment", p.environment)
	p.tryAddingDataForSelector(geoData, "note", Selector{"2113", "geography-note"}, geographyNote)
	if len(geoData.Keys()) == 0 {
		return geoData, NoValueErr
	}
	return geoData, nil
}

func (p *Page) people() (interface{}, error) {
	peopleData := orderedmap.New()
	p.tryAddingDataForSelector(peopleData, "population", Selector{"2119", "people-and-society-population"}, population)
	p.tryAddingDataForSelector(peopleData, "nationality", Selector{"2110", "people-and-society-nationality"}, nationality)
	p.tryAddingDataForSelector(peopleData, "ethnic_groups", Selector{"2075", "people-and-society-ethnic-groups"}, ethnicGroups)
	p.tryAddingDataForSelector(peopleData, "languages", Selector{"2098", "people-and-society-languages"}, languages)
	p.tryAddingDataForSelector(peopleData, "religions", Selector{"2122", "people-and-society-religions"}, religions)
	p.tryAddingDataForSelector(peopleData, "demographic_profile", Selector{"2257", "people-and-society-demographic-profile"}, demographicProfile)
	p.tryAddingDataForSelector(peopleData, "age_structure", Selector{"2010", "people-and-society-age-structure"}, ageStructure)
	p.tryAddingDataForSelector(peopleData, "dependency_ratios", Selector{"2261", "people-and-society-dependency-ratios"}, dependencyRatios)
	p.tryAddingDataForSelector(peopleData, "median_age", Selector{"2177", "people-and-society-median-age"}, medianAge)
	p.tryAddingDataForSelector(peopleData, "population_growth_rate", Selector{"2002", "people-and-society-population-growth-rate"}, populationGrowthRate)
	p.tryAddingDataForSelector(peopleData, "birth_rate", Selector{"2054", "people-and-society-birth-rate"}, birthRate)
	p.tryAddingDataForSelector(peopleData, "death_rate", Selector{"2066", "people-and-society-death-rate"}, deathRate)
	p.tryAddingDataForSelector(peopleData, "net_migration_rate", Selector{"2112", "people-and-society-net-migration-rate"}, netMigrationRate)
	p.tryAddingDataForSelector(peopleData, "population_distribution", Selector{"2267", "people-and-society-population-distribution"}, populationDistribution)
	p.tryAddingDataForSelector(peopleData, "urbanization", Selector{"2212", "people-and-society-urbanization"}, urbanization)
	p.tryAddingDataForSelector(peopleData, "major_urban_areas", Selector{"2219", "people-and-society-major-urban-areas-population"}, majorUrbanAreas)
	p.tryAddingDataForSelector(peopleData, "sex_ratio", Selector{"2018", "people-and-society-sex-ratio"}, sexRatio)
	p.tryAddingDataForSelector(peopleData, "mothers_mean_age_at_first_birth", Selector{"2256", "people-and-society-mother-s-mean-age-at-first-birth"}, mothersMeanAgeAtFirstBirth)
	p.tryAddingDataForSelector(peopleData, "maternal_mortality_rate", Selector{"2223", "people-and-society-maternal-mortality-rate"}, maternalMortalityRate)
	p.tryAddingDataForSelector(peopleData, "infant_mortality_rate", Selector{"2091", "people-and-society-infant-mortality-rate"}, infantMortalityRate)
	p.tryAddingDataForSelector(peopleData, "life_expectancy_at_birth", Selector{"2102", "people-and-society-life-expectancy-at-birth"}, lifeExpectancyAtBirth)
	p.tryAddingDataForSelector(peopleData, "total_fertility_rate", Selector{"2127", "people-and-society-total-fertility-rate"}, totalFertilityRate)
	p.tryAddingDataForSelector(peopleData, "contraceptive_prevalence_rate", Selector{"2258", "people-and-society-contraceptive-prevalence-rate"}, contraceptivePrevalenceRate)
	p.tryAddingDataForSelector(peopleData, "health_expenditures", Selector{"2225", "people-and-society-health-expenditures"}, healthExpenditures)
	p.tryAddingDataForSelector(peopleData, "physicians_density", Selector{"2226", "people-and-society-physicians-density"}, physiciansDensity)
	p.tryAddingDataForSelector(peopleData, "hospital_bed_density", Selector{"2227", "people-and-society-hospital-bed-density"}, hospitalBedDensity)
	p.tryAddingDataForSelector(peopleData, "drinking_water_source", Selector{"2216", "people-and-society-drinking-water-source"}, drinkingWaterSource)
	p.tryAddingDataForSelector(peopleData, "sanitation_facility_access", Selector{"2217", "people-and-society-sanitation-facility-access"}, sanitationFacilityAccess)
	tryAddingData(peopleData, "hiv_aids", p.hivAids)
	p.tryAddingDataForSelector(peopleData, "major_infectious_diseases", Selector{"2193", "people-and-society-major-infectious-diseases"}, majorInfectiousDiseases)
	p.tryAddingDataForSelector(peopleData, "adult_obesity", Selector{"2228", "people-and-society-obesity-adult-prevalence-rate"}, obesityAdultPrevalenceRate)
	p.tryAddingDataForSelector(peopleData, "underweight_children", Selector{"2224", "people-and-society-children-under-the-age-of-5-years-underweight"}, childrenUnderFiveYearsUnderweight)
	p.tryAddingDataForSelector(peopleData, "education_expenditures", Selector{"2206", "people-and-society-education-expenditures"}, educationExpenditures)
	p.tryAddingDataForSelector(peopleData, "literacy", Selector{"2103", "people-and-society-literacy"}, literacy)
	p.tryAddingDataForSelector(peopleData, "school_life_expectancy", Selector{"2205", "people-and-society-school-life-expectancy-primary-to-tertiary-education"}, schoolLifeExpectancy)
	p.tryAddingDataForSelector(peopleData, "child_labor", Selector{"2255", ""}, childLabor) // deprecated before id selectors came into use
	p.tryAddingDataForSelector(peopleData, "youth_unemployment", Selector{"2229", "people-and-society-unemployment-youth-ages-15-24"}, youthUnemployment)
	p.tryAddingDataForSelector(peopleData, "note", Selector{"2022", "people-and-society-people-note"}, peopleNote)
	if len(peopleData.Keys()) == 0 {
		return peopleData, NoValueErr
	}
	return peopleData, nil
}

func (p *Page) government() (interface{}, error) {
	governmentData := orderedmap.New()
	p.tryAddingDataForSelector(governmentData, "country_name", Selector{"2142", "government-country-name"}, countryName)
	p.tryAddingDataForSelector(governmentData, "union_name", Selector{"2189", "government-union-name"}, unionName)
	p.tryAddingDataForSelector(governmentData, "political_structure", Selector{"2190", "government-political-structure"}, politicalStructure)
	p.tryAddingDataForSelector(governmentData, "government_type", Selector{"2128", "government-government-type"}, governmentType)
	p.tryAddingDataForSelector(governmentData, "capital", Selector{"2057", "government-capital"}, capital)
	p.tryAddingDataForSelector(governmentData, "member_states", Selector{"2191", "government-member-states"}, memberStates)
	p.tryAddingDataForSelector(governmentData, "administrative_divisions", Selector{"2051", "government-administrative-divisions"}, administrativeDivisions)
	p.tryAddingDataForSelector(governmentData, "dependent_areas", Selector{"2068", "government-dependent-areas"}, dependentAreas)
	p.tryAddingDataForSelector(governmentData, "independence", Selector{"2088", "government-independence"}, independence)
	p.tryAddingDataForSelector(governmentData, "national_holidays", Selector{"2109", "government-national-holiday"}, nationalHoliday)
	p.tryAddingDataForSelector(governmentData, "constitution", Selector{"2063", "government-constitution"}, constitution)
	p.tryAddingDataForSelector(governmentData, "legal_system", Selector{"2100", "government-legal-system"}, legalSystem)
	p.tryAddingDataForSelector(governmentData, "international_law_organization_participation", Selector{"2220", "government-international-law-organization-participation"}, internationalLaw)
	p.tryAddingDataForSelector(governmentData, "citizenship", Selector{"2263", "government-citizenship"}, citizenship)
	p.tryAddingDataForSelector(governmentData, "suffrage", Selector{"2123", "government-suffrage"}, suffrage)
	p.tryAddingDataForSelector(governmentData, "executive_branch", Selector{"2077", "government-executive-branch"}, executiveBranch)
	p.tryAddingDataForSelector(governmentData, "legislative_branch", Selector{"2101", "government-legislative-branch"}, legislativeBranch)
	p.tryAddingDataForSelector(governmentData, "judicial_branch", Selector{"2094", "government-judicial-branch"}, judicialBranch)
	p.tryAddingDataForSelector(governmentData, "political_parties_and_leaders", Selector{"2118", "government-political-parties-and-leaders"}, politicalPartiesAndLeaders)
	p.tryAddingDataForSelector(governmentData, "political_pressure_groups_and_leaders", Selector{"2115", ""}, politicalPressureGroupsAndLeaders) // deprecated before id selectors came into use
	p.tryAddingDataForSelector(governmentData, "international_organization_participation", Selector{"2107", "government-international-organization-participation"}, internationalOrganizationParticipation)
	tryAddingData(governmentData, "diplomatic_representation", p.diplomaticRepresentation)
	p.tryAddingDataForSelector(governmentData, "flag_description", Selector{"2081", "government-flag-description"}, flagDescription)
	p.tryAddingDataForSelector(governmentData, "national_symbol", Selector{"2230", "government-national-symbol-s"}, nationalSymbol)
	tryAddingData(governmentData, "national_anthem", p.nationalAnthem)
	p.tryAddingDataForSelector(governmentData, "note", Selector{"2140", "government-government-note"}, governmentNote)
	if len(governmentData.Keys()) == 0 {
		return governmentData, NoValueErr
	}
	return governmentData, nil
}

func (p *Page) economy() (interface{}, error) {
	economyData := orderedmap.New()
	p.tryAddingDataForSelector(economyData, "overview", Selector{"2116", "economy-economy-overview"}, economyOverview)
	tryAddingData(economyData, "gdp", p.gdp)
	p.tryAddingDataForSelector(economyData, "gross_national_saving", Selector{"2260", "economy-gross-national-saving"}, grossNationalSaving)
	p.tryAddingDataForSelector(economyData, "agriculture_products", Selector{"2052", "economy-agriculture-products"}, agricultureProducts)
	p.tryAddingDataForSelector(economyData, "industries", Selector{"2090", "economy-industries"}, industries)
	p.tryAddingDataForSelector(economyData, "industrial_production_growth_rate", Selector{"2089", "economy-industrial-production-growth-rate"}, industrialProductionGrowthRate)
	tryAddingData(economyData, "labor_force", p.laborForce)
	p.tryAddingDataForSelector(economyData, "unemployment_rate", Selector{"2129", "economy-unemployment-rate"}, unemploymentRate)
	p.tryAddingDataForSelector(economyData, "population_below_poverty_line", Selector{"2046", "economy-population-below-poverty-line"}, populationBelowPovertyLine)
	p.tryAddingDataForSelector(economyData, "household_income_by_percentage_share", Selector{"2047", "economy-household-income-or-consumption-by-percentage-share"}, householdIncomeByPercentageShare)
	p.tryAddingDataForSelector(economyData, "distribution_of_family_income", Selector{"2172", "economy-distribution-of-family-income-gini-index"}, distributionOfFamilyIncome)
	p.tryAddingDataForSelector(economyData, "investment_gross_fixed", Selector{"2185", ""}, investmentGrossFixed) // deprecated before id selectors came into use
	p.tryAddingDataForSelector(economyData, "budget", Selector{"2056", "economy-budget"}, budget)
	p.tryAddingDataForSelector(economyData, "taxes_and_other_revenues", Selector{"2221", "economy-taxes-and-other-revenues"}, taxesAndOtherRevenues)
	p.tryAddingDataForSelector(economyData, "budget_surplus_or_deficit", Selector{"2222", "economy-budget-surplus-or-deficit"}, budgetSurplusOrDeficit)
	p.tryAddingDataForSelector(economyData, "public_debt", Selector{"2186", "economy-public-debt"}, publicDebt)
	p.tryAddingDataForSelector(economyData, "fiscal_year", Selector{"2080", "economy-fiscal-year"}, fiscalYear)
	p.tryAddingDataForSelector(economyData, "inflation_rate", Selector{"2092", "economy-inflation-rate-consumer-prices"}, inflationRate)
	p.tryAddingDataForSelector(economyData, "central_bank_discount_rate", Selector{"2207", "economy-central-bank-discount-rate"}, centralBankDiscountRate)
	p.tryAddingDataForSelector(economyData, "commercial_bank_prime_lending_rate", Selector{"2208", "economy-commercial-bank-prime-lending-rate"}, commercialBankPrimeLendingRate)
	p.tryAddingDataForSelector(economyData, "stock_of_money", Selector{"2209", ""}, stockOfMoney)            // deprecated before id selectors came into use
	p.tryAddingDataForSelector(economyData, "stock_of_quasi_money", Selector{"2210", ""}, stockOfQuasiMoney) // deprecated before id selectors came into use
	p.tryAddingDataForSelector(economyData, "stock_of_narrow_money", Selector{"2214", "economy-stock-of-narrow-money"}, stockOfNarrowMoney)
	p.tryAddingDataForSelector(economyData, "stock_of_broad_money", Selector{"2215", "economy-stock-of-broad-money"}, stockOfBroadMoney)
	p.tryAddingDataForSelector(economyData, "stock_of_domestic_credit", Selector{"2211", "economy-stock-of-domestic-credit"}, stockOfDomesticCredit)
	p.tryAddingDataForSelector(economyData, "market_value_of_publicly_traded_shares", Selector{"2200", "economy-market-value-of-publicly-traded-shares"}, marketValueOfPubliclyTradedShares)
	p.tryAddingDataForSelector(economyData, "current_account_balance", Selector{"2187", "economy-current-account-balance"}, currentAccountBalance)
	tryAddingData(economyData, "exports", p.exports)
	tryAddingData(economyData, "imports", p.imports)
	p.tryAddingDataForSelector(economyData, "reserves_of_foreign_exchange_and_gold", Selector{"2188", "economy-reserves-of-foreign-exchange-and-gold"}, reservesOfForeignExchangeAndGold)
	p.tryAddingDataForSelector(economyData, "external_debt", Selector{"2079", "economy-debt-external"}, externalDebt)
	tryAddingData(economyData, "stock_of_direct_foreign_investment", p.stockOfDirectForeignInvestment)
	p.tryAddingDataForSelector(economyData, "exchange_rates", Selector{"2076", "economy-exchange-rates"}, exchangeRates)
	//p.tryAddingDataForSelector(economyData, "economy_of_the_area_administered_by_turkish_cypriots", Selector{"2204", ""}, economyOfTurkishCypriots)
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
	p.tryAddingDataForSelector(energyData, "carbon_dioxide_emissions_from_consumption_of_energy", Selector{"2254", "energy-carbon-dioxide-emissions-from-consumption-of-energy"}, carbonDioxideEmissions)
	if len(energyData.Keys()) == 0 {
		return energyData, NoValueErr
	}
	return energyData, nil
}

func (p *Page) communications() (interface{}, error) {
	commsData := orderedmap.New()
	tryAddingData(commsData, "telephones", p.telephones)
	p.tryAddingDataForSelector(commsData, "broadcast_media", Selector{"2213", "communications-broadcast-media"}, broadcastMedia)
	p.tryAddingDataForSelector(commsData, "radio_broadcast_stations", Selector{"2013", ""}, radioBroacastStations)           // deprecated before id selectors came into use
	p.tryAddingDataForSelector(commsData, "television_broadcast_stations", Selector{"2015", ""}, televisionBroacastStations) // deprecated before id selectors came into use
	tryAddingData(commsData, "internet", p.internet)
	p.tryAddingDataForSelector(commsData, "note", Selector{"2138", "communications-communications-note"}, communicationsNote)
	if len(commsData.Keys()) == 0 {
		return commsData, NoValueErr
	}
	return commsData, nil
}

func (p *Page) transportation() (interface{}, error) {
	transportData := orderedmap.New()
	tryAddingData(transportData, "air_transport", p.airTransport)
	p.tryAddingDataForSelector(transportData, "pipelines", Selector{"2117", "transportation-pipelines"}, pipelines)
	p.tryAddingDataForSelector(transportData, "railways", Selector{"2121", "transportation-railways"}, railways)
	p.tryAddingDataForSelector(transportData, "roadways", Selector{"2085", "transportation-roadways"}, roadways)
	p.tryAddingDataForSelector(transportData, "waterways", Selector{"2093", "transportation-waterways"}, waterways)
	p.tryAddingDataForSelector(transportData, "merchant_marine", Selector{"2108", "transportation-merchant-marine"}, merchantMarine)
	p.tryAddingDataForSelector(transportData, "ports_and_terminals", Selector{"2120", "transportation-ports-and-terminals"}, portsAndTerminals)
	p.tryAddingDataForSelector(transportData, "shipyards_and_ship_building", Selector{"2231", ""}, shipyardsAndShipBuilding) // deprecated before id selectors came into use
	p.tryAddingDataForSelector(transportData, "note", Selector{"2008", "transportation-transportation-note"}, transportNote)
	if len(transportData.Keys()) == 0 {
		return transportData, NoValueErr
	}
	return transportData, nil
}

func (p *Page) militaryAndSecurity() (interface{}, error) {
	militaryData := orderedmap.New()
	p.tryAddingDataForSelector(militaryData, "expenditures", Selector{"2034", "military-and-security-military-expenditures"}, militaryExpenditures)
	p.tryAddingDataForSelector(militaryData, "branches", Selector{"2055", "military-and-security-military-branches"}, militaryBranches)
	tryAddingData(militaryData, "manpower", p.militaryManpower)
	p.tryAddingDataForSelector(militaryData, "service_age_and_obligation", Selector{"2024", "military-and-security-military-service-age-and-obligation"}, militaryServiceAgeAndObligation)
	p.tryAddingDataForSelector(militaryData, "terrorist_groups", Selector{"2265", ""}, terroristGroups) // moved into own terrorism section
	p.tryAddingDataForSelector(militaryData, "note", Selector{"2137", "military-and-security-military-note"}, militaryNote)
	if len(militaryData.Keys()) == 0 {
		return militaryData, NoValueErr
	}
	return militaryData, nil
}

func (p *Page) terrorism() (interface{}, error) {
	terrorismData := orderedmap.New()
	p.tryAddingDataForSelector(terrorismData, "home_based", Selector{"", "terrorism-terrorist-groups-home-based"}, terrorismHomeBased)          // came after fieldid selectors were in use
	p.tryAddingDataForSelector(terrorismData, "foreign_based", Selector{"", "terrorism-terrorist-groups-foreign-based"}, terrorismForeignBased) // came after fieldid selectors were in use
	if len(terrorismData.Keys()) == 0 {
		return terrorismData, NoValueErr
	}
	return terrorismData, nil
}

func (p *Page) transnationalIssues() (interface{}, error) {
	issuesData := orderedmap.New()
	p.tryAddingDataForSelector(issuesData, "disputes", Selector{"2070", "transnational-issues-disputes-international"}, disputes)
	p.tryAddingDataForSelector(issuesData, "refugees_and_iternally_displaced_persons", Selector{"2194", "transnational-issues-refugees-and-internally-displaced-persons"}, refugees)
	p.tryAddingDataForSelector(issuesData, "trafficking_in_persons", Selector{"2196", "transnational-issues-trafficking-in-persons"}, traffickingInPersons)
	p.tryAddingDataForSelector(issuesData, "illicit_drugs", Selector{"2086", "transnational-issues-illicit-drugs"}, illicitDrugs)
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
	areas, err := textForSelector(p.dom, Selector{"2147", "geography-area"})
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
	comparative, err := textForSelector(p.dom, Selector{"2023", "geography-area-comparative"})
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
		for _, key := range keys {
			if key == "total" {
				boundaryTotalStr, _ := boundaryMap.Get(key)
				boundaryTotalNum, err := stringToNumberWithUnits(boundaryTotalStr.(string))
				if err != nil {
					return boundaryMap, err
				}
				boundaryMap.Set(key, boundaryTotalNum)
			} else if key == "border_countries" {
				borderCountriesStr, _ := boundaryMap.Get(key)
				borderCountries, err := borderCountriesStringToSlice(borderCountriesStr.(string))
				if err != nil {
					return boundaryMap, err
				}
				boundaryMap.Set("border_countries", borderCountries)
			} else if key == "regional_borders" {
				regionalBordersStr, _ := boundaryMap.Get(key)
				regionalBorders, err := borderCountriesStringToSlice(regionalBordersStr.(string))
				if err != nil {
					return boundaryMap, err
				}
				boundaryMap.Set("regional_borders", regionalBorders)
			}
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
	currentIssues, err := textForSelector(p.dom, Selector{"2032", "geography-environment-current-issues"})
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
	internationalAgreements, err := textForSelector(p.dom, Selector{"2033", "geography-environment-international-agreements"})
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
	// See france
	value = strings.Replace(value, "Celtic and Latin with Teutonic", "Celtic, Latin, Teutonic", -1)
	value = strings.Replace(value, ", Basque minorities", ", Basque", -1)
	o, err := stringToPercentageList(value, "ethnicity")
	if err != nil {
		return o, err
	}
	// See france
	o.Delete("overseas_departments")
	return o, nil
}

func languages(value string) (interface{}, error) {
	o, err := stringToPercentageList(value, "language")
	if err != nil {
		return o, err
	}
	// See france
	o.Delete("overseas_departments")
	return o, nil
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
	o.Delete("africa")
	o.Delete("middle_east")
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
	// See Micronesia
	value = strings.Replace(value, "24total:", "total:", -1)
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
	keysToDelete := []string{}
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
		} else if v == "NA" {
			keysToDelete = append(keysToDelete, k)
		}
	}
	for _, k := range keysToDelete {
		o.Delete(k)
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
	s = removeCommasInNumbers(s)
	bits := splitByCommaOrSemicolon(s)
	for _, bit := range bits {
		bit = strings.TrimSpace(bit)
		placeMap := orderedmap.New()
		s, ps := removeParenthesis(bit)
		// get place name
		placeStr := ""
		placeBits := strings.Fields(s)
		populationStr := ""
		for _, placeBit := range placeBits {
			if startsWithNumber(placeBit) {
				populationStr = placeBit
			} else if placeBit == "million" {
				populationStr = populationStr + " " + placeBit
			} else {
				placeStr = placeStr + " " + placeBit
			}
		}
		placeStr = strings.TrimSpace(placeStr)
		placeStr = strings.ToLower(placeStr)
		placeStr = strings.Title(placeStr)
		if len(placeStr) == 0 {
			continue
		}
		// get place population
		population, err := stringToNumber(populationStr)
		hasPopulation := err == nil
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
	rateStr, err := textForSelector(p.dom, Selector{"2155", "people-and-society-hiv-aids-adult-prevalence-rate"})
	if err == nil {
		rate, err := stringToNumberWithGlobalRankAndDate(rateStr, "percent_of_adults")
		if err == nil {
			o.Set("adult_prevalence_rate", rate)
		}
	}
	// people living with hiv aids
	livingStr, err := textForSelector(p.dom, Selector{"2156", "people-and-society-hiv-aids-people-living-with-hiv-aids"})
	if err == nil {
		living, err := stringToNumberWithGlobalRankAndDate(livingStr, "total")
		if err == nil {
			o.Set("people_living_with_hiv_aids", living)
		}
	}
	// deaths
	deathsStr, err := textForSelector(p.dom, Selector{"2157", "people-and-society-hiv-aids-deaths"})
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
	value = strings.Replace(value, "notes:", "note:", -1)
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
	// pk.html
	value = strings.Replace(value, "no first order administrative divisions", "", -1)
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
			for len(typeNames) <= typeIndex {
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
		if len(typeNames) < typeNameIndex {
			typeName = typeNames[typeNameIndex]
		}
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
	// See Turkey
	value = strings.Replace(value, "highest court:", "highest courts:", -1)
	// See Vanuatu
	value = strings.Replace(value, ".highest court", "highest court", -1)
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
	diplomatInUs, err := textForSelector(p.dom, Selector{"2149", "government-diplomatic-representation-in-the-us"})
	if err == nil {
		person, err := stringToDiplomat(diplomatInUs)
		if err == nil {
			o.Set("in_united_states", person)
		}
	}
	// from the US
	diplomatFromUs, err := textForSelector(p.dom, Selector{"2007", "government-diplomatic-representation-from-the-us"})
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
	anthemStr, err := textForSelector(p.dom, Selector{"2218", "government-national-anthem"})
	if err != nil {
		return anthemStr, err
	}
	anthemStr = strings.Replace(anthemStr, "the Swiss anthem has four names:", "note: the Swiss anthem has four names - ", -1)
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
	p.tryAddingDataForSelector(gdp, "purchasing_power_parity", Selector{"2001", "economy-gdp-purchasing-power-parity"}, gdpPpp)
	p.tryAddingDataForSelector(gdp, "official_exchange_rate", Selector{"2195", "economy-gdp-official-exchange-rate"}, gdpOfficialExchangeRate)
	p.tryAddingDataForSelector(gdp, "real_growth_rate", Selector{"2003", "economy-gdp-real-growth-rate"}, gdpRealGrowthRate)
	p.tryAddingDataForSelector(gdp, "per_capita_purchasing_power_parity", Selector{"2004", "economy-gdp-per-capita-ppp"}, gdpPerCapitaPpp)
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
	p.tryAddingDataForSelector(composition, "by_end_use", Selector{"2259", "economy-gdp-composition-by-end-use"}, gdpCompositionByEndUse)
	p.tryAddingDataForSelector(composition, "by_sector_of_origin", Selector{"2012", "economy-gdp-composition-by-sector-of-origin"}, gdpCompositionBySector)
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
	p.tryAddingDataForSelector(force, "total_size", Selector{"2095", "economy-labor-force"}, laborForceTotal)
	p.tryAddingDataForSelector(force, "by_occupation", Selector{"2048", "economy-labor-force-by-occupation"}, laborForceByOccupation)
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
	p.tryAddingDataForSelector(exports, "total_value", Selector{"2078", "economy-exports"}, importExportsTotalValue)
	p.tryAddingDataForSelector(exports, "commodities", Selector{"2049", "economy-exports-commodities"}, importExportsCommodities)
	p.tryAddingDataForSelector(exports, "partners", Selector{"2050", "economy-exports-partners"}, importExportsPartners)
	keys := exports.Keys()
	if len(keys) == 0 {
		return exports, NoValueErr
	}
	return exports, nil
}

func (p *Page) imports() (interface{}, error) {
	exports := orderedmap.New()
	p.tryAddingDataForSelector(exports, "total_value", Selector{"2087", "economy-imports"}, importExportsTotalValue)
	// import commodities may be same as list of exorts, detect that here
	value, err := textForSelector(p.dom, Selector{"2058", "economy-imports-commodities"})
	if err == nil {
		isSameAsExports := strings.Index(value, "see listing for exports") > -1
		if isSameAsExports {
			p.tryAddingDataForSelector(exports, "commodities", Selector{"2049", "economy-exports-commodities"}, importExportsCommodities)
		} else {
			p.tryAddingDataForSelector(exports, "commodities", Selector{"2058", "economy-imports-commodities"}, importExportsCommodities)
		}
	}
	p.tryAddingDataForSelector(exports, "partners", Selector{"2061", "economy-imports-partners"}, importExportsPartners)
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
	p.tryAddingDataForSelector(stock, "at_home", Selector{"2198", "economy-stock-of-direct-foreign-investment-at-home"}, stockOfDirectForeignInvestmentAtHome)
	p.tryAddingDataForSelector(stock, "abroad", Selector{"2199", "economy-stock-of-direct-foreign-investment-abroad"}, stockOfDirectForeignInvestmentAbroad)
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
	p.tryAddingDataForSelector(electricity, "access", Selector{"2268", "energy-electricity-access"}, electricityAccess)
	p.tryAddingDataForSelector(electricity, "production", Selector{"2232", "energy-electricity-production"}, electricityTotalKwh)
	p.tryAddingDataForSelector(electricity, "consumption", Selector{"2233", "energy-electricity-consumption"}, electricityTotalKwh)
	p.tryAddingDataForSelector(electricity, "exports", Selector{"2234", "energy-electricity-exports"}, electricityTotalKwh)
	p.tryAddingDataForSelector(electricity, "imports", Selector{"2235", "energy-electricity-imports"}, electricityTotalKwh)
	p.tryAddingDataForSelector(electricity, "installed_generating_capacity", Selector{"2236", "energy-electricity-installed-generating-capacity"}, electricityTotalKw)
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
	p.tryAddingDataForSelector(from, "fossil_fuels", Selector{"2237", "energy-electricity-from-fossil-fuels"}, electricityPercent)
	p.tryAddingDataForSelector(from, "nuclear_fuels", Selector{"2239", "energy-electricity-from-nuclear-fuels"}, electricityPercent)
	p.tryAddingDataForSelector(from, "hydroelectric_plants", Selector{"2238", "energy-electricity-from-hydroelectric-plants"}, electricityPercent)
	p.tryAddingDataForSelector(from, "other_renewable_sources", Selector{"2240", "energy-electricity-from-other-renewable-sources"}, electricityPercent)
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
	p.tryAddingDataForSelector(oil, "production", Selector{"2241", "energy-crude-oil-production"}, crudeOilBblPerDay)
	p.tryAddingDataForSelector(oil, "exports", Selector{"2242", "energy-crude-oil-exports"}, crudeOilBblPerDay)
	p.tryAddingDataForSelector(oil, "imports", Selector{"2243", "energy-crude-oil-imports"}, crudeOilBblPerDay)
	p.tryAddingDataForSelector(oil, "proved_reserves", Selector{"2244", "energy-crude-oil-proved-reserves"}, crudeOilBbl)
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
	p.tryAddingDataForSelector(petrol, "production", Selector{"2245", "energy-refined-petroleum-products-production"}, crudeOilBblPerDay)
	p.tryAddingDataForSelector(petrol, "consumption", Selector{"2246", "energy-refined-petroleum-products-consumption"}, crudeOilBblPerDay)
	p.tryAddingDataForSelector(petrol, "exports", Selector{"2247", "energy-refined-petroleum-products-exports"}, crudeOilBblPerDay)
	p.tryAddingDataForSelector(petrol, "imports", Selector{"2248", "energy-refined-petroleum-products-imports"}, crudeOilBblPerDay)
	keys := petrol.Keys()
	if len(keys) == 0 {
		return petrol, NoValueErr
	}
	return petrol, nil
}

func (p *Page) naturalGas() (interface{}, error) {
	gas := orderedmap.New()
	p.tryAddingDataForSelector(gas, "production", Selector{"2249", "energy-natural-gas-production"}, naturalGasCuM)
	p.tryAddingDataForSelector(gas, "consumption", Selector{"2250", "energy-natural-gas-consumption"}, naturalGasCuM)
	p.tryAddingDataForSelector(gas, "exports", Selector{"2251", "energy-natural-gas-exports"}, naturalGasCuM)
	p.tryAddingDataForSelector(gas, "imports", Selector{"2252", "energy-natural-gas-imports"}, naturalGasCuM)
	p.tryAddingDataForSelector(gas, "proved_reserves", Selector{"2253", "energy-natural-gas-proved-reserves"}, naturalGasCuM)
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
	p.tryAddingDataForSelector(t, "fixed_lines", Selector{"2150", "communications-telephones-fixed-lines"}, telephonesFixedLines)
	p.tryAddingDataForSelector(t, "mobile_cellular", Selector{"2151", "communications-telephones-mobile-cellular"}, telephonesMobileCellular)
	p.tryAddingDataForSelector(t, "system", Selector{"2124", "communications-telephone-system"}, telephoneSystem)
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
	p.tryAddingDataForSelector(i, "country_code", Selector{"2154", "communications-internet-country-code"}, internetCountryCode)
	p.tryAddingDataForSelector(i, "hosts", Selector{"2184", ""}, internetHosts) // deprecated before id selectors came into use
	p.tryAddingDataForSelector(i, "users", Selector{"2153", "communications-internet-users"}, internetUsers)
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
	p.tryAddingDataForSelector(t, "national_system", Selector{"2269", "transportation-national-air-transport-system"}, nationalAirTransportSystem)
	p.tryAddingDataForSelector(t, "civil_aircraft_registration_country_code_prefix", Selector{"2270", "transportation-civil-aircraft-registration-country-code-prefix"}, civilAircraftRegistrationCountryCodePrefix)
	tryAddingData(t, "airports", p.airports)
	p.tryAddingDataForSelector(t, "heliports", Selector{"2019", "transportation-heliports"}, heliports)
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
	p.tryAddingDataForSelector(a, "total", Selector{"2053", "transportation-airports"}, airportsTotal)
	p.tryAddingDataForSelector(a, "paved", Selector{"2030", "transportation-airports-with-paved-runways"}, airportsRunways)
	p.tryAddingDataForSelector(a, "unpaved", Selector{"2031", "transportation-airports-with-unpaved-runways"}, airportsRunways)
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
	bits := splitByCommaOrSemicolon(value)
	for _, bit := range bits {
		o := orderedmap.New()
		bit = strings.TrimSpace(bit)
		b := strings.Split(bit, " ")
		if len(b) < 3 {
			continue
		}
		// get type, length and units
		pipeType := ""
		units := ""
		length := 0.0
		hasParsedLength := false
		lastItemWasLength := false
		var err error
		for _, thisBit := range b {
			if !hasParsedLength {
				length, err = stringToNumber(thisBit)
				if err == nil {
					hasParsedLength = true
					lastItemWasLength = true
				} else {
					lastItemWasLength = false
				}
			} else {
				if lastItemWasLength {
					units = thisBit
				} else {
					pipeType = pipeType + " " + thisBit
				}
				lastItemWasLength = false
			}
		}
		if !hasParsedLength {
			continue
		}
		// type
		o.Set("type", strings.TrimSpace(pipeType))
		// length
		o.Set("length", length)
		// units
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
			// See United Arab Emirates
			v = strings.Replace(v, "Papua New Guinea 6 (2010)", "Papua New Guinea 6) (2010)", -1)
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
	value = strings.Replace(value, "\nSaint Helena:", " Saint Helena - ", -1)
	value = strings.Replace(value, "\nAscension Island:", ", Ascension Island - ", -1)
	value = strings.Replace(value, "\nTristan da Cunha:", ", Tristan da Cunha - ", -1)
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
	p.tryAddingDataForSelector(manpower, "available_for_military_service", Selector{"2105", ""}, manpowerNumbers)               // deprecated before id selectors came into use
	p.tryAddingDataForSelector(manpower, "fit_for_military_service", Selector{"2025", ""}, manpowerNumbers)                     // deprecated before id selectors came into use
	p.tryAddingDataForSelector(manpower, "reaching_militarily_significant_age_annually", Selector{"2026", ""}, manpowerNumbers) // deprecated before id selectors came into use
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

// TODO consider parsing into a list
func terrorismHomeBased(value string) (interface{}, error) {
	return value, nil
}

// TODO consider parsing into a list
func terrorismForeignBased(value string) (interface{}, error) {
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
	// See France
	value = strings.Replace(value, "metropolitan France:", "note: metropolitan France - ", -1)
	value = strings.Replace(value, "French Guiana:", "note: French Guiana - ", -1)
	value = strings.Replace(value, "Martinique:", "note: Martinique - ", -1)
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
