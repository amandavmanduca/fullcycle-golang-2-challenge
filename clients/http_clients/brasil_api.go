package http_clients

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"

	"github.com/amandavmanduca/fullcycle-golang-2-challenge/interfaces"
	"github.com/amandavmanduca/fullcycle-golang-2-challenge/internal/structs"
)

type brasilApiClient struct {
	httpClient interfaces.HttpClientInterface
}

func NewBrasilApiClient(httpClient interfaces.HttpClientInterface) interfaces.BrasilApiInterface {
	return &brasilApiClient{
		httpClient: httpClient,
	}
}

func (c *brasilApiClient) GetAddress(ctx context.Context, cep string) (*structs.AddressResponse, error) {
	resp, err := c.httpClient.Get(ctx, fmt.Sprintf("/%s", cep))
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	res, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var response structs.BrasilAddressResponse
	err = json.Unmarshal(res, &response)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode == 200 {
		return response.ToAddressResponse(), nil
	}
	return nil, errors.New("ERROR_GETTING_ADDRESS")
}
