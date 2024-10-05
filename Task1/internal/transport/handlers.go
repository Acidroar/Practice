package transport

import (
	"fmt"
	"log"
	"net/http"

	"github.com/kenedyCO/Practice/internal/convertor"
	"github.com/kenedyCO/Practice/internal/models"
	"github.com/labstack/echo/v4"
)

type Service interface {
	GetRandomFilm() (*models.Film, error)
}

func (h *HttpServer) randFilm(c echo.Context) error {
	film, err := h.service.GetRandomFilm()
	if err != nil {
		log.Println("h.service.GetRandomFilm: " + err.Error())

		return c.HTML(http.StatusInternalServerError, "Internal Server Error")
	}

	resp := convertor.ToRandomFilmResponseFromFilm(film)
	// Backend end. Start Frontend.
	response := fmt.Sprintf(`
		<html>
			<head>
				<title>%v</title>
			</head>
			<body>
				<h1>%v</h1>
				<p>
					Ссылка на фильм: 
					<a href="https://www.kinopoisk.ru/film/%v">%v</a>
				</p>
				<img src="%v" alt="%v" />
			</body>
		</html>
	`, resp.Name, resp.Name, resp.Id, resp.Name, resp.Poster.URL, resp.Name)

	return c.HTML(http.StatusOK, response)
}

func (h *HttpServer) getTrailer(c echo.Context) error {
	return c.HTML(http.StatusOK, "Trailer")
}
