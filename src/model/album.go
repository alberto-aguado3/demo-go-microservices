package model

import (
	"fmt"
	"math"
)

// album represents data about a record album.
type Album struct {
	ID          int     `db:"id"`
	Title       string  `json:"title" binding:"required" db:"title"`
	Artist      string  `json:"artist" binding:"required" db:"artist"`
	Price       float64 `json:"price" db:"price"`
	Description string  `json:"albumDescription,omitempty" db:"album_description"`
}

func (a *Album) String() string {
	return fmt.Sprintf("Title: %s, Artist: %s\n", a.Title, a.Artist)
}

func (a *Album) ComplexCalculation(factor float64) float64 {
	//sqrt( price**2 ) * factor
	return math.Sqrt(math.Pow(a.Price, 2)) * factor
}
