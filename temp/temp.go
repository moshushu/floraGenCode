package temp

var PackName = "package %s \n"

var EntityContent = "func init() {\n" +
	"source := h.%s().DeclareModel()\n" +
	"source.SetDisplayName(\"%s\")\n" +
	"source.SetDescription(\"%s\")\n" +
	"source.AddFields(map[string]models.FieldDefinition{\n" +
	"%s\n" +
	"})\n" +
	"}"

var FieldContent = "\"%s\": models.%s{\n" +
	"String:      \"%s\",\n" +
	"Description: \"%s\",\n" +
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
