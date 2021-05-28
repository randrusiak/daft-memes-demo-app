package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"

	"github.com/dchest/uniuri"
	"github.com/gorilla/mux"
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
	a.Log.Printf("Handle POST meme")
	var m meme
	var imagePath string

	file, header, err := r.FormFile("file")
	title := r.FormValue("title")
	filePrefix := uniuri.New()

	if err != nil {
		a.Log.Fatal(err)
	}
	defer file.Close()

	if a.StorageType == "gcs" {

		imagePath, err = GCSUploadFile(file, filePrefix+header.Filename)
		if err != nil {
			a.Log.Fatal(err)
		}
	} else {
		imagePath = "./public/" + filePrefix + header.Filename
		f, err := os.OpenFile(imagePath, os.O_WRONLY|os.O_CREATE, 0666)
		defer f.Close()
		_, err = io.Copy(f, file)

		if err != nil {
			a.Log.Fatal(err)
		}
	}

	m.Title = title
	m.ImagePath = imagePath

	if err := m.createMeme(a.DB); err != nil {
		a.respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	a.respondWithJSON(w, http.StatusCreated, m)
}

func (a *App) deleteMeme(w http.ResponseWriter, r *http.Request) {
	a.Log.Printf("Handle DELETE meme")
	var m meme

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		msg := fmt.Sprintf("Invalid meme ID. Error: %s", err.Error())
		a.respondWithError(w, http.StatusBadRequest, msg)
		return
	}
	m.ID = id

	if err := m.deleteMeme(a.DB); err != nil {
		a.respondWithError(w, http.StatusInternalServerError, err.Error())
	}

	a.respondWithJSON(w, http.StatusOK, m)
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
