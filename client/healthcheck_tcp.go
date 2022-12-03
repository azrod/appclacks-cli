package client

import (
	"fmt"
	"net/http"

	apitypes "github.com/appclacks/go-types"
)

func (c *Client) CreateTCPHealthcheck(input apitypes.CreateTCPHealthcheckInput) (apitypes.Healthcheck, error) {
	var result apitypes.Healthcheck
	_, err := c.sendRequest("/api/v1/healthcheck/tcp", http.MethodPost, input, &result, nil, TokenAuth)
	if err != nil {
		return apitypes.Healthcheck{}, err
	}
	return result, nil
}

func (c *Client) UpdateTCPHealthcheck(input apitypes.UpdateTCPHealthcheckInput) (apitypes.Response, error) {
	var result apitypes.Response
	internalInput := internalUpdateHealthcheckInput{
		Name:        input.Name,
		Description: input.Description,
		Labels:      input.Labels,
		Interval:    input.Interval,
		Enabled:     input.Enabled,
	}
	payload, err := jsonMerge(internalInput, input.HealthcheckTCPDefinition)
	if err != nil {
		return result, err
	}
	_, err = c.sendRequest(fmt.Sprintf("/api/v1/healthcheck/tcp/%s", input.ID), http.MethodPut, payload, &result, nil, TokenAuth)
	if err != nil {
		return apitypes.Response{}, err
	}
	return result, nil
}