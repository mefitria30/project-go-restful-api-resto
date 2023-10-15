package database

import (
	"github.com/kelas.work/project-go-restful-api/internal/model"
	"github.com/kelas.work/project-go-restful-api/internal/model/constant"
	"gorm.io/gorm"
)

func seedDB(db * gorm.DB) {
	// Migrate the schema
	db.AutoMigrate(&model.MenuItem{}, &model.Order{}, &model.ProductOrder{}, &model.User{})

	foodMenu := []model.MenuItem{
		{
			Name:      "Bakmie",
			OrderCode: "bakmie",
			Price:     37500,
			Type:      constant.MenuTypeFood,
		},
		{
			Name:      "Ayam Rica Rica",
			OrderCode: "ayam_rica_rica",
			Price:     41250,
			Type:      constant.MenuTypeFood,
		},
	}

	drinkMenu := []model.MenuItem{
		{
			Name:      "Jus Mangga",
			OrderCode: "jus_mangga",
			Price:     15000,
			Type:      constant.MenuTypeDrink,
		},
		{
			Name:      "Coca Cola",
			OrderCode: "coca_cola",
			Price:     12000,
			Type:      constant.MenuTypeDrink,
		},
		{
			Name:      "Teh Manis",
			OrderCode: "teh_manis",
			Price:     5000,
			Type:      constant.MenuTypeDrink,
		},
		{
			Name:      "Air Mineral",
			OrderCode: "air_mineral",
			Price:     2500,
			Type:      constant.MenuTypeDrink,
		},
	}

	if err := db.First(&model.MenuItem{}).Error; err == gorm.ErrRecordNotFound {
		db.Create(&foodMenu)
		db.Create(&drinkMenu)
	}
}
