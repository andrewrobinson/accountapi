package client

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/andrewrobinson/accountapi/pkg/client/model"
	uuid "github.com/satori/go.uuid"
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

func (c *AccountRestClient) Fetch(id uuid.UUID) (model.FetchedAccountData, error) {
	var ret model.FetchedAccountData

	body, statusCode, err := c.fetchInternal(id)

	if err != nil {
		return ret, err
	}

	success := *statusCode == http.StatusOK

	fmt.Printf("fetchInternal response: %d, %s\n", *statusCode, string(body))

	// {
	// 	"data": {
	// 		"attributes": {
	// 			"account_classification": "Personal",
	// 			"account_number": "10000004",
	// 			"alternative_names": null,
	// 			"bank_id": "400302",
	// 			"bank_id_code": "GBDSC",
	// 			"base_currency": "GBP",
	// 			"bic": "NWBKGB42",
	// 			"country": "GB",
	// 			"iban": "GB28NWBK40030212764204",
	// 			"name": ["Samantha Holder"]
	// 		},
	// 		"created_on": "2021-08-04T21:20:48.639Z",
	// 		"id": "ad27e265-9605-4b4b-a0e5-3003ea9cc4dc",
	// 		"modified_on": "2021-08-04T21:20:48.639Z",
	// 		"organisation_id": "eb0bd6f5-c3f5-44b2-b677-acd23cdde73c",
	// 		"type": "accounts",
	// 		"version": 0
	// 	},
	// 	"links": {
	// 		"self": "/v1/organisation/accounts/ad27e265-9605-4b4b-a0e5-3003ea9cc4dc"
	// 	}
	// }

	if success {

		err := json.Unmarshal(body, &ret)

		if err != nil {
			return ret, err
		} else {
			// fmt.Printf("Fetch ret after unmarshall is: %+v\n", ret)
			return ret, nil
		}

	} else {
		return ret, errors.New("NOT FOUND")
	}

}

//the code in here is so similar to in Fetch ....
func (c *AccountRestClient) Create(data model.AccountDataForCreate) (model.AccountDataForCreate, error) {
	var ret model.AccountDataForCreate

	body, statusCode, err := c.createInternal(data)

	if err != nil {
		return ret, err
	}

	success := *statusCode == http.StatusCreated

	// fmt.Printf("createInternal response: %d, %s\n", *statusCode, string(body))

	if success {

		err := json.Unmarshal(body, &ret)

		if err != nil {
			return ret, err
		} else {
			// fmt.Printf("Create ret after unmarshall is: %+v\n", ret)
			return ret, nil
		}

	} else {
		return ret, errors.New(fmt.Sprintf("statusCode not 201:%d", *statusCode))
	}
}

func (c *AccountRestClient) Delete(id uuid.UUID, version int64) error {

	_, statusCode, err := c.deleteInternal(id, version)

	if err != nil {
		return err
	}

	success := *statusCode == http.StatusNoContent

	// fmt.Printf("deleteInternal response: %d, %s\n", *statusCode, string(body))

	if success {
		return nil
	} else {
		return errors.New(fmt.Sprintf("statusCode not 204:%d", *statusCode))
	}

}

//private methods

//TODO - prove these actually have the effect compared to inlining them
// a test assertion about headers?
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
