package utils

import (
	"encoding/json"
	"net/http"
)

func WriteJson(w http.ResponseWriter,status int ,response any)error{
	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(response)
}

func ReadJson(r *http.Request,result any)error{ 
	decoder:=json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	return decoder.Decode(result)
}

func CreateResponse(message string ,success bool,err error,data any) map[string]any{
	result:=map[string]any{
		"message":message,
		"success":success,
		"error":err,
		"data":data,
	}
	return result
}

func WriteJsonSuccessResponse(w http.ResponseWriter,status int ,message string,data any)error {
	response:=map[string]any{
		"message":message,
		"success":true,
		"error":nil,
		"data":data,
	}
	return WriteJson(w,status,response)
}
func WriteJsonErrorResponse(w http.ResponseWriter,status int ,message string,err error)error {
	response:=map[string]any{
		"message":message,
		"success":false,
		"error":err,
		"data":nil,
	}
	return WriteJson(w,status,response)
}

