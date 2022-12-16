package main

import (
	"floraGenCode/temp"
	"fmt"
	"os"

	"github.com/xuri/excelize/v2"
)

func main() {
	f, err := excelize.OpenFile("./实体.xlsx")
	if err != nil {
		fmt.Println("file open field,err:", err)
		return
	}
	defer f.Close()
	rows, err := f.GetRows("Sheet1")
	if err != nil {
		fmt.Println("file get rows field,err:", err)
		return
	}
	var entity, entityName, entityDate, filedDate string
	for _, row := range rows[1:] {
		if row[0] != "" && row[1] != "" {
			entity = row[0]
			entityName = row[1]
		}
		filedDate += filedDate + "\n" + fmt.Sprintf(temp.FieldContent, row[3], row[4], row[3], row[3])
	}
	entityDate = fmt.Sprintf(temp.PackName, entityName) + fmt.Sprintf(temp.EntityContent, entityName, entity, entity, filedDate)
	createGoFile(entityName, entityDate)
}

// createGoFile 创建Go文件
func createGoFile(name, data string) {
	fileName := "./student/" + name + ".go"
	file, err := os.Create(fileName)
	if err != nil {
		fmt.Println("createGoFile filed err:", err, "name is ", name)
		return
	}
	defer file.Close()

	num, err := file.WriteString(data)
	if err != nil {
		fmt.Println("write err", err)
		return
	}
	fmt.Println("num", num)
}

func readXLSXFile(name string) {
	f, err := excelize.OpenFile("name")
	if err != nil {
		fmt.Println("file open field,err:", err)
		return
	}
	defer f.Close()
	rows, err := f.GetRows("Sheet1")
	if err != nil {
		fmt.Println("file get rows field,err:", err)
		return
	}
}
