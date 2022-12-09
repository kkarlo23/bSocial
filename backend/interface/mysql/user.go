package mysql

import (
	"bSocial/domain"
	"bSocial/helpers"
	"errors"
)

// check if there is user with same email or username, if there is, user cannot be created
func ableToCreateUser(userData domain.User) (bool, error) {
	var count int64
	if result := MySql.Model(&userData).Where("username = ? OR email = ?", userData.Username, userData.Email).Count(&count); result.Error != nil {
		return false, result.Error
	}
	if count > 0 {
		return false, nil
	}

	return true, nil
}

func GetUserByUsernameAndPassword(username, password string) (*domain.User, error) {
	var users []domain.User
	var user *domain.User

	if result := MySql.Where("username = ?", username).Limit(1).Find(&users).Debug(); result.Error != nil {
		return nil, result.Error
	}
	if len(users) == 0 {
		return nil, errors.New("wrong username or password")
	}
	user = &users[0]
	isCorrectPassword := helpers.CheckPasswordHash(password, user.Password)
	if !isCorrectPassword {
		return nil, errors.New("wrong username or password")
	}
	return user, nil
}

func GetUsers() ([]domain.User, error) {
	var users []domain.User
	if result := MySql.Find(&users); result.Error != nil {
		return nil, result.Error
	}

	return users, nil
}

func CreateUser(userData domain.User) (*domain.User, error) {
	if isAble, err := ableToCreateUser(userData); err != nil || !isAble {
		if !isAble {
			return nil, errors.New("user already in use")
		}
		return nil, err
	}

	if resultCreate := MySql.Create(&userData); resultCreate.Error != nil {
		return nil, resultCreate.Error
	}

	var user domain.User

	// need to fetch user again because CreatedAt have no value after create
	if resultSelect := MySql.Where("id = ?", userData.ID).First(&user); resultSelect.Error != nil {
		return nil, resultSelect.Error
	}

	return &user, nil
}

func UserFollow(userID, toFollowID uint) error {
	userFollower := domain.UserFollower{FollowerID: userID, FollowingID: toFollowID}

	if resultCreate := MySql.Create(&userFollower); resultCreate.Error != nil {
		return resultCreate.Error
	}

	return nil
}
