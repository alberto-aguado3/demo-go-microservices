package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAlbum_ComplexCalculation(t *testing.T) {
	t.Run("Given a price greater than 0 with factor 1, should return same price", func(t *testing.T) {
		album := Album{
			Artist: "artist1",
			Title:  "title1",
			Price:  5.48,
		}
		var factor float64 = 1
		//factor := 1.0

		expected := 5.48

		actualNumber := album.ComplexCalculation(factor)

		assert.Equal(t, expected, actualNumber)
	})

	t.Run("Given a price greater than 0 with factor 3, should return same price * 3", func(t *testing.T) {
		album := Album{
			Artist: "artist1",
			Title:  "title1",
			Price:  10.2,
		}
		var factor float64 = 3
		//factor := 3.0

		expected := 30.6

		actualNumber := album.ComplexCalculation(factor)

		assert.InDelta(t, expected, actualNumber, 0.01)
	})

}
