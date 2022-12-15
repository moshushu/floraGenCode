package main

import (
	"fmt"
	"os"

	"github.com/xuri/excelize/v2"
)

var student = "package %s \n func %s() {}"

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
	for index, row := range rows {
		if index == 0 {
			data := fmt.Sprintf(student, row[1], row[3])
			createGoFile(row[1], data)
			fmt.Println(data)
		}
	}
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

// func writeGoFile(name, data string) {
// 	name = "./student/StuentInfo.txt"
// 	data = "去你妈的"
// 	file, err := os.OpenFile(name, os.O_WRONLY, 0777)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer file.Close()
// 	buffer := bufio.NewWriter(file)
// 	str, err := buffer.WriteString(data)
// 	if err != nil {
// 		log.Fatalln("writestring filed err:", err)
// 	}
// 	fmt.Println("str", str)
// }
