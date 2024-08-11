package common

import "google.golang.org/grpc/codes"

type APIErrorResource struct {
	error
	code codes.Code
}

func (apiError APIErrorResource) GetCode() codes.Code {
	return apiError.code
}

func NewAPIErrorResourceNotFound(err error) *APIErrorResource {
	return &APIErrorResource{
		error: err,
		code:  codes.NotFound,
	}
}
