package core

import (
	"context"

	"github.com/Zzocker/bookolab/model"
	"github.com/Zzocker/bookolab/pkg/errors"
	"github.com/Zzocker/bookolab/ports"
)

type bookCore struct {
	bStore ports.BookStore
}

func (b *bookCore) Add(ctx context.Context, in AddBookInput) errors.E {
	return nil // TODO
}
func (b *bookCore) Get(ctx context.Context, isbn string) (*model.Book, errors.E) {
	return nil, nil // TODO
}
func (b *bookCore) Update(ctx context.Context, book model.Book) errors.E {
	return nil // TODO
}
func (b *bookCore) Delete(ctx context.Context, isbn string) errors.E {
	return nil // TODO
}
func (b *bookCore) DeleteAll(ctx context.Context) errors.E {
	return nil // TODO
}
func (b *bookCore) Comment(ctx context.Context, isbn string, content string) errors.E {
	return nil // TODO
}
func (b *bookCore) Rate(ctx context.Context, isbn string, rating uint) errors.E {
	return nil // TODO
}
func (b *bookCore) SearchBookByName(ctx context.Context, name string, pageNumber uint) ([]model.Book, errors.E) {
	return nil, nil // TODO
}
func (b *bookCore) SearchBookByAuthor(ctx context.Context, author string, pageNumber uint) ([]model.Book, errors.E) {
	return nil, nil // TODO
}
func (b *bookCore) SearchBookByGenere(ctx context.Context, genere []string, pageNumber uint) ([]model.Book, errors.E) {
	return nil, nil // TODO
}
func (b *bookCore) GetAllOwned(ctx context.Context, username string) ([]model.Book, errors.E) {
	return nil, nil // TODO
}
func (b *bookCore) GetAllCurrent(ctx context.Context, username string) ([]model.Book, errors.E) {
	return nil, nil // TODO
}
func (b *bookCore) GetAllComment(ctx context.Context, isbn string) ([]model.Comment, errors.E) {
	return nil, nil // TODO
}
