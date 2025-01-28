package services

import (
	"errors"
	"time"

	"github.com/Lucasanim/shortly/internal/models"
	"github.com/Lucasanim/shortly/internal/repository"
	"github.com/Lucasanim/shortly/internal/utils"
)

type LinkService struct {
	linkRepository repository.LinkRepository
}

var LinkServiceImpl = &LinkService{
	linkRepository: repository.LinkRepository{},
}

func (ls *LinkService) Create(creationRequest models.CreateLink) models.Link {
	hash := utils.ToBase62(creationRequest.Url)

	link := ls.Get(hash)

	if (link != models.Link{}) {
		return link
	}

	link = models.Link{
		Url:          creationRequest.Url,
		Hash:         hash,
		CreationDate: time.Now(),
	}

	ls.linkRepository.Create(link)

	return link
}

func (ls *LinkService) Get(hash string) models.Link {
	return ls.linkRepository.Get(hash)
}

func (ls *LinkService) GetUrl(hash string) (string, error) {
	link := ls.Get(hash)

	if (link != models.Link{}) {
		return link.Url, nil
	}

	return "", errors.New("link not found")
}
