function deleteAccount() {
    window.location.href = "/profil/delete";
}

function uploadFile() {
    $('#file-type').click();
    $('#btn-upload-file').blur()
    return false;
}
