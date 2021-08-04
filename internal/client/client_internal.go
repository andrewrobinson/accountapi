package client

//this way works from here
// func fetchInternal(id uuid.UUID, c *AccountRestClient) ([]byte, *int, error) {

// 	url := fmt.Sprintf(c.getUrlFormatString, id)

// 	resp, err := c.doGet(url)
// 	if err != nil {
// 		return nil, nil, err
// 	}
// 	defer resp.Body.Close()

// 	body, err := ioutil.ReadAll(resp.Body)
// 	if err != nil {
// 		return nil, nil, err
// 	}

// 	return body, &resp.StatusCode, nil

// }

/*

What works inside pkg/client/client_internal.go doesn't work from here
The IDE is already not happy with the c *AccountRestClient receiver over there, but command line is fine
If I dropped the receiver I could probably get this working here?

*/

// admin@Admins-MacBook-Air-2 interview-accountapi % go run cmd/main.go
// # github.com/andrewrobinson/accountapi/pkg/client
// pkg/client/client.go:56:28: c.fetchInternal undefined (type *AccountRestClient has no field or method fetchInternal)
// admin@Admins-MacBook-Air-2 interview-accountapi %

//TODO - how to hide this? no receiver? Different struct? Different package? In int or in pkg?
// func (c *AccountRestClient) fetchInternal(id uuid.UUID) ([]byte, *int, error) {

// 	url := fmt.Sprintf(c.getUrlFormatString, id)

// 	resp, err := c.doGet(url)
// 	if err != nil {
// 		return nil, nil, err
// 	}
// 	defer resp.Body.Close()

// 	body, err := ioutil.ReadAll(resp.Body)
// 	if err != nil {
// 		return nil, nil, err
// 	}

// 	return body, &resp.StatusCode, nil

// }
