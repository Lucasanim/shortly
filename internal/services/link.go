package services

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/Lucasanim/shortly/internal/cache"
	"github.com/Lucasanim/shortly/internal/models"
	"github.com/Lucasanim/shortly/internal/repository"
	"github.com/Lucasanim/shortly/internal/utils"
)

type LinkService struct {
	linkRepository repository.LinkRepository
	cache          cache.Cache
}

var LinkServiceImpl = &LinkService{
	linkRepository: repository.LinkRepository{},
	cache:          cache.Cache{},
}

func (ls *LinkService) Create(creationRequest models.CreateLink) models.Link {
	hash := utils.ToBase62(creationRequest.Url)

	link := ls.GetFromCache(hash)

	if (link != models.Link{}) {
		return link
	}

	link = models.Link{
		ID:           int(time.Now().UnixNano()),
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
	link := ls.GetFromCache(hash)

	if (link != models.Link{}) {
		return link.Url, nil
	}

	return "", errors.New("link not found")
}

func (ls *LinkService) GetFromCache(hash string) models.Link {
	fmt.Println("Getting url for hash: ", hash)
	var link models.Link

	cacheData, err := ls.cache.Get(hash)
	if err == nil && cacheData != "" {
		err = json.Unmarshal([]byte(cacheData), &link)
		if err == nil {
			fmt.Println("Found link on cache:", link)
			return link
		}
		log.Println("Error unmarshalling cached data:", err)
	}

	fmt.Println("Getting url from db: ", hash)
	link = ls.Get(hash)

	jsonData, err := json.Marshal(link)
	if err == nil {
		ls.cache.Set(hash, string(jsonData))
	} else {
		log.Println("Error marshalling link for cache:", err)
	}

	return link
}
