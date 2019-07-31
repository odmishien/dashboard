package service

import (
	"github.com/gin-gonic/gin"

	"dashboard/db"
	"dashboard/entity"
)

type WishSongService struct{}

type WishSong entity.WishSong

func (s WishSongService) GetAll() ([]WishSong, error) {
	db := db.GetDB()
	var song []WishSong

	if err := db.Find(&song).Error; err != nil {
		return nil, err
	}

	return song, nil
}

func (s WishSongService) CreateModel(c *gin.Context) (WishSong, error) {
	db := db.GetDB()
	var song WishSong

	if err := c.BindJSON(&song); err != nil {
		return song, err
	}

	if err := db.Create(&song).Error; err != nil {
		return song, err
	}

	return song, nil
}

func (s WishSongService) GetByID(id string) (WishSong, error) {
	db := db.GetDB()
	var song WishSong

	if err := db.Where("id = ?", id).First(&song).Error; err != nil {
		return song, err
	}

	return song, nil
}

func (s WishSongService) UpdateByID(id string, c *gin.Context) (WishSong, error) {
	db := db.GetDB()
	var song WishSong

	if err := db.Where("id = ?", id).First(&song).Error; err != nil {
		return song, err
	}

	if err := c.BindJSON(&s); err != nil {
		return song, err
	}

	db.Save(&song)

	return song, nil
}

func (s WishSongService) DeleteByID(id string) error {
	db := db.GetDB()
	var song WishSong

	if err := db.Where("id = ?", id).Delete(&song).Error; err != nil {
		return err
	}

	return nil
}
