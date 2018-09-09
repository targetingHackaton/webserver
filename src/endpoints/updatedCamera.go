package endpoints

import (
	"net/http"
	"../storage"
	"../utils"
	"encoding/json"
)

type UpdatedCamera struct {
	Endpoint   string
	Storage	   *storage.Storage
}

func (ch UpdatedCamera) ServeHTTP(writer http.ResponseWriter, req *http.Request) {
	queryValues := req.URL.Query()
	showroomId := utils.StrToInt(queryValues.Get("showroomId"))
	cameraId := utils.StrToInt(queryValues.Get("cameraId"))

	hasUpdate := ch.Storage.HasUpdatedCamera(showroomId, cameraId)
	message, err := json.Marshal(hasUpdate)

	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		writer.Write(utils.GetErrorResponse())
		return
	}

	writer.WriteHeader(http.StatusOK)
	writer.Write(message)
}

func (ch UpdatedCamera) GetEndpoint() string {
	return ch.Endpoint
}

