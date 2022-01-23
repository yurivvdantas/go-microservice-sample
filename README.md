# go-microservice-sample

Project aim is share knowledge about golang, microservice, grpc, rest and database.

## What you will found here

This is a project with a simulation of a rest service exposes and with call internal grpc.
The context is a cryptcurrency service with is possible to get info of registered coins, upvotes, adding a new coin and have a streaming with the real time upvotes given a cryptocurrency.
In this way, its possible test grpc with a browser or with a rest client.

## Endpoints

Get infos given a crypto id:

```curl
curl --request GET \
  --url http://localhost:8080/crypto/1
```

Get all registred cryptos

```curl
curl --request GET \
  --url http://localhost:8080/crypto/
```

Upvote a crypto by id:

```curl
curl --request POST \
  --url http://localhost:8080/crypto/1/upvote 
```

Register a new crypto:

```curl
curl --request POST \
  --url http://localhost:8080/crypto/ \
  --header 'Content-Type: application/json' \
  --data '{
"name":"Ethereum",
"code":"ETC",
"description":"A amazing coin"
}'
```

Stream with a live crypto upvote by id:

```curl
curl --request GET \
  --url http://localhost:8080/crypto/1/upvote
```

## Struct of project

I have inspiration on [this repo](https://github.com/golang-standards/project-layout) .

## Architecture of project

![Demonstration of de architecture](docs/architecture.png)
This diagram was made on [Excalidraw](https://excalidraw.com/)

There is two service on this project (so far), a gateway rest service and a grpc service.
The gateway will receive a rest requisition, call grpc on another service and return a json to the client.

## Setup

### Database

A complete instrunctions how to install and do a initial config of a database using golang, docker and MySql can be found on [this link](https://go.dev/doc/tutorial/database-access) of the offcial documentation.

I have set de password "admin" for this project and the user "root".

After this, must have create a database on the docker console image of the MySql:
`create database cryptos;`

And then:
`use cryptos;`

Run the script [Create_crypto_table.sql](scripts/database/init/Create_crypto_table.sql) for init the database.

### Prerequists

For running this application, you have to:

-[Configure database](###database)

-[Installed golang](https://go.dev/learn/)

-[Installed protobuf](https://grpc.io/docs/languages/go/quickstart/)

## Commands

Use the next command for generating or updating of protobuf files:

```shell
protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    api/crypto.proto
```

## Tests

For run tests, use command from the root path of project:

```shell
go test ./test/
```

This will run all tests in location of tests folder.

For run all tests, the database must be on. Maybe in future i fix that making a mock of a database connection or something like that.

## Running

Must have running two application: [gateway_main](cmd/gateway/gateway_main.go) and [crypto_votes_service_main](cmd/crypto-votes-service/crypto_votes_service_main.go).
For run them just go to the path on a console and use the command:

```shell
go run gateway_main.go
```

```shell
go run crypto_votes_service_main.go
```

Or just use VScode navigate to the files and press F5. For execute a second service, you have to access Run menu and then `Start Debuggin`.
