package model

import (
	"context"
	"time"
	"tticket/pkg/dal"
)

type User struct {
	ID          int64
	Name        string
	Mail        string
	CreatedTime time.Time `gorm:"-"`
	UpdatedTime time.Time `gorm:"-"`
	DeletedTime time.Time `gorm:"-"`
}

func (b *User) TableName() string {
	return "user"
}

func FindUsers(ctx context.Context) ([]*User, error) {
	res := make([]*User, 0)
	if err := dal.DB.Model(&User{}).Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}
