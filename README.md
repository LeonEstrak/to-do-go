# To-Do-API

A simple CRUD compliant backend API with an in-memory database for a [to-do application](https://github.com/LeonEstrak/to-do-angular) written in [Go](https://golang.org/)

The data expected by the server is of the following shape

    type Todo struct {
      ID        string `json:"id"`
      Completed bool   `json:"completed"`
      Message   string `json:"message"`
    }

Following are the exposed API endpoints:

| **METHOD** | Endpoint | Expected Body   | Description                                     |
| ---------- | -------- | --------------- | ----------------------------------------------- |
| **GET**    | `/todo`  | `None`          | Retrieves the complete to-do list               |
| **POST**   | `/todo`  | A `Todo` object | Adds a todo item to the list                    |
| **PUT**    | `/todo`  | A `Todo` object | Updates a todo item already present in the list |
| **DELETE** | `/todo`  | A `Todo` object | Deletes a todo item                             |

## Installation
- Install Go by following the official installation steps from [here](https://golang.org/doc/install).
- Run `git clone https://github.com/LeonEstrak/to-do-go`
- `cd` into `to-do-go` directory 
- If you're in a Unix based system then you may use the ` make dev ` command in your terminal or else you may simply run it using `go run main.go` 