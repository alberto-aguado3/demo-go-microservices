package handler

import (
	"demo/pkg/errors"
	"demo/src/model"
	"demo/src/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

// ListAlbums - responds with the list of all albums as JSON.
func CreateAlbum(s service.AlbumService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var receivedAlbum model.Album
		//In the next line, the content received in the body will attempt to be mapped and assigned to variable "receivedAlbum".
		//Note the "&" operator, it's necessary
		if err := ctx.ShouldBindJSON(&receivedAlbum); err != nil {
			customErr := errors.NewFromError(err).WithMessage("Unexpected JSON received").WithStatusCode(http.StatusBadRequest)
			ctx.AbortWithStatusJSON(customErr.GetStatusCode(), customErr.ToJSON())
			return
		}

		err := s.StoreAlbum(ctx, receivedAlbum)
		if err != nil {
			ctx.AbortWithStatusJSON(err.GetStatusCode(), err.ToJSON())
			return
		}

		ctx.IndentedJSON(http.StatusCreated, nil)
	}
}
