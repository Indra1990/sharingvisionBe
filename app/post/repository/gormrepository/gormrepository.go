package gormrepository

import (
	"errors"
	"sharingvision_backendtest/entity"

	"gorm.io/gorm"
)

type GormRepositoryPost struct {
	db *gorm.DB
}

func NewGormRepositoryPost(db *gorm.DB) *GormRepositoryPost {
	return &GormRepositoryPost{db: db}
}

func (r *GormRepositoryPost) GetPost(limit int, offset int, status string) (ents []*entity.Post, err error) {
	query := r.db.Offset(offset).Limit(limit)
	if status != "" {
		query.Where("status = ?", status)
	}

	query.Find(&ents)
	if query.Error != nil {
		err = query.Error
		return
	}

	if query.RowsAffected == 0 {
		return
	}

	return
}

func (r *GormRepositoryPost) CreatePost(ent *entity.Post) (err error) {
	if err = r.db.Create(&ent).Error; err != nil {
		return
	}
	return
}

func (r *GormRepositoryPost) FindById(idPost uint64) (ent *entity.Post, err error) {
	query := r.db.Find(&ent, idPost)

	if errors.Is(query.Error, gorm.ErrRecordNotFound) {
		err = query.Error
		return
	}

	if query.Error != nil {
		err = query.Error
		return
	}

	if query.RowsAffected == 0 {
		err = gorm.ErrRecordNotFound
		return
	}

	return
}

func (r *GormRepositoryPost) UpdatePost(ent *entity.Post) (err error) {
	if err = r.db.Save(&ent).Error; err != nil {
		return
	}
	return
}

func (r *GormRepositoryPost) DeletePost(postId uint64) (err error) {
	var ent *entity.Post
	if err = r.db.Model(ent).Where("id = ?", postId).Update("status", "trashed").Error; err != nil {
		return
	}

	return
}
