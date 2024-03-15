package util

import "assignmentdua/model"

func CreateResponse(isSuccess bool, data any, errorMessage string) model.Response {
	return model.Response{
		Success: true,
		Data:    data,
		Error:   errorMessage,
	}
}