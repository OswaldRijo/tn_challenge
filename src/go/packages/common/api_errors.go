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

func NewAPIErrorInternal(err error) *APIErrorResource {
	return &APIErrorResource{
		error: err,
		code:  codes.Internal,
	}
}

func NewAPIErrorInvalidArgument(err error) *APIErrorResource {
	return &APIErrorResource{
		error: err,
		code:  codes.InvalidArgument,
	}
}

func NewAPIErrorPermissionsDenied(err error) *APIErrorResource {
	return &APIErrorResource{
		error: err,
		code:  codes.PermissionDenied,
	}
}
