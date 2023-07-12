package userRepository

import (
	"svetozar12/e-com/v2/apps/services/user/internal/app/databases/postgres"
	"svetozar12/e-com/v2/apps/services/user/internal/app/entities"
)

func GetUser(userId string) *entities.UserEntity {
	user := new(entities.UserEntity)
	postgres.DB.Where("id = ?", userId).First(user)
	return user
}

func getUserList(userIds []string) []entities.UserEntity {
	users := []entities.UserEntity{}
	postgres.DB.Where("id in (?)", userIds).Find(&users)
	return users
}

func CreateUser(user *entities.UserEntity) *entities.UserEntity {
	postgres.DB.Create(user)
	return user
}

func UpdateUser(userId string, user *entities.UserEntity) *entities.UserEntity {
	postgres.DB.Save(user)
	return user
}

func DeleteUser(userId string) *entities.UserEntity {
	user := new(entities.UserEntity)
	postgres.DB.Where("id = ?", userId).Delete(user)
	return user
}