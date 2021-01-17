package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateXdbPriceRequest(t *testing.T) {
	req, err := createXdbPriceRequest()
	assert.NoError(t, err)
	assert.Equal(t, "GET", req.Method)
	assert.Equal(t, stelExURL, req.URL.String())
}

func TestParseDigitalBitsExpertResponse(t *testing.T) {
	body := "hello"
	gotPrice, gotErr := parseDigitalBitsExpertLatestPrice(body)
	assert.EqualError(t, gotErr, "mis-formed response from digitalbits expert")

	body = "hello,"
	gotPrice, gotErr = parseDigitalBitsExpertLatestPrice(body)
	assert.EqualError(t, gotErr, "mis-formed price from digitalbits expert")

	body = "[[10001,hello]"
	gotPrice, gotErr = parseDigitalBitsExpertLatestPrice(body)
	assert.Error(t, gotErr)

	body = "[[100001,5.00],[100002,6.00]]"
	wantPrice := 5.00
	gotPrice, gotErr = parseDigitalBitsExpertLatestPrice(body)
	assert.NoError(t, gotErr)
	assert.Equal(t, wantPrice, gotPrice)
}
