FactbookCsv = function() {

    var DOM = {};
    DOM.csv = $("#csv textarea");
    DOM.csvTab = $("a[href='#csv']");

    DOM.csvTab.off("click");
    DOM.csvTab.on("shown.bs.tab", setSize);

    this.update = function(tableData) {
        var csvText = Papa.unparse(tableData);
        DOM.csv.text(csvText);
    }

    function setSize() {
        var height = DOM.csv[0].scrollHeight + 15;
        DOM.csv.css("height", height + "px");
    }

}
