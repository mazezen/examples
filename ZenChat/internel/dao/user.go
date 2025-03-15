package dao

import (
	"fmt"
	"github.com/mazezen/itools"
	"github.com/mazezen/zenchat/internel/models"
)

type UserDao struct{}

func NewUserDao() *UserDao {
	return new(UserDao)
}

func (u *UserDao) GetUserList() (int64, []*models.CUser, error) {
	var list []*models.CUser
	count, err := itools.Db.FindAndCount(&list)
	if err != nil {
		return 0, nil, err
	}
	return count, list, nil
}

func (u *UserDao) FindUserByNameAndPwd(username, pwd string) (*models.CUser, error) {
	var user = new(models.CUser)
	has, err := itools.Db.Where("username = ?", username).And("pwd = ?", pwd).Get(user)
	if err != nil {
		return nil, err
	}
	if !has {
		return nil, fmt.Errorf("user not found")
	}
	return user, nil
}

func (u *UserDao) FindUserByName(username string) (*models.CUser, error) {
	var user = new(models.CUser)
	has, err := itools.Db.Where("username = ?", username).Get(user)
	if err != nil {
		return nil, err
	}
	if !has {
		return nil, fmt.Errorf("user not found")
	}
	return user, nil
}

func (u *UserDao) FindUser(username string) (*models.CUser, error) {
	var user = new(models.CUser)
	has, err := itools.Db.Where("username = ?", username).Get(user)
	if err != nil {
		return nil, err
	}
	if has {
		return nil, fmt.Errorf("user is exists")
	}
	return user, nil
}

func (u *UserDao) FindUserById(id uint64) (*models.CUser, error) {
	var user = new(models.CUser)
	has, err := itools.Db.Where("id = ?", id).Get(user)
	if err != nil {
		return nil, err
	}
	if !has {
		return nil, fmt.Errorf("user not found")
	}
	return user, nil
}

func (u *UserDao) FindUserByEmail(email string) (*models.CUser, error) {
	var user = new(models.CUser)
	has, err := itools.Db.Where("email = ?", email).Get(user)
	if err != nil {
		return nil, err
	}
	if !has {
		return nil, fmt.Errorf("user not found")
	}
	return user, nil
}

func (u *UserDao) FindUserByPhone(phone string) (*models.CUser, error) {
	var user = new(models.CUser)
	has, err := itools.Db.Where("phone = ?", phone).Get(user)
	if err != nil {
		return nil, err
	}
	if !has {
		return nil, fmt.Errorf("user not found")
	}
	return user, nil
}

func (u *UserDao) CreateUser(user *models.CUser) error {
	affected, err := itools.Db.InsertOne(user)
	if err != nil {
		return err
	}
	if affected == 0 {
		return fmt.Errorf("create user failed")
	}
	return nil
}

func (u *UserDao) UpdateUser(user *models.CUser) error {
	affected, err := itools.Db.ID(user.ID).Update(user)
	if err != nil {
		return err
	}
	if affected == 0 {
		return fmt.Errorf("update user failed")
	}
	return nil
}

func (u *UserDao) DeleteUser(id uint64) error {
	_, err := u.FindUserById(id)
	if err != nil {
		return err
	}
	affected, err := itools.Db.ID(id).Delete(new(models.CUser))
	if err != nil {
		return err
	}
	if affected == 0 {
		return fmt.Errorf("delete user failed")
	}
	return nil
}
