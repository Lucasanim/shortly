package repository

import "github.com/Lucasanim/shortly/internal/models"

type LinkRepository struct{}

func (lr *LinkRepository) Create(link models.Link) {}

func (lr *LinkRepository) Get(hash string) models.Link {
	return models.Link{}
}
