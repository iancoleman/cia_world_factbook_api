ColumnList = function(app, factbookJson) {

    var self = this;

    var defaultColumns = [
        "data.economy.population_below_poverty_line.value",
        "data.people.literacy.total_population.value",
        "data.people.population.total",
    ];

    var DOM = {}
    DOM.columnList = $(".column-list");

    var columnItems = [];

    this.update = function() {
        var selectedColumnNames = [];
        for (var i=0; i<columnItems.length; i++) {
            var columnItem = columnItems[i];
            if (columnItem.checked) {
                var name = columnItem.name;
                selectedColumnNames.push(columnItem.name);
            }
        }
        app.updateColumns(selectedColumnNames);
    }

    function init() {
        DOM.columnList.empty();
        var uniqueKeys = [];
        if (!("countries" in factbookJson)) {
            return;
        }
        // get unique keys from every country
        for (var countryKey in factbookJson.countries) {
            var country = factbookJson.countries[countryKey];
            var allKeys = listAllKeysInObj(country, "");
            for (var i=0; i<allKeys.length; i++) {
                var keyName = allKeys[i];
                if (uniqueKeys.indexOf(keyName) == -1) {
                    uniqueKeys.push(keyName);
                }
            }
        }
        // sort the keys
        uniqueKeys.sort();
        // display the keys
        for (var i=0; i<uniqueKeys.length; i++) {
            var keyName = uniqueKeys[i];
            var isChecked = defaultColumns.indexOf(keyName) > -1;
            var columnItem = new ColumnListItem(self, keyName, isChecked);
            columnItems.push(columnItem);
            DOM.columnList.append(columnItem.el);
        }
        self.update();
    }

    function listAllKeysInObj(o, parentKey) {
        var keys = [];
        if (o instanceof Object && !(o instanceof Array)) {
            for (k in o) {
                var joiner = parentKey.length > 0 ? "." : "";
                var childKey = parentKey + joiner + k;
                var grandchildKeys = listAllKeysInObj(o[k], childKey);
                keys = keys.concat(grandchildKeys);
            }
        }
        else {
            keys.push(parentKey);
        }
        return keys;
    }

    init();

}
