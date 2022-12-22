package helpers

import "fmt"

type ErrorInfo struct {
	Message string `json:"message"`
}

func GenerateError(message string, err error) ErrorInfo {
	// TODO: Add log here
	fmt.Printf("-------ERROR-------\n\nmessage:%s\n\nerror:%v\n-------------------", err.Error(), err)
	return ErrorInfo{
		Message: message,
	}
}
