package errors

import (
	"fmt"
)

type GoApiError struct {
	Type    string
	Message string
	Code    int
}

func (e *GoApiError) Error() string {
	return fmt.Sprintf("Error: %s (Code: %d)", e.Message, e.Code)
}
