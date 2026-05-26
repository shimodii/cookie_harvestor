package utils

const (
	API_SENDING_PHONE_NUMBER = "https://api.divar.ir/v8/auth/open-initiate-page"
	API_SENDING_OTP_CODE     = "https://api.divar.ir/v8/auth/open-initiate-page" // same as phone number submitting, adding otp code in req body
	API_GETTING_TOKENS       = "https://api.divar.ir/v8/authenticate/signinup/code/consume"
)
