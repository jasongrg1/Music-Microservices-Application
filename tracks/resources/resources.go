package resources

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"tracks/repository"
)

func updateCell(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var c repository.Cell
	if err := json.NewDecoder(r.Body).Decode(&c); err == nil {
		if id == c.Id {
			if n := repository.Update(c); n > 0 {
				w.WriteHeader(204) /* No Content */
			} else if n := repository.Insert(c); n > 0 {
				w.WriteHeader(201) /* Created */
			} else {
				w.WriteHeader(500) /* Internal Server Error */
			}
		} else {
			w.WriteHeader(400) /* Bad Request */
		}
	} else {
		w.WriteHeader(400) /* Bad Request */
	}
}

func readCell(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	if c, n := repository.Read(id); n > 0 {
		d := repository.Cell{Id: c.Id, Audio: c.Audio}
		w.WriteHeader(200) /* OK */
		json.NewEncoder(w).Encode(d)
	} else if n == 0 {
		w.WriteHeader(404) /* Not Found */
	} else {
		w.WriteHeader(500) /* Internal Server Error */
	}
}

func readAllCells(w http.ResponseWriter, r *http.Request) {
	songs, n := repository.ReadAll()
	if n>= 0 {
		w.WriteHeader(200) /* OK */
		if n==0 {
			songs = []string{}
		}
		json.NewEncoder(w).Encode(songs)
	} else {
		w.WriteHeader(500) /* Internal Server Error */
	}
}

func deleteCell(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
    if _, n := repository.Read(id); n > 0 {
		if n := repository.Delete(id); n > 0{
			w.WriteHeader(204) /* No Content */
		} else {
			w.WriteHeader(500) /* Internal Server Error */
		}
	} else if n == 0 {
		w.WriteHeader(404) /* Not Found */
	} else {
		w.WriteHeader(500) /* Internal Server Error */
	}
}

func Router() http.Handler {
	r := mux.NewRouter()
	/* Store */
	r.HandleFunc("/tracks/{id}", updateCell).Methods("PUT")
	/* Document */
	r.HandleFunc("/tracks/{id}", readCell).Methods("GET")

	r.HandleFunc("/tracks", readAllCells).Methods("GET")

	r.HandleFunc("/tracks/{id}", deleteCell).Methods("DELETE")
	return r
}
