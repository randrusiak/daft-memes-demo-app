package main

import (
	"encoding/json"
	"io"
	"net/http"
	"os"

	"github.com/dchest/uniuri"
)

// send a payload of JSON content
func (a *App) respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

// send a JSON error message
func (a *App) respondWithError(w http.ResponseWriter, code int, message string) {
	a.respondWithJSON(w, code, map[string]string{"error": message})

	a.Log.Printf("error: code %d, message %s", code, message)
}

// addMeme adds meme to database
func (a *App) addMeme(w http.ResponseWriter, r *http.Request) {
	var m meme

	file, header, err := r.FormFile("file")
	title := r.FormValue("title")
	filePrefix := uniuri.New()
	imagePath := "./public/" + filePrefix + header.Filename

	if err != nil {
		a.Log.Fatal(err)
	}
	defer file.Close()

	// copy image to local dir
	f, err := os.OpenFile(imagePath, os.O_WRONLY|os.O_CREATE, 0666)
	defer f.Close()
	_, err = io.Copy(f, file)

	if err != nil {
		a.Log.Fatal(err)
	}

	m.Title = title
	m.ImagePath = imagePath

	if err := m.createProduct(a.DB); err != nil {
		a.respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	a.respondWithJSON(w, http.StatusCreated, m)
}

// getAllMemes returns all memes from the database
func (a *App) getAllMemes(w http.ResponseWriter, r *http.Request) {
	a.Log.Printf("Handle GET all memes")
	memes, err := getAllMemes(a.DB)
	if err != nil {
		a.respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	a.respondWithJSON(w, http.StatusOK, memes)
}
