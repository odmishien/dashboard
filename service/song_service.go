package service

import (
	"github.com/gin-gonic/gin"

	"dashboard/db"
	"dashboard/entity"
)

type SongService struct{}

type Song entity.Song

func (s SongService) GetAll() ([]Song, error) {
	db := db.GetDB()
	var song []Song

	if err := db.Find(&song).Error; err != nil {
		return nil, err
	}

	return song, nil
}

func (s SongService) CreateModel(c *gin.Context) (Song, error) {
	db := db.GetDB()
	var song Song

	if err := c.BindJSON(&song); err != nil {
		return song, err
	}

	if err := db.Create(&song).Error; err != nil {
		return song, err
	}

	return song, nil
}

func (s SongService) GetByID(id string) (Song, error) {
	db := db.GetDB()
	var song Song

	if err := db.Where("id = ?", id).First(&song).Error; err != nil {
		return song, err
	}

	return song, nil
}

func (s SongService) UpdateByID(id string, c *gin.Context) (Song, error) {
	db := db.GetDB()
	var song Song

	if err := db.Where("id = ?", id).First(&song).Error; err != nil {
		return song, err
	}

	if err := c.BindJSON(&s); err != nil {
		return song, err
	}

	db.Save(&song)

	return song, nil
}

func (s SongService) DeleteByID(id string) error {
	db := db.GetDB()
	var song Song

	if err := db.Where("id = ?", id).Delete(&song).Error; err != nil {
		return err
	}

	return nil
}
