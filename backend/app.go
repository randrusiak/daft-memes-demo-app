package main

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

type App struct {
	Router *mux.Router
	DB     *sql.DB
	Log    *log.Logger
}

func (a *App) Run(dbHost, dbPort, dbUser, dbPassword, dbName string) {
	var err error

	// create a new logger
	a.Log = log.New(os.Stdout, "meme-api ", log.LstdFlags)

	// prepare connection string
	connectionString :=
		fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", dbHost, dbPort, dbUser, dbPassword, dbName)

	a.DB, err = sql.Open("postgres", connectionString)
	defer a.DB.Close()

	// validate connection string
	if err != nil {
		a.Log.Fatal(err)
	}
	// validate db connection
	err = a.DB.Ping()
	if err != nil {
		a.Log.Fatal(err)
	}

	// create a new serve mux and register the handlers
	a.Router = mux.NewRouter()

	a.Router.HandleFunc("/memes", a.getAllMemes).Methods("GET")
	a.Router.HandleFunc("/meme", a.addMeme).Methods("POST")
	a.Router.HandleFunc("/meme/{id:[0-9]+}", a.deleteMeme).Methods("DELETE")

	if getEnv("STORAGE_TYPE", "local") == "local" {
		fs := http.FileServer(http.Dir("./public/"))
		a.Router.PathPrefix("/public/").Handler(http.StripPrefix("/public/", fs))
	}

	// // CORS
	headers := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	methods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"})
	origins := handlers.AllowedOrigins([]string{"*"})

	// create a new server
	s := http.Server{
		Addr:         ":8080",                                            // configure the bind address
		Handler:      handlers.CORS(headers, methods, origins)(a.Router), // set the default handler
		ErrorLog:     a.Log,                                              // set the logger for the server
		ReadTimeout:  5 * time.Second,                                    // max time to read request from the client
		WriteTimeout: 10 * time.Second,                                   // max time to write response to the client
		IdleTimeout:  120 * time.Second,                                  // max time for connections using TCP Keep-Alive
	}

	// start the server
	go func() {
		a.Log.Println("Starting server on", s.Addr)

		err := s.ListenAndServe()
		if err != nil {
			a.Log.Printf("Error starting server: %s\n", err)
			os.Exit(1)
		}
	}()

	// trap sigterm or interupt and gracefully shutdown the server
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, os.Kill)

	// Block until a signal is received.
	sig := <-c
	log.Println("Got signal:", sig)

	// gracefully shutdown the server, waiting max 30 seconds for current operations to complete
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	s.Shutdown(ctx)
}
