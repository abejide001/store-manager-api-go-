# Store-Manager
This is an that helps store owners manage sales and product inventory records. This application is meant for use in a single store.

## Introduction

Welcome to version 1 of Store Manager API. Below you will find a current list of available methods on different endpoints.

## Getting started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes.

## Prerequisites

To work with this project you need to have the following installed on your local machine

1. [Go](https://golang.org/dl/)
2. [Git](https://git-scm.com/downloads)
3. [Postman](https://www.getpostman.com/)

## Install and run locally

```bash
$ git clone https://github.com/abejide001/store-manager-api-go-.git
$ cd store-manager-api-go

# rename .env.sample to .env, and set your environment variables

$ go mod download
$ go build main.go
$ ./main.go
```
## API Usage

API BASE URL http://localhost:8080/api/v1. It's recommended to attach a `Authorization` Header containing the generated `token` from `/api/auth/login` to all access all requests.

### Products endpoints `/api/v1/products`

| method | route          | description             | data                                               |
| ------ | -------------- | ----------------------- | ---------------------------------------------------| 
| GET    | /products      | Get all store products  |                                                    |
| GET    | /products/:id  | Get a product           |                                                    |
| POST   | /products      | Create a product        |`{name, price, description, attendant_id}`          |
| PUT    | /products/:id  | Update a product        |                                                    |
| DELETE | /products/:id  | Delete a product        |                                                    |