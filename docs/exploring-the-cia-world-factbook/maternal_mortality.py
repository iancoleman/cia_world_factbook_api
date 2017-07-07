# coding: utf-8
import json

# World Health Organization data

f = open("who_maternal_mortality.json")
content = f.read()
content = content.replace("Côte", "Cote")
f.close()

who_countries = {}

# rename some WHO countries to CIA naming scheme
renames = {
    "united_states_of_america": "united_states",
    "united_kingdom_of_great_britain_and_northern_ireland": "united_kingdom",
    "congo": "congo_republic_of_the",
    "democratic_republic_of_the_congo": "congo_democratic_republic_of_the",
    "venezuela_bolivarian_republic_of": "venezuela",
    "bolivia_plurinational_state_of": "bolivia",
    "russian_federation": "russia",
    "united_republic_of_tanzania": "tanzania",
    "the_former_yugoslav_republic_of_macedonia": "macedonia",
    "syrian_arab_republic": "syria",
    "viet_nam": "vietnam",
    "gambia": "gambia_the",
    "iran_islamic_republic_of": "iran",
    "myanmar": "burma",
    "bahamas": "bahamas_the",
    "republic_of_korea": "korea_south",
    "lao_people's_democratic_republic": "laos",
    "democratic_people's_republic_of_korea": "korea_north",
    "cote_d'ivoire": "cote_d'_ivoire",
    "republic_of_moldova": "moldova",
    "brunei_darussalam": "brunei",
        }

w = json.loads(content)
for fact in w["fact"]:
    gho = fact["dim"]["GHO"]
    if gho != "Maternal mortality ratio (per 100 000 live births)":
        continue
    country = fact["dim"]["COUNTRY"]
    country = country.lower()
    country = country.replace(" ", "_")
    country = country.replace("-", "_")
    country = country.replace("(", "")
    country = country.replace(")", "")
    if country in renames:
        country = renames[country]
    year = int(fact["dim"]["YEAR"])
    if country in who_countries:
        if year < who_countries[country]["year"]:
            continue
    value = int(fact["Value"].split("[")[0].replace(" ", ""))
    who_countries[country] = {
            "year": year,
            "value": value,
            }

# CIA World Factbook data

f = open("2017-01-02_factbook.json")
content = f.read()
f.close()

cia_countries = {}

c = json.loads(content)
for country in c["countries"]:
    try:
        d = c["countries"][country]["data"]["people"]["maternal_mortality_rate"]
        cia_countries[country] = {
                "value": d["deaths_per_100k_live_births"],
                "year": int(d["date"]),
                "name": c["countries"][country]["data"]["name"],
                "source": c["countries"][country]["metadata"]["source"],
                }
    except:
        pass

# Match the countries up
matches = {}
totalMismatches = 0
largestDiff = 0
largestDiffCountry = ""
for country in who_countries:
    if country in cia_countries:
        wv = who_countries[country]["value"]
        cv = cia_countries[country]["value"]
        cn = cia_countries[country]["name"]
        wy = who_countries[country]["year"]
        cy = cia_countries[country]["year"]
        isMatch = wv == cv and wy == cy
        diff = abs(wv - cv) / min(wv, cv)
        if diff > largestDiff:
            largestDiff = diff
            largestDiffCountry = cn
        if not isMatch:
            totalMismatches = totalMismatches + 1
        matches[country] = {
                "who": who_countries[country],
                "cia": cia_countries[country],
                "is_match": isMatch,
                "diff_percent": diff,
                }

def showtableheading():
    print("  <tr>")
    print("    <th>Country</th>")
    print("    <th colspan=2>Deaths per 100K live births</th>")
    print("  </tr>")
    print("  <tr>")
    print("    <th></th>")
    print("    <th>CIA</th>")
    print("    <th>WHO</th>")
    print("  </tr>")

def showtablerow(m):
    cia = m["cia"]
    who = m["who"]
    cn = cia["name"]
    cv = cia["value"]
    cy = cia["year"]
    cs = cia["source"]
    wv = who["value"]
    wy = who["year"]
    rowstyle = "match" if m["is_match"] else "mismatch"
    print("  <tr class=\"%s\">" % rowstyle)
    print("    <td>")
    print("      <a href=\"%s\">%s</a></td>" % (cs, cn))
    print("    </td>")
    print("    <td>%s (%s)</td>" % (cv, cy))
    print("    <td>%s (%s)</td>" % (wv, wy))
    print("  </tr>")

# Full table

print("<table>")
showtableheading()
countryKeys = sorted(matches.keys())
for country in countryKeys:
    m = matches[country]
    showtablerow(m)
print("</table>")
print ("")

# Partial table

print("<table>")
showtableheading()
countryKeys = sorted(matches.keys())
for country in countryKeys[:5]:
    m = matches[country]
    showtablerow(m)
print("  <tr>")
print("    <td>...</td>")
print("    <td></td>")
print("    <td></td>")
print("  </tr>")
for country in countryKeys[-5:]:
    m = matches[country]
    showtablerow(m)
print("  <tr>")
print("    <td><a href=\"maternal-mortality-rate\">Full Table</a></td>")
print("    <td></td>")
print("    <td></td>")
print("  </tr>")
print("</table>")
print ("")

# Print any countries from CIA that are not in WHO
cia_extra_countries = []
for country in cia_countries:
    if country not in who_countries:
        cia_extra_countries.append(country)
if len(cia_extra_countries) > 0:
    print("Countries in CIA with MMR data but not in WHO")
    for country in cia_extra_countries:
        print(country)
    print("")

# Print any countries from WHO that are not in CIA
who_extra_countries = []
for country in who_countries:
    if country not in cia_countries:
        who_extra_countries.append(country)
if len(who_extra_countries) > 0:
    print("Countries in WHO with MMR data but not in CIA")
    for country in who_extra_countries:
        print(country)
    print("")

# Print stats
print("Countries in both CIA and WHO", len(matches.keys()))
print("Countries in WHO", len(who_countries.keys()))
print("Countries in CIA", len(c["countries"].keys()))
print("Countries in CIA with data", len(cia_countries.keys()))
print("Countries in CIA without data", len(c["countries"].keys()) - len(cia_countries.keys()))
print("Countries in CIA but not in WHO", len(c["countries"].keys()) - len(matches.keys()))
print("Countries in WHO but not in CIA", len(who_countries.keys()) - len(matches.keys()))
print("Countries with matching data", len(matches.keys()) - totalMismatches)
print("Countries with mismatching data", totalMismatches)
print("Country with largest diff", largestDiffCountry)
print("Largest diff", largestDiff)
print("")
