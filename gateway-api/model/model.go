package model

import (
	"encoding/json"
)

type TripResponse struct {
	Id int `json:"id"`
}

func (i *TripResponse) Marshal() []byte {
	jsonModel, _ := json.Marshal(i)
	return jsonModel
}

type Offer struct {
	PNR       string  `db:"pnr" json:"pnr"`
	Arrival   string  `db:"arrival" json:"arrival"`
	Departure string  `db:"departure" json:"departure"`
	Price     float32 `db:"price" json:"price"`
	TripId    int     `db:"tripId" json:"tripId"`
}

type OfferPrice struct {
	TripId int32   `ch:"tripId" json:"tripId"`
	Price  float64 `ch:"price" json:"price"`
}

func (op *OfferPrice) Marshal() []byte {
	jsonModel, _ := json.Marshal(op)
	return jsonModel
}

func (u Offer) MarshalBinary() ([]byte, error) {
	return json.Marshal(u)
}

func (u Offer) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, &u)
}
