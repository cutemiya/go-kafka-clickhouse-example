package queryService

import (
	"fmt"
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

func (i *Service) GetAll() ([]byte, error) {
	response, err := i.client.Get(fmt.Sprintf("%s/trips", i.baseUrl))
	if err != nil {
		return nil, err
	}

	return io.ReadAll(response.Body)
}

func (i *Service) GetById(id int) ([]byte, error) {
	response, err := i.client.Get(fmt.Sprintf("%s/trips/%d", i.baseUrl, id))
	if err != nil {
		return nil, err
	}

	return io.ReadAll(response.Body)
}
