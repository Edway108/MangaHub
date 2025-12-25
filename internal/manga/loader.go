package manga

import (
	"MangaHub/pkg/models"
	"database/sql"
	"encoding/json"
	"os"
)

func LoadManga(db *sql.DB, path string) error {
	data, err := os.ReadFile(path)
	if err != nil {
		return err
	}
	var mangas []models.Manga
	err = json.Unmarshal(data, &mangas)
	if err != nil {
		return err
	}

	for _, m := range mangas {
		_, _ = db.Exec(
			`INSERT OR IGNORE into manga
			(id , title , author , genres,status, total_chapters, description) VALUES (?,?,?,?,?,?,?)`,
			m.ID, m.Title, m.Author, m.Genres, m.Status, m.TotalChapters, m.Description,
		)
	}
	return nil
}
