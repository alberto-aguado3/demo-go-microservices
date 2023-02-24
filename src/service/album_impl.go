package service

import (
	"context"
	model "demo/src/model"
	"fmt"
	"net/http"
	"time"

	"demo/pkg/errors"

	log "github.com/rs/zerolog/log"
)

var (
	//unexported variable, package scope
	forbiddenNames = map[string]time.Time{
		"Quevedo":       time.Date(2001, 12, 7, 0, 0, 0, 0, time.Local),
		"Oques Grasses": time.Date(2010, 1, 1, 0, 0, 0, 0, time.Local),
	}
)

// example constant
var Pi = 3.14

// exported struct
type AlbumServiceImpl struct {
	//unexported field
	albumRepo  model.AlbumRepository
	httpClient model.HttpRepository
}

func NewAlbum(albumRepo model.AlbumRepository, client model.HttpRepository) AlbumService {
	return &AlbumServiceImpl{
		albumRepo:  albumRepo,
		httpClient: client,
	}
}

// GetAllAlbums - simple call to repository. Retrieve all albums in database. Log if an error happens.
func (a *AlbumServiceImpl) GetAllAlbums(ctx context.Context) ([]model.Album, error) {
	albums, err := a.albumRepo.FindAllAlbums(ctx)
	if err != nil {
		log.Error().Msgf("[AlbumService.GetAllAlbums] - error retrieving albums: %s", err.Error())
		return nil, err
	}

	return albums, nil
}

// GetAlbumsByArtistAndTitle - Return an album with provided artist and title, unless artist name belongs to forbidden names list.
func (a *AlbumServiceImpl) GetAlbumByArtistAndTitle(ctx context.Context, artistName, title string) (*model.Album, *errors.ApiError) {
	log.Trace().Msgf("[AlbumService.GetAlbumsByArtist] - artistName: %s, title: %s", artistName, title)

	//Manually iterate over each entry.
	// for name, _ := range forbiddenNames {
	// 	//if strings.ToLower(name) == strings.ToLower(artistName) {
	// 	if strings.EqualFold(name, artistName) {
	// 		log.Info().Msgf("[AlbumService.GetAlbumsByArtist] - tried to request information from banned author: %s", name)
	// 		return nil, fmt.Errorf("tried to request information from banned author: %s", name)
	// 	}
	// }

	//Alternative: Check quickly if a key exists in map. Cannot ignore string upper/lower casing in this example.
	_, ok := forbiddenNames[artistName]
	if ok {
		log.Info().Msgf("[AlbumService.GetAlbumsByArtist] - tried to request information from banned author: %s", artistName)
		return nil, errors.NewFromMessage(fmt.Sprintf("tried to request information from banned author: %s", artistName)).WithStatusCode(http.StatusUnauthorized)
	}

	album, err := a.albumRepo.FindAlbumByArtistAndTitle(ctx, artistName, title)
	if err != nil {
		log.Error().Msgf("[AlbumService.GetAlbumsByArtist] - error fetching album (%s, %s) from database: %s", artistName, title, err.GetOriginalMessage())

		return nil, err
	}

	if album.Artist == "" || album.Title == "" {
		log.Trace().Msgf("[AlbumService.GetAlbumsByArtist] - album not found")
		return nil, errors.NewFromMessage(fmt.Sprintf("not found album with title: %s and artist: %s", title, artistName)).WithStatusCode(http.StatusNotFound)
	}

	return &album, nil
}

// StoreAlbum - Persist the given album in the database
func (a *AlbumServiceImpl) StoreAlbum(ctx context.Context, album model.Album) *errors.ApiError {
	if err := a.albumRepo.StoreAlbum(ctx, album); err != nil {
		log.Error().Msgf("[AlbumService.StoreAlbum] - database error: %s", err.GetOriginalMessage())
		return err
	}
	return nil
}

func (a *AlbumServiceImpl) GetWeb(ctx context.Context) error {
	var response interface{}

	err := a.httpClient.Get("http://localhost:8080/albums", &response)
	if err != nil {
		log.Info().Msgf("[AlbumService.GetWeb] - err: %s", err.Error())
	}

	log.Info().Msgf("[AlbumService.GetWeb] - response: %+v", response)
	return nil
}
