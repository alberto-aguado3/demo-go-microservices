package handler

import (
	"demo/src/model"
	"demo/src/service"
	"demo/src/service/mocks"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func a() {
	var t *testing.T
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockAlbumRepository(ctrl)
	mockHttp := mocks.NewMockHttpRepository(ctrl)

	service := service.NewAlbum(mockRepo, mockHttp)

	gin.SetMode(gin.TestMode)
	router := gin.New()
	router.GET("/albums", ListAlbums(service))

	t.Run("A test description", func(t *testing.T) {
		mockRepo.EXPECT().FindAllAlbums(gomock.Any()).Return([]model.Album{}, nil)

		url := fmt.Sprintf("/albums")
		req := httptest.NewRequest(http.MethodGet, url, nil) //no body

		rec := httptest.NewRecorder()
		router.ServeHTTP(rec, req)

		res := rec.Result()
		defer res.Body.Close()

		var body []model.Album
		responseBytes, err := ioutil.ReadAll(res.Body)
		require.NoError(t, err)
		err = json.Unmarshal(responseBytes, &body)
		require.NoError(t, err)

		assert.True(t, len(body) > 0)

	})

}
