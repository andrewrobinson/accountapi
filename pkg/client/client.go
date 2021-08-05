package client

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/andrewrobinson/accountapi/pkg/client/model"
	uuid "github.com/satori/go.uuid"
)

type AccountClient interface {
	Fetch(id uuid.UUID) (model.FetchedAccountData, error)
	Create(data model.AccountDataForCreate) (model.AccountDataForCreate, error)
	Delete(id uuid.UUID, version int64) error
}

type accountRestClient struct {
	endpoint              string
	getUrlFormatString    string
	deleteUrlFormatString string
	httpClient            *http.Client
}

// returns the interface, could be used this way by the end user
// https://www.sohamkamani.com/golang/2018-06-20-golang-factory-patterns/
func NewAccountClient(endpoint string, httpClient *http.Client) AccountClient {
	getUrlFormatString := endpoint + "/%s"
	deleteUrlFormatString := endpoint + "/%s?version=%d"
	return &accountRestClient{endpoint, getUrlFormatString, deleteUrlFormatString, httpClient}

}

//returns the struct. useful at test time to get hold of the deleteInternal method
func NewAccountRestClient(endpoint string, httpClient *http.Client) *accountRestClient {
	getUrlFormatString := endpoint + "/%s"
	deleteUrlFormatString := endpoint + "/%s?version=%d"
	return &accountRestClient{endpoint, getUrlFormatString, deleteUrlFormatString, httpClient}
}

func (c *accountRestClient) Fetch(id uuid.UUID) (model.FetchedAccountData, error) {
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

func (c *accountRestClient) Create(data model.AccountDataForCreate) (model.AccountDataForCreate, error) {
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

func (c *accountRestClient) Delete(id uuid.UUID, version int64) error {

	_, statusCode, err := c.deleteInternal(id, version)

	if err != nil {
		return err
	}

	//Checking for 204 sort of makes sense but is debatable
	success := *statusCode == http.StatusNoContent

	if success {
		return nil
	} else {
		return fmt.Errorf("Delete statusCode:%d instead of 204", *statusCode)
	}

}
