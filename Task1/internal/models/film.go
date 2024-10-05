package models

type Film struct {
	Id     int    `json:"id"`
	Name   string `json:"name"`
	Poster Poster `json:"poster"`
	Videos Videos `json:"videos"`
}

type Poster struct {
	URL string `json:"url"`
}

type Videos struct {
	Trailers []Trailer `json:"trailers"`
}

type Trailer struct {
	URL  string `json:"url"`
	Name string `json:"name"`
}

type RandomFilmResponse struct {
	Id     int    `json:"id"`
	Name   string `json:"name"`
	Poster Poster `json:"poster"`
}
