package post

import (
	"net/http"
	"sharingvision_backendtest/app/post/usecase"
	"strconv"
	"strings"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

type PostHttpRouterRegistry struct {
	postService usecase.PostService
}

func NewPostHttpRouterRegistry(postService usecase.PostService) *PostHttpRouterRegistry {
	return &PostHttpRouterRegistry{
		postService: postService,
	}
}

func (p *PostHttpRouterRegistry) Articles(w http.ResponseWriter, r *http.Request) {

	limit, _ := strconv.Atoi(r.URL.Query().Get("limit"))
	offset, _ := strconv.Atoi(r.URL.Query().Get("offset"))

	paging := usecase.Pagination{
		Limit:  limit,
		Offset: offset,
		Status: r.URL.Query().Get("status"),
	}

	articles, err := p.postService.GetArticles(paging)
	if err != nil {
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, errorResponse(err))
		return
	}

	response := map[string]interface{}{
		"data":   articles,
		"limit":  limit,
		"offset": offset,
	}

	render.Status(r, http.StatusOK)
	render.JSON(w, r, response)
}

func (p *PostHttpRouterRegistry) CreatePost(w http.ResponseWriter, r *http.Request) {
	var requestArticle usecase.ArticlesRequest
	if err := render.DecodeJSON(r.Body, &requestArticle); err != nil {
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, errorResponse(err))
		return
	}

	if err := p.postService.PostCreate(requestArticle); err != nil {
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, errorResponse(err))
		return
	}
	render.Status(r, http.StatusOK)
	render.JSON(w, r, "article created")
}

func (p *PostHttpRouterRegistry) FindPostDetail(w http.ResponseWriter, r *http.Request) {
	postId := chi.URLParam(r, "id")
	parsePostId, parsePostIdErr := strconv.ParseUint(postId, 10, 64)

	if parsePostIdErr != nil {
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, errorResponse(parsePostIdErr))
		return
	}

	dtoPost, dtoPostErr := p.postService.FindById(parsePostId)
	if dtoPostErr != nil {
		if strings.Contains(dtoPostErr.Error(), "record not found") {
			render.Status(r, http.StatusNotFound)
			render.JSON(w, r, errorResponse(dtoPostErr))
			return
		}
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, errorResponse(dtoPostErr))
		return
	}
	render.Status(r, 200)
	render.JSON(w, r, dtoPost)
}

func (p *PostHttpRouterRegistry) UpddatePost(w http.ResponseWriter, r *http.Request) {
	var requestArticle usecase.ArticlesRequest
	postId := chi.URLParam(r, "id")
	parsePostId, parsePostIdErr := strconv.ParseUint(postId, 10, 64)

	if parsePostIdErr != nil {
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, errorResponse(parsePostIdErr))
		return
	}

	if err := render.DecodeJSON(r.Body, &requestArticle); err != nil {
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, errorResponse(err))
		return
	}

	if err := p.postService.PostUpdate(parsePostId, requestArticle); err != nil {
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, errorResponse(err))
		return
	}

	render.Status(r, 200)
	render.JSON(w, r, "updated article")
}

func (p *PostHttpRouterRegistry) DeletePost(w http.ResponseWriter, r *http.Request) {
	postId := chi.URLParam(r, "id")
	parsePostId, parsePostIdErr := strconv.ParseUint(postId, 10, 64)

	if parsePostIdErr != nil {
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, errorResponse(parsePostIdErr))
		return
	}

	if err := p.postService.PostDelete(parsePostId); err != nil {
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, errorResponse(err))
		return
	}

	render.Status(r, 200)
	render.JSON(w, r, "deleted article")
}

func errorResponse(err error) map[string]interface{} {
	errRspn := map[string]interface{}{
		"error": err.Error(),
	}
	return errRspn
}
