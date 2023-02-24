package mysql

import (
	"context"
	"database/sql"
	model "demo/src/model"
	"net/http"
	"time"

	"demo/pkg/errors"

	sqlbuilder "github.com/huandu/go-sqlbuilder"
)

// MysqlAlbumRepository implements AlbumRepository
type MysqlAlbumRepository struct {
	db      *sql.DB
	timeout time.Duration
}

func NewMysqlAlbumRepository(db *sql.DB, dbTimeout time.Duration) model.AlbumRepository {
	return &MysqlAlbumRepository{
		db:      db,
		timeout: dbTimeout,
	}
}

// FindAllAlbums - Return all albums stored in the database. A vanilla approach, without additional packages (fancier approach is commented out)
// It also returns a standard golang error.
func (r *MysqlAlbumRepository) FindAllAlbums(ctx context.Context) ([]model.Album, error) {
	query := "SELECT * FROM album"

	ctxTimeout, cancel := context.WithTimeout(ctx, r.timeout)
	defer cancel()

	rows, err := r.db.QueryContext(ctxTimeout, query)
	if err != nil {
		return nil, err
	}

	var albums []model.Album
	// sqlUserStruct := sqlbuilder.NewStruct(new(model.Album))
	for rows.Next() {
		var scannedAlbum model.Album
		// if err := rows.Scan(sqlUserStruct.Addr(&scannedAlbum)...); err != nil {
		// 	return nil, err
		// }

		if err := rows.Scan(&scannedAlbum.ID, &scannedAlbum.Artist, &scannedAlbum.Title, &scannedAlbum.Price, &scannedAlbum.Description); err != nil {
			return nil, err
		}

		albums = append(albums, scannedAlbum)
	}

	return albums, nil
}

// FindAlbumByArtistAndTitle - search an album by name of artist and title
func (r *MysqlAlbumRepository) FindAlbumByArtistAndTitle(ctx context.Context, artistName, title string) (model.Album, *errors.ApiError) {
	sqlAlbumStruct := sqlbuilder.NewStruct(new(model.Album))

	// TODO: maybe add "album" table name as a constant
	sb := sqlAlbumStruct.SelectFrom("album")
	query, args := sb.Where(sb.And(sb.Equal("artist", artistName), sb.Equal("title", title))).Build()

	ctxTimeout, cancel := context.WithTimeout(ctx, r.timeout)
	defer cancel()

	row, err := r.db.QueryContext(ctxTimeout, query, args...)
	if err != nil {
		//mask error returned only to consumer (avoid discolsing information about database, for example). They will only see the message "error storing album"
		//and the status code 500. The complete error will still be logged later.
		return model.Album{}, errors.NewFromError(err).WithMessage("error retrieving album").WithStatusCode(http.StatusInternalServerError)
	}

	var album model.Album
	row.Next()
	row.Scan(sqlAlbumStruct.Addr(&album)...)

	return album, nil
}

// StoreAlbum - saves the provided Album into the database
func (r *MysqlAlbumRepository) StoreAlbum(ctx context.Context, album model.Album) *errors.ApiError {
	sqlAlbumStruct := sqlbuilder.NewStruct(new(model.Album))
	query, args := sqlAlbumStruct.InsertInto("album", album).Build()

	ctxTimeout, cancel := context.WithTimeout(ctx, r.timeout)
	defer cancel()

	_, err := r.db.ExecContext(ctxTimeout, query, args...)
	if err != nil {
		return errors.NewFromError(err).WithMessage("error storing album").WithStatusCode(http.StatusInternalServerError)
	}
	return nil
}
