# invoice-marketplace

## Description

This is a golang RESTful API created with go-chi web framework and PostgreSQL as DB. It demonstrates an invoice bidding system that facilitates financing of invoices. It contains two users which are issuers and ivestors. Issuers creates invoices and investors can bid on them. The system enables create,view invoices and view and track balances of issuers and investors.

## Folder structure

```
├───cmd
├───internal
│   ├───auth
│   ├───db
│   ├───domain
│   ├───dto
│   ├───handlers
│   ├───middleware
│   ├───repositories
│   ├───service
│   └───tests
│       └───services
├───migrations
├───pg-data
├───pkg
│   └───errors
└───postman_collection

```

- cmd - contains the main application entrypoint
- internal - holds the core application projects.
  - db - contains database connection creation
  - domain - contains domain/models of the project
  - services - contains business logic and interfaces
  - repositories - contains interfaces & implements their methods that handles CRUD operations
  - middleware - contains JWT middleware & permission middleware
  - handlers - controllers that handles incoming requests and responses
  - auth - contains JWT validation & creation
  - dto - contains data transfer object types
  - tests - contains mocks of repos and test cases of service layer
- migrations - init SQL script to create database, tables and insert dummy data
- pg-data - persistent data volume of postgres db
- pkg - holds the packages using accross different layers
  - errors - customized package for handling errors throughout project
- postman_collection - contains postman collection (exported JSON file) to test endpoints

## Prerequisites

- Golang installed in your system (version 1.19 or higher)
- Docker installed in your system (This project will be built and run using Docker)

# Getting Started

1.  Clone the repository

```
$ git clone https://github.com/niluwats/invoice-marketplace.git
```

2. Navigate to the project directory

```
$ cd invoice-marketplace
```

3.  Build and run the application using Docker Compose:

```
$ docker-compose up --build
```

4. Access the applicatin

   Once the containers are up and running, you can access the application at http://localhost:8080

## Postgres Server Login

If you need to log in to the PostgreSQL server, follow these steps.

1.  Ensure that you have PostgreSQL installed and running on your system.
2.  Open a command-line interface or terminal.
3.  Use the following command to log in to the PostgreSQL server:

```
 $ psql -U mktplaceUser -W -h localhost -p 5432 invoice_marketplace
```

4.  When prompted, enter the password "inUser123!" for the specified username.
5.  If the credentials are correct, you should now be logged in to the PostgreSQL server via the command-line interface.

## Dummy Users to test the application

For testing purposes of endpoints, you can use the following dummy users:

### Login as Investors

 Investors can place bids on the invoices that issuers create.

- User 1

  - Username: `robert123@gmail.com`
  - Password: `Abc123!`

- User 2

  - Username: `will123@gmail.com`
  - Password: `Abc123!`

- User 3

  - Username: `sara123@gmail.com`
  - Password: `Abc123!`

This user represents an investor in the system. You can use these credentials to test the endpoints related to the investor's functionalities (bid on invoices & view functionalities).

### Login as Issuer

 Only issuers can create invoices and approve trades

- Username: `jane123@gmail.com`
- Password: `Abc123!`

This user represents an issuer in the system. You can use these credentials to test the endpoints related to the issuer's functionalities(create invoice,trade approvals & view functionalities).

Remember to include the appropriate authentication headers or tokens when making API requests to simulate the behavior of authenticated users.

## Endpoints

The following endpoints are available in this API:

| Method | Endpoint                | Description                  | Headers                                                               |
| ------ | ----------------------- | ---------------------------- | --------------------------------------------------------------------- |
| GET    | `/invoice`              | View all invoices            | `Authorization: Bearer <access_token>,`                               |
| POST   | `/invoice`              | Create new invoice           | `Authorization: Bearer <access_token>,`                               |
| GET    | `/invoice/{id}`         | View invoice by ID           | `Authorization: Bearer <access_token>,Content-Type: application/json` |
| POST   | `/bid`                  | Place bid                    | `Authorization: Bearer <access_token>,Content-Type: application/json` |
| GET    | `/bid/{id}`             | View bid                     | `Authorization: Bearer <access_token>`                                |
| PATCH  | `/invoice/{invoice_id}` | Approve trade                | `Authorization: Bearer <access_token>`                                |
| GET    | `/bids/{invoice_id}`    | View all bids for an invoice | `Authorization: Bearer <access_token>`                                |
| GET    | `/investor`             | View all investors           | `Authorization: Bearer <access_token>`                                |
| GET    | `/investor/{id}`        | View investors by ID         | `Authorization: Bearer <access_token>`                                |
| GET    | `/issuer`               | View all issuers             | `Authorization: Bearer <access_token>`                                |
| GET    | `/issuer/{id}`          | View issuer by ID            | `Authorization: Bearer <access_token>`                                |
| POST   | `/register`             | Register as an investor      | `Content-Type: application/json`                                      |
| POST   | `/auth`                 | Get access token             | `Content-Type: application/json`                                      |

## Mock Testing

The services in this application are tested using mock testing techniques to isolate dependencies and ensure reliable and efficient testing. Mock testing allows for the simulation of external dependencies, enabling focused testing of individual services.

The following libraries and frameworks are utilized for mock testing in this project:

- **Mockery**: Mockery is used to create mock objects for dependencies and define expected behaviors during unit testing.

- **Testify**: Testify is the testing framework used to create test fucntionalities

### Runnig Tests

1. Open a command-line interface or terminal and navigate to the tests directory of the project.

```
$ cd ./internal/tests/services
```

2. Run tests

```
$ go test -v
```

## Dependencies

- [go-chi](https://github.com/go-chi/chi) - web framework
- [pgx postgresql driver](https://github.com/jackc/pgx) - Golang PostgreSQL Driver and Toolkit
- [Jwt](https://github.com/golang-jwt/jwt) - JWT
- [Crypto](https://golang.org/x/crypto) - Hashing
- [Testify](https://github.com/stretchr/testify) - Testing services
- [godotenv](https://github.com/joho/godotenv) - Loading env variables

You can install above dependencies by running the following commands:

```
$ go get -u github.com/go-chi/chi/v5
$ go get github.com/jackc/pgx/v5
$ go get -u github.com/golang-jwt/jwt/v5
$ go get github.com/joho/godotenv
$ go get golang.org/x/crypto
$ go get github.com/stretchr/testify
```
