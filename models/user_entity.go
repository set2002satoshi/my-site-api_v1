package models

import (
	"time"

	"github.com/set2002satoshi/my-site-api/pkg/module/customs/errors"
	"github.com/set2002satoshi/my-site-api/pkg/module/customs/types"
	"golang.org/x/crypto/bcrypt"
)

type UserEntity struct {
	UserId    types.IDENTIFICATION `gorm:"primaryKey"`
	Email     string               `gorm:"unique;not null"`
	UserName  string               `gorm:"not null;max:16"`
	Password  []byte               `gorm:"not null;max:32"`
	Roll      types.AccessROLL
	Blogs     []BlogEntity `gorm:"foreignKey:UserId"`
	Revision  types.REVISION
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewUserEntity(
	userId int,
	Email string,
	userName,
	password string,
	roll string,
	revision int,
	createdAt,
	updatedAt time.Time,
) (*UserEntity, error) {
	ue := new(UserEntity)
	var err error
	err = errors.Combine(err, ue.setUserId(userId))
	err = errors.Combine(err, ue.setEmail(Email))
	err = errors.Combine(err, ue.setUserName(userName))
	err = errors.Combine(err, ue.setPassword(userName))
	err = errors.Combine(err, ue.setRoll(roll))
	err = errors.Combine(err, ue.setRevision(revision))
	err = errors.Combine(err, ue.setCreatedAt(createdAt))
	err = errors.Combine(err, ue.setUpdatedAt(updatedAt))
	if err != nil {
		return &UserEntity{}, err
	}
	return ue, nil
}

func (ue *UserEntity) GetUserId() types.IDENTIFICATION {
	return ue.UserId
}

func (ue *UserEntity) GetEmail() string {
	return ue.Email
}

func (ue *UserEntity) GetUserName() string {
	return ue.UserName
}

func (ue *UserEntity) GetPassword() string {
	return string(ue.Password)
}

func (ue *UserEntity) GetRoll() types.AccessROLL {
	return ue.Roll
}

func (ue *UserEntity) GetRevision() types.REVISION {
	return ue.Revision
}

func (ue *UserEntity) GetCreatedAt() time.Time {
	return ue.CreatedAt
}

func (ue *UserEntity) GetUpdatedAt() time.Time {
	return ue.UpdatedAt
}

func (ue *UserEntity) setUserId(id int) error {
	i, err := types.NewIDENTIFICATION(id)
	if err != nil {
		return errors.Wrap(errors.NewCustomError(), errors.EN0001, err.Error())
	}
	ue.UserId = i
	return nil
}

func (ue *UserEntity) setEmail(email string) error {
	ue.Email = email
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

func (ue *UserEntity) setRoll(roll string) error {
	rl, err := types.NewAccessROLL(roll)
	if err != nil {
		return errors.Wrap(errors.NewCustomError(), errors.EN0003, err.Error())
	}
	ue.Roll = rl
	return nil
}

func (ue *UserEntity) setRevision(revision int) error {
	ue.Revision = types.REVISION(revision)
	return nil
}

func (ue *UserEntity) CountUpRevision(currentNum types.REVISION) error {
	if ue.Revision != currentNum {
		return errors.Add(errors.NewCustomError(), errors.EN0004)
	}
	if err := ue.setRevision(int(currentNum) + 1); err != nil {
		return errors.Wrap(errors.NewCustomError(), errors.EN0005, err.Error())
	}
	return nil
}

func (ue *UserEntity) setCreatedAt(createdAt time.Time) error {
	ue.CreatedAt = createdAt
	return nil
}

func (ue *UserEntity) setUpdatedAt(updatedAt time.Time) error {
	ue.UpdatedAt = updatedAt
	return nil
}
