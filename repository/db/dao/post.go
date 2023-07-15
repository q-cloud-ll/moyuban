package dao

import (
	"context"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"project/repository/db/model"
)

var _ PostModel = (*customPostModel)(nil)

type (
	// PostModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUserModel.
	PostModel interface {
		JudgeCommunityIsExist(ctx context.Context, cid string) (bool, error)
		CreatePost(ctx context.Context, p *model.Post) error
		GetPostListByIds(ctx context.Context, ids []string) (resp interface{}, err error)
	}

	customPostModel struct {
		*gorm.DB
	}
)

// NewPostModel 初始化postModel
func NewPostModel() PostModel {
	return &customPostModel{
		DB: NewDBClient(),
	}
}

// GetPostListByIds 通过ids查询帖子列表
func (m *customPostModel) GetPostListByIds(ctx context.Context, ids []string) (resp interface{}, err error) {
	//TODO implement me
	var list []*model.Post

	err = m.DB.WithContext(ctx).Model(&model.Post{}).
		Where("post_id in (?)", ids).
		Clauses(clause.OrderBy{
			Expression: clause.Expr{
				SQL:                "FIELD(post_id, ?)",
				Vars:               []interface{}{ids},
				WithoutParentheses: true,
			},
		}).
		Find(&list).
		Error

	return list, err
}

// CreatePost 创建文章
func (m *customPostModel) CreatePost(ctx context.Context, p *model.Post) error {
	//TODO implement me
	return m.DB.WithContext(ctx).
		Model(&model.Post{}).
		Create(&p).
		Error
}

// JudgeCommunityIsExist 判断社区是否存在
func (m *customPostModel) JudgeCommunityIsExist(ctx context.Context, cid string) (bool, error) {
	//TODO implement me
	var total int64
	err := m.DB.WithContext(ctx).
		Model(&model.Community{}).
		Select("community_id").
		Where("community_id = ?", cid).
		Count(&total).
		Error

	if total == 0 {
		return true, err
	} else {
		return false, err
	}
}
