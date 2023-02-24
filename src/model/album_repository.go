package model

import (
	"context"
	"demo/pkg/errors"
)

type AlbumRepository interface {
	FindAllAlbums(ctx context.Context) ([]Album, error) //This option returns a normal golang error
	FindAlbumByArtistAndTitle(ctx context.Context, artistName, title string) (Album, *errors.ApiError)
	StoreAlbum(ctx context.Context, album Album) *errors.ApiError
	//These other options return a custom defined error, with custom methods.
	//It forces them to return an object of type ApiError (a struct), so for example,
	//you can't return directly an error of the sql library: you must, instead, wrap that error in a new ApiError struct.
}
