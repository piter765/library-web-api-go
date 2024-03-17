package services

import (
	"context"
	"library-web-api-go/api/dto"
	"library-web-api-go/config"
	"library-web-api-go/database"
	"library-web-api-go/models"
)

type BookService struct {
	base *BaseService[models.Book, dto.CreateBookRequest, dto.UpdateBookRequest, dto.BookResponse]
}

func NewBookService(cfg *config.Config) *BookService { //config needed for logger
	return &BookService{
		base: &BaseService[models.Book, dto.CreateBookRequest, dto.UpdateBookRequest, dto.BookResponse]{
			Database: database.GetDb(),
			Preloads: []preload{{string: "Author"}},
		},
	}
}

func (s *BookService) Create(ctx context.Context, req *dto.CreateBookRequest) (*dto.BookResponse, error) {
	return s.base.Create(ctx, req)
}

func (s *BookService) Update(ctx context.Context, id int, req *dto.UpdateBookRequest) (*dto.BookResponse, error) {
	return s.base.Update(ctx, id, req)
}

func (s *BookService) Delete(ctx context.Context, id int) error {
	return s.base.Delete(ctx, id)
}

func (s *BookService) GetById(ctx context.Context, id int) (*dto.BookResponse, error) {
	return s.base.GetById(id)
}
