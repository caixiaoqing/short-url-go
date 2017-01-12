package status

import "errors"

var ErrBadRequest = errors.New("HTTP 400: Bad Request")

type HTTPError struct {
	Code        int
	Description string
}