
# Modules

* A module has a unique identifier and can be distributed (e.g. library)
* Every Go project goes into a new module
* Projets can use (i.e. import from) multiple modules (custom and third-party)
* Created and managed via `go mod` commands and `go.mod` file


# Packages

* Every module contains at least one package (the "main" package)
* Multiple files can make up the same package (via package instruction - i.e. `package main`)
* A module may contain multiple packages, stored in subfolders
* Can be imported (it's ultimately always a module that is imported)

## Creating `go.mod`

```bash
# You name the module in a way that matches the folder path beginning with a github.com folder so the module matches the published path
 go mod init github.com/yourOrganization/moduleFolderPath
```

### Example go.mod
```go
module github.com/716green/academind/firstapp

go 1.16
```

### Run code in a folder with a `go.mod` file
```bash
go run .
```

## Sharing across files using packages
### This can be great for organization and readability
### There should be a subfolder with a one word name for each additional package
```go
// firstapp.go

package main

import "fmt"

func main() {
	fmt.Println("Hello World")

	// greetingText := "Hello World"
	luckyNumber := 17
	evenMoreLuckyNumber := luckyNumber + 5

	fmt.Println(greetingText)
	fmt.Println(luckyNumber)
	fmt.Println(evenMoreLuckyNumber)
}

```

```go
// Using the same package (main)
// variables.go
// package main

// As a separate package
// ./greeting/variables.go
package greeting

var greetingText string = "Hello World from variables.go"

```

## Importing Packages

```go
import (
	"fmt"
  
	"github.com/716green/academind/firstapp/greeting"
)
// or
import "fmt"
import "github.com/716green/academind/firstapp/greeting"

// Auto-format will change the second version to the first one
```