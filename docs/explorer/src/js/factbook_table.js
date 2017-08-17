FactbookTable = function() {

    var DOM = {};
    DOM.table = $("#tabulated .table");

    this.update = function(tableData) {
        DOM.table.empty();
        var thead = $("<thead>");
        // headings
        var headings = tableData[0];
        var rowEl = $("<tr>");
        for (var i=0; i<headings.length; i++) {
            if (i == 1) {
                continue;
            }
            var cell = $("<th>");
            cell.text(headings[i]);
            rowEl.append(cell);
        }
        thead.append(rowEl);
        DOM.table.append(thead);
        // data
        var tbody = $("<tbody>");
        for (var i=1; i<tableData.length; i++) {
            var rowEl = $("<tr>");
            var rowData = tableData[i];
            // country name
            var cell = $("<td>");
            var link = $("<a>");
            link.text(rowData[0]);
            link.attr("href", rowData[1]);
            link.attr("target", "_blank");
            link.attr("title", "Opens in new tab");
            cell.append(link);
            rowEl.append(cell);
            // other columns
            for (var j=2; j<rowData.length; j++) {
                var cell = $("<td>");
                var cellText = rowData[j];
                if (typeof(cellText) == "number") {
                    cell.addClass("text-right");
                    cellText = cellText.toLocaleString();
                }
                cell.text(cellText);
                rowEl.append(cell);
            }
            tbody.append(rowEl);
        }
        DOM.table.append(tbody);
    }

}
