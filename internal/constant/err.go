package constant

import "errors"

var (
	ErrGetRequestError      = errors.New("GET_REQUEST_ERROR")
	ErrConvertToStructError = errors.New("CONVERT_TO_STRUCT_ERROR")
)
