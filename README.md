# petstore-contract-first-go

This is a simple sample Go application following the contract-first approach, using [deepmaps/oapi-codegen](https://github.com/deepmap/oapi-codegen) for OpenAPI specification code generation, and [sqlc](https://github.com/sqlc-dev/sqlc) to compile and generate SQL code.

## Why Use the Contract-First Approach?

The contract-first approach to API development involves designing the API contract (e.g., OpenAPI specification) before writing any code. This approach offers several benefits:

- **Clarity and Consistency**: By defining the API contract upfront, you establish clear guidelines for the API's behavior and structure, promoting consistency across development teams and reducing ambiguity.

- **Client-Server Separation**: With a well-defined API contract, client and server teams can work independently. Clients can start development based on the contract before the server implementation is complete.

- **Reduced Integration Risks**: Since clients and servers adhere to a common contract, integration issues are minimized, leading to smoother and more predictable integration processes.

- **Code Generation**: OpenAPI generators offer a diverse selection of tools to generate essential objects and boilerplate code for both clients and servers. This flexibility accelerates development and enables developers to prioritize the most crucial aspects of implementation, ensuring alignment with the defined API "contract".

## Tools Used

#### OpenAPI Specification and oapi-codegen

The OpenAPI Specification (OAS) is a standard format for describing RESTful APIs. oapi-codegen is a tool that generates Go code from OpenAPI specifications, enabling type-safe HTTP client and server implementations.

#### SQLC

sqlc is a tool for generating type-safe Go code from SQL. It allows you to write SQL queries alongside your Go code and generate Go code with type-safe methods for executing those queries.


### Package structure


This project follows a simple and encapsulated package structure to facilitate easy testing of different implementation layers.

Below is an overview of the directory structure:

```markdown
.
├── .kubernetes             # Kubernetes manifests
├── build                   # Build-related files (e.g., Dockerfile)
├── cmd                     # Server main executable
├── docs                    # Documentation files containing the openapi.yaml file for our application
│   └── openapi.yaml        # OpenAPI specification for the application
├── internal                # Root directory for application sources
    ├── http                # HTTP-related components (Controllers, Handlers, Middlewares, Server)
    │   └── model.gen.go    # Generated model files from oapi-codegen
    │   └── router.gen.go   # Generated server from oapi-codegen
    │   └── controller.go   # Implementation of generated code
    └── persistence         # Persistence layer implementations
        └── postgresql      # The postgresql implementation of the applications persistence layer
            └── .schema     # sqlc schema files for generation
            └── sqlcgen     # sqlc generated source code
```



### Configuration

You can configure the following environment variables:

| Variable                | Description                                      |
|-------------------------|--------------------------------------------------|
| `APP_API_PORT`          | The port on which the API server will run.       |
| `APP_API_BASEURL`       | The base URL for API endpoints.                  |
| `APP_POSTGRESQL_URL`    | The URL of the PostgreSQL database.              |
| `APP_POSTGRESQL_DATABASE` | The name of the PostgreSQL database to use.     |


## Usage

#### Local dev with postgresql docker container

```shell
# start postgresql local
docker run -e POSTGRES_PASSWORD=admin -e POSTGRES_DB=petstore -p 5432:5432 -d postgres
 
# start application
go run cmd/server/main.go
```

#### Run application with docker compose

```shell
docker compose up --build
```