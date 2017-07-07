fs = require('fs');
fs.readFile("2017-01-02_factbook.json", function(err, content) {

function forKey(obj, key) {
    var keyBits = key.split(".")
    var resp = obj;
    for (var i=0; i<keyBits.length; i++) {
        keyBit = keyBits[i];
        if (!(keyBit in resp)) {
            return;
        }
        resp = resp[keyBit];
    }
    return resp;
}

if (err) {
    console.log(err);
    return;
}

d = JSON.parse(content);

var key = "data.people.major_urban_areas.places"
var countries = d.countries;

var buckets = [];

totalCountries = 0;
totalCountriesWithPopulation = 0;
totalPlaces = 0;
for (var countryKey in countries) {
    totalCountries = totalCountries + 1;
    var country = countries[countryKey];
    var places = forKey(country, key);
    if (!places) {
        console.log(countryKey);
        console.log(country.metadata.source);
        continue;
    }
    for (var i=0; i<places.length; i++) {
        var place = places[i];
        if (!place.population) {
            continue
        }
        if (place.population < 1000) {
            console.log("Less than a thousand people:");
            console.log(countryKey);
            console.log(place);
        }
        var bucketIndex = Math.floor(Math.log10(place.population));
        while (buckets.length <= bucketIndex) {
            buckets.push(0);
        }
        buckets[bucketIndex] = buckets[bucketIndex] + 1;
        totalPlaces = totalPlaces + 1;
    }
    totalCountriesWithPopulation = totalCountriesWithPopulation + 1;
}

console.log("");
console.log("Total countries", totalCountries);
console.log("Total countries with population", totalCountriesWithPopulation);
console.log("Total countries without population", totalCountries - totalCountriesWithPopulation);
console.log("Total places", totalPlaces);
console.log("Histogram:");
for (var i=0; i<buckets.length; i++) {
    var startRange = Math.pow(10, i) / 1e3;
    var endRange = Math.pow(10, i+1) / 1e3;
    var x = startRange.toLocaleString() + "-" + endRange.toLocaleString();
    var y = buckets[i];
    console.log(x + ";" + y);
}
console.log("");

})
