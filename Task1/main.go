package main

import (
	"encoding/json"
	"fmt"
	"github.com/labstack/echo/v4"
	"log"
	"math"
	"math/rand"
	"net/http"
	"time"
)

func main() {
	httpServer := echo.New()

	httpServer.GET("/", randFilm)

	httpServer.Logger.Fatal(httpServer.Start(":8080"))

}

type Film struct {
	Name   string `json:"name"`
	Id     int    `json:"id"`
	Year   int    `json:"year"`
	Desc   string `json:"description"`
	Rating rating `json:"rating"`
	Poster Poster `json:"poster"`
}

type rating struct {
	Kp   float64 `json:"kp"`
	Imdb float64 `json:"imdb"`
}

type Poster struct {
	Url string `json:"url"`
}

func randFilm(c echo.Context) error {
	id := random(298, 400)
	req, err := http.NewRequest("GET", fmt.Sprintf("https://api.kinopoisk.dev/v1.4/movie/%v", id), nil)
	if err != nil {
		return err
	}
	req.Header.Set("X-API-KEY", "JQVDNCF-AE245RW-NQX8G9S-VSQ9REH")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer func() {
		err = resp.Body.Close()
		if err != nil {
			log.Println(err)
		}
	}()

	var film Film
	err = json.NewDecoder(resp.Body).Decode(&film)
	if err != nil {
		return err
	}
	fmt.Println(film)
	R(film.Rating.Kp)
	R(film.Rating.Imdb)
	response := fmt.Sprintf("	Случайный фильм - %v\n \r\n	Год - %v\n \r\n	Оценка Кинопоиска: %.1f\n \r\n	Оценка Imdb: %.1f\n \r\n	Описание: %v\n \r\n	Ссылка - https://www.kinopoisk.ru/film/%v\n	Постер фильма %v", film.Name, film.Year, film.Rating.Kp, film.Rating.Imdb, film.Desc, film.Id, film.Poster.Url)
	return c.String(http.StatusOK, response)
}

func random(min, max int) int {
	rand.New(rand.NewSource(time.Now().UnixNano()))
	return rand.Intn(max-min) + min
}

func R(x float64) float64 {
	x = math.Round(x*10) / 10
	return x
}
