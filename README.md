# init project

```
mkdir MY_PROJECT_NAME
cd MY_PROJECT_NAME
go mod init github.com/oarriet/MY_PROJECT_NAME
```

## create .go file
```bash
touch example.go
```

## main example
```go
package main

import "fmt"

func main() {
	fmt.Println("Hello, world!")
}
```

## run example
```bash 
go run example.go
```