package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/andrewrobinson/accountapi/pkg/client/model"
)

type AccountRestClient struct {
	endpoint              string
	getUrlFormatString    string
	deleteUrlFormatString string
	httpClient            *http.Client
}

// const (
// 	ACCEPT_HEADER = "application/vnd.api+json"
// 	GAME_ID2 = 67890
// )

func initHTTPClient() *http.Client {
	return &http.Client{
		Timeout: time.Second * 30,
		Transport: &http.Transport{
			MaxIdleConns:        5,
			MaxIdleConnsPerHost: 1,
		},
	}
}

//TODO - a fake fn for a test test
func ReturnFive(in int) int {
	return 5
}

func NewAccountRestClient(endpoint string) *AccountRestClient {
	httpClient := initHTTPClient()

	//TODO - any better way of storing and not exposing these?
	getUrlFormatString := endpoint + "/%s"
	deleteUrlFormatString := endpoint + "/%s?version=%d"

	return &AccountRestClient{endpoint, getUrlFormatString, deleteUrlFormatString, httpClient}
}

// TODO - make it return a struct
func (c *AccountRestClient) Fetch(id string) ([]byte, *int, error) {

	url := fmt.Sprintf(c.getUrlFormatString, id)

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

//TODO - return the body, or marshall the returned out as a struct or return the input struct?
func (c *AccountRestClient) Create(data model.AccountData) ([]byte, *int, error) {

	json, err := json.Marshal(data)

	// fmt.Printf("CreateAccount body json:%s\n\n", json)

	if err != nil {
		return nil, nil, err
	}

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

//TODO - not sure if we actually need the body returned? return it?
func (c *AccountRestClient) Delete(id string, version int64) ([]byte, *int, error) {

	deleteUrl := fmt.Sprintf(c.deleteUrlFormatString, id, version)

	// fmt.Printf("DeleteAccount passed id:%s, gives deleteUrl:%s\n", id, deleteUrl)

	resp, err := c.doDelete(deleteUrl)
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

//private methods

//TODO - prove these actually have the effect compared to inlining them
func setCommonHeaders(req *http.Request) {
	req.Header.Set("Accept", "application/vnd.api+json")
	req.Header.Set("Date", "{request_date}")
}

//TODO - could move this out of the "class" and just pass in the c.httpClient?
func (c *AccountRestClient) doGet(url string) (*http.Response, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	//TODO - I got these from the postman collection? Seems to work
	setCommonHeaders(req)

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

	setCommonHeaders(req)

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
	setCommonHeaders(req)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
