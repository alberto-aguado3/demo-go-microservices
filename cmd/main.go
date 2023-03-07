package main

import (
	"demo/cmd/configuration"
	"demo/cmd/di"
	mysql "demo/src/repository/mysql"
	"demo/src/server/handler"
	"demo/src/server/mapper"
	"demo/src/service"
	goErrors "errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
)

func main() {
	conf, err := configuration.ReadEnv()
	if err != nil {
		panic(err)
	}

	sqlDb := mysql.NewMysqlDatabase(&conf.Sql)
	// if err := mysql.MigrateSQL(*sqlDb); err != nil {
	// 	panic(err)
	// }

	interactors := di.Bootstrap(sqlDb, conf)

	//set log level
	zerolog.SetGlobalLevel(zerolog.DebugLevel)

	//Set gin mode
	gin.SetMode(conf.GinMode)

	router := gin.Default()

	// router.Handle(http.MethodGet, "albums", handler.ListAlbums(interactors.Album))

	rAlbums := router.Group("albums/")

	//router.GET("albums", handler.ListAlbums(interactors.Album))
	rAlbums.GET("", handler.ListAlbums(interactors.Album))
	rAlbums.GET("artist/:artist/title/:title", auth, handler.AlbumByTitleAndArtist(interactors.Album))
	router.POST("albums", handler.CreateAlbum(interactors.Album))

	router.GET("demoHttp", handler.DemoRequest(interactors.Album))

	//From here, example with middleware for DI
	routerWithServices := router.Group("")

	routerWithServices.Use(ServiceMiddleware(interactors))

	routerWithServices.GET("ping", ping)
	routerWithServices.GET("pong", auth, pong)

	// interactors.Transactions

	router.Run("localhost:8080")

}

func ServiceMiddleware(services di.Services) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Set("albumService", services.Album)

		ctx.Next()
	}
}

// Function handler to showcase that interactors can be recovered from context.
var ping = func(ctx *gin.Context) {
	albumService := ctx.MustGet("albumService").(service.AlbumService)

	albums, err := albumService.GetAllAlbums(ctx)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.IndentedJSON(http.StatusOK, mapper.FromAlbums(albums))
}

// An example of a middleware, that protects an endpoint (see main)
var auth = func(ctx *gin.Context) {
	password := ctx.GetHeader("Authorization")

	if password == "" {
		ctx.AbortWithError(http.StatusUnauthorized, goErrors.New("no password provided"))
	}

	ctx.Next()
}

// An endpoint that is protected by auth middleware.
var pong = func(ctx *gin.Context) {
	ctx.String(http.StatusOK, "Pong")
}
