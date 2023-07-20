package dao

import (
	"context"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"project/repository/db/model"
	"project/types"
)

var _ CommentModel = (*customCommentModel)(nil)

type (
	// CommentModel is an interface to be customized, add more methods here,
	// and implement the added methods in customCommentModel.
	CommentModel interface {
		CreateComment(ctx context.Context, comment *model.Comment) error
		GetCommentsListByIds(ctx context.Context, ids []string) (commentList []*model.Comment, err error)
		GetUserInfoByCommentId(ctx context.Context, id int64) (user *types.UserCommentInfo, err error)
	}

	customCommentModel struct {
		*gorm.DB
	}
)

func NewCommentModel() CommentModel {
	return &customCommentModel{
		DB: NewDBClient(),
	}
}

func (m *customCommentModel) GetUserInfoByCommentId(ctx context.Context, commentId int64) (user *types.UserCommentInfo, err error) {
	err = m.WithContext(ctx).Model(&model.Comment{}).
		Select("fu.nickname as nickname, fu.avatar as avatar, fu.user_id as user_id").
		Joins("join frm_users fu on fu.user_id = fc.user_id").
		Where("fc.comment_id = ?", commentId).
		Find(&user).Error

	return
}

func (m *customCommentModel) GetCommentsListByIds(ctx context.Context, ids []string) (commentList []*model.Comment, err error) {
	//TODO implement me
	err = m.WithContext(ctx).Model(&model.Comment{}).
		Where("comment_id in (?)", ids).
		Clauses(clause.OrderBy{
			Expression: clause.Expr{
				SQL:                "FIELD(comment_id, ?)",
				Vars:               []interface{}{ids},
				WithoutParentheses: true,
			},
		}).
		Find(&commentList).
		Error

	return
}

func (m *customCommentModel) CreateComment(ctx context.Context, comment *model.Comment) error {
	//TODO implement me
	return m.WithContext(ctx).Model(&model.Comment{}).Create(&comment).Error
}
