package main

import (
	"fmt"

	"github.com/xuri/excelize/v2"
)

var student = `package %s \n func %s() {}`

func main() {
	f, err := excelize.OpenFile("./实体.xlsx")
	if err != nil {
		fmt.Println("file open field,err:", err)
		return
	}
	defer func() {
		if err := f.Close(); err != nil {
			fmt.Println("file close field,err:", err)
		}
	}()
	rows, err := f.GetRows("Sheet1")
	if err != nil {
		fmt.Println("file get rows field,err:", err)
		return
	}
	for _, row := range rows {
		// for _, r := range row {
		// // 	fmt.Printf("%s,", r)
		// }
		// fmt.Println()
		fmt.Printf(student, row[1], row[3])
	}
}
