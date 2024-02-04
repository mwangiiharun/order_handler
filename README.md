# Introduction

## pre-requisites

- make sure you have the go installed in your system. If not, you can download it from [here](https://golang.org/dl/)

- Mak sure you have a firestore project created in your google cloud console. If not, you can create it from [here](https://console.cloud.google.com/firestore)

## Usage

- Clone the repository
- Create a `.config.yaml` file in the root of the project and add your config to it. Here is an example of the config file

``` yaml
server:
  name: "my-server"
  port: 8080
  version: 1.0.0
  logger:
    level: "info"
storage:
  credentials_file: "path/to/credentials.json"
  project_id: "my-project"
  database_id: "my-database"
services:
  - name: "service-1"
    type: "Http"
    parent: 1.0.0
    collectionId: "my-collection"
  - name: "service-2"
    type: "Http"
    parent: 1.0.0
    collectionId: "my-collection"
  - name: "service-3"
    type: "Http"
    parent: 1.0.0
    collectionId: "my-collection"
  - name: "service-4"
    type: "Http"
    parent: 1.0.0
    collectionId: "my-collection"

```

To run the server, run the following command

``` bash
go run cmd/server/main.go
```

Or alernativelly where you can specify the config file

``` bash
go run cmd/server/main.go -config .config.yaml
```
