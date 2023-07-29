package dao

import (
	"context"
	"project/consts"
	"project/repository/db/model"

	"gorm.io/gorm"
)

var _ MessageModel = (*customMessageModel)(nil)

type (
	// MessageModel is an interface to be customized, add more methods here,
	// and implement the added methods in customMessageModel.
	MessageModel interface {
		CreateMessage(ctx context.Context, msg *model.Message) error
		ExistsMessage(ctx context.Context, mid int64) bool
		GetMessageListByPage(ctx context.Context, messageType int, uid, page, pageSize int64) ([]*model.Message, error)
		MarkRead(ctx context.Context, userId int64) error
	}

	customMessageModel struct {
		*gorm.DB
	}
)

func (m *customMessageModel) MarkRead(ctx context.Context, userId int64) error {
	//TODO implement me
	return m.WithContext(ctx).Model(&model.Message{}).
		Where("user_id = ? and status = ?", userId, consts.StatusUnread).
		UpdateColumn("status", consts.StatusHaveRead).Error
}

func (m *customMessageModel) GetMessageListByPage(ctx context.Context, messageType int, uid, page, pageSize int64) ([]*model.Message, error) {
	//TODO implement me
	var messages []*model.Message
	limit := int(pageSize)
	offset := int(pageSize * (page - 1))
	err := m.WithContext(ctx).Model(&model.Message{}).
		Where("user_id = ? and type = ?", uid, messageType).
		Limit(limit).Offset(offset).
		Find(&messages).Error

	return messages, err
}

func (m *customMessageModel) ExistsMessage(ctx context.Context, mid int64) bool {
	//TODO implement me
	var count int64
	m.WithContext(ctx).Model(&model.Message{}).Select("message_id").Where("message_id = ?", mid).Count(&count)
	if count != 0 {
		return true
	}
	return false
}

func (m *customMessageModel) CreateMessage(ctx context.Context, msg *model.Message) error {
	//TODO implement me
	return m.WithContext(ctx).Model(&model.Message{}).Create(&msg).Error
}

func NewMessageModel() MessageModel {
	return &customMessageModel{
		DB: NewDBClient(),
	}
}
