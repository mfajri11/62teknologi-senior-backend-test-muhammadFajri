# 62teknologi-senior-backend-test

This repository is intended as an answer to 62teknologi-senior-backend-test challenge.

## how to use

**NOTE**: import defined schema/run the query at `./IMPORTHIS.sql`


```shell
$ git clone https://github.com/mfajri11/62teknologi-senior-backend-test-muhammadFajri

$ cd 62teknologi-senior-backend-test-muhammadFajri
# running the program
$ make run
```
endpoints

- `[GET] /business_search/search` 
- `[POST] /business_search`
- `[PUT] /business_search/:id`
- `[DEL] /business_search/:id`

default host:port `localhost:8080`.

## example
### GET
request example
```sh
curl -X GET \
  'http://localhost:8080/business_search?location=los%20angeles&term=pizza&latitude=41.7873382568359&longitude=-123.051551818848' \
  --header 'Accept: */*' \
  --header 'User-Agent: Thunder Client (https://www.thunderclient.com)' \
  --header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiI2MnRla25vbG9naSIsInJvbGUiOjEsImlhdCI6MTY4MTU2MTg4MCwiZXhwIjoxODc5NTYwNjc5fQ.2SDZCXjGYNIPhJj5lJxFQQmg8O8TJ6eqrDW2GXlpYJk'
```

response
```json
{
  "businesses": [
    {
      "ID": "0d4bc1a2-dc4d-11ed-b037-14f6d817012e",
      "Categories": [
        {
          "title": "Pizza",
          "alias": "pizza"
        },
        {
          "title": "Food",
          "alias": "food"
        }
      ],
      "Latitude": 41.7873382568359,
      "Longitude": -123.051551818848,
      "Coordinates": {
        "longitude": -123.051551818848,
        "latitude": 41.7873382568359
      },
      "Location": {
        "address": "James street 68M",
        "district": "los angeles",
        "province": "California",
        "country_code": "",
        "zipcode": "22399",
        "DisplayAddress": [
          "James street 68M",
          "los angeles",
          "California",
          "22399"
        ]
      },
      "Name": "golden-boy-pizza-hamburg",
      "Phone": "+14159829738",
      "Price": 20,
      "PriceRange": "$$",
      "Rating": 903,
      "RatingCount": 0
    }
  ]
}
```

### POST
request example
```sh
curl -X POST \
  'http://localhost:8080/business_search' \
  --header 'Accept: */*' \
  --header 'User-Agent: Thunder Client (https://www.thunderclient.com)' \
  --header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiI2MnRla25vbG9naSIsInJvbGUiOjEsImlhdCI6MTY4MTU2MTg4MCwiZXhwIjoxODc5NTYwNjc5fQ.2SDZCXjGYNIPhJj5lJxFQQmg8O8TJ6eqrDW2GXlpYJk' \
  --header 'Content-Type: application/json' \
  --data-raw '{
  "name": "golden-boy-pizza-hamburg",
  "phone": "+14159829738",
  "price": 20,
  "open_at":"09:00:00",
  "categories": [
      {
        "alias": "pizza",
        "title":"Pizza"
      },
      {
        "alias":"food",
        "title":"Food"
      }
    ],
    "address": "James street 68M",
    "district": "los angeles",
    "province":"California",
    "Country_code": "US",
    "zip_code": "22399",
    "latitude": 41.7873382568359,
    "longitude": -123.051551818848,
    "rating": 4,
    "rating_count": 903
}'
```
reponse example
```json
{
  "id": "0d4bc1a2-dc4d-11ed-b037-14f6d817012e",
  "name": "golden-boy-pizza-hamburg",
  "phone": "+14159829738",
  "open_at": "09:00:00",
  "price": 20,
  "categories": [
    {
      "title": "Pizza",
      "alias": "pizza"
    },
    {
      "title": "Food",
      "alias": "food"
    }
  ],
  "address": "James street 68M",
  "district": "los angeles",
  "province": "California",
  "country_code": "US",
  "zip_code": "22399",
  "latitude": 41.7873382568359,
  "longitude": -123.051551818848,
  "rating": 4,
  "rating_count": 903,
  "display_address": [
    "James street 68M",
    "los angeles",
    "California",
    "22399",
    "US"
  ],
  "price_range": "$$"
}
```
### PUT
request example
```sh
curl -X PUT \
  'http://localhost:8080/business_search/04880121-dc4c-11ed-b037-14f6d817012e' \
  --header 'Accept: */*' \
  --header 'User-Agent: Thunder Client (https://www.thunderclient.com)' \
  --header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiI2MnRla25vbG9naSIsInJvbGUiOjEsImlhdCI6MTY4MTU2MTg4MCwiZXhwIjoxODc5NTYwNjc5fQ.2SDZCXjGYNIPhJj5lJxFQQmg8O8TJ6eqrDW2GXlpYJk' \
  --header 'Content-Type: application/json' \
  --data-raw '{
  "name": "golden-boy-pizza-hamburg updated"
}'
```

response example
```json
{
  "id": "04880121-dc4c-11ed-b037-14f6d817012e",
  "name": "golden-boy-pizza-hamburg updated",
  "phone": "",
  "open_at": "",
  "price": 0,
  "categories": null,
  "address": "",
  "district": "",
  "province": "",
  "country_code": "",
  "zip_code": "",
  "latitude": 0,
  "longitude": 0,
  "rating": 0,
  "rating_count": 0,
  "display_address": [
    "",
    "",
    "",
    "",
    ""
  ],
  "price_range": "$"
}
```
### DEL
request example
```sh
curl -X DELETE \
  'http://localhost:8080/business_search/04880121-dc4c-11ed-b037-14f6d817012e' \
  --header 'Accept: */*' \
  --header 'User-Agent: Thunder Client (https://www.thunderclient.com)' \
  --header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiI2MnRla25vbG9naSIsInJvbGUiOjEsImlhdCI6MTY4MTU2MTg4MCwiZXhwIjoxODc5NTYwNjc5fQ.2SDZCXjGYNIPhJj5lJxFQQmg8O8TJ6eqrDW2GXlpYJk'
```

response with status `OK` if success.

token used is jwt with algo `HS256` and have additional claim `role`. secret used to signing jwt can be found at `.env.example`.
example of jwt payload
```json
{
  "iss": "62teknologi",
  "role":1,
  "iat": 1681561880,
  "exp":1879560679
}
```

