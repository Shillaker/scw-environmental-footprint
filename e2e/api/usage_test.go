package usage

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"testing"

	"github.com/shillaker/scw-environmental-footprint/e2e"
	"github.com/shillaker/scw-environmental-footprint/util"
	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func makeJsonApiRequest(t *testing.T, relativeUrl string, jsonStr []byte) string {
	fullUrl := fmt.Sprintf("%v/%v", e2e.TestGatewayUrl, relativeUrl)

	log.Infof("Making request to %v", fullUrl)

	req, err := http.NewRequest("POST", fullUrl, bytes.NewBuffer(jsonStr))
	require.NoError(t, err)

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	if resp.Status != "200 OK" {
		assert.Fail(t, fmt.Sprintf("Failed JSON request, status: %v", resp.Status))
	}

	body, _ := io.ReadAll(resp.Body)
	return string(body)
}

func TestE2EUsageQuery(t *testing.T) {
	// One year of usage
	queryJson := []byte(`
{
  "instance": {
	"type": "play2-pico"
  },

  "usage": {
    "count": 2,
	"loadPercentage": 25,
	"timeSeconds": 31536000,
	"region": "fra"
  }
}`)

	actualResponseJson := makeJsonApiRequest(t, "v1/impact/instance", queryJson)

	fmt.Println(actualResponseJson)

	assert.NotNil(t, actualResponseJson)

	expected := `
{
	"impacts": {
		"adp": {
			"manufacture":0.0009453125,"use":3.59375e-7,"unit":"kgSbeq"
		},
		"gwp": {
			"manufacture":4.765625,"use":0.71875,"unit":"kgCO2eq"
		},
		"pe": {
			"manufacture":64.84375,"use":82.8125,"unit":"MJ"
		}
	},
	"equivalentsManufacture": [
		{"amount":0.005295139,"thing":"flights from London to New York"},
		{"amount":0.0390625,"thing":"flights from London to Paris"},
		{"amount":34.04018,"thing":"kms driven in a petrol car"},
		{"amount":5.295139,"thing":"kgs of cement manufactured"},
		{"amount":366.58652,"thing":"grams of beef eaten"},
		{"amount":794.2708,"thing":"litres of water boiled in a kettle"}
	],
	"equivalentsUse": [
	    {"amount":0.0007986111,"thing":"flights from London to New York"},
		{"amount":0.0058913934,"thing":"flights from London to Paris"},
		{"amount":5.133929,"thing":"kms driven in a petrol car"},
		{"amount":0.7986111,"thing":"kgs of cement manufactured"},
		{"amount":55.28846,"thing":"grams of beef eaten"},
		{"amount":119.791664,"thing":"litres of water boiled in a kettle"}
	]
}
`

	assert.Equal(t, util.CompressJsonString(expected), util.CompressJsonString(actualResponseJson))
}
