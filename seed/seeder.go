package seed

import (
	"goginCasbin/model"
	"goginCasbin/utils"
	"gorm.io/gorm"
	"log"
)

var Users = model.User{
	Name: "admin",
	Email: "admin@mail.com",
	Password: "secret",
	Role: "admin",
}

func Load(db *gorm.DB){
	db.AutoMigrate(&model.User{})
	utils.HashPassword(&Users.Password)
	data := db.Debug().Model(&model.User{}).Create(&Users).Error
	if data != nil {
		log.Fatalf("cannot seed users table: %v", data)
	}
}