package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/andrewrobinson/accountapi/pkg/client/model"
	uuid "github.com/satori/go.uuid"
)

func (c *accountRestClient) fetchInternal(id uuid.UUID) ([]byte, *int, error) {

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

func (c *accountRestClient) createInternal(data model.AccountDataForCreate) ([]byte, *int, error) {

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
func (c *accountRestClient) deleteInternal(id uuid.UUID, version int64) ([]byte, *int, error) {

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

//TODO - prove these actually have the effect compared to inlining them
// a test assertion about headers?
func setCommonHeaders(req *http.Request) {
	req.Header.Set("Accept", "application/vnd.api+json")
	req.Header.Set("Date", "{request_date}")
}

func (c *accountRestClient) doGet(url string) (*http.Response, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	setCommonHeaders(req)

	resp, err := c.Client.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *accountRestClient) doDelete(url string) (*http.Response, error) {
	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		return nil, err
	}

	setCommonHeaders(req)

	resp, err := c.Client.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *accountRestClient) doPost(url string, json []byte) (*http.Response, error) {
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(json))
	if err != nil {
		return nil, err
	}

	setCommonHeaders(req)

	resp, err := c.Client.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
