package di

import (
	"demo/cmd/configuration"
	"demo/src/model"
	myHttp "demo/src/repository/http"
	mysql "demo/src/repository/mysql"
	"demo/src/service"
)

type Repositories struct {
	Album model.AlbumRepository
	Http  model.HttpRepository
}

type Services struct {
	Album service.AlbumService
}

func Bootstrap(db *mysql.SqlDatabase, config *configuration.Configuration) Services {
	return setupServices(db)
}

func setupServices(sqlDb *mysql.SqlDatabase) Services {
	repositories := setupRepositories(sqlDb)

	albumService := service.NewAlbum(repositories.Album, repositories.Http)

	return Services{
		Album: albumService,
	}
}

func setupRepositories(sql *mysql.SqlDatabase) Repositories {
	albumRepo := mysql.NewMysqlAlbumRepository(sql.Db, sql.Timeout)
	httpClient := myHttp.NewClient()

	return Repositories{
		Album: albumRepo,
		Http:  httpClient,
	}
}
