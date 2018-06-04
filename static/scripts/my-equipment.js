function extend(itemId) {
    $.get("/my-equipment/extend", { id: itemId });
    var oldDate = $("#date-" + itemId).text();
    var newDate = moment(oldDate, "DD.MM.YYYY").add('days', 1);
    $("#date-" + itemId).text(newDate.format("DD.MM.YYYY"));
    $("#btn-" + itemId).blur()
}

function remove(itemId) {
    $.get("/my-equipment/remove", { id: itemId });
    $("#row-" + itemId).remove();
    $("#hr-" + itemId).remove();
}
