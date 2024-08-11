package common

import (
	"errors"

	"google.golang.org/grpc/codes"
)

func HandleApiError(err error) codes.Code {
	var e APIErrorResource
	switch {
	case errors.As(err, &e):
		return e.GetCode()
	default:
		return codes.InvalidArgument
	}
}
