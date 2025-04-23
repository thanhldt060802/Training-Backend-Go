package dto

import "training-project/internal/model"

type CategoryResDTO struct {
	Id          int64  `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type CategoryCreateReqDTO struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type CategoryUpdateReqDTO struct {
	Name        *string `json:"name"`
	Description *string `json:"description"`
}

type CategoryDeleteReqDTO struct {
	Id int64 `json:"id"`
}

func ToCategoryResDTO(category *model.Category) *CategoryResDTO {
	return &CategoryResDTO{
		Id:          category.Id,
		Name:        category.Name,
		Description: category.Description,
	}
}

func ToCategoryResDTOs(categories []model.Category) []CategoryResDTO {
	categoryResDTOs := make([]CategoryResDTO, len(categories))
	for i, category := range categories {
		categoryResDTOs[i] = *ToCategoryResDTO(&category)
	}
	return categoryResDTOs
}
