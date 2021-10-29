package response

import (
	"AltaStore/business/category"
)

type AllCategory struct {
	Categories []OneCategory
}

func GetAllCategory(categories *[]category.Category) *AllCategory {
	var allCategory AllCategory
	var oneCategory OneCategory // temporery format

	for _, value := range *categories {
		oneCategory.ID = value.ID
		oneCategory.Code = value.Code
		oneCategory.Name = value.Name
		oneCategory.UpdatedAt = value.UpdatedAt

		allCategory.Categories = append(allCategory.Categories, oneCategory)
	}

	if allCategory.Categories == nil {
		allCategory.Categories = []OneCategory{}
	}

	return &allCategory
}
