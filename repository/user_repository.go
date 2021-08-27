package repository

import (
	"goginCasbin/model"
	"gorm.io/gorm"
)


type userRepository struct {
	DB *gorm.DB
}

type UserRepository interface {
	AddUser(model.User) (model.User, error)
	GetUser(int) (model.User, error)
	GetByEmail(string) (model.User, error)
	GetAllUser() ([]model.User, error)
	UpdateUser(model.User) (model.User, error)
	DeleteUser(model.User) (model.User, error)
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return userRepository{
		DB: db,
	}
}

func (u userRepository) GetUser(id int) (user model.User, err error) {
	return user, u.DB.First(&user, id).Error
}

func (u userRepository) GetAllUser() (users []model.User, err error) {
	return users, u.DB.Find(&users).Error
}

func (u userRepository) GetByEmail(email string) (user model.User, err error) {
	//return user, u.DB.Where(&user, "email=?", email).Error
	//return user, u.DB.Find(&user, "email = ?", email).Order("name email password").Error
	return user, u.DB.First(&user, "email=?", email).Error
}

func (u userRepository) AddUser(user model.User) (model.User, error) {
	return user, u.DB.Model(&user).Create(&user).Error
}

func (u userRepository) UpdateUser(user model.User) (model.User , error) {
	return user, u.DB.Save(&user).Error
}

func (u userRepository) DeleteUser(user model.User) (model.User, error) {
	if err := u.DB.First(&user, user.ID).Error; err != nil {
		return user, err
	}
	return user, u.DB.Delete(&user).Error
}

//func (u userRepository) Migrate() error {
//	log.Print("[UserRepository]...Migrate")
//	return u.DB.AutoMigrate(&model.User{})
//}
//
//func (u userRepository) AdminUser() (user model.User, err error) {
//	user.Name = "admin"
//	user.Email = "admin@mail.com"
//	user.Password = "secret"
//	user.Role = "admin"
//	utils.HashPassword(&user.Password)
//
//	return user, u.DB.Model(&user).Create(&user).Error
//}
//
