package models

import (
	"github.com/set2002satoshi/my-site-api/pkg/module/customs/errors"
	"github.com/set2002satoshi/my-site-api/pkg/module/customs/types"
	"golang.org/x/crypto/bcrypt"
)

type UserEntity struct {
	UserId   types.IDENTIFICATION `gorm:"primaryKey"`
	UserName string
	Password []byte
	UserRoll types.Roll
	Blogs    []BlogEntity
}

func NewUserEntity(
	userId int,
	userName string,
	password []byte,
	roll int,
) (*UserEntity, error) {
	ue := new(UserEntity)
	var err error
	err = errors.Combine(err, ue.setUserId(userId))
	err = errors.Combine(err, ue.setUserName(userName))
	err = errors.Combine(err, ue.setPassword(userName))
	err = errors.Combine(err, ue.setUserRoll(roll))
	if err != nil {
		return &UserEntity{}, err
	}
	return ue, nil
}

func (ue *UserEntity) GetUserId() types.IDENTIFICATION {
	return ue.UserId
}

func (ue *UserEntity) GetUserName() string {
	return ue.UserName
}

func (ue *UserEntity) GetPassword() string {
	return string(ue.Password)
}

func (ue *UserEntity) GetUserRoll() types.Roll {
	return ue.UserRoll
}

func (ue *UserEntity) setUserId(id int) error {
	i, err := types.NewIDENTIFICATION(id)
	if err != nil {
		return err
	}
	ue.UserId = i
	return nil
}

func (ue *UserEntity) setUserName(userName string) error {
	ue.UserName = userName
	return nil
}

func (ue *UserEntity) setPassword(password string) error {
	pass, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return err
	}
	ue.Password = []byte(pass)
	return nil
}

func (ue *UserEntity) setUserRoll(roll int) error {
	rl, err := types.NewRoll(roll)
	if err != nil {
		return err
	}
	ue.UserRoll = rl
	return nil
}
