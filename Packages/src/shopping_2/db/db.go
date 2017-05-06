package db

import(
	"shopping_2/models"
)

func LoadItem(id int) *models.Item {
	return &models.Item{
		Price: 9.002,
	}
}