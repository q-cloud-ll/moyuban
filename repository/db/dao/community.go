package dao

import (
	"context"
	"gorm.io/gorm"
	"project/consts"
	"project/repository/db/model"
)

var _ CommunityModel = (*customCommunityModel)(nil)

type (
	// CommunityModel is an interface to be customized, add more methods here,
	// and implement the added methods in customCommunityModel.
	CommunityModel interface {
		GetCommunityDetailById(ctx context.Context, id int64) (cd *model.Community, err error)
	}

	customCommunityModel struct {
		*gorm.DB
	}
)

func NewCommunity() CommunityModel {
	return &customCommunityModel{
		DB: NewDBClient(),
	}
}

func (m *customCommunityModel) GetCommunityDetailById(ctx context.Context, id int64) (cd *model.Community, err error) {
	//TODO implement me
	if err := m.DB.WithContext(ctx).
		Model(&model.Community{}).
		Where("community_id = ?", id).
		Find(&cd).
		Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			err = consts.InvalidIDErr
		}
	}

	return
}
