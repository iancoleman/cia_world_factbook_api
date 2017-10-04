package main

import (
	"country"
	"encoding/json"
	"io/ioutil"
	"logger"
	"os"
	"path"
	"runtime"
	"sync"
)

// Converts html files into json files.
// Expects html files in dirs pages/YYYY-MM-DD/*.html
func main() {
	// prepare for concurrent parsing
	numCpus := runtime.NumCPU()
	runtime.GOMAXPROCS(numCpus)
	var wg sync.WaitGroup
	// read config
	configBytes, err := ioutil.ReadFile("config.json")
	if err != nil {
		logger.Stderr("Error reading config.json")
		logger.Stderr(err)
		return
	}
	// parse config
	config := map[string]string{}
	err = json.Unmarshal(configBytes, &config)
	countryHtmlRoot, exists := config["country_html_root"]
	if !exists {
		logger.Stderr("Missing config value: country_html_root")
	}
	countryJsonRoot, exists := config["country_json_root"]
	if !exists {
		logger.Stderr("Missing config value: country_json_root")
	}
	// Get directories
	dirs, err := ioutil.ReadDir(countryHtmlRoot)
	if err != nil {
		logger.Stderr("Error reading country_html_root directory")
		logger.Stderr(countryHtmlRoot)
		logger.Stderr(err)
		return
	}
	// iterate over date directories
	for _, dir := range dirs {
		// ignore files
		if !dir.IsDir() {
			continue
		}
		// get the path and files for this directory
		filesRoot := path.Join(countryHtmlRoot, dir.Name())
		files, err := ioutil.ReadDir(filesRoot)
		if err != nil {
			logger.Stderr("Error reading directory")
			logger.Stderr(filesRoot)
			logger.Stderr(err)
			return
		}
		// iterate over page files
		for _, file := range files {
			// ignore directories
			if file.IsDir() {
				logger.Stderr("Unexpected directory")
				logger.Stderr(file.Name())
				return
			}
			filelocation := path.Join(filesRoot, file.Name())
			// check if already parsed
			dstDir := path.Join(countryJsonRoot, dir.Name())
			dst := path.Join(dstDir, file.Name()+".json")
			wg.Add(1)
			go func(src, dst string, wg *sync.WaitGroup) {
				defer wg.Done()
				_, err = os.Stat(dst)
				if !os.IsNotExist(err) {
					//logger.Stdout("Already parsed", filelocation)
					return
				}
				// parse the file
				logger.Stdout("Parsing", filelocation)
				p, err := country.NewPage(filelocation)
				if err != nil {
					logger.Stderr("Error parsing file")
					logger.Stderr(filelocation)
					logger.Stderr(err)
					return
				}
				// save the parsed json
				content, err := json.MarshalIndent(p.ParsedData, "", "  ")
				if err != nil {
					logger.Stderr("Error marshalling parsed page to json")
					logger.Stderr(filelocation)
					logger.Stderr(err)
					return
				}
				err = os.MkdirAll(dstDir, 0775)
				if err != nil {
					logger.Stderr("Error creating json directory")
					logger.Stderr(dstDir)
					logger.Stderr(err)
					return
				}
				err = ioutil.WriteFile(dst, content, 0664)
				if err != nil {
					logger.Stderr("Error saving content for json file")
					logger.Stderr(dstDir)
					logger.Stderr(err)
					return
				}
			}(filelocation, dst, &wg)
		}
	}
	wg.Wait()
}
