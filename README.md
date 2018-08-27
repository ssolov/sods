# sods
** About
A simple parser for ODS tables
** Install
go get -u github.com/ssolov/sods
** Example
Read and print to stdout all cell from ODS table
#+BEGIN_SRC go
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
		for _, row :=  range sh.Rows {
			for _, cell :=  range row.Cells {
				fmt.Print(cell.Text)
			}
			fmt.Println()
		}
	}
}
#+END_SRC