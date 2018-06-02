function updateQuantity(id, value) {
    $.get('/cart/updatequantity', { id: id, quantity: value });
}

function updateDate(id, value) {
    $.get('/cart/updatedate', { id: id, date: value });
}

function remove(itemId) {
    $.get("/cart/remove", { id: itemId });
    $("#row-" + itemId).remove();
    $("#hr-" + itemId).remove();
    var counter = $("#cart-item-counter")
    var count = parseInt(counter.text()) + -1
    if (count < 1) {
        counter.addClass("display-none");
        $("#img-cart").attr("src", "/img/shopping-cart.svg")
    }
    counter.text(count)
}
