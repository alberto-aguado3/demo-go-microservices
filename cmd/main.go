package main

import (
	"demo/cmd/configuration"
	"demo/cmd/di"
	mysql "demo/src/repository/mysql"
	"demo/src/server/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	conf, err := configuration.ReadEnv()
	if err != nil {
		panic(err)
	}

	sqlDb := mysql.NewMysqlDatabase(&conf.Sql)
	interactors := di.Bootstrap(sqlDb, conf)

	router := gin.Default()
	router.GET("albums", handler.ListAlbums(interactors.Album))
	router.GET("albums/artist/:artist/title/:title", handler.AlbumByTitleAndArtist(interactors.Album))
	router.POST("albums", handler.CreateAlbum(interactors.Album))

	router.GET("demoHttp", handler.DemoRequest(interactors.Album))

	router.Run("localhost:8080")
}
