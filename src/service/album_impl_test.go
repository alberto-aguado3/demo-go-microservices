package service

import (
	"context"
	"demo/src/model"
	"demo/src/service/mocks"
	"fmt"
	http "net/http"
	"strings"
	"testing"

	"demo/pkg/errors"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestAlbumService_GetAllAlbums(t *testing.T) {
	//Note: if you read documentation of "NewController", it will tell you that you don't need the "defer" statement anymore
	ctrl := gomock.NewController(t)
	//defer ctrl.Finish()

	ar := mocks.NewMockAlbumRepository(ctrl)
	http := mocks.NewMockHttpRepository(ctrl)
	service := NewAlbum(ar, http)

	//Either define in tests, in some other package, read json files...
	mockedAlbums := []model.Album{
		{
			ID:     1,
			Title:  "title1",
			Artist: "artist1",
			Price:  20.4,
		},
		{
			ID:     4,
			Title:  "title4",
			Artist: "artist4",
			Price:  50.7,
		},
	}

	t.Run("Given a valid request, should return a list of albums", func(t *testing.T) {
		ar.EXPECT().FindAllAlbums(gomock.Any()).Return(mockedAlbums, nil)

		actualAlbums, err := service.GetAllAlbums(context.Background())
		assert.NoError(t, err)
		assert.Equal(t, mockedAlbums, actualAlbums)
	})

	t.Run("Given a random database error, should return error", func(t *testing.T) {
		expectedError := errors.NewFromMessage("Random database error")
		ar.EXPECT().FindAllAlbums(gomock.Any()).Return(nil, expectedError)

		actualAlbums, err := service.GetAllAlbums(context.Background())
		assert.ErrorAs(t, expectedError, &err)
		assert.Nil(t, actualAlbums)
	})
}

func TestAlbumService_GetAlbumByArtistAndTitle(t *testing.T) {
	ctrl := gomock.NewController(t)

	ar := mocks.NewMockAlbumRepository(ctrl)
	httpRepo := mocks.NewMockHttpRepository(ctrl)
	service := NewAlbum(ar, httpRepo)

	title := "title1"
	artist := "artist1"
	mockedAlbum := &model.Album{
		ID:     1,
		Title:  title,
		Artist: artist,
		Price:  20.4,
	}

	t.Run("Given a valid album with same artist and title, should return that album", func(t *testing.T) {
		ar.EXPECT().FindAlbumByArtistAndTitle(gomock.Any(), artist, title).Return(*mockedAlbum, nil)

		actualAlbum, err := service.GetAlbumByArtistAndTitle(context.Background(), artist, title)
		assert.Nil(t, err)
		assert.Equal(t, mockedAlbum, actualAlbum)
	})

	t.Run("Given a random database error, should return error", func(t *testing.T) {
		expectedError := errors.NewFromMessage("Random database error")
		ar.EXPECT().FindAlbumByArtistAndTitle(gomock.Any(), artist, title).Return(model.Album{}, expectedError)

		actualAlbum, err := service.GetAlbumByArtistAndTitle(context.Background(), artist, title)
		assert.ErrorAs(t, expectedError, &err)
		assert.Nil(t, actualAlbum)
	})

	t.Run("Given a forbidden artist name, should return error", func(t *testing.T) {
		bannedAuthor := "Quevedo"
		expectedError := errors.NewFromMessage(fmt.Sprintf("tried to request information from banned author: %s", bannedAuthor)).WithStatusCode(http.StatusUnauthorized)

		actualAlbum, err := service.GetAlbumByArtistAndTitle(context.Background(), bannedAuthor, title)
		assert.Equal(t, expectedError, err)
		assert.Nil(t, actualAlbum)
	})

	//Developer decides if an album not found should be an error or not
	t.Run("Given an album not found, should return error", func(t *testing.T) {
		ar.EXPECT().FindAlbumByArtistAndTitle(gomock.Any(), artist, title).Return(model.Album{}, nil)

		actualAlbum, err := service.GetAlbumByArtistAndTitle(context.Background(), artist, title)
		assert.True(t, strings.HasPrefix(err.Error(), "not found"), "Actual returned error: '%s'", err.Error())
		assert.Nil(t, actualAlbum)
	})
}

func TestAlbumService_StoreAlbum(t *testing.T) {
	ctrl := gomock.NewController(t)
	//defer ctrl.Finish()

	ar := mocks.NewMockAlbumRepository(ctrl)
	http := mocks.NewMockHttpRepository(ctrl)
	service := NewAlbum(ar, http)

	//Either define in tests, in some other package, read json files...
	albumToStore := model.Album{
		Title:  "title1",
		Artist: "artist1",
		Price:  20.4,
	}

	t.Run("Given an album, should add to database without problem", func(t *testing.T) {
		ar.EXPECT().StoreAlbum(gomock.Any(), albumToStore).Return(nil)

		err := service.StoreAlbum(context.Background(), albumToStore)
		assert.Nil(t, err)
	})

	t.Run("Given a random database error, should return error", func(t *testing.T) {
		expectedError := errors.NewFromMessage("Random database error")
		ar.EXPECT().StoreAlbum(gomock.Any(), albumToStore).Return(expectedError)

		err := service.StoreAlbum(context.Background(), albumToStore)
		if assert.Error(t, err) {
			//Check for a particular error by content (better if checking by error type)
			assert.Equal(t, "Random database error", err.Error())
		}
	})
}
