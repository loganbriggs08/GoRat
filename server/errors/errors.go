package errors

type MethodError struct {
	ErrorCode    uint64 `json:"error_code"`
	ErrorMessage string `json:"error_message"`
}
