package services

import (
	"context"
	"library-web-api-go/api/dto"
	"library-web-api-go/config"
	"library-web-api-go/database"
	"library-web-api-go/models"
)

type AuthorService struct {
	base *BaseService[models.Author, dto.CreateAuthorRequest, dto.UpdateAuthorRequest, dto.AuthorResponse]
}

func NewAuthorService(cfg *config.Config) *AuthorService { //config needed for logger
	return &AuthorService{
		base: &BaseService[models.Author, dto.CreateAuthorRequest, dto.UpdateAuthorRequest, dto.AuthorResponse]{
			Database: database.GetDb(),
			Preloads: []preload{{string: "Books"}},
		},
	}
}

func (s *AuthorService) Create(ctx context.Context, req *dto.CreateAuthorRequest) (*dto.AuthorResponse, error) {
	return s.base.Create(ctx, req)
}

func (s *AuthorService) Update(ctx context.Context, id int, req *dto.UpdateAuthorRequest) (*dto.AuthorResponse, error) {
	return s.base.Update(ctx, id, req)
}

func (s *AuthorService) Delete(ctx context.Context, id int) error {
	return s.base.Delete(ctx, id)
}

func (s *AuthorService) GetById(ctx context.Context, id int) (*dto.AuthorResponse, error) {
	return s.base.GetById(id)
}
