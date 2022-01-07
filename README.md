# go-microservice-sample

Project aim is share knowledge about golang, microservice, grpc, rest and database.

## Struct of project

I have inspiration on [this repo](https://github.com/golang-standards/project-layout) .

## Architecture of project

![Demonstration of de architecture](docs/architecture.png)
This diagram was made o [Excalidraw](https://excalidraw.com/)

## Commands

Use the next command for generating or updating of protobuf files:

```cmd
protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    api/helloworld.proto
```
