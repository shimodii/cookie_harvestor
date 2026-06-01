package interface

import (
	"divar_cookie_harvestor/utils"
)

func Make_phone_sending_request(phone_number string) struct {
	req := utils.Request[SendPhoneData]{
		Specification: Specification[SendPhoneData]{
			Page: 1,
			Data: SpecData[SendPhoneData]{
				Data: SendPhoneData{
					Phone: StrField{Str: StrValue{Value: phone_number}},
				},
			},
		},
	}

	return req
}

func Make_otp_sending_request(phone_number string, otp_code string, preAuth string, deviceId string) struct {
	req := utils.Request[VerifyData]{
		Specification: Specification[VerifyData]{
			Page: 1,
			Data: SpecData[VerifyData]{
				Data: VerifyData{
					Phone: 			   StrField{Str: StrValue{Value: phone_number}},
					Code:  			   StrField{Str: StrValue{Value: otp_code}},
					PreAuthSessionID:  StrField{Str: StrValue{Value: preAuth}},
					DeviceID:		   StrField{Str: StrValue{Value: deviceId}},
				},
			},
		},
	}

	return req
}
