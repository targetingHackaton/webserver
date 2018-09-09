package endpoints

import (
	"net/http"
	"../storage"
	"../utils"
	"encoding/json"
)

type ShowroomCounter struct {
	Endpoint string
	Storage  *storage.Storage
}

type RelevantAgeAndGender struct {
	AgeInterval int    `json:"ageInterval"`
	Gender      string `json:"gender"`
}

func (ch ShowroomCounter) ServeHTTP(writer http.ResponseWriter, req *http.Request) {
	queryValues := req.URL.Query()
	showroomId := utils.StrToInt(queryValues.Get("showroomId"))

	ageInterval, gender := ch.Storage.GetRelevantAgeAndGender(showroomId)
	message, err := json.Marshal(RelevantAgeAndGender{AgeInterval: ageInterval, Gender: gender})

	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		writer.Write(utils.GetErrorResponse())
		return
	}

	writer.WriteHeader(http.StatusOK)
	writer.Write(message)
}

func (ch ShowroomCounter) GetEndpoint() string {
	return ch.Endpoint
}
