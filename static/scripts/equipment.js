function add(itemId, buttonId) {
    $.get('/equipment/add', { id: itemId });
    var counter = $("#cart-item-counter")
    var count = parseInt(counter.text()) + 1
    if (count > 0) {
        counter.removeClass("display-none");
        $("#img-cart").attr("src", "/img/shopping-cart-loaded.svg")
    }
    counter.text(count)
    $("#" + buttonId).blur()
}

function mark(itemId, buttonId) {
    $.get('/equipment/mark', { id: itemId });
    $("#" + buttonId).blur()
}
