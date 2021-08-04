package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/andrewrobinson/accountapi/client"
)

//package names clashing with module name ... ?

func main() {
	fmt.Println("hello world, sleeping for 5")

	time.Sleep(5 * time.Second)

	get()
	create()
	get()
	delete()
	// getAll()
}

func get() {

	//use this when running locally from go run
	// endpoint := "http://localhost:8080/v1/organisation/accounts"

	//use this one when running from docker-compose/script/run-tests.sh
	endpoint := "http://accountapi:8080/v1/organisation/accounts"

	c := client.NewAccountRestClient(endpoint, "")

	body, statusCode, err := c.GetAccount("ad27e265-9605-4b4b-a0e5-3003ea9cc4dc")

	if err != nil {
		fmt.Printf("%+v", err)
		os.Exit(1)
	}

	success := *statusCode == http.StatusOK

	fmt.Printf("GetAccount response: %d, %s\n\n", *statusCode, string(body))

	if !success {
		fmt.Printf("response from GetAccount HTTP request not 200 : %d, body: %s", *statusCode, string(body))
	}

}

func delete() {

	//may need version, this is hardcoded deeper currently
	// endpoint := "http://localhost:8080/v1/organisation/accounts"
	endpoint := "http://accountapi:8080/v1/organisation/accounts"

	c := client.NewAccountRestClient(endpoint, "")

	body, statusCode, err := c.DeleteAccount("ad27e265-9605-4b4b-a0e5-3003ea9cc4dc")

	if err != nil {
		fmt.Printf("%+v", err)
		os.Exit(1)
	}

	success := *statusCode == http.StatusNoContent

	fmt.Printf("DeleteAccount response: %d, %s\n\n", *statusCode, string(body))

	if !success {
		fmt.Printf("response from DeleteAccount HTTP request not 204 : %d, body: %s", *statusCode, string(body))
	}

}

func create() {

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

	// endpoint := "http://localhost:8080/v1/organisation/accounts"
	endpoint := "http://accountapi:8080/v1/organisation/accounts"

	c := client.NewAccountRestClient(endpoint, "")

	body, statusCode, err := c.CreateAccount(data)

	if err != nil {
		fmt.Printf("%+v", err)
		os.Exit(1)
	}

	success := *statusCode == http.StatusCreated

	fmt.Printf("CreateAccount response for ID: %s, %d, %s\n\n", m.ID, *statusCode, string(body))

	if !success {
		fmt.Printf("response from CreateAccount HTTP request not 201 : %d, body: %s", *statusCode, string(body))
	}

}

func getAll() {

	// endpoint := "http://localhost:8080/v1/organisation/accounts"
	endpoint := "http://accountapi:8080/v1/organisation/accounts"

	c := client.NewAccountRestClient(endpoint, "")

	body, statusCode, err := c.GetAccounts()

	if err != nil {
		fmt.Printf("%+v", err)
		os.Exit(1)
	}

	success := *statusCode == http.StatusOK

	fmt.Printf("GetAccounts response: %d, %s\n", *statusCode, string(body))

	if !success {
		fmt.Printf("response from GetAccounts HTTP request not 200 : %d, body: %s", *statusCode, string(body))
	}

}
