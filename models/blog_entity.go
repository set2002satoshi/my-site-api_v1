package models

import (
	"time"

	"github.com/set2002satoshi/my-site-api/pkg/module/customs/errors"
	"github.com/set2002satoshi/my-site-api/pkg/module/customs/types"
)

type BlogEntity struct {
	BlogId   types.IDENTIFICATION `gorm:"primaryKey"`
	UserId   types.IDENTIFICATION
	UserName string `gorm:"not null"`
	// UserICON string `gorm:"not null"`
	Title       string                  `gorm:"not null;max:26"`
	Content     string                  `gorm:"not null;max:100"`
	CategoryIds []BlogAndCategoryEntity `gorm:"foreignKey:BlogId"`
	categories  []CategoryEntity        `gorm:"-:migration"`
	Revision    types.REVISION
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func NewBlogEntity(
	blogId int,
	userId int,
	userName string,
	title string,
	content string,
	blogIdAndCategories []BlogAndCategoryEntity,
	categories []CategoryEntity,
	revision int,
	createdAt time.Time,
	updatedAt time.Time,
) (*BlogEntity, error) {
	be := new(BlogEntity)
	var err error
	err = errors.Combine(err, be.setBlogId(blogId))
	err = errors.Combine(err, be.setUserId(userId))
	err = errors.Combine(err, be.setUserName(userName))
	err = errors.Combine(err, be.setTitle(title))
	err = errors.Combine(err, be.setContent(content))
	err = errors.Combine(err, be.setBlogAndCategories(blogIdAndCategories))
	err = errors.Combine(err, be.setCategories(categories))
	err = errors.Combine(err, be.setRevision(revision))
	err = errors.Combine(err, be.setCreatedAt(createdAt))
	err = errors.Combine(err, be.setUpdatedAt(updatedAt))
	if err != nil {
		return new(BlogEntity), err
	}
	return be, err
}

func (be *BlogEntity) GetBlogId() types.IDENTIFICATION {
	return be.BlogId
}

func (be *BlogEntity) GetUserId() types.IDENTIFICATION {
	return be.UserId
}

func (be *BlogEntity) GetUserName() string {
	return be.UserName
}

func (be *BlogEntity) GetTitle() string {
	return be.Title
}

func (be *BlogEntity) GetContent() string {
	return be.Content
}

func (be *BlogEntity) GetBlogAndCategories() []BlogAndCategoryEntity {
	return be.CategoryIds
}

func (be *BlogEntity) GetCategories() []CategoryEntity {
	return be.categories
}

func (be *BlogEntity) GetRevision() types.REVISION {
	return be.Revision
}

func (be *BlogEntity) GetCreatedAt() time.Time {
	return be.CreatedAt
}

func (be *BlogEntity) GetUpdatedAt() time.Time {
	return be.UpdatedAt
}

func (be *BlogEntity) setBlogId(id int) error {
	i, err := types.NewIDENTIFICATION(id)
	if err != nil {
		return err
	}
	be.BlogId = i
	return nil
}

func (be *BlogEntity) setUserId(id int) error {
	i, err := types.NewIDENTIFICATION(id)
	if err != nil {
		return err
	}
	be.UserId = i
	return nil
}

func (be *BlogEntity) setUserName(name string) error {
	be.UserName = name
	return nil
}

func (be *BlogEntity) setTitle(title string) error {
	be.Title = title
	return nil
}

func (be *BlogEntity) setContent(content string) error {
	be.Content = content
	return nil
}

func (be *BlogEntity) setBlogAndCategories(categories []BlogAndCategoryEntity) error {
	be.CategoryIds = categories
	return nil
}

func (be *BlogEntity) setCategories(objs []CategoryEntity) error {
	be.categories = objs
	return nil
}

func (be *BlogEntity) setRevision(revision int) error {
	be.Revision = types.REVISION(revision)
	return nil
}

func (be *BlogEntity) CountUpRevision(currentNum types.REVISION) error {
	if be.Revision != currentNum {
		return errors.Add(errors.NewCustomError(), errors.EN0004)
	}
	if err := be.setRevision(int(currentNum) + 1); err != nil {
		return errors.Wrap(errors.NewCustomError(), errors.EN0005, err.Error())
	}
	return nil
}

func (be *BlogEntity) setCreatedAt(createdAt time.Time) error {
	be.CreatedAt = createdAt
	return nil
}

func (be *BlogEntity) setUpdatedAt(updatedAt time.Time) error {
	be.UpdatedAt = updatedAt
	return nil
}
