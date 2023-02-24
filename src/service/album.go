package service

import (
	"context"

	"demo/pkg/errors"
	"demo/src/model"
)

type AlbumService interface {
	GetAllAlbums(ctx context.Context) ([]model.Album, error)
	GetAlbumByArtistAndTitle(ctx context.Context, artistName string, title string) (*model.Album, *errors.ApiError)
	StoreAlbum(ctx context.Context, album model.Album) *errors.ApiError
	GetWeb(ctx context.Context) error
}
