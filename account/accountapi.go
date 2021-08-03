package account

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
	// log "github.com/sirupsen/logrus"
)

type AccountRestClient struct {
	endpoint   string
	apiKey     string
	httpClient *http.Client
}

func initHTTPClient() *http.Client {
	return &http.Client{
		Timeout: time.Second * 30,
		Transport: &http.Transport{
			MaxIdleConns:        5,
			MaxIdleConnsPerHost: 1,
		},
	}
}

func NewAccountRestClient(endpoint string, apiKey string) *AccountRestClient {

	httpClient := initHTTPClient()

	return &AccountRestClient{endpoint, apiKey, httpClient}
}

// serviceURL := fmt.Sprintf("https://%s:%s/api/v1", os.Getenv("API_CLIENT_INTERACTION_SERVICE_HOST"), os.Getenv("API_CLIENT_INTERACTION_SERVICE_PORT"))
// clientInteractionClient := gateway.NewRestClientInteractionClient(serviceURL, httpClient)

func (c *AccountRestClient) CreateAccount(data Data) ([]byte, *int, error) {
	// url := fmt.Sprintf("https://%s.eyerys.co.za/api/services/data_integrator/v1.0/datasources/allan_gray", c.account)

	json, err := json.Marshal(data)

	fmt.Printf("json:%s", json)

	if err != nil {
		return nil, nil, err
	}

	// log.Debugf("trigger JSON payload: %s", string(json))

	resp, err := c.doPost(c.endpoint, json)
	if err != nil {
		return nil, nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, nil, err
	}

	return body, &resp.StatusCode, nil
}

func (c *AccountRestClient) doPost(url string, json []byte) (*http.Response, error) {
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(json))
	if err != nil {
		return nil, err
	}
	req.Header.Set("apiKey", c.apiKey)
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
