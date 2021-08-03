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

func (c *AccountRestClient) DeleteAccount(id string) ([]byte, *int, error) {

	urlFmt := c.endpoint + "/%s?version=0"

	url := fmt.Sprintf(urlFmt, id)

	// fmt.Printf("DeleteAccount passed id:%s, gives urlFmt: %s, url:%s\n", id, urlFmt, url)

	resp, err := c.doDelete(url)
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

func (c *AccountRestClient) GetAccount(id string) ([]byte, *int, error) {

	urlFmt := c.endpoint + "/%s"

	url := fmt.Sprintf(urlFmt, id)

	// fmt.Printf("GetAccount passed id:%s, gives urlFmt: %s, url:%s\n", id, urlFmt, url)

	resp, err := c.doGet(url)
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

func (c *AccountRestClient) CreateAccount(data Data) ([]byte, *int, error) {

	json, err := json.Marshal(data)

	fmt.Printf("json:%s\n\n", json)

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

func (c *AccountRestClient) GetAccounts() ([]byte, *int, error) {

	resp, err := c.doGet(c.endpoint)
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

func (c *AccountRestClient) doGet(url string) (*http.Response, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	//TODO - I got these from the postman collection? Seems to work
	req.Header.Set("Accept", "application/vnd.api+json")
	req.Header.Set("Date", "{request_date}")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *AccountRestClient) doDelete(url string) (*http.Response, error) {
	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		return nil, err
	}

	//TODO - I got these from the postman collection? Seems to work
	req.Header.Set("Accept", "application/vnd.api+json")
	req.Header.Set("Date", "{request_date}")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *AccountRestClient) doPost(url string, json []byte) (*http.Response, error) {
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(json))
	if err != nil {
		return nil, err
	}

	//TODO - I got these from the postman collection? Seems to work
	req.Header.Set("Accept", "application/vnd.api+json")
	req.Header.Set("Date", "{request_date}")

	//TODO - play with  this
	//https://medium.com/orijtech-developers/taming-net-http-b946edfda562
	// blob, err := httputil.DumpRequestOut(req, true)
	// if err != nil {
	// 	panic(err)
	// }
	// println("blob:" + string(blob))

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
