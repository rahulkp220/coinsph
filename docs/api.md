# API Documentation


## Create an account
#### `POST /accounts`


Sample Request
```
{
        "username":"alpha",
        "currency":"USD",
        "balance":"1000.00"
}
```

Sample Response
```
    {
        "ok": "success"
    }
```

## Get an account by ID
#### `GET /accounts/:id`

Sample Response
```
{
    "result": [
        {
            "id": "c7153ced-9c7c-4f76-99a9-6b1b782d6046",
            "username": "alpha",
            "currency": "USD",
            "balance": "550.000000"
        },
    ]
}
```

## Get all accounts
#### `GET /accounts`

Sample Response
```
{
    "result": [
        {
            "id": "c7153ced-9c7c-4f76-99a9-6b1b782d6046",
            "username": "alpha",
            "currency": "USD",
            "balance": "550.000000"
        },
        {
            "id": "b45e8df7-a84e-44f5-b3f3-ea44bf20d770",
            "username": "beta",
            "currency": "USD",
            "balance": "1450.000000"
        },
        {
            "id": "2e1f9347-5b67-4e52-8d61-23bf03957a18",
            "username": "gamma",
            "currency": "USD",
            "balance": "1000"
        }
    ]
}

```

## Delete an account
#### `DELETE /accounts/:id`

Sample Response
```
{
    "ok": "success"
}
```

## Make a transafer
#### `POST /payments`

Sample Request
```
{
    "sender":"c7153ced-9c7c-4f76-99a9-6b1b782d6046",
    "reciever":"b45e8df7-a84e-44f5-b3f3-ea44bf20d770",
    "amount":450
}

```

Sample Response 
```
{
    "ok": "success"
}
```

## Get all transfers
#### `GET /payments`

Sample Response
```
{
    "result": [
        {
            "id": "16155630-ac59-4516-a249-9340ffa08e23",
            "sender": "c7153ced-9c7c-4f76-99a9-6b1b782d6046",
            "reciever": "b45e8df7-a84e-44f5-b3f3-ea44bf20d770",
            "amount": 450,
            "initiated": "0000-01-01T16:56:04.084817Z",
            "completed": "0000-01-01T16:56:04.108362Z"
        }
    ]
}
```