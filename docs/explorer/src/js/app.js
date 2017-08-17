(function() {

    var self = this;

    var DOM = {};
    DOM.fileLoader = $(".file-loader");
    DOM.filename = $(".filename");
    DOM.results = $(".results");
    DOM.loadHardcodedButton = $(".load-hardcoded");
    DOM.loadHardcodedLoading = $(".loading-hardcoded");
    DOM.loadHardcodedError = $(".error-hardcoded");

    var dropzone = null;
    var displayedData = null;
    var factbookJson = null;
    var factbookTable = new FactbookTable();
    var factbookCsv = new FactbookCsv();

    this.updateColumns = function(columnNames) {
        // create headings
        var rows = [];
        var headingRow = [
            "country",
            "source",
        ];
        for (var i=0; i<columnNames.length; i++) {
            headingRow.push(columnNames[i]);
        }
        rows.push(headingRow);
        // add other rows
        for (var countryKey in factbookJson.countries) {
            var countryJson = factbookJson.countries[countryKey];
            // name and source
            var countryRow = [
                countryJson.data.name,
                countryJson.metadata.source,
            ];
            for (var i=0; i<columnNames.length; i++) {
                var columnName = columnNames[i];
                var keyBits = columnName.split(".");
                var datapoint = countryJson;
                for (var j=0; j<keyBits.length; j++) {
                    var keyBit = keyBits[j];
                    if (!(keyBit in datapoint)) {
                        countryRow.push("");
                        break;
                    }
                    datapoint = datapoint[keyBit];
                    if (j == keyBits.length - 1) {
                        cellValue = datapoint;
                        if (datapoint instanceof Object) {
                            cellValue = JSON.stringify(datapoint);
                        }
                        countryRow.push(cellValue);
                        break;
                    }
                    if (!(datapoint instanceof Object && !(datapoint instanceof Array))) {
                        countryRow.push("");
                        break;
                    }
                }
            }
            rows.push(countryRow);
        }
        // set displayed data
        displayedData = rows;
        updateTabPanels();
    }

    function init() {
        // file loading
        dropzone = new Dropzone("#file-dropzone", {
            url: "#",
            acceptedFiles: ".json",
            accept: proccessFile,
            clickable: "#file-dropzone,#file-dropzone *",
        });
        // hardcoded file
        DOM.loadHardcodedButton.on("click", loadHardcoded);
    }

    function loadHardcoded() {
        showHardcodedLoading();
        $.ajax({
            url: "https://raw.githubusercontent.com/iancoleman/cia_world_factbook_api/d5c11e4689e388ed006acc0cd4e71cda282c775c/data/2017-07-31_factbook.json",
            success: loadHardcodedSuccess,
            error: loadHardcodedError,
            complete: hideHardcodedLoading,
        });
    }

    function loadHardcodedSuccess(jsonStr) {
        // clear existing dropzone files
        dropzone.removeAllFiles();
        // show the filename as the heading for results
        DOM.filename.text("2017-07-31_factbook.json");
        // clear existing data
        factbookJson = {};
        displayedData = [];
        // show the data
        parseFactbookJson(jsonStr);
    }

    function showHardcodedLoading() {
        DOM.loadHardcodedButton.addClass("hidden");
        DOM.loadHardcodedLoading.removeClass("hidden");
        DOM.loadHardcodedError.addClass("hidden");
    }

    function loadHardcodedError() {
        DOM.loadHardcodedButton.removeClass("hidden");
        DOM.loadHardcodedLoading.addClass("hidden");
        DOM.loadHardcodedError.removeClass("hidden");
    }

    function hideHardcodedLoading() {
        DOM.loadHardcodedButton.removeClass("hidden");
        DOM.loadHardcodedLoading.addClass("hidden");
        DOM.loadHardcodedError.addClass("hidden");
    }

    function proccessFile(file, done) {
        // don't show the file details in the dropzone
        dropzone.removeAllFiles();
        // show the filename as the heading for results
        DOM.filename.text(file.name);
        // clear existing data
        factbookJson = {};
        displayedData = [];
        // Read the file
        var reader = new FileReader();
        reader.addEventListener("loadend", function(event) {
            var jsonStr = event.target.result;
            parseFactbookJson(jsonStr);
        });
        reader.readAsText(file);
    }

    function parseFactbookJson(jsonStr) {
        factbookJson = JSON.parse(jsonStr);
        var columnList = new ColumnList(self, factbookJson);
        DOM.fileLoader.addClass("hidden");
        DOM.results.removeClass("hidden");
    }

    function updateTabPanels() {
        factbookTable.update(displayedData);
        factbookCsv.update(displayedData);
    }

    init();

})()
