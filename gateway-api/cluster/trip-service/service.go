package tripService

import (
	"bytes"
	"encoding/json"
	"fmt"
	"gateway-api/model"
	"io"
	"net/http"
)

type Service struct {
	client  *http.Client
	baseUrl string
}

func NewService(client *http.Client, baseUrl string) *Service {
	return &Service{
		client:  client,
		baseUrl: baseUrl,
	}
}

func (i *Service) CreateTrip() (int, error) {
	r := bytes.NewReader([]byte(""))
	resp, err := i.client.Post(fmt.Sprintf("%s/trip", i.baseUrl), "application/json", r)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return 0, err
	}

	var userResponse model.TripResponse
	if err := json.Unmarshal(respBody, &userResponse); err != nil {
		return 0, err
	}

	return userResponse.Id, nil
}
