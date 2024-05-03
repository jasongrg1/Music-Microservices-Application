package resources

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"net/url"
    "bytes"
	"cooltown/repository"
)


func retrieveSong(w http.ResponseWriter, r *http.Request) {

	requestBody := map[string]interface{} {}
	if err := json.NewDecoder(r.Body).Decode(&requestBody); err == nil {
		song, ok := requestBody["Audio"]
		
		if !ok {
			w.WriteHeader(400) /* Bad Request */
			return
		}
		payload := map[string]interface{}{
	    	"Audio":     song,
	    }
		
		jsonPayload, err := json.Marshal(payload)
		if err != nil {
            w.WriteHeader(500) /* Internal Server Error */
		}

		client := &http.Client{}

	    searchReq, err := http.NewRequest("POST", "http://127.0.0.1:3001/search", bytes.NewBuffer(jsonPayload))
        if err != nil {
            w.WriteHeader(500)
            return
		}
		searchReq.Header.Set("Content-Type", "application/json")
		
		searchResponse, err := client.Do(searchReq)
		if err != nil {
			w.WriteHeader(500) /* Internal Server Error */
			return
		}

		if searchResponse.StatusCode != http.StatusOK {
			w.WriteHeader(searchResponse.StatusCode)
			return
		}

		defer searchResponse.Body.Close()

		searchRes := repository.TitleStruct{}
		err = json.NewDecoder(searchResponse.Body).Decode(&searchRes)
		if err != nil {
			w.WriteHeader(500) /* Internal Server Error */
			return
		}

		//tracks request
		
		songID := searchRes.Title

		escapedID := url.QueryEscape(songID)


		trackReq, err := http.NewRequest("GET", "http://127.0.0.1:3000/tracks/" + escapedID, nil)
        if err != nil {
            w.WriteHeader(500) /* Internal Server Error */
            return
		}

		trackResponse, err := client.Do(trackReq)
		if err != nil {
			w.WriteHeader(500) /* Internal Server Error */
			return
		}

		if trackResponse.StatusCode != http.StatusOK {
			w.WriteHeader(trackResponse.StatusCode)
			return
		}

		defer trackResponse.Body.Close()

		trackRes := repository.AudioStruct{}
		err = json.NewDecoder(trackResponse.Body).Decode(&trackRes)
		if err != nil {
			w.WriteHeader(500) /* Internal Server Error */
			return
		}
		w.WriteHeader(200) /* OK */
		json.NewEncoder(w).Encode(trackRes)


	} else {
		w.WriteHeader(500) /* Internal Server Error */
	}
}



func Router() http.Handler {
	r := mux.NewRouter()
	r.HandleFunc("/cooltown", retrieveSong).Methods("POST")
	return r
}
