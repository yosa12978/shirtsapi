package dtos

import (
	"errors"
	"strings"

	"github.com/yosa12978/MyShirts/internal/models"
)

type ShirtCreateDTO struct {
	Color   string `json:"color"`
	Size    string `json:"size"`
	Pattern string `json:"pattern"`
	Photo   string `json:"photo"`
}

func (scd *ShirtCreateDTO) Map() (models.Shirt, error) {
	if strings.Replace(scd.Color, " ", "", -1) != "" {
		return models.Shirt{}, errors.New("color can't be empty ")
	}
	if strings.Replace(scd.Size, " ", "", -1) != "" {
		return models.Shirt{}, errors.New("size can't be empty ")
	}
	if strings.Replace(scd.Pattern, " ", "", -1) != "" {
		return models.Shirt{}, errors.New("pattern can't be empty ")
	}

	return models.Shirt{
		Color:   scd.Color,
		Size:    scd.Size,
		Pattern: scd.Pattern,
		Photo:   scd.Photo,
	}, nil
}

type ShirtUpdateDTO struct {
	Color   string `json:"color"`
	Size    string `json:"size"`
	Pattern string `json:"pattern"`
	Photo   string `json:"photo"`
}

func (sud *ShirtUpdateDTO) Map(parent *models.Shirt) *models.Shirt {
	if sud.Color != "" {
		parent.Color = sud.Color
	}
	if sud.Size != "" {
		parent.Size = sud.Size
	}
	if sud.Pattern != "" {
		parent.Pattern = sud.Pattern
	}
	if sud.Photo != "" {
		parent.Photo = sud.Photo
	}
	return parent
}
