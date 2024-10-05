package convertor

import "github.com/kenedyCO/Practice/internal/models"

func ToRandomFilmResponseFromFilm(film *models.Film) *models.RandomFilmResponse {
	return &models.RandomFilmResponse{
		Id:     film.Id,
		Name:   film.Name,
		Poster: film.Poster,
	}
}
