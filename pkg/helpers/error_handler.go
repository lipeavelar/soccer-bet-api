package helpers

import (
	"fmt"
	"io"
)

type ErrorInfo struct {
	Message string `json:"message"`
}

func GenerateError(message string, err error, logger io.Writer) ErrorInfo {
	logger.Write([]byte(fmt.Sprintf("%s\n", err.Error())))
	return ErrorInfo{
		Message: message,
	}
}
