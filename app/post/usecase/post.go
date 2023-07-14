package usecase

type PostService interface {
	GetArticles(cmd Pagination) (dto []*ArticlesResponse, err error)
	PostCreate(cmd ArticlesRequest) (err error)
	FindById(idPost uint64) (dto *ArticlesResponse, err error)
	PostUpdate(idPost uint64, cmd ArticlesRequest) (err error)
	PostDelete(dPost uint64) (err error)
}
