# STD

This project contains a bunch of functionality I consider would've been nice to have been included in the Go standard library.

## WORK IN PROGRESS

## Strings

Contains functionality related to string manipulation.

```go
package main

import (
    "fmt"
    "github.com/calini/std/strings"
)

func main() {
    str := strings.Before("abcd", "c")
    fmt.Printf("%s\n", str) // "ab" 
}
```