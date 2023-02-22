package helpers

import "strings"

type ResponseDataStruct struct {
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Error   interface{} `json:"error"`
	Data    interface{} `json:"data"`
}

type EmptyObjStruct struct {
}

func BuildResponse(status bool, message string, data interface{}) ResponseDataStruct {
	res := ResponseDataStruct{
		Status:  status,
		Message: message,
		Error:   nil,
		Data:    data,
	}
	return res
}

func BuildErrorResponse(message string, err string, data interface{}) ResponseDataStruct {
	splitedErr := strings.Split(err, "\n")
	res := ResponseDataStruct{
		Status:  false,
		Message: message,
		Error:   splitedErr,
		Data:    data,
	}
	return res
}
