package util

type ResponseCode int

const (
	SuccessCode    ResponseCode = 200 // success
	ErrorCode      ResponseCode = 400 // failed
	TokenErrorCode ResponseCode = 401 // token error
	SignErrorCode  ResponseCode = 402 // sign error
	ScopeErrorCode ResponseCode = 403 // scope
	NoFound        ResponseCode = 404 // not found
	Verification   ResponseCode = 422 // Parameter error
	SystemError    ResponseCode = 500 // System error
)
