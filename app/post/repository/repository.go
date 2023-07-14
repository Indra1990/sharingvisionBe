package repository

import "sharingvision_backendtest/entity"

type PostRepository interface {
	GetPost(limit int, offset int, status string) (ents []*entity.Post, err error)
	CreatePost(ent *entity.Post) (err error)
	FindById(idPost uint64) (ent *entity.Post, err error)
	UpdatePost(ent *entity.Post) (err error)
	DeletePost(postId uint64) (err error)
}
