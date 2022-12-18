package temp

var PackName = "package %s \n"

// 实体内容
var EntityContent = "func init() {\n" +
	"source := h.%s().DeclareModel()\n" +
	"source.SetDisplayName(\"%s\")\n" +
	"source.SetDescription(\"%s\")\n" +
	"source.AddFields(map[string]models.FieldDefinition{\n" +
	"%s\n" +
	"})\n" +
	"}"

// 字段内容
var FieldContent = "\"%s\": models.%s{\n" +
	"String:      \"%s\",\n" +
	"Description: \"%s\",\n" +
	"%s\n" +
	"},"

var FieldType = map[string]string{
	"整数":    "IntegerField",
	"布尔":    "BooleanField",
	"日期":    "DateField",
	"文本":    "TextField",
	"单选":    "SelectionField",
	"多选":    "MultiSelectionField",
	"浮点数":   "FloatField",
	"字符串":   "CharField",
	"一对一":   "One2OneField",
	"一对多":   "One2ManyField",
	"多对多":   "Many2ManyField",
	"多对一":   "Many2OneField",
	"反向一对一": "Rev2OneField",
}

// 选项
var Selections = "Selections: []types.SelectionOption{\n" +
	"%s\n" +
	"},"

var Selection = "{Name: \"%s\", Value: \"%s\"},"

// 关联关系
var Relation = "RelationModel: h.%s(),"

var ReverseFK = "RelationModel: h.%s(),\n" +
	"ReverseFK:     \"%s\","
