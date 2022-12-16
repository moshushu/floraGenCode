package StuentInfo

import "ma.applysquare.net/eng/flora/pkg/core/models"

func init() {
	source := h.StuentInfo().DeclareModel()
	source.SetDisplayName("学生信息")
	source.SetDescription("学生信息")
	source.AddFields(map[string]models.FieldDefinition{

		"Name": models.字符串{
			String:      "Name",
			Description: "Name",
		},
		"Name": models.字符串{
			String:      "Name",
			Description: "Name",
		},
		"Age": models.字符串{
			String:      "Age",
			Description: "Age",
		},
	})
}
