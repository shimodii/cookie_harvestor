package utils

type StrValue struct {
	Value string `json:"value"`
}

type StrField struct {
	Str StrValue `json:"str"`
}

type BoolValue struct {
	Value bool `json:"value"`
}

type BoolField struct {
	Boolean BoolValue `json:"boolean"`
}

type SendPhoneData struct {
	Phone StrField `json:"phone"`
}

type VerifyData struct {
	Phone            StrField  `json:"phone"`
	Code             StrField  `json:"code"`
	IsLegacy         BoolField `json:"is_legacy"`
	PreAuthSessionID StrField  `json:"pre_auth_session_id"`
	DeviceID         StrField  `json:"device_id"`
}

type SpecData[T any] struct {
	Data                      T                      `json:"data"`
	OnlineRequestResponseData map[string]interface{} `json:"online_request_response_data"`
}

type Specification[T any] struct {
	Page     int         `json:"page"`
	Data     SpecData[T] `json:"data"`
	IsReload bool        `json:"is_reload"`
}

type Request[T any] struct {
	Specification Specification[T] `json:"specification"`
}
