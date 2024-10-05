package clients

import (
	"log"
	"net/http"

	"github.com/kenedyCO/Practice/internal/clients/models"
)

func (c *Client) GetRequest(url string) (*http.Response, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Println("http.NewRequest: " + err.Error())

		return nil, err
	}

	req.Header.Set(models.TokenKinopoiskKey, models.TokenKinopoiskValue)

	response, err := c.Do(req)
	if err != nil {
		log.Println("client.Do: " + err.Error())

		return nil, err
	}

	return response, nil
}
