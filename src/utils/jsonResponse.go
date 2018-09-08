package utils

import "encoding/json"

type JsonResponse struct {
	Success bool    `json:"success"`
	Data    []int64 `json:"data"`
}

func GetErrorResponse() []byte {
	var response = JsonResponse{Success: false, Data: []int64{}}
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		return []byte(``)
	}

	return jsonResponse
}

func GetSuccessResponse(data []int64) []byte {
	var response = JsonResponse{Success: true, Data: data}
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		return []byte(``)
	}

	return jsonResponse
}
