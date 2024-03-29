package services

import (
	"restfulalta/part-4-middleware/config"
	"restfulalta/part-4-middleware/models"
	"restfulalta/part-4-middleware/util/jwt"
)

func GetUsers() ([]models.UserAPI, error) {
	var users []models.UserAPI

	res := config.Db.Model(&models.User{}).Find(&users)

	if res.Error != nil {
		return nil, res.Error
	}
	return users, nil
}


func GetUserById(targetId int) (models.UserAPI, int, error) {
	var user models.UserAPI

	res := config.Db.Model(&models.User{}).Find(&user, targetId)

	if res.Error != nil {
		return models.UserAPI{}, 0, res.Error
	}

	if res.RowsAffected == 0 {
		return models.UserAPI{}, 0, nil
	}

	return user, 1, nil
}

func AddUser(newUser *models.User) (models.UserAPI, error) {
	res := config.Db.Create(newUser)
	if res.Error != nil {
		return models.UserAPI{}, res.Error
	}
	newUserAPI := models.UserAPI{}
	newUserAPI.Id = newUser.Id
	newUserAPI.Name = newUser.Name
	newUserAPI.Email = newUser.Email
	
	return newUserAPI, nil
}

func EditUser(newData models.User, targetId int) (models.UserAPI ,int, error) {
	targetUser := models.User{}

	res := config.Db.Where(`id = ?`, targetId).Find(&targetUser).Omit("password", "id").Updates(newData)

	if res.Error != nil {
		return models.UserAPI{}, 0, res.Error
	}

	if res.RowsAffected == 0 {
		return models.UserAPI{}, 0, nil
	}

	edittedUser := models.UserAPI{}
	edittedUser.Id = targetUser.Id
	edittedUser.Name = targetUser.Name
	edittedUser.Email = targetUser.Email

	return edittedUser, 1, nil
}

func DeleteUser(targetId int) (int, error) {	
	targetUser := models.User{}
	res := config.Db.Find(&targetUser, targetId)

	if res.Error != nil {
		return 0, res.Error
	}

	if res.RowsAffected == 0 {
		return 0, nil
	}

	res = config.Db.Unscoped().Delete(&targetUser)

	if res.Error != nil {
		return 0, res.Error
	}

	if res.RowsAffected == 0 {
		return 0, nil
	}

	return 1, nil
}

func LoginUser(user *models.User) (string ,error) {
	res := config.Db.Where("email = ? AND password = ?", user.Email, user.Password).First(user)

	if res.Error != nil {
		return "", res.Error
	}

	token, err := implementjwt.CreateToken(user.Id)

	if err != nil {
		return "", err
	}

	return token, nil

}