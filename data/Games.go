package data

import (
	"time"
)

type Game struct {
	Name        string `json:"name"`
	Publisher   string `json:"publisher"`
	ReleaseDate time.Time `json:"time"`
	Description string `json:"description"`
}
