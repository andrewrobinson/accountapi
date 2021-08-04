package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"

	"github.com/andrewrobinson/accountapi/pkg/client"
	"github.com/andrewrobinson/accountapi/pkg/client/model"
	uuid "github.com/satori/go.uuid"
)

func main() {

	endpointFlag := flag.String("endpoint", "http://localhost:8080/v1/organisation/accounts", "")
	flag.Parse()

	fmt.Printf("main running against endpoint:%s\n", *endpointFlag)

	accountRestClient := client.NewAccountRestClient(*endpointFlag)

	id := uuid.FromStringOrNil("ad27e265-9605-4b4b-a0e5-3003ea9cc4dc")

	fetch(accountRestClient, id)
	create(accountRestClient, id)
	fetch(accountRestClient, id)
	delete(accountRestClient, id)

}

func fetch(accountRestClient *client.AccountRestClient, id uuid.UUID) {

	_, statusCode, err := accountRestClient.Fetch(id)

	if err != nil {
		fmt.Printf("%+v", err)
		os.Exit(1)
	}

	success := *statusCode == http.StatusOK

	fmt.Printf("GetAccount statusCode: %d, success: %v\n", *statusCode, success)
	// fmt.Printf("GetAccount response: %d, %s\n\n", *statusCode, string(body))

}

func buildAccountData(id uuid.UUID) model.AccountData {

	country := "GB"
	accountClassification := "Personal"

	att := model.AccountAttributes{Name: []string{"Samantha Holder"},
		Country: &country, BaseCurrency: "GBP", BankID: "400302", BankIDCode: "GBDSC",
		AccountNumber: "10000004", CustomerID: "234", Iban: "GB28NWBK40030212764204", Bic: "NWBKGB42", AccountClassification: &accountClassification,
	}

	m := model.Account{ID: id,
		OrganisationID: "eb0bd6f5-c3f5-44b2-b677-acd23cdde73c",
		Type:           "accounts", Attributes: &att}

	return model.AccountData{Data: &m}
}

func create(accountRestClient *client.AccountRestClient, id uuid.UUID) {

	data := buildAccountData(id)

	_, statusCode, err := accountRestClient.Create(data)

	if err != nil {
		fmt.Printf("%+v", err)
		os.Exit(1)
	}

	success := *statusCode == http.StatusCreated

	fmt.Printf("CreateAccount statusCode: %d, success: %v\n", *statusCode, success)

}

func delete(accountRestClient *client.AccountRestClient, id uuid.UUID) {

	_, statusCode, err := accountRestClient.Delete(id, 0)

	if err != nil {
		fmt.Printf("%+v", err)
		os.Exit(1)
	}

	success := *statusCode == http.StatusNoContent

	fmt.Printf("DeleteAccount statusCode: %d, success: %v\n", *statusCode, success)

}
