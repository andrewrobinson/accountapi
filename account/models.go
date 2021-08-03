package account

type Data struct {
	Data *AccountData `json:"data,omitempty"`
}

// Account represents an account in the form3 org section.
// See https://api-docs.form3.tech/api.html#organisation-accounts for
// more information about fields.
type AccountData struct {
	Attributes     *AccountAttributes `json:"attributes,omitempty"`
	ID             string             `json:"id,omitempty"`
	OrganisationID string             `json:"organisation_id,omitempty"`
	Type           string             `json:"type,omitempty"`
	Version        *int64             `json:"version,omitempty"`
}

// "account_number": "10000004",
// //         "customer_id": "234",

type AccountAttributes struct {
	AccountClassification   *string  `json:"account_classification,omitempty"`
	AccountMatchingOptOut   *bool    `json:"account_matching_opt_out,omitempty"`
	AccountNumber           string   `json:"account_number,omitempty"`
	CustomerID              string   `json:"customer_id,omitempty"`
	AlternativeNames        []string `json:"alternative_names,omitempty"`
	BankID                  string   `json:"bank_id,omitempty"`
	BankIDCode              string   `json:"bank_id_code,omitempty"`
	BaseCurrency            string   `json:"base_currency,omitempty"`
	Bic                     string   `json:"bic,omitempty"`
	Country                 *string  `json:"country,omitempty"`
	Iban                    string   `json:"iban,omitempty"`
	JointAccount            *bool    `json:"joint_account,omitempty"`
	Name                    []string `json:"name,omitempty"`
	SecondaryIdentification string   `json:"secondary_identification,omitempty"`
	Status                  *string  `json:"status,omitempty"`
	Switched                *bool    `json:"switched,omitempty"`
}

// curl --location --request POST 'http://localhost:8080/v1/organisation/accounts' \
// --header 'Content-Type: application/json' \
// --header 'Date: {{request_date}}' \
// --data-raw '{
//   "data": {
//     "id": "ad27e265-9605-4b4b-a0e5-3003ea9cc4dc",
//     "organisation_id": "eb0bd6f5-c3f5-44b2-b677-acd23cdde73c",
//     "type": "accounts",
//     "attributes": {
// "name": [
//     "Samantha Holder"
//   ],
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
