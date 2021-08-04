package client

import (
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

	// fmt.Printf("fetchInternal response: %d, %s\n", *statusCode, string(body))

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
		return ret, fmt.Errorf("statusCode not 201:%d", *statusCode)
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
		return fmt.Errorf("statusCode not 204:%d", *statusCode)
	}

}
