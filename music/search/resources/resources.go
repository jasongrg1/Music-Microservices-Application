package resources

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"bytes"
	"search/repository"
)


const (
KEY = "test"
url = "https://api.audd.io"
)


func searchSong(w http.ResponseWriter, r *http.Request) {
	requestBody := map[string]interface{} {}
	if err := json.NewDecoder(r.Body).Decode(&requestBody); err == nil {
		song, ok := requestBody["Audio"]

		if !ok {
			w.WriteHeader(400) /* Bad Request */
			return
		}

	    apiBody := map[string]interface{}{
	    	"api_token": KEY,
	    	"audio":     song,
	    }

		apiBodyBytes, err := json.Marshal(apiBody)
		if err != nil {
            w.WriteHeader(500) /* Internal Server Error */
			return
		}

		// create post request

		request, err := http.NewRequest("POST", url, bytes.NewBuffer(apiBodyBytes))
		if err != nil {
			w.WriteHeader(500) /* Internal Server Error */
			return
		}
		request.Header.Set("Content-Type", "application/json")

		// send request
		client := &http.Client{}
		response, err := client.Do(request)
		if err != nil {
			w.WriteHeader(500) /* Internal Server Error */
			return
		}
		defer response.Body.Close()


		responseRes := repository.ResponseStruct{}
		err = json.NewDecoder(response.Body).Decode(&responseRes)
		if err != nil {
			w.WriteHeader(500) /* Internal Server Error */
			return
		}

		if responseRes.Status != "success" {
			w.WriteHeader(500) /* Internal Server Error */
			return
		}

		if responseRes.Result.Title == ""{
			w.WriteHeader(404) /* Not Found */
			return
		}

		finalResponse := map[string]interface{} {"Id" : responseRes.Result.Title}
		w.WriteHeader(200)
		json.NewEncoder(w).Encode(finalResponse)

	} else { 
		w.WriteHeader(500) /* Internal Server Error */
	}
}




func Router() http.Handler {
	r := mux.NewRouter()
	r.HandleFunc("/search", searchSong).Methods("POST")
	return r
}
