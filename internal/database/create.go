package database

import (
	"fmt"
	"log"
	"mangathrV2/ent"
)

func (d *Driver) CreateManga(mangaID, title, source, mapping string) (*ent.Manga, error) {
	u, err := d.client.Manga.
		Create().
		SetMangaID(mangaID).
		SetTitle(title).
		SetSource(source).
		SetMapping(mapping).
		Save(d.ctx)
	if err != nil {
		return nil, fmt.Errorf("failed creating user: %w", err)
	}
	log.Println("user was created: ", u)
	return u, nil
}

func (d *Driver) CreateChapter(chapterID, num string, manga *ent.Manga) error {
	err := d.client.Chapter.
		Create().
		SetChapterID(chapterID).
		SetNum(num).
		SetManga(manga).
		Exec(d.ctx)
	if err != nil {
		return fmt.Errorf("failed creating user: %w", err)
	}
	return nil
}
