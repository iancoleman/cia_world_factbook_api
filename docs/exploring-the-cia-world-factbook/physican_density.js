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

function showRow(rank, name, density, source) {
    console.log("  <tr>");
    console.log("    <td>" + rank + "</td>");
    console.log("    <td>");
    console.log("      <a href=\"" + source + "\">");
    console.log("        " + name);
    console.log("      </a>");
    console.log("    </td>");
    console.log("    <td>" + density + "</td>");
    console.log("  </tr>");
}

if (err) {
    console.log(err);
    return;
}

d = JSON.parse(content);

var key = "data.people.physicians_density.physicians_per_1000_population"
var populationKey = "data.people.population.total"
var countries = d.countries;

var densities = [];
var totalCountries = 0;
var totalCountriesWithStat = 0;
var totalPopulationWithoutStat = 0;
var largestPopulationWithoutStat = 0;
var largestCountryWithoutStat = "";

for (var countryKey in countries) {
    if (countryKey == "european_union" || countryKey == "world") {
        continue
    }
    var country = countries[countryKey];
    var countryName = country.data.name;
    totalCountries = totalCountries + 1;
    var density = forKey(country, key);
    if (!density) {
        var population = forKey(country, populationKey);
        if (!!population) {
            if (population > largestPopulationWithoutStat) {
                largestPopulationWithoutStat = population;
                largestCountryWithoutStat = countryName
            }
            totalPopulationWithoutStat = totalPopulationWithoutStat + population;
        }
        continue;
    }
    totalCountriesWithStat = totalCountriesWithStat + 1;
    densities.push({
        name: countryName,
        density: density,
        source: country.metadata.source,
    });
}

densities.sort(function(a,b) {
    return b.density - a.density;
});

console.log("<table>");
console.log("  <tr>");
console.log("    <th>Rank</th>");
console.log("    <th>Country</th>");
console.log("    <th>Physicians per 1000 population</th>");
console.log("  </tr>");
console.log("  <tr>");
console.log("    <td colspan=3>");
console.log("      <a href=\"#TODO\">Full Table</a>");
console.log("    </td>");
console.log("  </tr>");
for (var i=0; i<5; i++) {
    var density = densities[i];
    showRow(i+1, density.name, density.density, density.source);
}
console.log("    <td colspan=3>...</td>");
for (var i=densities.length-5; i<densities.length; i++) {
    var density = densities[i];
    showRow(i+1, density.name, density.density, density.source);
}
console.log("</table>");
console.log("");

console.log("<table>");
console.log("  <tr>");
console.log("    <th>Rank</th>");
console.log("    <th>Country</th>");
console.log("    <th>Physicians per 1000 population</th>");
console.log("  </tr>");
for (var i=0; i<densities.length; i++) {
    var density = densities[i];
    showRow(i+1, density.name, density.density, density.source);
}
console.log("</table>");
console.log("");
console.log("Total countries", totalCountries);
console.log("Total countries with stat", totalCountriesWithStat);
console.log("Total countries without stat", totalCountries - totalCountriesWithStat);
console.log("Total population without stat", totalPopulationWithoutStat.toLocaleString());
console.log("Largest country without stat", largestCountryWithoutStat);
console.log("Largest population without stat", largestPopulationWithoutStat.toLocaleString());
console.log("");

});
