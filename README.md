### Project Structure
- `internal`
  - `api`: defined api endpoints
  - `config`: application configuration
  - `controller`: business logic and orchestration source code
  - `middleware`: all middlewares executed before processing a request.
  - `repository`: external integrations like databases and other services.
  - `runner`: main program execution configuration
  - `util`: utility code
- `scripts`: files needed for local execution

### Dependencies
In order to run the project you need to install the following tools:
- Docker Desktop.

### How to run locally
From the project's root directory, execute `make run-local` from a terminal window to 
run the project in docker. This will perform the following steps:
- Build the main project.
- Configure local container with localstack.
- Configure local container with mockServer (Core Banking simulation).
- Execute `deps.sh` that will create default tables in DynamoDB.
- Run main application service.

There is a postman collection inside tests directory. An example of the request execution:
1. Call Login -> Get the token.
2. Go to Create -> Replace the token in the header `Authorization` -> Call Create.
3. Go to Get -> Set the Path param -> Call Get
4. Go to Refund -> Set the Payment Id in the request body -> Call Refund