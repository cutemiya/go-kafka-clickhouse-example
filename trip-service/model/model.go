package model

import "encoding/json"

type Status int

const (
	BookedStatus    Status = 0
	DeletedStatus   Status = 1
	CancelledStatus Status = 2
)

type TripResponse struct {
	Id int `json:"id"`
}

func (i *TripResponse) Marshal() []byte {
	jsonModel, _ := json.Marshal(i)
	return jsonModel
}
