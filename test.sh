# create:

curl -v --location --request POST 'http://localhost:8080/v1/organisation/accounts' \
--header 'Content-Type: application/json' \
--header 'Date: {{request_date}}' \
--data-raw '{
  "data": {
    "id": "ad27e265-9605-4b4b-a0e5-3003ea9cc4dc",
    "organisation_id": "eb0bd6f5-c3f5-44b2-b677-acd23cdde73c",
    "type": "accounts",
    "attributes": {
        "name": [
        "Samantha Holder"
      ],
        "country": "GB",
        "base_currency": "GBP",
        "bank_id": "400302",
        "bank_id_code": "GBDSC",
        "account_number": "10000004",
        "customer_id": "234",
        "iban": "GB28NWBK40030212764204",
        "bic": "NWBKGB42",
        "account_classification": "Personal"
    }
  }
}'

# fetch all:

curl -v --location --request GET 'http://localhost:8080/v1/organisation/accounts' \
--header 'Accept: application/vnd.api+json' \
--header 'Date: {{request_date}}'

# fetch one:

curl -v --location -g --request GET 'http://localhost:8080/v1/organisation/accounts/ad27e265-9605-4b4b-a0e5-3003ea9cc4dc' \
--header 'Accept: application/vnd.api+json' \
--header 'Date: {{request_date}}'

# delete one

curl -v --location -g --request DELETE 'http://localhost:8080/v1/organisation/accounts/ad27e265-9605-4b4b-a0e5-3003ea9cc4dc?version=0' \
--header 'Authorization: {{authorization}}' \
--header 'Date: {{request_date}}' \
--data-raw ''

# fetch one: (shouldn't be there any more)

curl -v --location -g --request GET 'http://localhost:8080/v1/organisation/accounts/ad27e265-9605-4b4b-a0e5-3003ea9cc4dc' \
--header 'Accept: application/vnd.api+json' \
--header 'Date: {{request_date}}'