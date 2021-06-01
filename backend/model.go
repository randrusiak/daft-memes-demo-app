package main

import (
	"database/sql"
	"errors"
)

type meme struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	ImagePath string `json:"imagePath"`
}

func (m *meme) createMemesTable(db *sql.DB) error {
	_, err := db.Exec(`CREATE TABLE IF NOT EXISTS memes (
		id SERIAL,
		title TEXT NOT NULL,
		image_path TEXT NOT NULL,
		CONSTRAINT memes_pkey PRIMARY KEY (id))`)
	return err
}

func (m *meme) getMeme(db *sql.DB) error {
	return errors.New("Not implemented")
}

func (m *meme) updateMeme(db *sql.DB) error {
	return errors.New("Not implemented")
}

func (m *meme) deleteMeme(db *sql.DB) error {
	_, err := db.Exec("DELETE FROM memes WHERE id=$1", m.ID)
	return err
}

func (m *meme) createMeme(db *sql.DB) error {
	err := db.QueryRow("INSERT INTO memes(title, image_path) VALUES($1, $2) RETURNING id", m.Title, m.ImagePath).Scan(&m.ID)
	if err != nil {
		return err
	}
	return nil
}

func getAllMemes(db *sql.DB) ([]meme, error) {
	rows, err := db.Query("SELECT id, title, image_path FROM memes")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	memes := []meme{}

	for rows.Next() {
		var m meme
		if err := rows.Scan(&m.ID, &m.Title, &m.ImagePath); err != nil {
			return nil, err
		}
		memes = append(memes, m)
	}

	return memes, nil
}
