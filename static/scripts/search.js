$(document).ready(function(){
    $("#search-box").on("keyup", function() {
        var value = $(this).val().toLowerCase();
        $(".search-row").filter(function() {
            $(this).toggle($(this).attr("name").toLowerCase().indexOf(value) > -1)
        });
    });
});
