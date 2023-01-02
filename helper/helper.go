package helper

import "reflect"

type Response struct {
	Code    int         `json:"code"`
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Size    int         `json:"size"`
	Data    interface{} `json:"data"`
}

type ResponseWithoutData struct {
	Code    int    `json:"code"`
	Status  string `json:"status"`
	Message string `json:"message"`
}

func APIResponse(message string, code int, status string, data interface{}, dataSize_optional ...interface{}) Response {

	dataSize := 0
	if len(dataSize_optional) > 0 {
		dataSize = checkSize(dataSize_optional[0])
	} else {
		dataSize = checkSize(data)
	}

	jsonResponse := Response{
		Code:    code,
		Status:  status,
		Message: message,
		Size:    0,
		Data:    data,
	}

	if data != nil {
		jsonResponse = Response{
			Code:    code,
			Status:  status,
			Message: message,
			Size:    dataSize,
			Data:    data,
		}
	}

	return jsonResponse
}

func APIResponseWithoutData(message string, code int, status string) ResponseWithoutData {

	jsonResponse := ResponseWithoutData{
		Code:    code,
		Status:  status,
		Message: message,
	}

	return jsonResponse
}

func checkSize(data interface{}) int {
	switch reflect.TypeOf(data).Kind() {
	case reflect.Slice:
		s := reflect.ValueOf(data)
		return s.Len()
	default:
		return 1
	}
}
