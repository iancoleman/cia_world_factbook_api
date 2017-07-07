from bs4 import BeautifulSoup
import datetime
import json
import os
import requests
import time
import urllib

def writeStdout(s):
    print(s)

def countryListForFile(filelocation):
    countryPages = []
    f = open(filelocation)
    content = f.read()
    f.close()
    soup = BeautifulSoup(content, 'html.parser')
    countryEls = soup.select("select option")
    for countryEl in countryEls:
        if "value" in countryEl.attrs and countryEl.attrs["value"]:
            page = countryEl.attrs["value"].split("/")[-1]
            if not page.endswith(".html"):
                continue
            countryPages.append(page)
    return countryPages

# Saves the current world factbook as raw html

ajaxUrl = "https://web.archive.org/__wb/calendarcaptures?url=https%%3A%%2F%%2Fwww.cia.gov%%2Flibrary%%2Fpublications%%2Fthe-world-factbook%%2Fgeos%%2F%s.html&selected_year=%s"

f = open("config.json")
configStr = f.read()
f.close()
config = json.loads(configStr)
dstRoot = config["country_html_root"]
blacklistRoot = config["country_html_blacklist"]
archiveRoot = config["country_html_yearly_summaries"]

goodCitizenDelay = 2

yearForToday = datetime.datetime.utcnow().year

def getPage(pageFilename):
    # set the date as the end of this year
    currentYear = datetime.datetime.utcnow().year
    currentDate = datetime.datetime.utcnow().date()
    minYear = 2007
    while currentYear >= minYear:
        pageCode = pageFilename.replace(".html", "")
        yearlySummaryUrl = ajaxUrl % (pageCode, currentYear)
        # if year is this year, remove yearly summary from local cache
        if currentYear == yearForToday:
            removeArchiveYearlySummary(yearlySummaryUrl)
        # get the archive history for that year
        yearlySummaryContent = saveArchiveYearlySummary(yearlySummaryUrl)
        if yearlySummaryContent is None:
            currentDate = datetime.datetime(currentDate.year - 1, 12, 31).date()
            continue
        data = json.loads(yearlySummaryContent)
        # get links to archive pages
        # specifically get only
        # * https pages - http links to redirect with no content
        # * from the popup on the date (not the date itself)
        # * the earliest time from the popup
        links = []
        for month in data:
            for week in month:
                for day in week:
                    if day and "ts" in day:
                        for i, st in enumerate(day["st"]):
                            if st != 200:
                                continue
                            url = "https://web.archive.org/web/%s/https://www.cia.gov/library/publications/the-world-factbook/geos/%s.html"
                            ts = day["ts"][i]
                            link = url % (ts, pageCode)
                            links.append(link)
                            break
        # iterate backward through the year getting archive files
        while currentDate.year == currentYear:
            # look for previous monday in archive
            currentDate = getPrevMonday(currentDate)
            # find file that's most recent and before this date
            latestLinkDate = None
            latestLinkHref = None
            for link in links:
                linkHref = link
                dateStr = linkHref.split("/")[4][:8]
                linkDate = datetime.datetime.strptime(dateStr, "%Y%m%d").date()
                isEarlierThanPrevMonday = linkDate <= currentDate
                isLaterThanLatestDate = latestLinkDate is None or linkDate > latestLinkDate
                if isEarlierThanPrevMonday and isLaterThanLatestDate:
                    latestLinkDate = linkDate
                    latestLinkHref = linkHref
            if latestLinkHref is not None:
                # save file for this page
                pageContent = savePageForUrl(latestLinkHref, latestLinkDate)
            # ensure next fetch is prior to this file date
            if latestLinkDate is not None:
                currentDate = latestLinkDate
        # set year to previous year
        currentYear = currentDate.year

def yearlySummaryFilenameForUrl(url):
    pageFilename = urlToFilename(url)
    dstFilename = os.path.join(archiveRoot, pageFilename)
    return dstFilename

def urlToFilename(s):
    return urllib.parse.quote(s, safe='')

def removeArchiveYearlySummary(url):
    filename = yearlySummaryFilenameForUrl(url)
    # only remove if created more than 24h ago
    if not os.path.isfile(filename):
        return
    modified = os.path.getmtime(filename)
    cacheExpiry = int(time.time()) - 3 * 24 * 60 * 60 # 3 days
    if modified < cacheExpiry:
        writeStdout("Removing outdated yearly summary: %s" % filename)
        os.remove(filename)
    else:
        writeStdout("Using recently cached yearly summary: %s" % filename)

