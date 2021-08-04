package client

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/andrewrobinson/accountapi/pkg/client/model"
	uuid "github.com/satori/go.uuid"
)

func (c *AccountRestClient) fetchInternal(id uuid.UUID) ([]byte, *int, error) {

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

func (c *AccountRestClient) createInternal(data model.AccountData) ([]byte, *int, error) {

	json, err := json.Marshal(data)

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
