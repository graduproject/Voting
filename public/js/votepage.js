window.onload = function () {
    // Listening for auth state changes.

    document.getElementById('voter_info_num').addEventListener('keyup', updateSignInButtonUI);
    document.getElementById('voter_info_num').addEventListener('change', updateSignInButtonUI);

    window.recaptchaVerifier = new firebase.auth.RecaptchaVerifier('sign-in-button', {
        'size': 'invisible',
        'callback': function (response) {
            // reCAPTCHA solved, allow signInWithPhoneNumber.
            onSignInSubmit();
        }
    });

    recaptchaVerifier.render().then(function (widgetId) {
        window.recaptchaWidgetId = widgetId;
        updateSignInButtonUI();
    });
};

function getPhoneNumberFromUserInput() {
    return document.getElementById('voter_info_num').value;
}

function onSignInSubmit() {
    if (isUserInfoValid()) {
        window.signingIn = true;
        updateSignInButtonUI();
        var phoneNumber = getPhoneNumberFromUserInput();
        var appVerifier = window.recaptchaVerifier;
        firebase.auth().signInWithPhoneNumber(phoneNumber, appVerifier)
            .then(function (confirmationResult) {
                // SMS sent. Prompt user to type the code from the message, then sign the
                // user in with confirmationResult.confirm(code).
                window.confirmationResult = confirmationResult;
                window.signingIn = false;
                updateSignInButtonUI();
                updateSignInFormUI();
            })
    }
}


function isUserInfoValid() {
    var pattern = "[0-9]{13}";
    var phoneNumber = getPhoneNumberFromUserInput();
    return phoneNumber.search(pattern) !== -1;
}

function updateSignInFormUI() {
    if (firebase.auth().currentUser || window.confirmationResult) {
        document.getElementById('sign-in-form').style.display = 'none';
    } else {
        resetReCaptcha();
        document.getElementById('sign-in-form').style.display = 'block';
    }
}

function updateSignInButtonUI() {
    document.getElementById('sign-in-button').disabled =
        !isPhoneNumberValid()
        || !!window.signingIn;
}

function updateVerifyCodeButtonUI() {
    document.getElementById('verify-code-button').disabled =
        !!window.verifyingCode
        || !getCodeFromUserInput();
}