def saveArchiveYearlySummary(url):
    # create the filename for this page
    dstFilename = yearlySummaryFilenameForUrl(url)
    # create the directory to store this page
    os.makedirs(archiveRoot, exist_ok=True)
    if not os.path.isfile(dstFilename):
        print("Fetching", url)
        r = requests.get(url)
        yearlySummaryContent = r.text
        f = open(dstFilename, 'w')
        f.write(yearlySummaryContent)
        f.close()
        time.sleep(goodCitizenDelay)
    else:
        print("Reading", url)
        f = open(dstFilename)
        yearlySummaryContent = f.read()
        f.close()
    return yearlySummaryContent

def savePageForUrl(url, date):

    # get the filename for the page
    pageFilename = None
    bits = url.split("?")[0].split("/")
    bits.reverse()
    for bit in bits:
        if bit.endswith(".html"):
            pageFilename = bit
            break
    if pageFilename is None:
        print("Not saving blank page for", url)
        return

    # Create the directory for this date to store this set of pages
    dateStr = date.isoformat()
    dstDir = os.path.join(dstRoot, dateStr)
    os.makedirs(dstDir, exist_ok=True)

    blacklistDir = os.path.join(blacklistRoot, dateStr)

    # Create the filename for this page
    pageFilename = urlToFilename(url)
    dstFilename = os.path.join(dstDir, pageFilename)

    # Prepare the blacklist filename in case the page is blacklisted
    dstBlacklist = os.path.join(blacklistDir, pageFilename)

    # Fetch it if required
    if not os.path.isfile(dstFilename) and not os.path.isfile(dstBlacklist):
        print("Fetching %s" % url)
        r = requests.get(url)
        content = r.text
        # decide whether or not to blacklist this file based on content
        if shouldBlacklist(content):
            os.makedirs(blacklistDir, exist_ok=True)
            dstFilename = dstBlacklist
            print("Blacklisting %s" % url)
        # save the page content
        f = open(dstFilename, 'w')
        f.write(content)
        f.close()
        time.sleep(goodCitizenDelay)
    else:
        if os.path.isfile(dstFilename):
            print("Reading %s" % url)
            f = open(dstFilename)
            content = f.read()
            f.close()
        elif os.path.isfile(dstBlacklist):
            print("Blacklisted %s" % url)
            f = open(dstBlacklist)
            content = f.read()
            f.close()
    return content

# some files have no content, in which case they should be blacklisted
def shouldBlacklist(content):
    # detect illegal strings
    mustNotInclude = [
        "HTTP 301",
        "404 Not Found",
        "404 - Not Found",
        "Access Denied",
        "meta http-equiv=\"refresh\"",
        "Connection Failure",
        "Connection Timeout",
        "coldfusion.bootstrap",
    ]
    isBlacklisted = False
    for s in mustNotInclude:
        foundForbidden = content.find(s) > -1
        if foundForbidden:
            print("Found forbidden blacklist phrase: %s" % s)
        isBlacklisted = isBlacklisted or foundForbidden
    # detect missing strings
    mustInclude = [
        #"<!-- InstanceEnd -->",
    ]
    for s in mustInclude:
        missingRequired = content.find(s) == -1
        if missingRequired:
            print("Missing required blacklist phrase: %s" % s)
        isBlacklisted = isBlacklisted or missingRequired
    return isBlacklisted

def getPrevMonday(d):
    daysAfterPrevMonday = d.isoweekday() - 1 % 7
    if daysAfterPrevMonday == 0:
        daysAfterPrevMonday = 7
    return d - datetime.timedelta(days=daysAfterPrevMonday)


worldPage = "xx.html"
worldContent = getPage(worldPage)

print("Getting country list")
# get all other countries from all world pages
# could use set instead of list, but want to preserve order
countryPages = []
for dirDate in os.listdir(dstRoot):
    worldDir = os.path.join(dstRoot, dirDate)
    possibleFiles = os.listdir(worldDir)
    worldFiles = [x for x in possibleFiles if x.endswith(worldPage)]
    if len(worldFiles) > 1:
        print("WARNING found multiple world files in", worldDir, worldFiles)
    if len(worldFiles) == 1:
        # get list of countries for the world page on this date
        worldFile = os.path.join(worldDir, worldFiles[0])
        countries = countryListForFile(worldFile)
        # add any new countries to countryPages
        for country in countries:
            if country not in countryPages:
                countryPages.append(country)
# don't repeat fetch for the world
countryPages.remove(worldPage)
# don't fetch Baker Island
countryPages.remove("fq.html")
# fetch the historical files for the rest of the countries
print("Parsing %s countries" % len(countryPages))
for i, countryPage in enumerate(countryPages):
    print("Parsing %s of %s" % ((i+1), len(countryPages)))
    getPage(countryPage)
