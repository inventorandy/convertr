# CONVERTR Common Adapter Functions in Go

This package provides some common methods for adapting between different types in Go.

## Installation

```bash
go get github.com/inventorandy/convertr
```

## Adapters

### Adapt

Use the `adapt` package to convert between common structures.

#### Using JSON Struct Tags

This converts between structs using the JSON struct tags on each object.

```go
package main

import "github.com/inventorandy/convertr/adapt"

type StructA struct {
  ID int64 `json:"id"`
  Name string `json:"name"`
}

type StructB struct {
  ID int `json:"id"`
  Name *string `json:"name"`
}

func main() {
  // Struct to Convert From
  from := StructA{1234, "Hello, World!"}

  // Struct to Convert To
  var to StructB

  // Convert the Structs
  if err := adapt.UsingJSON(from, &to); err != nil {
    log.Fatalln(err)
  }
}
```

### Pointers

Use the `ptr` package to easily convert between types and pointers.

#### Example Usage

```go
package main

import "github.com/inventorandy/convertr/ptr"

func main() {
  // Create a Pointer to a String
  str := "Hello, World"
  strPtr := ptr.To(str) // Returns *string

  // Convert back to a String
  newStr := ptr.From(strPtr) // Returns string
}
```
