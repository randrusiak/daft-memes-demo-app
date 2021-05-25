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

func (m *meme) getMeme(db *sql.DB) error {
	return errors.New("Not implemented")
}

func (m *meme) updateMeme(db *sql.DB) error {
	return errors.New("Not implemented")
}

func (m *meme) deleteMeme(db *sql.DB) error {
	return errors.New("Not implemented")
}

func (m *meme) createProduct(db *sql.DB) error {
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
