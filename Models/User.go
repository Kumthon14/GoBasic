package Models

import (
	"GoPractice/Config"
	"fmt"

	_ "gorm.io/driver/sqlserver"
)

func CreateUser(user *User) (err error) {
	if err = Config.DB.Create(user).Error; err != nil {
		return err
	}
	return nil
}

func GetAllUsers(user *[]User) error {
	if err := Config.DB.Find(user).Error; err != nil {
		return err
	}
	return nil
}

func GetUserById(user *User, id string) error {
	if err := Config.DB.Where("id = ?", id).First(user).Error; err != nil {
		return err
	}
	return nil
}

func UpdateUser(user *User, id string) error {
	fmt.Println(user)
	Config.DB.Save(user)
	return nil
}

func DeleteUser(user *User, id string) error {
	Config.DB.Where("id = ?", id).Delete(user)
	return nil
}

func Login(user *User) error {
	if err := Config.DB.Where(&User{Username: user.Username, Password: user.Password}).First(user).Error; err != nil {
		return err
	}
	return nil
}
