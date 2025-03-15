package service

import (
	"errors"
	"github.com/mazezen/itools"
	"github.com/mazezen/zenchat/common/sdk"
	"github.com/mazezen/zenchat/internel/dao"
	"github.com/mazezen/zenchat/internel/models"
	"time"
)

type UserService struct{}

func NewUserService() *UserService {
	return new(UserService)
}

func (u *UserService) Register(username, password, repeatPassword, email, phone, identity string) error {
	user, err := dao.NewUserDao().FindUser(username)
	if err != nil {
		return err
	}
	if password != repeatPassword {
		return errors.New("two passwords are inconsistent")
	}
	sha256Pwd := itools.Sha256(password)
	user.Username = username
	user.Password = sha256Pwd
	user.Email = email
	user.Phone = phone
	user.Identity = identity
	t := time.Now()
	user.LoginTime = t
	user.LoginOutTime = t
	user.HearBeatTime = t
	return dao.NewUserDao().CreateUser(user)
}

func (u *UserService) LgnPwd(username, password string) (string, error) {
	user, err := dao.NewUserDao().FindUserByName(username)
	if err != nil {
		return "", err
	}
	if user.Password != itools.Sha256(password) {
		return "", errors.New("wrong password")
	}
	var lgnPwdInfo struct {
		Id       uint64 `json:"id"`
		Username string `json:"username"`
		Password string `json:"password"`
		Email    string `json:"email"`
		Phone    string `json:"phone"`
	}
	lgnPwdInfo.Id = user.ID
	lgnPwdInfo.Username = user.Username
	lgnPwdInfo.Password = itools.Sha256(password)
	lgnPwdInfo.Email = user.Email
	lgnPwdInfo.Phone = user.Phone
	ex := time.Duration(sdk.GetConf().Jwt.Expire) * time.Second
	newJwt := itools.NewJwt(ex, sdk.GetConf().Jwt.Secret)
	token, err := newJwt.GenerateToken(lgnPwdInfo)
	if err != nil {
		return "", err
	}
	return token, nil
}

func (u *UserService) Update(
	username,
	password,
	repeatPassword,
	email,
	phone,
	identity,
	avatar,
	gender string) error {
	user, err := dao.NewUserDao().FindUserByPhone(phone)
	if err != nil {
		return err
	}
	if password != repeatPassword {
		return errors.New("two passwords are inconsistent")
	}
	sha256Pwd := itools.Sha256(password)
	user.Username = username
	user.Password = sha256Pwd
	user.Email = email
	user.Phone = phone
	user.Identity = identity
	user.Avatar = avatar
	if gender != "" {
		user.Gender = gender
	}
	return dao.NewUserDao().UpdateUser(user)
}

func (u *UserService) Delete(id uint64) error {
	return dao.NewUserDao().DeleteUser(id)
}

func (u *UserService) List() (int64, []*models.CUser, error) {
	return dao.NewUserDao().GetUserList()
}
