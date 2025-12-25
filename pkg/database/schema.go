package database

import "database/sql"

func TableCreate(db *sql.DB) error {
	queries := []string{
		`CREATE TABLE IF NOT EXISTS manga (
		id TEXT PRIMARY KEY ,
		title TEXT ,
		author TEXT,
		genres TEXT, 
		status TEXT,
		total_chapters INTERGER,
		description TEXT 
		)`,
		`CREATE TABLE IF NOT EXISTS user(
		id TEXT PRIMARY KEY ,
		username TEXT UNIQUE,
		password_hash TEXT,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		)`,
		`CREATE TABLE IF NOT EXISTS user_progress(
		user_id TEXT,
   		manga_id TEXT,
    	current_chapters INTEGER,
    	status TEXT,
    	updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    	PRIMARY KEY (user_id, manga_id)
	)`,
	}
	for _, q := range queries {
		_, err := db.Exec(q)
		if err != nil {
			return err
		}
	}

	return nil

}
