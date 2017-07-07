# CIA World Factbook API

Converts the [CIA World Factbook](https://www.cia.gov/library/publications/the-world-factbook/index.html) into a json data structure.

## Data

* [Latest](https://www.github.com/iancoleman/cia_world_factbook_api/master/data) - approx 3 MB
* [Historical](#TODO) - approx 30 MB
* [Html Archives](#TODO) - approx 240 MB

## Usage

If you just want the latest data, get it using the Latest link above.

If you also want to get the full historical data set, use the Historical link above.

If you want to parse the factbook html into json for yourself:

* clone this repository to your local machine.
* download the Html Archives above.
* edit `config.json` with the paths to your downloaded html archives.
* run `go run parse_html_to_json.go` to convert each country html to a json structure.
* run `go run create_weekly_json_files.go` to combine each individual country into a week-by-week data file.

If you want to fetch the html files yourself and then parse them:

* clone this repository to your local machine.
* edit `config.json` with the paths to use for the downloaded html archives.
* run `python fetch.py` to fetch the historical html files from archive.org (will take several days).
* run `go run parse_html_to_json.go` to convert each country html to a json structure.
* run `go run create_weekly_json_files.go` to combine each individual country into a week-by-week data file.

## Tests

* clone this repository to your local machine.
* `cd cwf/src/country`
* `go test`

## Contributing

Contributions are most welcome.

### Reporting Issues

Please report issues using the Issues tab at the top of this page.

### Pull Requests

If you modify the code please submit a pull request for review.

Most of the parsing logic is in `src/country` in the files `page.go` and `string_conversions.go`.

If the parser is modified, please update the `VERSION ` contant in `country/page.go`.

## License

MIT - see LICENSE
