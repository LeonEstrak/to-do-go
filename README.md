# To-Do-API

A simple CRUD compliant backend API with an in-memory database for a [to-do application](https://github.com/LeonEstrak/to-do-angular) written in [Go](https://golang.org/)

Following are the exposed API endpoints:

| **METHOD** | Endpoint   | Description                                     |
| ---------- | ---------- | ----------------------------------------------- |
| **GET**    | `/todo`    | Retrieves the complete to-do list               |
| **POST**   | `/todo`    | Adds a todo item to the list                    |
| **PUT**    | `/todo`    | Updates a todo item already present in the list |
| **POST**   | `/deltodo` | Deletes a todo item                             |

## Installation
- Install Go by following the official installation steps from [here](https://golang.org/doc/install).
- Run `git clone https://github.com/LeonEstrak/to-do-go`
- `cd` into `to-do-go` directory 
- If you're in a Unix based system then you may use the ` make dev ` command in your terminal or else you may simply run it using `go run main.go` 