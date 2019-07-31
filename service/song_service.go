package service

import (
	"fmt"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/sclevine/agouti"

	"dashboard/entity"
)

type SongService struct{}

type Song entity.Song

func (s SongService) FindByName(name string) ([]entity.Song, error) {
	url := fmt.Sprintf("https://www.joysound.com/web/search/song?match=1&keyword=%s", name)
	songs, err := getSongFromURL(url)
	if err != nil {
		return nil, err
	}

	return songs, nil
}

func getSongFromURL(url string) ([]entity.Song, error) {
	var sl []entity.Song

	// 動的なサイトなのでChromeで実際に読む
	driver := agouti.ChromeDriver()
	err := driver.Start()
	if err != nil {
		fmt.Printf("Failed to start driver: %v", err)
	}
	defer driver.Stop()

	page, err := driver.NewPage(agouti.Browser("chrome"))
	if err != nil {
		fmt.Printf("Failed to open page: %v", err)
	}

	err = page.Navigate(url)
	if err != nil {
		fmt.Printf("Failed to navigate: %v", err)
	}

	// contentにHTMLが入る
	content, err := page.HTML()
	if err != nil {
		fmt.Printf("Failed to get html: %v", err)
	}

	reader := strings.NewReader(content)
	doc, _ := goquery.NewDocumentFromReader(reader)

	doc.Find("h3.jp-cmp-music-title-001").Each(func(_ int, s *goquery.Selection) {
		var song entity.Song
		fmt.Println(s)
		text := s.Find("h3.jp-cmp-music-title-001").Text()
		slice := strings.Split(text, "／")
		song.Name = slice[0]
		song.Artist = slice[1]
		sl = append(sl, song)
	})
	return sl, nil
}
