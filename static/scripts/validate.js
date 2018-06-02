function validateForm() {
    if ($('#password').val() != $('#confirm-password').val()) {
        alert("Das Passwort stimmt nicht überein");
        return false
    }
}
