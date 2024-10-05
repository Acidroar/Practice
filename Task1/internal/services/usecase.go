package services

import (
	"encoding/json"
	"fmt"
	"log"
	rand2 "math/rand/v2"
	"net/http"

	"github.com/kenedyCO/Practice/internal/models"
)

type Client interface {
	GetRequest(url string) (*http.Response, error)
}

func (s *Service) GetRandomFilm() (*models.Film, error) {
	id := rand2.IntN(100) + 300

	response, err := s.client.GetRequest(fmt.Sprintf("https://api.kinopoisk.dev/v1.4/movie/%v", id))
	if err != nil {
		log.Println("GetRandomFilm: " + err.Error())

		return nil, err
	}
	defer func() {
		err = response.Body.Close()
		if err != nil {
			log.Println(err)
		}
	}()

	// Обработать ответ
	var film models.Film
	err = json.NewDecoder(response.Body).Decode(&film)
	if err != nil {
		log.Println(" json.NewDecoder: " + err.Error())

		return nil, err
	}

	return &film, nil
}
