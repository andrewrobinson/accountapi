package main

import (
	"fmt"
	"net/http"

	"github.com/andrewrobinson/accountapi/account"
)

//package names clashing with module name ... ?

func main() {
	fmt.Println("hello world")
	get()
	getAll()
	create()
}

func get() {

}

func getAll() {

}

// {
// 	"attributes": {
// 		"account_classification": "Personal",
// 		"account_number": "10000004",
// 		"customer_id": "234",
// 		"bank_id": "400302",
// 		"bank_id_code": "GBDSC",
// 		"base_currency": "GBP",
// 		"bic": "NWBKGB42",
// 		"country": "GB",
// 		"iban": "GB28NWBK40030212764204",
// 		"name": ["Samantha Holder"]
// 	},
// 	"id": "ad27e265-9605-4b4b-a0e5-3003ea9cc4dc",
// 	"organisation_id": "eb0bd6f5-c3f5-44b2-b677-acd23cdde73c",
// 	"type": "accounts",
// 	"version": 1
// }

func create() {
	// 	curl -v --location --request POST 'http://localhost:8080/v1/organisation/accounts' \
	// --header 'Content-Type: application/json' \
	// --header 'Date: {{request_date}}' \
	// --data-raw '{
	//   "data": {
	//     "id": "ad27e265-9605-4b4b-a0e5-3003ea9cc4dc",
	//     "organisation_id": "eb0bd6f5-c3f5-44b2-b677-acd23cdde73c",
	//     "type": "accounts",
	//     "attributes": {
	//         "name": [
	//         "Samantha Holder"
	//       ],
	//         "country": "GB",
	//         "base_currency": "GBP",
	//         "bank_id": "400302",
	//         "bank_id_code": "GBDSC",
	//         "account_number": "10000004",
	//         "customer_id": "234",
	//         "iban": "GB28NWBK40030212764204",
	//         "bic": "NWBKGB42",
	//         "account_classification": "Personal"
	//     }
	//   }
	// }'

	country := "GB"
	accountClassification := "Personal"

	//not in the curl
	version := int64(1)

	//customer_id isn't in the model
	att := account.AccountAttributes{Name: []string{"Samantha Holder"},
		Country: &country, BaseCurrency: "GBP", BankID: "400302", BankIDCode: "GBDSC",
		AccountNumber: "10000004", CustomerID: "234", Iban: "GB28NWBK40030212764204", Bic: "NWBKGB42", AccountClassification: &accountClassification,
	}

	m := account.AccountData{ID: "ad27e265-9605-4b4b-a0e5-3003ea9cc4dc",
		OrganisationID: "eb0bd6f5-c3f5-44b2-b677-acd23cdde73c",
		Type:           "accounts", Attributes: &att, Version: &version}

	data := account.Data{Data: &m}

	fmt.Printf("model data: %+v", data)

	endpoint := "http://localhost:8080/v1/organisation/accounts"

	c := account.NewAccountRestClient(endpoint, "")

	body, statusCode, err := c.CreateAccount(data)

	if err != nil {
		fmt.Printf("%v", err)
		// return err
	}

	success := *statusCode == http.StatusOK

	fmt.Printf("CreateAccount response for ParentInteractionID: %s, %d, %s", m.ID, *statusCode, string(body))

	if !success {
		fmt.Printf("response from CreateAccount HTTP request not 200 : %d, body: %s", *statusCode, string(body))
	}

	fmt.Printf("%v", c)
}
