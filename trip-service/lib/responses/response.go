package responses

import (
	"encoding/json"
	"net/http"
)

type Message struct {
	Error string
}

func (u Message) MarshalBinary() []byte {
	j, _ := json.Marshal(u)
	return j
}

func SendResponse(w http.ResponseWriter, httpStatus int, jsonResponse []byte) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(httpStatus)
	_, _ = w.Write(jsonResponse)
}

func SendInternalServerError(w http.ResponseWriter) {
	marshal, _ := json.Marshal(Message{
		Error: "Произошла внутрення ошибка сервиса",
	})

	SendResponse(w, 500, marshal)
}
