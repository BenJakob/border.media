function deleteItem(itemId) {
    $.get("/admin/equipment/delete", { id: itemId });
    $("#row-" + itemId).remove();
    $("#hr-" + itemId).remove();
}
