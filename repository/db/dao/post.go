package dao

import (
	"context"
	"project/repository/db/model"
	"project/types"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

var _ PostModel = (*customPostModel)(nil)

type (
	// PostModel is an interface to be customized, add more methods here,
	// and implement the added methods in customPostModel.
	PostModel interface {
		JudgeCommunityIsExist(ctx context.Context, cid string) (bool, error)
		CreatePost(ctx context.Context, p *model.Post) error
		GetPostListByIds(ctx context.Context, ids []string) (list []*model.Post, err error)
		GetPostDetailById(ctx context.Context, pid string) (post *model.Post, err error)
		IncrViewCount(ctx context.Context, pid string) error
		UpdatePostEdit(ctx context.Context, p *types.EditPostReq) error
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

func (m *customPostModel) UpdatePostEdit(ctx context.Context, p *types.EditPostReq) error {
	//TODO implement me
	return m.WithContext(ctx).Model(&model.Post{}).
		Where("post_id = ?", p.PostId).
		Updates(&p).Error
}

func (m *customPostModel) IncrViewCount(ctx context.Context, pid string) error {
	//TODO implement me
	return m.WithContext(ctx).Model(&model.Post{}).
		Where("post_id = ?", pid).
		UpdateColumn("view_count", gorm.Expr("view_count + ?", 1)).
		Error
}

// GetPostDetailById 查看帖子详情
func (m *customPostModel) GetPostDetailById(ctx context.Context, pid string) (post *model.Post, err error) {
	//TODO implement me
	err = m.DB.WithContext(ctx).
		Model(&model.Post{}).
		Where("post_id = ?", pid).
		Find(&post).
		Error

	return
}

// GetPostListByIds 通过ids查询帖子列表
func (m *customPostModel) GetPostListByIds(ctx context.Context, ids []string) (list []*model.Post, err error) {
	//TODO implement me

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

	return
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
