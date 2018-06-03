$(document).ready(function(){
    $("#search-box").on("keyup", function() {
        var value = $(this).val().toLowerCase();

        $(".filter-element").filter(function() {
            $(this).toggle($(this).attr("data-name").toLowerCase().indexOf(value) > -1)
        });
    });

    $("#select-sortorder").on("change", function() {
        var value = $(this).val()
        var $wrapper = $(".elements");

        $wrapper.find(".filter-element").sort(function (a, b) {
            return ($(b).attr("data-" + value).toLowerCase()) < ($(a).attr("data-" + value).toLowerCase()) ? 1 : -1;
        })
        .appendTo( $wrapper );
    });

    $("#select-category").on("change", function() {
        var value = $(this).val()
        if (value == "all") {
            $(".filter-element").filter(function() {
                $(this).toggle(true)
            });
        } else {
            $(".filter-element").filter(function() {
                $(this).toggle($(this).attr("data-category").indexOf(value) > -1)
            });
        }
    });
});
