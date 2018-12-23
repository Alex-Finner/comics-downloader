package sites

import (
	"github.com/Girbons/comics-downloader/pkg/core"
	"github.com/Girbons/mangarock"
	log "github.com/sirupsen/logrus"
)

func findChapterName(chapterID string, chapters []*mangarock.Chapter) (string, bool) {
	for _, chapter := range chapters {
		if chapter.OID == chapterID {
			return chapter.Name, true
		}
	}
	return "", false
}

func SetupMangaRock(c *core.Comic) {
	series := c.SplitURL()[4]
	chapterID := c.SplitURL()[6]

	client := mangarock.NewClient()
	if _, ok := c.Options["country"]; ok {
		client.SetOptions(c.Options)
	}
	// get info about the manga
	info, infoErr := client.Info(series)
	if infoErr != nil {
		log.Error("[MANGAROCK] Cannot retrieve info for series: ", series, infoErr)
	}
	// retrieve pages
	pages, pagesErr := client.Pages(chapterID)
	if pagesErr != nil {
		log.Error("[MANGAROCK] Cannot retrieve pages for chapter: ", chapterID, pagesErr)
	}

	name := info.Data.Name
	chapter, found := findChapterName(chapterID, info.Data.Chapters)

	if !found {
		log.Info("[MANGAROCK] Chapter not found")
		chapter = chapterID
	}

	c.SetInfo(name, chapter, "")
	c.SetImageLinks(pages.Data)
}