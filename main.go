package main

import (
	"floraGenCode/temp"
	"fmt"
	"os"

	"github.com/xuri/excelize/v2"
)

type nameAndData struct {
	name string
	data string
}

var (
	entityName = make(map[string]string, 0)
	entityData = make(map[string][][]string, 0)
)

func main() {
	err := readXLSXFile("./实体.xlsx", "Sheet1")
	if err != nil {
		panic(err)
	}
	datas := splitTemplate()
	for _, data := range datas {
		createAndWriteGoFile(data)
	}
}

// createGoFile 创建Go文件
func createAndWriteGoFile(data nameAndData) {
	fileName := "./student/" + data.name + ".go"
	file, err := os.Create(fileName)
	if err != nil {
		fmt.Println("createGoFile filed err:", err, "name is ", data.name)
		return
	}
	defer file.Close()

	num, err := file.WriteString(data.data)
	if err != nil {
		fmt.Println("write err", err)
		return
	}
	fmt.Println("num", num)
}

// readXLSXFile 读取xlsx文件内容
func readXLSXFile(name, sheet string) error {
	f, err := excelize.OpenFile(name)
	if err != nil {
		fmt.Println("file open field,err:", err)
		return err
	}
	defer f.Close()
	rows, err := f.GetRows(sheet)
	if err != nil {
		fmt.Println("file get rows field,err:", err)
		return err
	}
	var cn, en string
	for _, row := range rows[1:] {
		if cn != row[0] && en != row[1] {
			cn = row[0]
			en = row[1]
			entityName[row[1]] = row[0]
			entityData[row[1]] = append(entityData[row[1]], row[2:])
		} else {
			entityData[row[1]] = append(entityData[row[1]], row[2:])
		}
	}
	fmt.Println("entityData", entityData)
	return err
}

// splitTemplate 拼接需要写入的内容
func splitTemplate() []nameAndData {
	n := []nameAndData{}
	for name, datas := range entityData {
		entityDate := ""
		filedDate := ""
		nameCn := entityName[name]
		for _, data := range datas {
			filedDate += fmt.Sprintf(temp.FieldContent, data[1], temp.FieldType[data[2]], data[0], data[0]) + "\n"
		}
		entityDate = fmt.Sprintf(temp.PackName, name) +
			fmt.Sprintf(temp.EntityContent, name, nameCn, nameCn, filedDate)

		n = append(n, nameAndData{
			name: name,
			data: entityDate,
		})
	}
	return n
}
