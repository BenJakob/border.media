function lockAccount(userID) {
    $.get("/admin/edit-client/lock", { id: userID });
    $("#status").text(" Gesperrt");
    $("#btn-delete").blur()
}

function uploadFile() {
    $('#file-type').click();
    $('#btn-upload-file').blur()
    return false;
}
