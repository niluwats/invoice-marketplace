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
├───pg-data
└───pkg
    └───errors

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
- pg-data - persistent data volume of postgres db
- pkg - holds the packages using accross different layers
  - errors - customized package for handling errors throughout project

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

Once the containers are up and running, you can access the application at <localhost:8080>
