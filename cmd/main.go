package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"

	"github.com/andrewrobinson/accountapi/pkg/client"
)

func main() {

	endpointFlag := flag.String("endpoint", "http://localhost:8080/v1/organisation/accounts", "")

	flag.Parse()
	endpoint := *endpointFlag

	fmt.Printf("main running against endpoint:%s\n", endpoint)

	c := client.NewAccountRestClient(endpoint, "")

	get(c)
	create(c)
	get(c)
	delete(c)
	// getAll()
}

func get(c *client.AccountRestClient) {

	_, statusCode, err := c.GetAccount("ad27e265-9605-4b4b-a0e5-3003ea9cc4dc")

	if err != nil {
		fmt.Printf("%+v", err)
		os.Exit(1)
	}

	success := *statusCode == http.StatusOK

	fmt.Printf("GetAccount statusCode: %d, success: %v\n", *statusCode, success)

	// fmt.Printf("GetAccount response: %d, %s\n\n", *statusCode, string(body))

	// if !success {
	// 	fmt.Printf("response from GetAccount HTTP request not 200 : %d, body: %s", *statusCode, string(body))
	// }

}

func delete(c *client.AccountRestClient) {

	_, statusCode, err := c.DeleteAccount("ad27e265-9605-4b4b-a0e5-3003ea9cc4dc")

	if err != nil {
		fmt.Printf("%+v", err)
		os.Exit(1)
	}

	success := *statusCode == http.StatusNoContent

	fmt.Printf("DeleteAccount statusCode: %d, success: %v\n", *statusCode, success)

	// fmt.Printf("DeleteAccount response: %d, %s\n\n", *statusCode, string(body))

	// if !success {
	// 	fmt.Printf("response from DeleteAccount HTTP request not 204 : %d, body: %s", *statusCode, string(body))
	// }

}

func create(c *client.AccountRestClient) {

	country := "GB"
	accountClassification := "Personal"
	// version := int64(0)

	att := client.AccountAttributes{Name: []string{"Samantha Holder"},
		Country: &country, BaseCurrency: "GBP", BankID: "400302", BankIDCode: "GBDSC",
		AccountNumber: "10000004", CustomerID: "234", Iban: "GB28NWBK40030212764204", Bic: "NWBKGB42", AccountClassification: &accountClassification,
	}

	m := client.AccountData{ID: "ad27e265-9605-4b4b-a0e5-3003ea9cc4dc",
		OrganisationID: "eb0bd6f5-c3f5-44b2-b677-acd23cdde73c",
		Type:           "accounts", Attributes: &att}

	data := client.Data{Data: &m}

	// fmt.Printf("model data: %+v", data)

	_, statusCode, err := c.CreateAccount(data)

	if err != nil {
		fmt.Printf("%+v", err)
		os.Exit(1)
	}

	success := *statusCode == http.StatusCreated

	fmt.Printf("CreateAccount statusCode: %d, success: %v\n", *statusCode, success)

	// fmt.Printf("CreateAccount response for ID: %s, %d, %s\n\n", m.ID, *statusCode, string(body))

	// if !success {
	// 	fmt.Printf("response from CreateAccount HTTP request not 201 : %d, body: %s", *statusCode, string(body))
	// }

}

// func getAll(c *client.AccountRestClient) {

// 	body, statusCode, err := c.GetAccounts()

// 	if err != nil {
// 		fmt.Printf("%+v", err)
// 		os.Exit(1)
// 	}

// 	success := *statusCode == http.StatusOK

// 	fmt.Printf("GetAccounts response: %d, %s\n", *statusCode, string(body))

// 	if !success {
// 		fmt.Printf("response from GetAccounts HTTP request not 200 : %d, body: %s", *statusCode, string(body))
// 	}

// }
