package common

import (
	"errors"

	"google.golang.org/grpc/codes"
)

func HandleApiError(err error) codes.Code {
	var e *APIErrorResource
	isApiErr := errors.As(err, &e)
	if isApiErr {
		return e.GetCode()
	}

	return codes.InvalidArgument
}
