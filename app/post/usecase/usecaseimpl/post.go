package usecaseimpl

import (
	"sharingvision_backendtest/app/post/repository"
	"sharingvision_backendtest/app/post/usecase"
	"sharingvision_backendtest/entity"
	"time"
)

type PostService struct {
	postRepository repository.PostRepository
}

func NewPostImpl(repo repository.PostRepository) *PostService {
	return &PostService{repo}
}

func (p *PostService) GetArticles(cmd usecase.Pagination) (dto []*usecase.ArticlesResponse, err error) {

	// if err = articleGetValidate(cmd); err != nil {
	// 	return
	// }

	ents, entsErr := p.postRepository.GetPost(cmd.Limit, cmd.Offset, cmd.Status)
	if entsErr != nil {
		err = entsErr
		return
	}

	article := make([]*usecase.ArticlesResponse, len(ents))
	for i, ent := range ents {
		var updateAt string
		if !ent.UpdatedDate.IsZero() {
			updateAt = ent.UpdatedDate.Format("2006-01-02 15:04:05")
		}

		post := &usecase.ArticlesResponse{
			ID:        ent.ID,
			Title:     ent.Title,
			Content:   ent.Content,
			Category:  ent.Category,
			Status:    ent.Status,
			CreatedAt: ent.CreatedDate.Format("2006-01-02 15:04:05"),
			UpdateAt:  updateAt,
		}
		article[i] = post
	}

	dto = article
	return
}

func (p *PostService) PostCreate(cmd usecase.ArticlesRequest) (err error) {
	if err = articleCreatedValidate(cmd); err != nil {
		return
	}

	post := entity.Post{
		Title:       cmd.Title,
		Content:     cmd.Content,
		Category:    cmd.Category,
		Status:      cmd.Status,
		CreatedDate: time.Now(),
	}

	if err = p.postRepository.CreatePost(&post); err != nil {
		return
	}

	return
}

func (p *PostService) FindById(idPost uint64) (dto *usecase.ArticlesResponse, err error) {
	article, articleErr := p.postRepository.FindById(idPost)
	if articleErr != nil {
		err = articleErr
		return
	}

	var updateAt string
	if !article.UpdatedDate.IsZero() {
		updateAt = article.UpdatedDate.Format("2006-01-02 15:04:05")
	}

	dto = &usecase.ArticlesResponse{
		ID:        article.ID,
		Title:     article.Title,
		Content:   article.Content,
		Category:  article.Category,
		Status:    article.Status,
		CreatedAt: article.CreatedDate.Format("2006-01-02 15:04:05"),
		UpdateAt:  updateAt,
	}

	return
}

func (p *PostService) PostUpdate(idPost uint64, cmd usecase.ArticlesRequest) (err error) {
	if err = articleCreatedValidate(cmd); err != nil {
		return
	}

	post := entity.Post{
		ID:          idPost,
		Title:       cmd.Title,
		Content:     cmd.Content,
		Category:    cmd.Category,
		Status:      cmd.Status,
		UpdatedDate: time.Now(),
	}

	if err = p.postRepository.UpdatePost(&post); err != nil {
		return
	}

	return
}

func (p *PostService) PostDelete(idPost uint64) (err error) {
	if err = p.postRepository.DeletePost(idPost); err != nil {
		return
	}

	return
}
