# DaftAcademy Demo App

Prosta aplikacja demo przygotowana na potrzeby zajęć DaftAcademy - Google Cloud Computing Foundations. 

Aplikacja składa się z dwóch części:
- frontend (React)
- backend  (Golang)


Podczas zajęć nie będzie konieczności uruchamiana aplikacji lokalnie. Jeśli jednak chciałbyś przetestować aplikacje możesz to zrobić wykonujac polecenie:  `docker-compose up`


Więcej szczegółów: 
 - [daftacademy.pl](https://daftacademy.pl/courses/ZPptVZ)
 - [daftacademy-gcp-terraform](https://github.com/randrusiak/daftacademy-gcp-terraform)

## Frontend

Na czas zajęć obraz aplikacji będzie publicznie dostępny pod adresem: `eu.gcr.io/daftacademy-tf-intro/demo-frontend`

### Zmiennie środowiskowe

    REACT_APP_API_URL - adres do backendu

## Backend

Na czas zajęć obraz aplikacji będzie publicznie dostępny pod adresem: `eu.gcr.io/daftacademy-tf-intro/demo-backend`

### Zmiennie środowiskowe

    STORAGE_TYPE - typ storage (local/gcs)
    GCS_BUCKET_NAME - nazwa bucketa (wymagane jeśli STORAGE_TYPE=gcs)
    DB_NAME - nazwa bazy danych
    DB_USER - nazwa użytkownika 
    DB_PASS - hasło do użytkownika
    DB_HOST - adres bazy danych


