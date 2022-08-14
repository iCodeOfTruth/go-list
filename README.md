# go-list
List Manager created in Go

# How to install
```
go get -u github.com/iCodeOfTruth/go-list
```

# Usage

```golang
package main

import golist "github.com/iCodeOfTruth/go-list"

var RandomData = []interface{}{1, 2, 3}

func main() {
	listM := golist.NewListManager(RandomData)
	listM.Random()                    // Get Random Element
	listM.Next()                      // Get Next Element
	listM.RemoveIndex(1)              // Remove Element at Index
	listM.Remove(2)                   // Remove Element
	listM.Append(4)                   // Append Element
	listM.Extend([]interface{}{5, 6}) // Extend List
}

```
