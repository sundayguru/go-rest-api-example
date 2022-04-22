# Piggy Bank

API that manipulates the basic operations of a Bank. This is my first project with Golang.

# Installation

To execute project in your local machine, you need to have installed docker and docker compose in your machine.

- Copy environment file: `cp .env.example .env`
- Build the docker image: `docker-compose build`
- Run the project: `docker-compose up`

If everything is successful, you can access the application at http://localhost:8080

# Testing

To execute the tests, you need to have installed docker and docker compose in your machine.

- Build the docker image: `docker-compose --f docker-compoose-test.yml build`
- Run the project: `docker-compose --f docker-compoose-test.yml up`
- Run the tests: `docker exec -it api-piggybanktest go test ./...`

# API Documentation

## Create Account

### POST /bank-account

Creates a new account

### Inputs

| Name     | Descriptions                                                                          |
| -------- | ------------------------------------------------------------------------------------- |
| name     | The full name of the account owner (\*Required)                                       |
| email    | The email address of the account owner (\*Required and must be a valid email address) |
| username | The username of the account owner (\*Required)                                        |

### Sample Request

```
{
    "name": "testing",
    "email": "testing@example.com",
    "username": "testing"
}
```

### Sample Output

```
{
    "ID": 9,
    "CreatedAt": "2022-04-21T15:37:53.0036643Z",
    "UpdatedAt": "2022-04-21T15:37:53.0036643Z",
    "DeletedAt": null,
    "name": "adasdsa",
    "email": "testing@example.com",
    "username": "testing",
    "BankAccountTransactions": null
}
```

## Deposit

### POST /bank-account/:username/deposit

Deposit cash to your account

### Inputs

| Name   | Descriptions                                |
| ------ | ------------------------------------------- |
| amount | The amount you want to deposit (\*Required) |
| note   | Optional note to describe the transaction   |

### Sample Request

```
{
    "amount": 100,
    "note": "test"
}
```

### Sample Output

```
{
    "ID": 17,
    "CreatedAt": "2022-04-21T15:47:48.8521191Z",
    "UpdatedAt": "2022-04-21T15:47:48.8521191Z",
    "DeletedAt": null,
    "note": "test",
    "amount": 100,
    "type": "CREDIT",
    "BankAccountID": 9
}
```

## Withdraw

### POST /bank-account/:username/withdraw

Withdraw cash from your account

### Inputs

| Name   | Descriptions                                |
| ------ | ------------------------------------------- |
| amount | The amount you want to deposit (\*Required) |
| note   | Optional note to describe the transaction   |

### Sample Request

```
{
    "amount": 50,
    "note": "test"
}
```

### Sample Output

```
{
    "ID": 18,
    "CreatedAt": "2022-04-21T15:50:08.833657Z",
    "UpdatedAt": "2022-04-21T15:50:08.833657Z",
    "DeletedAt": null,
    "note": "test",
    "amount": -50,
    "type": "DEBIT",
    "BankAccountID": 9
}
```

## Check Account Balance

### GET /bank-account/:username/balance

Check your account balance

### Sample Output

```
{
    "balance": 50
}
```

# Go Packages Used

- github.com/gorilla/mux
- github.com/jinzhu/gorm
- github.com/qor/validations
