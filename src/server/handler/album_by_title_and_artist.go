package handler

import (
	"demo/src/server/mapper"
	"demo/src/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

// ListAlbums - responds with the list of all albums as JSON.
func AlbumByTitleAndArtist(s service.AlbumService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// albums/artist/:artist/title/:title
		artist := ctx.Param("artist")
		title := ctx.Param("title")

		album, err := s.GetAlbumByArtistAndTitle(ctx, artist, title)
		if err != nil {
			ctx.AbortWithStatusJSON(err.GetStatusCode(), err.ToJSON())
			return
		}

		ctx.IndentedJSON(http.StatusOK, mapper.FromAlbum(*album))
	}
}
