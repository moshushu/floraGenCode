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
	// 选项
	optionData = make(map[string][][]string, 0)
	// 关联关系
	relation = make(map[string][]string, 0)
)

func main() {
	readXLSXFile("./实体.xlsx", "实体")
	readXLSXFile("./实体.xlsx", "选项")
	readXLSXFile("./实体.xlsx", "关联关系")
	datas := splitTemplate()
	for _, data := range datas {
		createAndWriteGoFile(data)
	}
}

// createGoFile 创建Go文件
func createAndWriteGoFile(data nameAndData) {
	name := splitName(data.name)
	fileName := "./student/" + name + "_models" + ".go"
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
func readXLSXFile(name, sheet string) {
	f, err := excelize.OpenFile(name)
	if err != nil {
		fmt.Println("file open field,err:", err)
		panic(err)
	}
	defer f.Close()
	rows, err := f.GetRows(sheet)
	if err != nil {
		fmt.Println("file get rows field,err:", err)
		panic(err)
	}
	var cn, en string
	for _, row := range rows[1:] {
		switch sheet {
		case "实体":
			{
				if cn != row[0] && en != row[1] {
					cn = row[0]
					en = row[1]
					entityName[row[1]] = row[0]
					entityData[row[1]] = append(entityData[row[1]], row[2:])
				} else {
					entityData[row[1]] = append(entityData[row[1]], row[2:])
				}
			}
		case "选项":
			{
				if en != row[0] {
					en = row[0] + "_" + row[1]
					optionData[en] = append(optionData[en], row[2:])
				} else {
					optionData[en] = append(optionData[en], row[2:])
				}
			}
		case "关联关系":
			{
				if en != row[0] {
					en = row[0] + "_" + row[1]
					relation[en] = row[2:]
				} else {
					relation[en] = row[2:]
				}
			}
		}
	}
}

// splitTemplate 拼接需要写入的内容
func splitTemplate() []nameAndData {
	n := []nameAndData{}
	for name, datas := range entityData {
		entityDate := ""
		filedDate := ""
		exterData := ""
		option := ""
		nameCn := entityName[name]
		for _, data := range datas {
			switch data[2] {
			case "单选", "多选":
				{
					option = ""
					for _, d := range optionData[name+"_"+data[1]] {
						option += fmt.Sprintf(temp.Selection, d[0], d[1]) + "\n"
					}
					exterData = fmt.Sprintf(temp.Selections, option)
				}
			case "一对一", "多对一", "多对多":
				{
					d := relation[name+"_"+data[1]]
					exterData = fmt.Sprintf(temp.Relation, d[0])
				}
			case "反一对一", "一对多":
				{
					d := relation[name+"_"+data[1]]
					exterData = fmt.Sprintf(temp.ReverseFK, d[0], d[1])
				}
			default:
				{
					exterData = ""
				}
			}
			filedDate += fmt.Sprintf(temp.FieldContent, data[1], temp.FieldType[data[2]], data[0], data[0], exterData) + "\n"
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

// splitName 将文件名字转换成下划线形式
func splitName(name string) string {
	fileName := ""
	for i, x := range name {
		if (65 <= x) && (x <= 90) {
			if i == 0 {
				fileName += string(x + 32)
			} else {
				fileName += "_" + string(x+32)
			}
		} else {
			fileName += string(x)
		}
	}
	return fileName
}
