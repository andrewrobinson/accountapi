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

func (c *AccountRestClient) createInternal(data model.AccountDataForCreate) ([]byte, *int, error) {

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

//this returns the body but Delete discards it
func (c *AccountRestClient) deleteInternal(id uuid.UUID, version int64) ([]byte, *int, error) {

	deleteUrl := fmt.Sprintf(c.deleteUrlFormatString, id, version)

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
