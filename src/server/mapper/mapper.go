package mapper

import "demo/src/model"

type AlbumResponse struct {
	Title       string  `json:"title"`
	Artist      string  `json:"artist"`
	Price       float64 `json:"price"`
	Description string  `json:"albumDescription"`
}

func FromAlbum(album model.Album) AlbumResponse {
	return AlbumResponse{
		Title:       album.Title,
		Artist:      album.Artist,
		Price:       album.Price,
		Description: album.Description,
	}
}

func FromAlbums(albums []model.Album) []AlbumResponse {
	var response []AlbumResponse
	for _, album := range albums {
		response = append(response, FromAlbum(album))
	}

	return response
}
