package dao

import (
	"context"
	"project/repository/db/model"

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
		GetPostListByIds(ctx context.Context, ids []string) (resp interface{}, err error)
		GetPostDetailById(ctx context.Context, pid string) (resp interface{}, err error)
		IncrViewCount(ctx context.Context, pid string) error
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

func (m *customPostModel) IncrViewCount(ctx context.Context, pid string) error {
	//TODO implement me
	err := m.WithContext(ctx).Model(&model.Post{}).Update("view_count", gorm.Expr("view_count + ?", 1)).Where("pid = ?", pid).Error
	return err
}

// GetPostDetailById 查看帖子详情
func (m *customPostModel) GetPostDetailById(ctx context.Context, pid string) (resp interface{}, err error) {
	//TODO implement me
	var post model.Post
	err = m.DB.WithContext(ctx).
		Model(&model.Post{}).
		Select("post_id,content,like_num,author_id,title,community_id, created_at").
		Where("post_id = ?", pid).
		Find(&post).
		Error

	return post, err
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
