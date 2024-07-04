package usage

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"testing"

	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"github.com/shillaker/scw-environmental-footprint/e2e"
	"github.com/shillaker/scw-environmental-footprint/util"
)

func makeJsonApiRequest(t *testing.T, relativeUrl string, jsonStr []byte) string {
	fullUrl := fmt.Sprintf("%v/%v", e2e.TestGatewayUrl, relativeUrl)

	log.Infof("Making request to %v", fullUrl)

	req, err := http.NewRequest("POST", fullUrl, bytes.NewBuffer(jsonStr))
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
	assert.NotNil(t, actualResponseJson)

	expected := `
{
  "impacts": {
    "adp": {
      "manufacture": 0.00125,
      "use": 0.00000434375,
      "unit": "kgSbeq"
    },
    "gwp": {
      "manufacture": 6.25,
      "use": 8.75,
      "unit": "kgCO2eq"
    },
    "pe": {
      "manufacture": 125,
      "use": 1008.75,
      "unit": "MJ"
    }
  },
  "equivalentsManufacture": [
    {
      "amount": 0.0069444445,
      "thing": "flights from London to New York"
    },
    {
      "amount": 0.051229507,
      "thing": "flights from London to Paris"
    },
    {
      "amount": 44.642857,
      "thing": "kms driven in a petrol car"
    },
    {
      "amount": 6.9444447,
      "thing": "kgs of cement manufactured"
    },
    {
      "amount": 480.76923,
      "thing": "grams of beef eaten"
    },
    {
      "amount": 1041.6666,
      "thing": "litres of water boiled in a kettle"
    }
  ],
  "equivalentsUse": [
    {
      "amount": 0.009722223,
      "thing": "flights from London to New York"
    },
    {
      "amount": 0.07172131,
      "thing": "flights from London to Paris"
    },
    {
      "amount": 62.5,
      "thing": "kms driven in a petrol car"
    },
    {
      "amount": 9.722222,
      "thing": "kgs of cement manufactured"
    },
    {
      "amount": 673.0769,
      "thing": "grams of beef eaten"
    },
    {
      "amount": 1458.3334,
      "thing": "litres of water boiled in a kettle"
    }
  ]
}
`
	assert.Equal(t, util.CompressJsonString(expected), util.CompressJsonString(actualResponseJson))
}
