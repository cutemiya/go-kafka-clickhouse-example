package model

type OfferPrice struct {
	TripId int32   `ch:"tripId" json:"tripId"`
	Price  float64 `ch:"price" json:"price"`
}
