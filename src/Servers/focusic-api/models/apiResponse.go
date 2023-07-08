package models

import "net/http"

type baseApi struct {
	Success    bool `json:"success"`
	StatusCode int  `json:"statusCode"`
}

type BaseApiSuccess struct {
	baseApi
	Data any `json:"data"`
}

type BaseApiFailed struct {
	baseApi
	Errors string `json:"errors"`
}

func ApiSuccessResult(data any) BaseApiSuccess {
	base := baseApi{true, http.StatusOK}
	return BaseApiSuccess{base, data}
}

func ApiSuccessResultWithStatusCode(data any, statusCode int) BaseApiSuccess {
	base := baseApi{true, statusCode}
	return BaseApiSuccess{base, data}
}

func ApiFailedResult(errors string) BaseApiFailed {
	base := baseApi{false, http.StatusBadRequest}
	return BaseApiFailed{base, errors}
}
