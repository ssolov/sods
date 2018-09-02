# sods
## About
A simple parser for ODS tables
## Installation
```shell
go get -u github.com/ssolov/sods
```
## Example
Read and print all cell from ODS table to stdout
```go
package main

import (
	"github.com/ssolov/sods"
)

func main() {
	ods, err := sods.Read("filename.ods")
	if err != nil {
		panic(err)
	}

	for _, sh := range ods.Sheets {
		fmt.Println("Table name: " + sh.Name)
		for _, row := range sh.Rows {
			for _, cell := range row {
				fmt.Print(cell)
			}
			fmt.Println()
		}
	}
}
```