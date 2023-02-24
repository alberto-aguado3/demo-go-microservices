package handler

import (
	"demo/src/server/mapper"
	"demo/src/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

// ListAlbums - responds with the list of all albums as JSON.
func ListAlbums(s service.AlbumService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		albums, err := s.GetAllAlbums(ctx)
		if err != nil {
			ctx.AbortWithError(http.StatusInternalServerError, err)
			return
		}

		ctx.IndentedJSON(http.StatusOK, mapper.FromAlbums(albums))
	}
}
