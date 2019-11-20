# Recombee

Recombee is a Go client library for accessing the Recombee V2 API.

## Install
```sh
go get github.com/mavvverick/recombee
```

## Usage

```go
import "github.com/mavvverick/recombee"
```


## Initialize

To get recomended item to a specific user:

```
package main
import (
	"context"
	"github.com/mavvverick/recombee"
)


func main() {
	client := recombee.NewClient(nil)
}
```


## Examples

To get recommended item to a specific user:

```
u := User{ID: "1"}
opts := &recombee.ListOptions{
    Count:        10,
    RotationRate: 1,
}

recoms, _, err := client.Reco.ItemsToUser(ctx, u, opts)
if err != nil {
    t.Fatalf("unexpected error: %s", err)
}

fmt.Println(recoms)
```

To get default recommendation

```
u := User{ID: "1"}
l := logics["recombee:default"]

recoms, _, err := client.Reco.GetPreset(ctx, u, l)
if err != nil {
    t.Fatalf("unexpected error: %s", err)
}

fmt.Println(recoms)
```


## Versioning

Each version of the client is tagged and the version is updated accordingly.

To see the list of past versions, run `git tag`.


## Documentation

For a comprehensive list of examples, check out the [API documentation]().

For details on all the functionality in this library, see the [GoDoc]() documentation.


## Contributing

We love pull requests! Please see the [contribution guidelines](CONTRIBUTING.md).
