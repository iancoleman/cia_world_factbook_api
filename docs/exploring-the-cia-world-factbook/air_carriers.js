fs = require('fs');
fs.readFile("2017-01-02_factbook.json", function(err, content) {

function median(m) {
    var middle = Math.floor((m.length - 1) / 2); // NB: operator precedence
        if (m.length % 2) {
        return m[middle];
    } else {
        return (m[middle] + m[middle + 1]) / 2.0;
    }
}

function mode(array) {
    if(array.length == 0)
        return null;
    var modeMap = {};
    var maxEl = array[0], maxCount = 1;
    for(var i = 0; i < array.length; i++)
    {
        var el = array[i];
        if(modeMap[el] == null)
            modeMap[el] = 1;
        else
            modeMap[el]++;  
        if(modeMap[el] > maxCount)
        {
            maxEl = el;
            maxCount = modeMap[el];
        }
    }
    return maxEl;
}

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

var key = "data.transportation.air_transport.national_system.number_of_registered_air_carriers"
var countries = d.countries;

var carriers = [];
var totalAirCarriers = 0;
var totalCountries = 0;
var totalCountriesWithAirCarriers = 0;
var minCarriers = 99999999;
var maxCarriers = 0;
var maxCarriersCountry = "";

for (var countryKey in countries) {
    if (countryKey == "european_union" || countryKey == "world") {
        continue
    }
    totalCountries = totalCountries + 1;
    var country = countries[countryKey];
    var carriersCountForThisCountry = forKey(country, key);
    if (!carriersCountForThisCountry) {
        carriersCountForThisCountry = 0;
    } else {
        totalCountriesWithAirCarriers = totalCountriesWithAirCarriers + 1;
    }
    totalAirCarriers = totalAirCarriers + carriersCountForThisCountry;
    if (carriersCountForThisCountry > maxCarriers) {
        maxCarriers = carriersCountForThisCountry;
        maxCarriersCountry = countryKey
    }
    if (carriersCountForThisCountry < minCarriers) {
        minCarriers = carriersCountForThisCountry;
    }
    carriers.push(carriersCountForThisCountry);
}

carriers.sort();

var median = median(carriers);
var mean = totalAirCarriers / carriers.length;
var mode = mode(carriers);

console.log("Total countries", totalCountries);
console.log("Total countries with air carriers", totalCountriesWithAirCarriers);
console.log("Total countries without air carriers", totalCountries - totalCountriesWithAirCarriers);
console.log("Total air carriers", totalAirCarriers);
console.log("Minimum air carriers", minCarriers);
console.log("Maximum air carriers", maxCarriers + " (" + maxCarriersCountry + ")");
console.log("Median air carriers", median);
console.log("Mean air carriers", mean.toFixed(1));
console.log("Mode air carriers", mode);
console.log("");

});
