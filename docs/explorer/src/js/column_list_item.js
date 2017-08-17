ColumnListItem = function(columnList, name, isChecked) {

    var self = this;

    var template = $("#column-list-item-template").html();

    this.el = $(template);
    this.el.find(".name").text(name);

    this.name = name;
    this.checked = isChecked;

    var checkbox = this.el.find("input[type=checkbox]");

    if (isChecked) {
        checkbox.prop("checked", true);
    }

    checkbox.on("change", function() {
        self.checked = checkbox.prop("checked");
        columnList.update();
    });
}
