function deleteItem(itemId) {
    window.location.href = "/admin/equipment/delete?id=" + itemId;
}

function uploadFile() {
    $('#file-type').click();
    $('#btn-upload-file').blur()
    return false;
}
