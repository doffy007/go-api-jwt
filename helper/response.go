package helper

import "strings"

//Response is used for static shape json return
type Response struct {
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Errors  interface{} `json:"errors"` //type data interface digunakan untuk menampung banyak data,
	Data    interface{} `json:"data"`   //semisal terjadi 2 atau lebih field error, maka interface akan menampung 2 atau lebih error tersebut
}

//EmptyObj Object is used when data doesnt want to be null on json
type EmptyObj struct{} 

//BuildResponse method is to inject data value to dynamic success response
func BuildResponse(status bool, message string, errors interface{}, data interface{}) Response {
	r := Response{
		Status:  status,
		Message: message,
		Errors:  nil, 
		Data:    data,
	}

	return r
}

//BuildErrorResponse method is to inject data value to dynamic failed response
func BuildErrorResponse(message string, err string, data interface{}) Response {
	splitedError := strings.Split(err, "\n") //split digunakan untuk memecah menjadi beberapa
	r := Response{
		Status:  false,
		Message: message,
		Errors:  splitedError,
		Data:    data,
	}

	return r
}
