package dto

type CreateBookRequest struct {
	Name          string `json:"name" binding:"required,alpha,min=1,max=50"`
	NumberOfPages int    `json:"numberOfPages" binding:"max=5000"`
	Description   string `json:"description" binding:"alpha,min=10,max=250"`
	AuthorId      int    `json:"authorId" binding:"required:true"`
}

type UpdateBookRequest struct {
	Name          string `json:"name"`
	NumberOfPages int    `json:"numberOfPages"`
	Description   string `json:"description"`
	AuthorId      int    `json:"authorId"`
}

type BookResponse struct {
	Id            int    `json:"id"`
	Name          string `json:"name"`
	NumberOfPages int    `json:"numberOfPages"`
	Description   string `json:"description,omitempty"`
	AuthorId      int    `json:"authorId"`
}
