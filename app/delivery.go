package app

import (
	"sharingvision_backendtest/app/post"
	"sharingvision_backendtest/app/post/repository/gormrepository"
	"sharingvision_backendtest/app/post/usecase/usecaseimpl"

	"gorm.io/gorm"
)

func deliveryPostApp(db *gorm.DB) *post.PostHttpRouterRegistry {
	repoPost := gormrepository.NewGormRepositoryPost(db)
	impPost := usecaseimpl.NewPostImpl(repoPost)
	postRoute := post.NewPostHttpRouterRegistry(impPost)
	return postRoute
}
