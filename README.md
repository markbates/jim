# Jim

Jim is a task runner, similiar to [https://github.com/markbates/grift](https://github.com/markbates/grift).

### Requirements

* Go 1.13+
* Modules

## Installation

```bash
$ go get github.com/markbates/jim/cmd/jim
```

---

## Usage

To create a new task that you can run with jim. You need to create a package level function with the following API in your module:

```go
func <name>(ctx context.Context, args []string) error {
	return nil
}
```

You can place these functions anywhere in your application in any package. You don't need to follow any naming conventions or use any special build tags. Your function need only take a [`context#Context`](https://godoc.org/context#Context) and a slice of `string`, and return an error.


See the [./examples/ref](./examples/ref) application for the full code.

```text
.
├── db
│   ├── seed
│   │   └── users.go
│   └── seed.go
├── go.mod
├── go.sum
├── ref.go
└── task
    └── task.go

3 directories, 6 files
```

The `go.mod` file:

```go
module ref

go 1.13
```

Example task `ref/db/seed/users.go`

```go
package seed

import (
	"context"
	"fmt"
)

// Users puts all of the users into all of the databases
func Users(ctx context.Context, args []string) error {
	fmt.Println("loading users", args)
	return nil
}
```

### Running a Task via CLI

To run the above mentioned task we can use the `jim` command:

```bash
$ jim db:seed:Users 1 2 3 4

loading users [1 2 3 4]
```

Let's break down the `db:seed:Users` bit, shall we? The last part, `Users` is the name of the function that will be run. This **MUST** match capitalization.

The `db:seed` part is converted to the package `<module path>/db/seed` which, hopefully, contains a `Users` function that matches the correct API.

### Running a Task via API

To run your task programmatically in your application, give it a `context.Context`, some arguments, and you're good to go. :)

### Getting Task Help

The `jim -h` flag, followed by the task will print the GoDoc for that function.

```bash
$ jim -h db:seed:Users

package seed // import "ref/db/seed"

func Users(ctx context.Context, args []string) error
    Users puts all of the users into all of the databases
```
